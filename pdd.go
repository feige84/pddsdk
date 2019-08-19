package pddsdk

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"

	"github.com/tidwall/gjson"
)

const (
	//普通http请求网关
	HttpRouter = "https://gw-api.pinduoduo.com/api/router"

	//设置API读取失败时重试的次数,可以提高API的稳定性,默认为2次
	RestNumeric = 2
)

type ApiReq struct {
	ClientId     string
	ClientSecret string
	AccessToken  string
	Version      string
	DataType     string
	CacheLife    int64
	ReqCount     int64
	//GetCache    func(...interface{}) string  //FILE
	//SetCache    func(...interface{}) bool //FILE
	GetCache    func(string) string              //REDIS
	SetCache    func(string, interface{}, int64) //REDIS
	WriteErrLog func(ApiLog)
}

type ApiLog struct {
	Id          int64  `json:"id"`
	ClientId    string `json:"client_id"`
	AccessToken string `json:"access_token"`
	Version     string `json:"version"`
	DateType    string `json:"date_type"`
	Timestamp   string `json:"timestamp"`
	Sign        string `json:"sign"`
	Method      string `json:"method"`
	ParamJson   string `json:"param_json"`
	ApiCount    int64  `json:"api_count"`
	ApiErrorInfo
	Result string `json:"result"`
}

type ApiErrorInfo struct {
	ErrorCode int64  `json:"error_code"`
	ErrorMsg  string `json:"error_msg"`
	SubCode   int64  `json:"sub_code"`
	SubMsg    string `json:"sub_msg"`
	RequestId string `json:"request_id"`
}

//请求业务参数
type ApiParams map[string]string

func NewClient(clientId, clientSecret string) *ApiReq {
	return &ApiReq{
		ClientId:     clientId,
		ClientSecret: clientSecret,
	}
}

func (client *ApiReq) Execute(method string, params ApiParams) (*gjson.Result, *ApiErrorInfo) {

	apiErrInfo := &ApiErrorInfo{}

	// system params
	value := url.Values{}
	value.Add("type", method) //方法名，接口名。
	value.Add("client_id", client.ClientId)
	if client.AccessToken != "" {
		value.Add("access_token", client.AccessToken)
	}
	if client.Version != "" {
		value.Add("version", client.Version)
	} else {
		value.Add("version", "V1")
	}
	if client.DataType != "" {
		value.Add("data_type", client.DataType)
	} else {
		value.Add("data_type", "JSON")
	}
	value.Add("timestamp", fmt.Sprint(GetNow().Unix()))

	// api params
	for k, v := range params {
		value.Add(k, v)
	}

	args := []string{}
	cacheParams := []string{}
	for k, v := range value {
		args = append(args, k+v[0])
		if k != "timestamp" {
			cacheParams = append(cacheParams, v[0])
		}
	}

	//缓存key组合
	sort.Strings(cacheParams)
	cacheId := strings.Join(cacheParams, "_")
	cacheId = method + "." + MD5(cacheId)

	//先获取缓存
	var cacheData string
	if client.CacheLife > 0 && client.GetCache != nil {
		cacheData = client.GetCache(cacheId)
	}
	if cacheData == "" {
		// make sign
		sort.Strings(args)
		argsStr := strings.Join(args, "")
		value.Add("sign", MD5(client.ClientSecret+argsStr+client.ClientSecret))

		//开始请求
		client.ReqCount++ //请求次数+1
		fmt.Println("value.Encode():", value.Encode())
		response, err := client.httpSend("POST", HttpRouter, value.Encode())
		fmt.Println("response:", string(response))
		if err != nil {
			//重试N次
			if RestNumeric > 0 && client.ReqCount < RestNumeric {
				//fmt.Println("尝试重新请求", client.ReqCount)
				return client.Execute(method, params)
			}
			apiErrInfo.SubMsg = err.Error()
			return nil, apiErrInfo
		}

		apiError := gjson.GetBytes(response, "error_response")
		if apiError.Exists() {
			//如果存在error_response.说明出错了。
			apiErrInfo.ErrorCode = apiError.Get("error_code").Int()
			apiErrInfo.ErrorMsg = apiError.Get("error_msg").String()
			apiErrInfo.SubCode = apiError.Get("sub_code").Int()
			apiErrInfo.SubMsg = apiError.Get("sub_msg").String()
			apiErrInfo.RequestId = apiError.Get("request_id").String()

			if client.WriteErrLog != nil {
				//把错误日志记录到表里。
				//因为本项目会定时查询优惠券有效性。下架商品太多。所以把下架的错误信息排除记录。
				errLog := ApiLog{}
				errLog.ClientId = value.Get("client_id")
				errLog.Version = value.Get("version")
				errLog.DateType = value.Get("data_type")
				errLog.Timestamp = value.Get("timestamp")
				errLog.AccessToken = value.Get("access_token")
				errLog.Sign = value.Get("sign")
				errLog.Method = method
				errLog.ParamJson = value.Get("param_json")
				errLog.ApiCount = client.ReqCount
				errLog.ErrorCode = apiErrInfo.ErrorCode
				errLog.ErrorMsg = apiErrInfo.ErrorMsg
				errLog.SubCode = apiErrInfo.SubCode
				errLog.SubMsg = apiErrInfo.SubMsg
				errLog.Result = string(response)

				client.WriteErrLog(errLog)
			}
			if apiErrInfo.ErrorCode == 70031 || apiErrInfo.ErrorCode == 70032 {
				//调用超限
			}
			if apiErrInfo.ErrorCode == 65 || apiErrInfo.ErrorCode == 66 {
				//远程服务调用超时
				time.Sleep(500 * time.Millisecond)
				//if RestNumeric > 0 && client.ReqCount < RestNumeric {
				//	//fmt.Println("尝试重新请求1", client.ReqCount)
				//	return client.Execute(method, params)
				//}
			}
			return &apiError, apiErrInfo
		} else {
			//没有error_response，说明是正常的。
			if client.CacheLife > 0 && client.SetCache != nil {
				client.SetCache(cacheId, string(response), client.CacheLife)
			}
			returnResp := gjson.ParseBytes(response)
			client.ReqCount = 0 //还原请求次数
			return &returnResp, nil
		}
	} else {
		returnResp := gjson.Parse(cacheData)
		client.ReqCount = 0 //还原请求次数
		return &returnResp, nil
	}
}

//检查是否有错误码
func (client *ApiReq) CheckApiErr(resp *gjson.Result) *ApiErrorInfo {
	//这里是xxx_yyy_zzz_response.result节点内的内容检查.
	if resp != nil {
		if resp.Get("error_response").Exists() {
			errInfo := ApiErrorInfo{}
			if err := json.Unmarshal([]byte(resp.Raw), &errInfo); err != nil {
				errInfo.ErrorCode = 66666
				errInfo.ErrorMsg = ApiErrInfo[errInfo.ErrorCode].Error()
			}
			return &errInfo
		}
	}
	return nil
}

func (client *ApiReq) httpSend(method, router, param string) ([]byte, error) {
	var req *http.Request
	var err error
	if method == "POST" {
		req, err = http.NewRequest(method, router, strings.NewReader(param))
	} else {
		req, err = http.NewRequest(method, router+"?"+param, nil)
	}
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")
	httpClient := &http.Client{}
	httpClient.Timeout = 3 * time.Second
	response, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("请求错误:%d", response.StatusCode)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func MD5(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%X", has)
	return md5str
}
