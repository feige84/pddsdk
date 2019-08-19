package pddsdk

type GoodsSearch struct {
	GoodsList  []GoodsList `json:"Goods_list"`
	TotalCount int64       `json:"total_count"`
	RequestId  string      `json:"request_id"`
}

type GoodsList struct {
	HasMallCoupon               bool     `json:"has_mall_coupon"`                //是否有店铺券
	MallCouponId                int64    `json:"MallCoupon_id"`                  //店铺券id
	MallCouponDiscountPct       int64    `json:"MallCoupon_discount_pct"`        //店铺券折扣
	MallCouponMinOrderAmount    int64    `json:"MallCoupon_min_order_amount"`    //最小使用金额
	MallCouponMaxDiscountAmount int64    `json:"MallCoupon_max_discount_amount"` //最大使用金额
	MallCouponTotalQuantity     int64    `json:"MallCoupon_total_quantity"`      //店铺券总量
	MallCouponRemainQuantity    int64    `json:"MallCoupon_remain_quantity"`     //店铺券余量
	MallCouponStartTime         int64    `json:"MallCoupon_start_time"`          //店铺券开始使用时间
	MallCouponEndTime           int64    `json:"MallCoupon_end_time"`            //店铺券结束使用时间
	CreateAt                    int64    `json:"create_at"`                      //创建时间（unix时间戳）
	GoodsId                     int64    `json:"Goods_id"`                       //商品id
	GoodsName                   string   `json:"Goods_name"`                     //商品名称
	GoodsDesc                   string   `json:"Goods_desc"`                     //商品描述
	GoodsThumbnailUrl           string   `json:"Goods_thumbnail_url"`            //商品缩略图
	GoodsImageUrl               string   `json:"Goods_image_url"`                //商品主图
	GoodsGalleryUrls            []string `json:"Goods_gallery_urls"`             //商品轮播图
	MinGroupPrice               int64    `json:"min_group_price"`                //最小拼团价（单位为分）
	MinNormalPrice              int64    `json:"min_normal_price"`               //最小单买价格（单位为分）
	MallName                    string   `json:"mall_name"`                      //店铺名字
	MerchantType                int64    `json:"merchant_type"`                  //店铺类型，1-个人，2-企业，3-旗舰店，4-专卖店，5-专营店，6-普通店
	CategoryId                  int64    `json:"category_id"`                    //商品类目ID，使用pdd.goods.cats.get接口获取
	CategoryName                string   `json:"category_name"`                  //商品类目名
	OptId                       int64    `json:"opt_id"`                         //商品标签ID，使用pdd.goods.opts.get接口获取
	OptName                     string   `json:"opt_name"`                       //商品标签名
	OptIds                      []int64  `json:"opt_ids"`                        //商品标签id
	CatIds                      []int64  `json:"cat_ids"`                        //商品类目id
	MallCps                     int64    `json:"mall_cps"`                       //该商品所在店铺是否参与全店推广，0：否，1：是
	HasCoupon                   bool     `json:"has_coupon"`                     //商品是否有优惠券 true-有，false-没有
	CouponMinOrderAmount        int64    `json:"coupon_min_order_amount"`        //优惠券门槛价格，单位为分
	CouponDiscount              int64    `json:"coupon_discount"`                //优惠券面额，单位为分
	CouponTotalQuantity         int64    `json:"coupon_total_quantity"`          //优惠券总数量
	CouponRemainQuantity        int64    `json:"coupon_remain_quantity"`         //优惠券剩余数量
	CouponStartTime             int64    `json:"coupon_start_time"`              //优惠券生效时间，UNIX时间戳
	CouponEndTime               int64    `json:"coupon_end_time"`                //优惠券失效时间，UNIX时间戳
	PromotionRate               int64    `json:"promotion_rate"`                 //佣金比例，千分比
	GoodsEvalCount              int64    `json:"Goods_eval_count"`               //商品评价数量
	SalesTip                    string   `json:"sales_tip"`                      //已售卖件数
	ActivityType                int64    `json:"activity_type"`                  //活动类型，0-无活动;1-秒杀;3-限量折扣;12-限时折扣;13-大促活动;14-名品折扣;15-品牌清仓;16-食品超市;17-一元幸运团;18-爱逛街;19-时尚穿搭;20-男人帮;21-9块9;22-竞价活动;23-榜单活动;24-幸运半价购;25-定金预售;26-幸运人气购;27-特色主题活动;28-断码清仓;29-一元话费;30-电器城;31-每日好店;32-品牌卡;101-大促搜索池;102-大促品类分会场;
	ServiceTags                 []int64  `json:"service_tags"`                   //服务标签: 4-送货入户并安装,5-送货入户,6-电子发票,9-坏果包赔,11-闪电退款,12-24小时发货,13-48小时发货,17-顺丰包邮,18-只换不修,19-全国联保,20-分期付款,24-极速退款,25-品质保障,26-缺重包退,27-当日发货,28-可定制化,29-预约配送,1000001-正品发票,1000002-送货入户并安装
	CltCpnBatchSn               string   `json:"CltCpn_batch_sn"`                //店铺收藏券id
	CltCpnStartTime             int64    `json:"CltCpn_start_time"`              //店铺收藏券起始时间
	CltCpnEndTime               int64    `json:"CltCpn_end_time"`                //店铺收藏券截止时间
	CltCpnQuantity              int64    `json:"CltCpn_quantity"`                //店铺收藏券总量
	CltCpnRemainQuantity        int64    `json:"CltCpn_remain_quantity"`         //店铺收藏券剩余量
	CltCpnDiscount              int64    `json:"CltCpn_discount"`                //店铺收藏券面额，单位为分
	CltCpnMinAmt                int64    `json:"CltCpn_min_amt"`                 //店铺收藏券使用门槛价格，单位为分
	DescTxt                     string   `json:"desc_txt"`                       //描述分
	ServTxt                     string   `json:"serv_txt"`                       //服务分
	LgstTxt                     string   `json:"lgst_txt"`                       //物流分
	PlanType                    int64    `json:"plan_type"`                      //推广计划类型 3:定向 4:招商
	ZsDuoId                     int64    `json:"zs_duo_id"`                      //招商团长id
}

type GoodsDetail struct {
	MallCouponId                int64    `json:"mall_coupon_id"`                  //店铺优惠券id
	MallCouponDiscountPct       int64    `json:"mall_coupon_discount_pct"`        //店铺折扣
	MallCouponMinOrderAmount    int64    `json:"mall_coupon_min_order_amount"`    //最小使用金额
	MallCouponMaxDiscountAmount int64    `json:"mall_coupon_max_discount_amount"` //最大使用金额
	MallCouponTotalQuantity     int64    `json:"mall_coupon_total_quantity"`      //店铺券总量
	MallCouponRemainQuantity    int64    `json:"mall_coupon_remain_quantity"`     //店铺券余量
	MallCouponStartTime         int64    `json:"mall_coupon_start_time"`          //店铺券使用开始时间
	MallCouponEndTime           int64    `json:"mall_coupon_end_time"`            //店铺券使用结束时间
	GoodsId                     int64    `json:"goods_id"`                        //参与多多进宝的商品ID
	GoodsName                   string   `json:"goods_name"`                      //参与多多进宝的商品标题
	GoodsDesc                   string   `json:"goods_desc"`                      //参与多多进宝的商品描述
	GoodsImageUrl               string   `json:"goods_image_url"`                 //多多进宝商品主图
	GoodsGalleryUrls            []string `json:"goods_gallery_urls"`              //商品轮播图
	MinGroupPrice               int64    `json:"min_group_price"`                 //最低价sku的拼团价，单位为分
	MinNormalPrice              int64    `json:"min_normal_price"`                //最低价sku的单买价，单位为分
	MallName                    string   `json:"mall_name"`                       //店铺名称
	OptId                       int64    `json:"opt_id"`                          //商品标签ID，使用pdd.goods.opt.get接口获取
	OptName                     string   `json:"opt_name"`                        //商品标签名称
	OptIds                      []int64  `json:"opt_ids"`                         //商品标签ID
	CatIds                      []int64  `json:"cat_ids"`                         //商品一~四级类目ID列表
	CouponMinOrderAmount        int64    `json:"coupon_min_order_amount"`         //优惠券门槛金额，单位为分
	CouponDiscount              int64    `json:"coupon_discount"`                 //优惠券面额，单位为分
	CouponTotalQuantity         int64    `json:"coupon_total_quantity"`           //优惠券总数量
	CouponRemainQuantity        int64    `json:"coupon_remain_quantity"`          //优惠券剩余数量
	CouponStartTime             int64    `json:"coupon_start_time"`               //优惠券生效时间，UNIX时间戳
	CouponEndTime               int64    `json:"coupon_end_time"`                 //优惠券失效时间，UNIX时间戳
	PromotionRate               int64    `json:"promotion_rate"`                  //佣金比例，千分比
	GoodsEvalCount              int64    `json:"goods_eval_count"`                //商品评价数
	CatId                       int64    `json:"cat_id"`                          //商品类目ID，使用pdd.goods.cats.get接口获取
	SalesTip                    string   `json:"sales_tip"`                       //已售卖件数
	MallId                      int64    `json:"mall_id"`                         //商家id
	ServiceTags                 []int64  `json:"service_tags"`                    //服务标签: 4-送货入户并安装,5-送货入户,6-电子发票,9-坏果包赔,11-闪电退款,12-24小时发货,13-48小时发货,17-顺丰包邮,18-只换不修,19-全国联保,20-分期付款,24-极速退款,25-品质保障,26-缺重包退,27-当日发货,28-可定制化,29-预约配送,1000001-正品发票,1000002-送货入户并安装
	CltCpnBatchSn               string   `json:"clt_cpn_batch_sn"`                //店铺收藏券id
	CltCpnStartTime             int64    `json:"clt_cpn_start_time"`              //店铺收藏券起始时间
	CltCpnEndTime               int64    `json:"clt_cpn_end_time"`                //店铺收藏券截止时间
	CltCpnQuantity              int64    `json:"clt_cpn_quantity"`                //店铺收藏券总量
	CltCpnRemainQuantity        int64    `json:"clt_cpn_remain_quantity"`         //店铺收藏券剩余量
	CltCpnDiscount              int64    `json:"clt_cpn_discount"`                //店铺收藏券面额，单位为分
	CltCpnMinAmt                int64    `json:"clt_cpn_min_amt"`                 //店铺收藏券使用门槛价格，单位为分
	DescTxt                     string   `json:"desc_txt"`                        //描述分
	ServTxt                     string   `json:"serv_txt"`                        //服务分
	LgstTxt                     string   `json:"lgst_txt"`                        //物流分
	PlanType                    int64    `json:"plan_type"`                       //推广计划类型
	ZsDuoId                     int64    `json:"zs_duo_id"`                       //招商团长id
}
