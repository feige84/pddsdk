package pddsdk

import (
	"fmt"
	"strings"
)

//pdd.ddk.goods.search（多多进宝商品查询）
func (client *ApiReq) DdkGoodsSearch(keyword string, goodsIds interface{}, sortType, page, pageSize, catId, optId, merchantType int64, withCoupon, isBrandGoods bool) (*GoodsSearch, *ApiErrorInfo) {
	params := ApiParams{}
	//如果是链接。goods_id
	if goodsIds != nil {
		var goodsIdList string
		switch result := goodsIds.(type) {
		case int, int64:
			goodsIdList = "[" + fmt.Sprint(result) + "]"
		case []int:
			strArr := []string{}
			for _, v := range result {
				strArr = append(strArr, fmt.Sprint(v))
			}
			goodsIdList = "[" + strings.Join(strArr, ",") + "]"
		case []int64:
			strArr := []string{}
			for _, v := range result {
				strArr = append(strArr, fmt.Sprint(v))
			}
			goodsIdList = "[" + strings.Join(strArr, ",") + "]"
		case string:
			goodsIdList = "[" + result + "]"
		case []string:
			goodsIdList = "[" + strings.Join(result, ",") + "]"
		}
		//传入ID精准查询
		params["goods_id_list"] = goodsIdList //商品ID列表。例如：[123456,123]，当入参带有goods_id_list字段，将不会以opt_id、 cat_id、keyword维度筛选商品
	} else {
		params["keyword"] = keyword //商品关键词，与opt_id字段选填一个或全部填写
		if optId > 0 {
			params["opt_id"] = fmt.Sprint(optId) //商品标签类目ID，使用pdd.goods.opt.get获取
		}
		if catId > 0 {
			params["cat_id"] = fmt.Sprint(catId) //商品类目ID，使用pdd.goods.cats.get接口获取
		}
	}
	if page <= 0 {
		page = 1
	}
	params["page"] = fmt.Sprint(page) //默认值1，商品分页数
	if pageSize <= 0 || pageSize >= 100 {
		pageSize = 20
	}
	params["page_size"] = fmt.Sprint(pageSize) //默认100，每页商品数量
	if sortType <= 0 || sortType > 32 {
		sortType = 0
	}
	params["sort_type"] = fmt.Sprint(sortType)
	//排序方式:
	// 0-综合排序;
	// 1-按佣金比率升序;
	// 2-按佣金比例降序;
	// 3-按价格升序;
	// 4-按价格降序;
	// 5-按销量升序;
	// 6-按销量降序;
	// 7-优惠券金额排序升序;
	// 8-优惠券金额排序降序;
	// 9-券后价升序排序;
	// 10-券后价降序排序;
	// 11-按照加入多多进宝时间升序;
	// 12-按照加入多多进宝时间降序;
	// 13-按佣金金额升序排序;
	// 14-按佣金金额降序排序;
	// 15-店铺描述评分升序;
	// 16-店铺描述评分降序;
	// 17-店铺物流评分升序;
	// 18-店铺物流评分降序;
	// 19-店铺服务评分升序;
	// 20-店铺服务评分降序;
	// 27-描述评分击败同类店铺百分比升序;
	// 28-描述评分击败同类店铺百分比降序;
	// 29-物流评分击败同类店铺百分比升序;
	// 30-物流评分击败同类店铺百分比降序;
	// 31-服务评分击败同类店铺百分比升序;
	// 32-服务评分击败同类店铺百分比降序;

	if withCoupon {
		params["with_coupon"] = "true" //是否只返回优惠券的商品，false返回所有商品，true只返回有优惠券的商品
	} else {
		params["with_coupon"] = "false" //是否只返回优惠券的商品，false返回所有商品，true只返回有优惠券的商品
	}
	//params["range_list"] = keyword
	// 筛选范围列表 样例：[{"range_id":0,"range_from":1,"range_to":1500},{"range_id":1,"range_from":1,"range_to":1500}]
	// range_id枚举及描述：
	// 0，最小成团价
	// 1，券后价
	// 2，佣金比例
	// 3，优惠券价格
	// 4，广告创建时间
	// 5，销量
	// 6，佣金金额
	// 7，店铺描述分
	// 8，店铺物流分
	// 9，店铺服务分
	// 10，店铺描述分击败同行业百分比
	// 11，店铺物流分击败同行业百分比
	// 12，店铺服务分击败同行业百分比
	// 13，商品分
	// 17，优惠券/最小团购价
	// 18，过去两小时pv
	// 19，过去两小时销量

	if merchantType > 0 {
		params["merchant_type"] = fmt.Sprint(merchantType) //店铺类型，1-个人，2-企业，3-旗舰店，4-专卖店，5-专营店，6-普通店（未传为全部）
	}
	params["pid"] = keyword //推广位id
	//params["custom_parameters"] = keyword  //自定义参数
	//params["merchant_type_list"] = keyword //店铺类型数组
	if isBrandGoods {
		params["is_brand_goods"] = "true" //是否为品牌商品
	} else {
		params["is_brand_goods"] = "false" //是否为品牌商品
	}

	resp, err := client.Execute("pdd.ddk.goods.search", params)
	if err != nil {
		return nil, err
	}

	if apiErrInfo := client.CheckApiErr(resp); apiErrInfo != nil {
		return nil, apiErrInfo
	}

	var goodsSearch GoodsSearch

	goodsResp := resp.Get("goods_search_response")
	if goodsResp.Exists() && goodsResp.Get("goods_list").IsArray() {
		goodsSearch.RequestId = goodsResp.Get("request_id").String()
		goodsSearch.TotalCount = goodsResp.Get("total_count").Int()
		goodsList := []GoodsList{} //创建一个以goodsId为索引的字典
		for _, value := range goodsResp.Get("goods_list").Array() {
			goods := GoodsList{}
			goods.HasMallCoupon = value.Get("has_mall_coupon").Bool()
			goods.MallCouponId = value.Get("mall_coupon_id").Int()
			goods.MallCouponDiscountPct = value.Get("mall_coupon_discount_pct").Int()              //店铺券折扣
			goods.MallCouponMinOrderAmount = value.Get("mall_coupon_min_order_amount").Int()       //最小使用金额
			goods.MallCouponMaxDiscountAmount = value.Get("mall_coupon_max_discount_amount").Int() //最大使用金额
			goods.MallCouponTotalQuantity = value.Get("mall_coupon_total_quantity").Int()          //店铺券总量
			goods.MallCouponRemainQuantity = value.Get("mall_coupon_remain_quantity").Int()        //店铺券余量
			goods.MallCouponStartTime = value.Get("mall_coupon_start_time").Int()                  //店铺券开始使用时间
			goods.MallCouponEndTime = value.Get("mall_coupon_end_time").Int()                      //店铺券结束使用时间
			goods.CreateAt = value.Get("create_at").Int()                                          //创建时间（unix时间戳）
			goods.GoodsId = value.Get("goods_id").Int()                                            //商品id
			goods.GoodsName = value.Get("goods_name").String()                                     //商品名称
			goods.GoodsDesc = value.Get("goods_desc").String()                                     //商品描述
			goods.GoodsThumbnailUrl = value.Get("goods_thumbnail_url").String()                    //商品缩略图
			goods.GoodsImageUrl = value.Get("goods_image_url").String()                            //商品主图
			if value.Get("goods_gallery_urls").IsArray() {
				for _, v := range value.Get("goods_gallery_urls").Array() {
					goods.GoodsGalleryUrls = append(goods.GoodsGalleryUrls, v.String()) //商品轮播图
				}
			}
			goods.MinGroupPrice = value.Get("min_group_price").Int()   //最小拼团价（单位为分）
			goods.MinNormalPrice = value.Get("min_normal_price").Int() //最小单买价格（单位为分）
			goods.MallName = value.Get("mall_name").String()           //店铺名字
			goods.MerchantType = value.Get("merchant_type").Int()      //店铺类型，1-个人，2-企业，3-旗舰店，4-专卖店，5-专营店，6-普通店
			goods.CategoryId = value.Get("category_id").Int()          //商品类目ID，使用pdd.goods.cats.get接口获取
			goods.CategoryName = value.Get("category_name").String()   //商品类目名
			goods.OptId = value.Get("opt_id").Int()                    //商品标签ID，使用pdd.goods.opts.get接口获取
			goods.OptName = value.Get("opt_name").String()             //商品标签名
			if value.Get("opt_ids").IsArray() {
				for _, v := range value.Get("opt_ids").Array() {
					goods.OptIds = append(goods.OptIds, v.Int()) //商品标签id
				}
			}
			if value.Get("cat_ids").IsArray() {
				for _, v := range value.Get("cat_ids").Array() {
					goods.CatIds = append(goods.CatIds, v.Int()) //商品类目id
				}
			}
			goods.MallCps = value.Get("mall_cps").Int()                             //该商品所在店铺是否参与全店推广，0：否，1：是
			goods.HasCoupon = value.Get("has_coupon").Bool()                        //商品是否有优惠券 true-有，false-没有
			goods.CouponMinOrderAmount = value.Get("coupon_min_order_amount").Int() //优惠券门槛价格，单位为分
			goods.CouponDiscount = value.Get("coupon_discount").Int()               //优惠券面额，单位为分
			goods.CouponTotalQuantity = value.Get("coupon_total_quantity").Int()    //优惠券总数量
			goods.CouponRemainQuantity = value.Get("coupon_remain_quantity").Int()  //优惠券剩余数量
			goods.CouponStartTime = value.Get("coupon_start_time").Int()            //优惠券生效时间，UNIX时间戳
			goods.CouponEndTime = value.Get("coupon_end_time").Int()                //优惠券失效时间，UNIX时间戳
			goods.PromotionRate = value.Get("promotion_rate").Int()                 //佣金比例，千分比
			goods.GoodsEvalCount = value.Get("goods_eval_count").Int()              //商品评价数量
			goods.SalesTip = value.Get("sales_tip").String()                        //已售卖件数
			goods.ActivityType = value.Get("activity_type").Int()                   //活动类型，0-无活动;1-秒杀;3-限量折扣;12-限时折扣;13-大促活动;14-名品折扣;15-品牌清仓;16-食品超市;17-一元幸运团;18-爱逛街;19-时尚穿搭;20-男人帮;21-9块9;22-竞价活动;23-榜单活动;24-幸运半价购;25-定金预售;26-幸运人气购;27-特色主题活动;28-断码清仓;29-一元话费;30-电器城;31-每日好店;32-品牌卡;101-大促搜索池;102-大促品类分会场;
			if value.Get("service_tags").IsArray() {
				for _, v := range value.Get("service_tags").Array() {
					goods.ServiceTags = append(goods.ServiceTags, v.Int()) //服务标签: 4-送货入户并安装,5-送货入户,6-电子发票,9-坏果包赔,11-闪电退款,12-24小时发货,13-48小时发货,17-顺丰包邮,18-只换不修,19-全国联保,20-分期付款,24-极速退款,25-品质保障,26-缺重包退,27-当日发货,28-可定制化,29-预约配送,1000001-正品发票,1000002-送货入户并安装
				}
			}
			goods.CltCpnBatchSn = value.Get("clt_cpn_batch_sn").String()            //店铺收藏券id
			goods.CltCpnStartTime = value.Get("clt_cpn_start_time").Int()           //店铺收藏券起始时间
			goods.CltCpnEndTime = value.Get("clt_cpn_end_time").Int()               //店铺收藏券截止时间
			goods.CltCpnQuantity = value.Get("clt_cpn_quantity").Int()              //店铺收藏券总量
			goods.CltCpnRemainQuantity = value.Get("clt_cpn_remain_quantity").Int() //店铺收藏券剩余量
			goods.CltCpnDiscount = value.Get("clt_cpn_discount").Int()              //店铺收藏券面额，单位为分
			goods.CltCpnMinAmt = value.Get("clt_cpn_min_amt").Int()                 //店铺收藏券使用门槛价格，单位为分
			goods.DescTxt = value.Get("desc_txt").String()                          //描述分
			goods.ServTxt = value.Get("serv_txt").String()                          //服务分
			goods.LgstTxt = value.Get("lgst_txt").String()                          //物流分
			goods.PlanType = value.Get("plan_type").Int()                           //推广计划类型 3:定向 4:招商
			goods.ZsDuoId = value.Get("zs_duo_id").Int()                            //招商团长id
			goodsList = append(goodsList, goods)
		}
		goodsSearch.GoodsList = goodsList
		return &goodsSearch, nil
	} else {
		errInfo := ApiErrorInfo{}
		errInfo.ErrorCode = 77777
		errInfo.SubMsg = ApiErrInfo[errInfo.ErrorCode].Error()
		return nil, &errInfo
	}
}

//pdd.ddk.goods.detail（多多进宝商品详情查询）
func (client *ApiReq) DdkGoodsDetail(goodsIds int64, pid, customParameters string, zsDuoId, planType int64) (*GoodsDetail, *ApiErrorInfo) {
	params := ApiParams{}
	params["goods_id_list"] = "[" + fmt.Sprint(goodsIds) + "]" //商品ID列表。例如：[123456,123]，当入参带有goods_id_list字段，将不会以opt_id、 cat_id、keyword维度筛选商品
	if pid != "" {
		params["pid"] = pid //推广位id
	}
	if customParameters != "" {
		params["custom_parameters"] = customParameters //自定义参数
	}
	if zsDuoId > 0 {
		params["zs_duo_id"] = fmt.Sprint(zsDuoId) //招商多多客ID
	}
	if planType > 0 {
		params["plan_type"] = fmt.Sprint(planType) //佣金优惠券对应推广类型，3：专属 4：招商
	}
	resp, err := client.Execute("pdd.ddk.goods.detail", params)
	if err != nil {
		return nil, err
	}

	if apiErrInfo := client.CheckApiErr(resp); apiErrInfo != nil {
		return nil, apiErrInfo
	}

	goods := &GoodsDetail{}

	goodsResp := resp.Get("goods_detail_response")
	if goodsResp.Exists() && goodsResp.Get("goods_details").IsArray() {
		if len(goodsResp.Get("goods_details").Array()) > 0 {
			detail := goodsResp.Get("goods_details").Array()[0]
			goods.HasMallCoupon = detail.Get("has_mall_coupon").Bool()
			goods.MallCouponId = detail.Get("mall_coupon_id").Int()
			goods.MallCouponDiscountPct = detail.Get("mall_coupon_discount_pct").Int()              //店铺券折扣
			goods.MallCouponMinOrderAmount = detail.Get("mall_coupon_min_order_amount").Int()       //最小使用金额
			goods.MallCouponMaxDiscountAmount = detail.Get("mall_coupon_max_discount_amount").Int() //最大使用金额
			goods.MallCouponTotalQuantity = detail.Get("mall_coupon_total_quantity").Int()          //店铺券总量
			goods.MallCouponRemainQuantity = detail.Get("mall_coupon_remain_quantity").Int()        //店铺券余量
			goods.MallCouponStartTime = detail.Get("mall_coupon_start_time").Int()                  //店铺券开始使用时间
			goods.MallCouponEndTime = detail.Get("mall_coupon_end_time").Int()                      //店铺券结束使用时间
			goods.CreateAt = detail.Get("create_at").Int()                                          //创建时间（unix时间戳）
			goods.GoodsId = detail.Get("goods_id").Int()                                            //商品id
			goods.GoodsName = detail.Get("goods_name").String()                                     //商品名称
			goods.GoodsDesc = detail.Get("goods_desc").String()                                     //商品描述
			goods.GoodsThumbnailUrl = detail.Get("goods_thumbnail_url").String()                    //商品缩略图
			goods.GoodsImageUrl = detail.Get("goods_image_url").String()                            //商品主图
			if detail.Get("goods_gallery_urls").IsArray() {
				for _, v := range detail.Get("goods_gallery_urls").Array() {
					goods.GoodsGalleryUrls = append(goods.GoodsGalleryUrls, v.String()) //商品轮播图
				}
			}
			goods.MinGroupPrice = detail.Get("min_group_price").Int()   //最小拼团价（单位为分）
			goods.MinNormalPrice = detail.Get("min_normal_price").Int() //最小单买价格（单位为分）
			goods.MallName = detail.Get("mall_name").String()           //店铺名字
			goods.MerchantType = detail.Get("merchant_type").Int()      //店铺类型，1-个人，2-企业，3-旗舰店，4-专卖店，5-专营店，6-普通店
			goods.CategoryId = detail.Get("category_id").Int()          //商品类目ID，使用pdd.goods.cats.get接口获取
			goods.CategoryName = detail.Get("category_name").String()   //商品类目名
			goods.OptId = detail.Get("opt_id").Int()                    //商品标签ID，使用pdd.goods.opts.get接口获取
			goods.OptName = detail.Get("opt_name").String()             //商品标签名
			if detail.Get("opt_ids").IsArray() {
				for _, v := range detail.Get("opt_ids").Array() {
					goods.OptIds = append(goods.OptIds, v.Int()) //商品标签id
				}
			}
			if detail.Get("cat_ids").IsArray() {
				for _, v := range detail.Get("cat_ids").Array() {
					goods.CatIds = append(goods.CatIds, v.Int()) //商品类目id
				}
			}
			goods.MallCps = detail.Get("mall_cps").Int()                             //该商品所在店铺是否参与全店推广，0：否，1：是
			goods.HasCoupon = detail.Get("has_coupon").Bool()                        //商品是否有优惠券 true-有，false-没有
			goods.CouponMinOrderAmount = detail.Get("coupon_min_order_amount").Int() //优惠券门槛价格，单位为分
			goods.CouponDiscount = detail.Get("coupon_discount").Int()               //优惠券面额，单位为分
			goods.CouponTotalQuantity = detail.Get("coupon_total_quantity").Int()    //优惠券总数量
			goods.CouponRemainQuantity = detail.Get("coupon_remain_quantity").Int()  //优惠券剩余数量
			goods.CouponStartTime = detail.Get("coupon_start_time").Int()            //优惠券生效时间，UNIX时间戳
			goods.CouponEndTime = detail.Get("coupon_end_time").Int()                //优惠券失效时间，UNIX时间戳
			goods.PromotionRate = detail.Get("promotion_rate").Int()                 //佣金比例，千分比
			goods.GoodsEvalCount = detail.Get("goods_eval_count").Int()              //商品评价数量
			goods.SalesTip = detail.Get("sales_tip").String()                        //已售卖件数
			goods.ActivityType = detail.Get("activity_type").Int()                   //活动类型，0-无活动;1-秒杀;3-限量折扣;12-限时折扣;13-大促活动;14-名品折扣;15-品牌清仓;16-食品超市;17-一元幸运团;18-爱逛街;19-时尚穿搭;20-男人帮;21-9块9;22-竞价活动;23-榜单活动;24-幸运半价购;25-定金预售;26-幸运人气购;27-特色主题活动;28-断码清仓;29-一元话费;30-电器城;31-每日好店;32-品牌卡;101-大促搜索池;102-大促品类分会场;
			if detail.Get("service_tags").IsArray() {
				for _, v := range detail.Get("service_tags").Array() {
					goods.ServiceTags = append(goods.ServiceTags, v.Int()) //服务标签: 4-送货入户并安装,5-送货入户,6-电子发票,9-坏果包赔,11-闪电退款,12-24小时发货,13-48小时发货,17-顺丰包邮,18-只换不修,19-全国联保,20-分期付款,24-极速退款,25-品质保障,26-缺重包退,27-当日发货,28-可定制化,29-预约配送,1000001-正品发票,1000002-送货入户并安装
				}
			}
			goods.CltCpnBatchSn = detail.Get("clt_cpn_batch_sn").String()            //店铺收藏券id
			goods.CltCpnStartTime = detail.Get("clt_cpn_start_time").Int()           //店铺收藏券起始时间
			goods.CltCpnEndTime = detail.Get("clt_cpn_end_time").Int()               //店铺收藏券截止时间
			goods.CltCpnQuantity = detail.Get("clt_cpn_quantity").Int()              //店铺收藏券总量
			goods.CltCpnRemainQuantity = detail.Get("clt_cpn_remain_quantity").Int() //店铺收藏券剩余量
			goods.CltCpnDiscount = detail.Get("clt_cpn_discount").Int()              //店铺收藏券面额，单位为分
			goods.CltCpnMinAmt = detail.Get("clt_cpn_min_amt").Int()                 //店铺收藏券使用门槛价格，单位为分
			goods.DescTxt = detail.Get("desc_txt").String()                          //描述分
			goods.ServTxt = detail.Get("serv_txt").String()                          //服务分
			goods.LgstTxt = detail.Get("lgst_txt").String()                          //物流分
			goods.PlanType = detail.Get("plan_type").Int()                           //推广计划类型 3:定向 4:招商
			goods.ZsDuoId = detail.Get("zs_duo_id").Int()                            //招商团长id
		}
		return goods, nil
	} else {
		errInfo := ApiErrorInfo{}
		errInfo.ErrorCode = 77777
		errInfo.SubMsg = ApiErrInfo[errInfo.ErrorCode].Error()
		return nil, &errInfo
	}
}
