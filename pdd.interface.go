package pddsdk

import (
	"fmt"
	"strings"
)

//pdd.ddk.goods.search（多多进宝商品查询）
func (client *ApiReq) DdkGoodsSearch(keyword string, goodsIds interface{}, pid, rangeList string, sortType, page, pageSize, catId, optId, merchantType int64, withCoupon, isBrandGoods bool) (*GoodsSearch, *ApiErrorInfo) {
	params := ApiParams{}
	//如果是链接。goods_id

	if goodsIdList, hasData := joinArr(goodsIds); hasData {
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

	if rangeList != "" {
		params["range_list"] = rangeList
	}
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
	if pid != "" {
		params["pid"] = pid //推广位id
	}
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

			goods.MallCouponId = detail.Get("mall_coupon_id").Int()                                 //店铺优惠券id
			goods.MallCouponDiscountPct = detail.Get("mall_coupon_discount_pct").Int()              //店铺折扣
			goods.MallCouponMinOrderAmount = detail.Get("mall_coupon_min_order_amount").Int()       //最小使用金额
			goods.MallCouponMaxDiscountAmount = detail.Get("mall_coupon_max_discount_amount").Int() //最大使用金额
			goods.MallCouponTotalQuantity = detail.Get("mall_coupon_total_quantity").Int()          //店铺券总量
			goods.MallCouponRemainQuantity = detail.Get("mall_coupon_remain_quantity").Int()        //店铺券余量
			goods.MallCouponStartTime = detail.Get("mall_coupon_start_time").Int()                  //店铺券使用开始时间
			goods.MallCouponEndTime = detail.Get("mall_coupon_end_time").Int()                      //店铺券使用结束时间
			goods.GoodsId = detail.Get("goods_id").Int()                                            //参与多多进宝的商品ID
			goods.GoodsName = detail.Get("goods_name").String()                                     //参与多多进宝的商品标题
			goods.GoodsDesc = detail.Get("goods_desc").String()                                     //参与多多进宝的商品描述
			goods.GoodsImageUrl = detail.Get("goods_image_url").String()                            //多多进宝商品主图
			goods.GoodsImageUrl = detail.Get("goods_image_url").String()                            //商品主图
			if detail.Get("goods_gallery_urls").IsArray() {
				for _, v := range detail.Get("goods_gallery_urls").Array() {
					goods.GoodsGalleryUrls = append(goods.GoodsGalleryUrls, v.String()) //商品轮播图
				}
			}
			goods.MinGroupPrice = detail.Get("min_group_price").Int()   //最低价sku的拼团价，单位为分
			goods.MinNormalPrice = detail.Get("min_normal_price").Int() //最低价sku的单买价，单位为分
			goods.MallName = detail.Get("mall_name").String()           //店铺名称
			goods.OptId = detail.Get("opt_id").Int()                    //商品标签ID，使用pdd.goods.opt.get接口获取
			goods.OptName = detail.Get("opt_name").String()             //商品标签名称
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
			goods.CouponMinOrderAmount = detail.Get("coupon_min_order_amount").Int() //优惠券门槛金额，单位为分
			goods.CouponDiscount = detail.Get("coupon_discount").Int()               //优惠券面额，单位为分
			goods.CouponTotalQuantity = detail.Get("coupon_total_quantity").Int()    //优惠券总数量
			goods.CouponRemainQuantity = detail.Get("coupon_remain_quantity").Int()  //优惠券剩余数量
			goods.CouponStartTime = detail.Get("coupon_start_time").Int()            //优惠券生效时间，UNIX时间戳
			goods.CouponEndTime = detail.Get("coupon_end_time").Int()                //优惠券失效时间，UNIX时间戳
			goods.PromotionRate = detail.Get("promotion_rate").Int()                 //佣金比例，千分比
			goods.GoodsEvalCount = detail.Get("goods_eval_count").Int()              //商品评价数
			goods.CatId = detail.Get("cat_id").Int()                                 //商品类目ID，使用pdd.goods.cats.get接口获取
			goods.SalesTip = detail.Get("sales_tip").String()                        //已售卖件数
			goods.MallId = detail.Get("mall_id").Int()                               //商家id
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
			goods.PlanType = detail.Get("plan_type").Int()                           //推广计划类型
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

//pdd.ddk.goods.promotion.url.generate（多多进宝推广链接生成）
func (client *ApiReq) DdkGoodsPromotionUrlGenerate(goodsIds int64, pid, customParameters string, zsDuoId int64, generateShortUrl, multiGroup, generateWeAppWebview, generateWeApp, GenerateWeiboAppWebview bool) (*GoodsPromotionUrl, *ApiErrorInfo) {
	params := ApiParams{}
	params["goods_id_list"] = "[" + fmt.Sprint(goodsIds) + "]" //商品ID列表。例如：[123456,123]，当入参带有goods_id_list字段，将不会以opt_id、 cat_id、keyword维度筛选商品
	params["p_id"] = pid                                       //推广位id
	if customParameters != "" {
		params["custom_parameters"] = customParameters //自定义参数
	}
	if zsDuoId > 0 {
		params["zs_duo_id"] = fmt.Sprint(zsDuoId) //招商多多客ID
	}
	if generateShortUrl {
		params["generate_short_url"] = "true" //是否生成短链接，true-是，false-否
	}
	if multiGroup {
		params["multi_group"] = "true" //true--生成多人团推广链接 false--生成单人团推广链接（默认false）1、单人团推广链接：用户访问单人团推广链接，可直接购买商品无需拼团。2、多人团推广链接：用户访问双人团推广链接开团，若用户分享给他人参团，则开团者和参团者的佣金均结算给推手
	}
	if generateWeAppWebview {
		params["generate_weapp_webview"] = "true" //是否生成唤起微信客户端链接，true-是，false-否，默认false
	}
	if generateWeApp {
		params["generate_we_app"] = "true" //是否生成小程序推广
	}
	if GenerateWeiboAppWebview {
		params["generate_weiboapp_webview"] = "true" //是否生成微博推广链接
	}

	resp, err := client.Execute("pdd.ddk.goods.promotion.url.generate", params)
	if err != nil {
		return nil, err
	}

	if apiErrInfo := client.CheckApiErr(resp); apiErrInfo != nil {
		return nil, apiErrInfo
	}

	goods := GoodsPromotionUrl{}

	goodsResp := resp.Get("goods_promotion_url_generate_response")
	if goodsResp.Exists() && goodsResp.Get("goods_promotion_url_list").IsArray() {
		if len(goodsResp.Get("goods_promotion_url_list").Array()) > 0 {

			urlInfo := goodsResp.Get("goods_promotion_url_list").Array()[0]
			/*
				err := json.Unmarshal([]byte(detail.Raw), &goods)
				if err != nil {
					fmt.Println("err:", err.Error())
				}

				fmt.Println("goods:", goods)
			*/
			goods.WeAppWebViewShortUrl = urlInfo.Get("we_app_web_view_short_url").String()       //唤起微信app推广短链接
			goods.WeAppWebViewUrl = urlInfo.Get("we_app_web_view_url").String()                  //唤起微信app推广链接
			goods.MobileShortUrl = urlInfo.Get("mobile_short_url").String()                      //唤醒拼多多app的推广短链接
			goods.MobileUrl = urlInfo.Get("mobile_url").String()                                 //唤醒拼多多app的推广长链接
			goods.ShortUrl = urlInfo.Get("short_url").String()                                   //推广短链接
			goods.Url = urlInfo.Get("url").String()                                              //推广长链接
			goods.WeiboAppWebViewShortUrl = urlInfo.Get("weibo_app_web_view_short_url").String() //微博推广短链接
			goods.WeiboAppWebViewUrl = urlInfo.Get("weibo_app_web_view_url").String()            //微博推广链接

			if urlInfo.Get("we_app_info").Exists() && urlInfo.Get("we_app_info").IsObject() {
				weapp := urlInfo.Get("we_app_info")
				goods.WeAppInfo = &WeAppInfo{}
				goods.WeAppInfo.WeAppIconUrl = weapp.Get("we_app_icon_url").String()          //小程序图片
				goods.WeAppInfo.BannerUrl = weapp.Get("banner_url").String()                  //Banner图
				goods.WeAppInfo.Desc = weapp.Get("desc").String()                             //描述
				goods.WeAppInfo.SourceDisplayName = weapp.Get("source_display_name").String() //来源名
				goods.WeAppInfo.PagePath = weapp.Get("page_path").String()                    //小程序path值
				goods.WeAppInfo.UserName = weapp.Get("user_name").String()                    //用户名
				goods.WeAppInfo.Title = weapp.Get("title").String()                           //小程序标题
				goods.WeAppInfo.AppId = weapp.Get("app_id").String()                          //拼多多小程序id
			}
			if urlInfo.Get("goods_detail").Exists() && urlInfo.Get("goods_detail").IsObject() {
				detail := urlInfo.Get("goods_detail")
				goods.GoodsDetail = &PromotionUrlGoodsDetail{}
				goods.GoodsDetail.HasMallCoupon = detail.Get("has_mall_coupon").Bool()
				goods.GoodsDetail.MallCouponId = detail.Get("mall_coupon_id").Int()
				goods.GoodsDetail.MallCouponDiscountPct = detail.Get("mall_coupon_discount_pct").Int()              //店铺券折扣
				goods.GoodsDetail.MallCouponMinOrderAmount = detail.Get("mall_coupon_min_order_amount").Int()       //最小使用金额
				goods.GoodsDetail.MallCouponMaxDiscountAmount = detail.Get("mall_coupon_max_discount_amount").Int() //最大使用金额
				goods.GoodsDetail.MallCouponTotalQuantity = detail.Get("mall_coupon_total_quantity").Int()          //店铺券总量
				goods.GoodsDetail.MallCouponRemainQuantity = detail.Get("mall_coupon_remain_quantity").Int()        //店铺券余量
				goods.GoodsDetail.MallCouponStartTime = detail.Get("mall_coupon_start_time").Int()                  //店铺券开始使用时间
				goods.GoodsDetail.MallCouponEndTime = detail.Get("mall_coupon_end_time").Int()                      //店铺券结束使用时间
				goods.GoodsDetail.CreateAt = detail.Get("create_at").Int()                                          //创建时间（unix时间戳）
				goods.GoodsDetail.GoodsId = detail.Get("goods_id").Int()                                            //商品id
				goods.GoodsDetail.GoodsName = detail.Get("goods_name").String()                                     //商品名称
				goods.GoodsDetail.GoodsDesc = detail.Get("goods_desc").String()                                     //商品描述
				goods.GoodsDetail.GoodsThumbnailUrl = detail.Get("goods_thumbnail_url").String()                    //商品缩略图
				goods.GoodsDetail.GoodsImageUrl = detail.Get("goods_image_url").String()                            //商品主图
				if detail.Get("goods_gallery_urls").IsArray() {
					for _, v := range detail.Get("goods_gallery_urls").Array() {
						goods.GoodsDetail.GoodsGalleryUrls = append(goods.GoodsDetail.GoodsGalleryUrls, v.String()) //商品轮播图
					}
				}
				goods.GoodsDetail.MinGroupPrice = detail.Get("min_group_price").Int()   //最小拼团价（单位为分）
				goods.GoodsDetail.MinNormalPrice = detail.Get("min_normal_price").Int() //最小单买价格（单位为分）
				goods.GoodsDetail.MallName = detail.Get("mall_name").String()           //店铺名字
				goods.GoodsDetail.MerchantType = detail.Get("merchant_type").Int()      //店铺类型，1-个人，2-企业，3-旗舰店，4-专卖店，5-专营店，6-普通店
				goods.GoodsDetail.CategoryId = detail.Get("category_id").Int()          //商品类目ID，使用pdd.goods.GoodsDetail.cats.get接口获取
				goods.GoodsDetail.CategoryName = detail.Get("category_name").String()   //商品类目名
				goods.GoodsDetail.OptId = detail.Get("opt_id").Int()                    //商品标签ID，使用pdd.goods.GoodsDetail.opts.get接口获取
				goods.GoodsDetail.OptName = detail.Get("opt_name").String()             //商品标签名
				if detail.Get("opt_ids").IsArray() {
					for _, v := range detail.Get("opt_ids").Array() {
						goods.GoodsDetail.OptIds = append(goods.GoodsDetail.OptIds, v.Int()) //商品标签id
					}
				}
				if detail.Get("cat_ids").IsArray() {
					for _, v := range detail.Get("cat_ids").Array() {
						goods.GoodsDetail.CatIds = append(goods.GoodsDetail.CatIds, v.Int()) //商品类目id
					}
				}
				goods.GoodsDetail.MallCps = detail.Get("mall_cps").Int()                             //该商品所在店铺是否参与全店推广，0：否，1：是
				goods.GoodsDetail.HasCoupon = detail.Get("has_coupon").Bool()                        //商品是否有优惠券 true-有，false-没有
				goods.GoodsDetail.CouponMinOrderAmount = detail.Get("coupon_min_order_amount").Int() //优惠券门槛价格，单位为分
				goods.GoodsDetail.CouponDiscount = detail.Get("coupon_discount").Int()               //优惠券面额，单位为分
				goods.GoodsDetail.CouponTotalQuantity = detail.Get("coupon_total_quantity").Int()    //优惠券总数量
				goods.GoodsDetail.CouponRemainQuantity = detail.Get("coupon_remain_quantity").Int()  //优惠券剩余数量
				goods.GoodsDetail.CouponStartTime = detail.Get("coupon_start_time").Int()            //优惠券生效时间，UNIX时间戳
				goods.GoodsDetail.CouponEndTime = detail.Get("coupon_end_time").Int()                //优惠券失效时间，UNIX时间戳
				goods.GoodsDetail.PromotionRate = detail.Get("promotion_rate").Int()                 //佣金比例，千分比
				goods.GoodsDetail.GoodsEvalCount = detail.Get("goods_eval_count").Int()              //商品评价数量
				goods.GoodsDetail.SalesTip = detail.Get("sales_tip").String()                        //已售卖件数
				goods.GoodsDetail.ActivityType = detail.Get("activity_type").Int()                   //活动类型，0-无活动;1-秒杀;3-限量折扣;12-限时折扣;13-大促活动;14-名品折扣;15-品牌清仓;16-食品超市;17-一元幸运团;18-爱逛街;19-时尚穿搭;20-男人帮;21-9块9;22-竞价活动;23-榜单活动;24-幸运半价购;25-定金预售;26-幸运人气购;27-特色主题活动;28-断码清仓;29-一元话费;30-电器城;31-每日好店;32-品牌卡;101-大促搜索池;102-大促品类分会场;
				if detail.Get("service_tags").IsArray() {
					for _, v := range detail.Get("service_tags").Array() {
						goods.GoodsDetail.ServiceTags = append(goods.GoodsDetail.ServiceTags, v.Int()) //服务标签: 4-送货入户并安装,5-送货入户,6-电子发票,9-坏果包赔,11-闪电退款,12-24小时发货,13-48小时发货,17-顺丰包邮,18-只换不修,19-全国联保,20-分期付款,24-极速退款,25-品质保障,26-缺重包退,27-当日发货,28-可定制化,29-预约配送,1000001-正品发票,1000002-送货入户并安装
					}
				}
				goods.GoodsDetail.CltCpnBatchSn = detail.Get("clt_cpn_batch_sn").String()            //店铺收藏券id
				goods.GoodsDetail.CltCpnStartTime = detail.Get("clt_cpn_start_time").Int()           //店铺收藏券起始时间
				goods.GoodsDetail.CltCpnEndTime = detail.Get("clt_cpn_end_time").Int()               //店铺收藏券截止时间
				goods.GoodsDetail.CltCpnQuantity = detail.Get("clt_cpn_quantity").Int()              //店铺收藏券总量
				goods.GoodsDetail.CltCpnRemainQuantity = detail.Get("clt_cpn_remain_quantity").Int() //店铺收藏券剩余量
				goods.GoodsDetail.CltCpnDiscount = detail.Get("clt_cpn_discount").Int()              //店铺收藏券面额，单位为分
				goods.GoodsDetail.CltCpnMinAmt = detail.Get("clt_cpn_min_amt").Int()                 //店铺收藏券使用门槛价格，单位为分
				goods.GoodsDetail.DescTxt = detail.Get("desc_txt").String()                          //描述分
				goods.GoodsDetail.ServTxt = detail.Get("serv_txt").String()                          //服务分
				goods.GoodsDetail.LgstTxt = detail.Get("lgst_txt").String()                          //物流分
				goods.GoodsDetail.PlanType = detail.Get("plan_type").Int()                           //推广计划类型 3:定向 4:招商
				goods.GoodsDetail.ZsDuoId = detail.Get("zs_duo_id").Int()                            //招商团长id
			}
		}
		return &goods, nil
	} else {
		errInfo := ApiErrorInfo{}
		errInfo.ErrorCode = 77777
		errInfo.SubMsg = ApiErrInfo[errInfo.ErrorCode].Error()
		return nil, &errInfo
	}
}

//pdd.ddk.order.list.increment.get（最后更新时间段增量同步推广订单信息）
/*
{"order_list_get_response":{"total_count":1,"order_list":[{"match_channel":5,"goods_price":1300,"promotion_rate":80,"type":0,"order_status":1,"order_create_time":1566463230,"order_settle_time":null,"order_verify_time":null,"order_group_success_time":1566463236,"order_amount":1300,"order_modify_at":1566463244,"auth_duo_id":0,"cpa_new":0,"goods_name":"高品质金属指陀螺钥匙扣开瓶器男士腰挂扣创意汽车钥匙链礼品","batch_no":"","goods_quantity":1,"goods_id":213159472,"goods_thumbnail_url":"http://t09img.yangkeduo.com/images/2018-05-17/fe69fd2298c5a492897168b4cf265079.jpeg","order_receive_time":null,"custom_parameters":"sEe7fkOiB4TWogQImRRfHFs3DTMsnERL5PsoTHfyrtNWIly67iR13PlVeqqPqAoUfLViamh2XWmJlB6sXZn/sw==","promotion_amount":104,"order_pay_time":1566463236,"group_id":859337392764094000,"duo_coupon_amount":0,"scene_at_market_fee":0,"order_status_desc":"已成团","fail_reason":null,"order_id":"diJCAbih9WeZFkMQXJklbg==","order_sn":"190822-337392764093890","p_id":"1817118_107021838","zs_duo_id":0}],"request_id":"15664634761188237"}}
*/
func (client *ApiReq) DdkOrderListGet(startUpdateTime, endUpdateTime, page, pageSize int64, returnCount bool) (*OrderList, *ApiErrorInfo) {
	params := ApiParams{}
	params["start_update_time"] = fmt.Sprint(startUpdateTime) //最近90天内多多进宝商品订单更新时间--查询时间开始。note：此时间为时间戳，指格林威治时间 1970 年01 月 01 日 00 时 00 分 00 秒(北京时间 1970 年 01 月 01 日 08 时 00 分 00 秒)起至现在的总秒数
	params["end_update_time"] = fmt.Sprint(endUpdateTime)     //查询结束时间，和开始时间相差不能超过24小时。note：此时间为时间戳，指格林威治时间 1970 年01 月 01 日 00 时 00 分 00 秒(北京时间 1970 年 01 月 01 日 08 时 00 分 00 秒)起至现在的总秒数
	if page <= 0 {
		page = 1
	}
	params["page"] = fmt.Sprint(page) //第几页，从1到10000，默认1，注：使用最后更新时间范围增量同步时，必须采用倒序的分页方式（从最后一页往回取）才能避免漏单问题。
	if pageSize <= 0 || pageSize > 100 {
		pageSize = 50
	}
	params["page_size"] = fmt.Sprint(pageSize) //返回的每页结果订单数，默认为100，范围为10到100，建议使用40~50，可以提高成功率，减少超时数量。

	if returnCount {
		params["return_count"] = "true"
	} else {
		params["return_count"] = "false" //是否返回总数，默认为true，如果指定false, 则返回的结果中不包含总记录数，通过此种方式获取增量数据，效率在原有的基础上有80%的提升。
	}

	resp, err := client.Execute("pdd.ddk.order.list.increment.get", params)
	if err != nil {
		return nil, err
	}

	if apiErrInfo := client.CheckApiErr(resp); apiErrInfo != nil {
		return nil, apiErrInfo
	}

	orderList := OrderList{}
	orderResp := resp.Get("order_list_get_response")
	if orderResp.Exists() && orderResp.Get("order_list").IsArray() {
		orderDetails := []OrderDetail{}
		for _, order := range orderResp.Get("order_list").Array() {
			orderDetail := OrderDetail{}
			orderDetail.OrderSn = order.Get("order_sn").String()                            //推广订单编号
			orderDetail.GoodsId = order.Get("goods_id").Int()                               //商品ID
			orderDetail.GoodsName = order.Get("goods_name").String()                        //商品标题
			orderDetail.GoodsThumbnailUrl = order.Get("goods_thumbnail_url").String()       //商品缩略图
			orderDetail.GoodsQuantity = order.Get("goods_quantity").Int()                   //购买商品的数量
			orderDetail.GoodsPrice = order.Get("goods_price").Int()                         //订单中sku的单件价格，单位为分
			orderDetail.OrderAmount = order.Get("order_amount").Int()                       //实际支付金额，单位为分
			orderDetail.PId = order.Get("p_id").String()                                    //推广位ID
			orderDetail.PromotionRate = order.Get("promotion_rate").Int()                   //佣金比例，千分比
			orderDetail.PromotionAmount = order.Get("promotion_amount").Int()               //佣金金额，单位为分
			orderDetail.OrderStatus = order.Get("order_status").Int()                       //订单状态： -1 未支付; 0-已支付；1-已成团；2-确认收货；3-审核成功；4-审核失败（不可提现）；5-已经结算；8-非多多进宝商品（无佣金订单）
			orderDetail.OrderStatusDesc = order.Get("order_status_desc").String()           //订单状态描述
			orderDetail.OrderCreateTime = order.Get("order_create_time").Int()              //订单生成时间，UNIX时间戳
			orderDetail.OrderPayTime = order.Get("order_pay_time").Int()                    //支付时间
			orderDetail.OrderGroupSuccessTime = order.Get("order_group_success_time").Int() //成团时间
			orderDetail.OrderVerifyTime = order.Get("order_verify_time").Int()              //审核时间
			orderDetail.OrderModifyAt = order.Get("order_modify_at").Int()                  //最后更新时间
			orderDetail.CustomParameters = order.Get("custom_parameters").String()          //自定义参数
			orderDetail.CpaNew = order.Get("cpa_new").Int()                                 //是否是 cpa 新用户，1表示是，0表示否
			orderDetails = append(orderDetails, orderDetail)
		}
		orderList.OrderList = orderDetails
		orderList.TotalCount = orderResp.Get("total_count").Int()
		return &orderList, nil
	} else {
		errInfo := ApiErrorInfo{}
		errInfo.ErrorCode = 77777
		errInfo.SubMsg = ApiErrInfo[errInfo.ErrorCode].Error()
		return nil, &errInfo
	}
}

//pdd.ddk.merchant.list.get（多多客查店铺列表接口）这里只是为了获取店铺信息写的接口。要获取完整信息。重新写一个。
func (client *ApiReq) DdkMallListGet(mallIds, merchantType interface{}, catId, hasCoupon, page, pageSize, queryRangeStr int64, hasCltCpn bool) (*MallList, *ApiErrorInfo) {
	params := ApiParams{}
	if mallIdList, hasData := joinArr(mallIds); hasData {
		//传入ID精准查询
		params["mall_id_list"] = mallIdList //店铺id
	}
	if merchantTypeList, hasData := joinArr(merchantType); hasData {
		//传入ID精准查询
		params["merchant_type_list"] = merchantTypeList //店铺id
	}
	if catId > 0 {
		params["cat_id"] = fmt.Sprint(catId) //商品类目ID，使用pdd.goods.cats.get接口获取
	}
	if hasCoupon > 0 {
		params["has_coupon"] = fmt.Sprint(hasCoupon) //是否有优惠券 （0 所有；1 必须有券。）
	}
	if queryRangeStr > 0 {
		params["query_range_str"] = fmt.Sprint(queryRangeStr)
		// 查询范围
		// 0----商品拼团价格区间；
		// 1----商品券后价价格区间；
		// 2----佣金比例区间；
		// 3----优惠券金额区间；
		// 4----加入多多进宝时间区间；
		// 5----销量区间；
		// 6----佣金金额区间
	}
	if page <= 0 {
		page = 1
	}
	params["page"] = fmt.Sprint(page)
	if pageSize <= 0 || pageSize > 100 {
		pageSize = 50
	}
	params["page_size"] = fmt.Sprint(pageSize)
	if hasCltCpn {
		params["has_clt_cpn"] = "true" //是否有店铺收藏券 （0 所有；1 必须有券）
	}

	//range_vo_list 这个参数文档说是筛选范围。并没有给样例。

	resp, err := client.Execute("pdd.ddk.merchant.list.get", params)
	if err != nil {
		return nil, err
	}

	if apiErrInfo := client.CheckApiErr(resp); apiErrInfo != nil {
		return nil, apiErrInfo
	}

	mallList := MallList{}
	mallResp := resp.Get("merchant_list_response")
	if mallResp.Exists() && mallResp.Get("mall_search_info_vo_list").IsArray() {
		mallInfoList := []MallInfo{}
		for _, mall := range mallResp.Get("mall_search_info_vo_list").Array() {
			mallInfo := MallInfo{}
			mallInfo.GoodsNum = mall.Get("goods_num").Int()         //商品数
			mallInfo.ImgUrl = mall.Get("img_url").String()          //店铺logo
			mallInfo.MallId = mall.Get("mall_id").Int()             //店铺id
			mallInfo.MallName = mall.Get("mall_name").String()      //店铺名称
			mallInfo.MallRate = mall.Get("mall_rate").Int()         //全店推广佣金
			mallInfo.MerchantType = mall.Get("merchant_type").Int() //店铺类型，1-个人，2-企业，3-旗舰店，4-专卖店，5-专营店，6-普通店
			mallInfo.SalesTip = mall.Get("sales_tip").String()      //销量
			mallInfo.DescTxt = mall.Get("desc_txt").String()        //描述评分
			mallInfo.ServTxt = mall.Get("serv_txt").String()        //服务评分
			mallInfo.LgstTxt = mall.Get("lgst_txt").String()        //物流评分
			mallInfoList = append(mallInfoList, mallInfo)
		}
		mallList.MallList = mallInfoList
		mallList.Total = mallResp.Get("total_count").Int()
		return &mallList, nil
	} else {
		errInfo := ApiErrorInfo{}
		errInfo.ErrorCode = 77777
		errInfo.SubMsg = ApiErrInfo[errInfo.ErrorCode].Error()
		return nil, &errInfo
	}
}

func joinArr(value interface{}) (string, bool) {
	var returnStr string
	if value != nil {
		switch result := value.(type) {
		case int:
			if result > 0 {
				returnStr = "[" + fmt.Sprint(result) + "]"
			}
		case int64:
			if result > 0 {
				returnStr = "[" + fmt.Sprint(result) + "]"
			}
		case []int:
			strArr := []string{}
			for _, v := range result {
				if v > 0 {
					strArr = append(strArr, fmt.Sprint(v))
				}
			}
			returnStr = "[" + strings.Join(strArr, ",") + "]"
		case []int64:
			strArr := []string{}
			for _, v := range result {
				if v > 0 {
					strArr = append(strArr, fmt.Sprint(v))
				}
			}
			if len(strArr) > 0 {
				returnStr = "[" + strings.Join(strArr, ",") + "]"
			}
		case string:
			if result != "" {
				returnStr = "[" + result + "]"
			}
		case []string:
			returnStr = "[" + strings.Join(result, ",") + "]"
		}
	}
	return returnStr, returnStr != ""
}
