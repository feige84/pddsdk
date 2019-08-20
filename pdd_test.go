package pddsdk

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestExecute(t *testing.T) {
	//0: {duoId: 1817118, pid: "1817118_107021838", pidName: "花生小宝安卓"}
	//1: {duoId: 1817118, pid: "1817118_107021792", pidName: "花生小宝IOS"

	client := NewClient("7ae0ebbb60e6460d9444a56460a177db", "eef17075f318f9a3f1acfe04b42c9eb62453bc18")
	//client.GetCache = FileGetCache
	//client.SetCache = FileSetCache
	//client.WriteErrLog = InsertErrLog
	client.CacheLife = 0

	//jd.union.open.category.goods.get 商品类目查询
	//data, apiErr := client.DdkGoodsPromotionUrlGenerate(1116130524, "1817118_107021838", "", 0, true, true, true, true, true)
	//data, apiErr := client.DdkGoodsPromotionUrlGenerate(1116130524, "1817118_107021838", "", 0, false, false, false, false, false)
	//jsonData, jsonErr := json.Marshal(data)
	//fmt.Println("jsonData:", string(jsonData), jsonErr, apiErr)

	//jd.union.open.order.query 订单查询接口
	//data2, apiErr2 := client.DdkGoodsDetail(25579270053, "", "", 0, 0)
	//jsonData, jsonErr := json.Marshal(data2)
	//fmt.Println("jsonData:", string(jsonData), jsonErr, apiErr2)

	//jd.union.open.promotion.common.get 获取通用推广链接//这个pid一直都是这样，提示错误，用不了。
	//data3, apiErr3 := client.DdkGoodsSearch("", nil, "", 0, 1, 10, 0, 0, 0, false, false)
	//jsonData, jsonErr := json.Marshal(data3)
	//fmt.Println("jsonData:", string(jsonData), jsonErr, apiErr3)
	//fmt.Println("return:", data3, apiErr3)

	//jd.union.open.goods.promotiongoodsinfo.query 获取通用推广链接,批量的
	//data4, apiErr4 := client.DdkOrderListGet(1566277200, 1566277800, 1, 50, false)
	//jsonData, jsonErr := json.Marshal(data4)
	//fmt.Println("jsonData:", string(jsonData), jsonErr, apiErr4)

	//jd.union.open.goods.promotiongoodsinfo.query 获取通用推广链接,单个的
	data5, apiErr5 := client.DdkMallListGet(515020177, nil, 0, 0, 1, 1, 0, false)
	jsonData, jsonErr := json.Marshal(data5)
	fmt.Println("jsonData:", string(jsonData), jsonErr, apiErr5)

	//jd.union.open.goods.jingfen.query 京粉精选商品查询接口
	//data6, apiErr6 := client.NormalGetJFGoods(1, 1, 20, "inOrderCount30DaysSku", "desc")
	//fmt.Println("return:", data6, apiErr6)

	//jd.union.open.user.pid.get 获取PID //母账号无权限
	//data7, apiErr7 := client.NormalGetPid(1000218958, 1224241220, 1, "", "花生小宝")
	//fmt.Println("return:", data7, apiErr7)
}
