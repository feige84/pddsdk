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

	data6, apiErr6 := client.DdkMallGoodsListGet(879556298, 3, 20)
	jsonData, jsonErr := json.Marshal(apiErr6)
	fmt.Println("jsonData:", string(jsonData), jsonErr, data6)

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
	//2019/8/22 16:40:0   1566463200   1566463320
	//1566547200, 1566548700

	//data4, apiErr4 := client.DdkOrderListGet(1566748800, 1566799200, 1, 10, true)
	//jsonData, jsonErr := json.Marshal(data4)
	//fmt.Println("jsonData:", string(jsonData), jsonErr, apiErr4)

	//jd.union.open.goods.promotiongoodsinfo.query 获取通用推广链接,单个的
	//data5, apiErr5 := client.DdkMallListGet(515020177, nil, 0, 0, 1, 1, 0, false)
	//jsonData, jsonErr := json.Marshal(data5)
	//fmt.Println("jsonData:", string(jsonData), jsonErr, apiErr5)

	//jd.union.open.goods.jingfen.query 京粉精选商品查询接口
	//data6, apiErr6 := client.DdkOrderListRangeGet("2019-08-22 00:00:00", "2019-08-24 00:00:00", "", 100)
	//jsonData, jsonErr := json.Marshal(data6)
	//fmt.Println("jsonData:", string(jsonData), jsonErr, apiErr6)

	//jd.union.open.user.pid.get 获取PID //母账号无权限
	//data7, apiErr7 := client.NormalGetPid(1000218958, 1224241220, 1, "", "花生小宝")
	//fmt.Println("return:", data7, apiErr7)
}

/*
已成团
{"order_list_get_response":{"total_count":2,"order_list":[{"match_channel":5,"goods_price":890,"promotion_rate":100,"type":0,"order_status":1,"order_create_time":1566548624,"order_settle_time":null,"order_verify_time":null,"order_group_success_time":1566548631,"order_amount":890,"order_modify_at":1566548639,"auth_duo_id":0,"cpa_new":0,"goods_name":"磁吸数据线苹果vivo三合一安卓手机快充车载磁铁吸头强磁充电线器","batch_no":"","goods_quantity":1,"goods_id":11048105047,"goods_thumbnail_url":"http://t00img.yangkeduo.com/goods/images/2019-06-02/6c308ca422f995999940f298001374db.jpeg","order_receive_time":null,"custom_parameters":"sEe7fkOiB4TWogQImRRfHFs3DTMsnERL5PsoTHfyrtNWIly67iR13PlVeqqPqAoU5xbESXAWFf4o5YOxnagNx5On+KeDcZYU9KJ8BbS/QL4=","promotion_amount":89,"order_pay_time":1566548631,"group_id":860461037241593915,"duo_coupon_amount":0,"scene_at_market_fee":0,"order_status_desc":"已成团","fail_reason":null,"order_id":"UCmODJJCnlrspQTc9X9Kyg==","order_sn":"190823-461037241593915","p_id":"1817118_107021838","zs_duo_id":0},{"match_channel":5,"goods_price":1100,"promotion_rate":200,"type":0,"order_status":1,"order_create_time":1566548613,"order_settle_time":null,"order_verify_time":null,"order_group_success_time":1566548620,"order_amount":1100,"order_modify_at":1566548628,"auth_duo_id":0,"cpa_new":0,"goods_name":"【30款任选】指尖陀螺金属儿童手指陀螺螺旋陀螺成人创意减压玩具","batch_no":"","goods_quantity":1,"goods_id":75704543,"goods_thumbnail_url":"http://t00img.yangkeduo.com/goods/images/2018-08-01/aaa0839917734dd6b96ea4e0cfa6c7ad.jpeg","order_receive_time":null,"custom_parameters":"sEe7fkOiB4TWogQImRRfHFs3DTMsnERL5PsoTHfyrtNWIly67iR13PlVeqqPqAoU5xbESXAWFf4o5YOxnagNx5On+KeDcZYU9KJ8BbS/QL4=","promotion_amount":220,"order_pay_time":1566548620,"group_id":860059208172513890,"duo_coupon_amount":0,"scene_at_market_fee":0,"order_status_desc":"已成团","fail_reason":null,"order_id":"XB8n0J2CigpxntbXMdeElg==","order_sn":"190823-059208172513890","p_id":"1817118_107021838","zs_duo_id":0}],"request_id":"15667868556439988"}}

已收货
{"order_list_get_response":{"total_count":1,"order_list":[{"match_channel":5,"goods_price":1300,"promotion_rate":80,"type":0,"order_status":2,"order_create_time":1566463230,"order_settle_time":null,"order_verify_time":null,"order_group_success_time":1566463236,"order_amount":1300,"order_modify_at":1566797306,"auth_duo_id":0,"cpa_new":0,"goods_name":"高品质金属指陀螺钥匙扣开瓶器男士腰挂扣创意汽车钥匙链礼品","batch_no":"","goods_quantity":1,"goods_id":213159472,"goods_thumbnail_url":"http://t09img.yangkeduo.com/images/2018-05-17/fe69fd2298c5a492897168b4cf265079.jpeg","order_receive_time":1566797296,"custom_parameters":"sEe7fkOiB4TWogQImRRfHFs3DTMsnERL5PsoTHfyrtNWIly67iR13PlVeqqPqAoUfLViamh2XWmJlB6sXZn/sw==","promotion_amount":104,"order_pay_time":1566463236,"group_id":859337392764093890,"duo_coupon_amount":0,"scene_at_market_fee":0,"order_status_desc":"确认收货","fail_reason":null,"order_id":"diJCAbih9WeZFkMQXJklbg==","order_sn":"190822-337392764093890","p_id":"1817118_107021838","zs_duo_id":0},{"match_channel":5,"goods_price":1100,"promotion_rate":200,"type":0,"order_status":2,"order_create_time":1566548613,"order_settle_time":null,"order_verify_time":null,"order_group_success_time":1566548620,"order_amount":1100,"order_modify_at":1566797217,"auth_duo_id":0,"cpa_new":0,"goods_name":"【30款任选】指尖陀螺金属儿童手指陀螺螺旋陀螺成人创意减压玩具","batch_no":"","goods_quantity":1,"goods_id":75704543,"goods_thumbnail_url":"http://t00img.yangkeduo.com/goods/images/2018-08-01/aaa0839917734dd6b96ea4e0cfa6c7ad.jpeg","order_receive_time":1566797207,"custom_parameters":"sEe7fkOiB4TWogQImRRfHFs3DTMsnERL5PsoTHfyrtNWIly67iR13PlVeqqPqAoU5xbESXAWFf4o5YOxnagNx5On+KeDcZYU9KJ8BbS/QL4=","promotion_amount":220,"order_pay_time":1566548620,"group_id":860059208172513890,"duo_coupon_amount":0,"scene_at_market_fee":0,"order_status_desc":"确认收货","fail_reason":null,"order_id":"XB8n0J2CigpxntbXMdeElg==","order_sn":"190823-059208172513890","p_id":"1817118_107021838","zs_duo_id":0}],"request_id":"15667973546286085"}}
*/
