package pddsdk

import "errors"

var (
	ApiErrInfo = make(map[int64]error)
)

func init() {
	ApiErrInfo[10000] = errors.New("参数错误")                     //参数值有误，按照文档要求填写请求参数
	ApiErrInfo[10001] = errors.New("公共参数错误")                   //请检查请求的公共参数
	ApiErrInfo[10002] = errors.New("请求方法错误，仅支持POST")           //请使用POST请求
	ApiErrInfo[10010] = errors.New("应用不存在")                    //您的应用不存在
	ApiErrInfo[10011] = errors.New("应用已被驳回")                   //请前往拼多多开放平台查看应用驳回的原因，及时修改并重新提交应用，或者创建新的应用
	ApiErrInfo[10014] = errors.New("授权已被取消")                   //商家和您的授权关系已经取消了
	ApiErrInfo[10016] = errors.New("client_id不正确")             //请核查您的client_id是否正确
	ApiErrInfo[10019] = errors.New("access_token已过期")          //刷新access_token或者重新授权再次获取access_token
	ApiErrInfo[10035] = errors.New("access_token已过期")          //请刷新access_token或者重新授权获取access_token
	ApiErrInfo[20004] = errors.New("签名sign校验失败")               //请按照接入指南第三部分指导，生成签名
	ApiErrInfo[20005] = errors.New("ip无权访问接口，请加入ip白名单")        //把ip白名单加入白名单
	ApiErrInfo[20007] = errors.New("缺少必填请求参数")                 //请查看接入指南第三部分和API文档，核对公共参数和业务必填参数是否正确
	ApiErrInfo[20031] = errors.New("用户没有授权访问此接口")              //您创建的应用中不包含此接口，请查看API文档，了解相关权限包
	ApiErrInfo[20032] = errors.New("access_token或client_id错误") //检查access_token或client_id
	ApiErrInfo[50000] = errors.New("系统内部错误")                   //系统内部错误，请加群联系相关负责人
	ApiErrInfo[50001] = errors.New("业务服务错误")                   //请根据子错误判断错误详情，无法解决请联系相关负责人
	ApiErrInfo[50012] = errors.New("此API已下线")                  //此API已下线
	ApiErrInfo[70031] = errors.New("调用过于频繁，请调整调用频率")           //调用过于频繁，请调整调用频率
	ApiErrInfo[70032] = errors.New("调用过于频繁，请调整调用频率")           //调用过于频繁，请调整调用频率
	ApiErrInfo[70033] = errors.New("当前接口因系统维护，暂时下线，请稍后再试！")    //当前接口因系统维护，暂时下线，请稍后再试！
	ApiErrInfo[70034] = errors.New("当前用户存在风险接，禁止调用！")          //当前用户存在风险接，禁止调用！

	ApiErrInfo[66666] = errors.New("数据解析错误") //不是有效的json数据。自定义的。
	ApiErrInfo[77777] = errors.New("无结果")    //无错误的情况下。没有结果。自定义的。
}
