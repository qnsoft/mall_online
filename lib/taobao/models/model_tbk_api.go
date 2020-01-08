package taobao_model

/**
 * 淘宝客实体对象
 */
//提交返回参数
type Error_Response struct {
	ErrorResponse Error_Response_Model `json:"error_response"`
}

//提交返回参数详情实体
type Error_Response_Model struct {
	Code      int    `json:"code"`
	Msg       string `json:"msg"`
	RequestId string `json:"request_id"`
}

/* -----淘宝客实体开始------- */
/*-------------------淘宝客商品查询开始---------------------------------------------------------------------------------------------------*/
//淘宝客商品查询
type Tbk_item_get_model struct {
	Fields      string `json:"fields"`        //需返回的字段列表
	Q           string `json:"q"`             //查询词(女装)
	Cat         string `json:"cat"`           //后台类目ID，用,分割，最大10个，该ID可以通过taobao.itemcats.get接口获取到
	Itemloc     string `json:"itemloc"`       //所在地(杭州)
	Sort        string `json:"sort"`          //排序_des（降序），排序_asc（升序），销量（total_sales），淘客佣金比率（tk_rate）， 累计推广量（tk_total_sales），总支出佣金（tk_total_commi）
	IsTmall     string `json:"is_tmall"`      //是否商城商品，设置为true表示该商品是属于淘宝商城商品，设置为false或不设置表示不判断这个属性
	IsOverseas  string `json:"is_overseas"`   //是否海外商品，设置为true表示该商品是属于海外商品，设置为false或不设置表示不判断这个属性
	StartPrice  string `json:"start_price"`   //折扣价范围下限，单位：元
	EndPrice    string `json:"end_price"`     //折扣价范围上限，单位：元
	StartTkRate string `json:"start_tk_rate"` //淘客佣金比率上限，如：1234表示12.34%
	EndTkRate   string `json:"end_tk_rate"`   //淘客佣金比率下限，如：1234表示12.34%
	Platform    string `json:"platform"`      //链接形式：1：PC，2：无线，默认：１
	PageNo      string `json:"page_no"`       //第几页，默认：１
	PageSize    string `json:"page_size"`     //页大小，默认20，1~100
}

//淘宝客商品查询 返回值
type Tbk_item_get_return_model struct {
	TbkItemGetResponse Tbk_item_get_response_model `json:"tbk_item_get_response"`
}

//淘宝客商品查询 返回值实体
type Tbk_item_get_response_model struct {
	Results      Results_model1 `json:"results"`
	TotalResults int            `json:"total_results"`
	RequestId    string         `json:"request_id"`
}
type Results_model1 struct {
	NtbkItem []N_tbk_item_model `json:"n_tbk_item"`
}
type N_tbk_item_model struct {
	ItemUrl      string            `json:"item_url"`
	Nick         string            `json:"nick"`
	NumIid       string            `json:"num_iid"`
	PictUrl      string            `json:"pict_url"`
	Provcity     string            `json:"provcity"`
	ReservePrice string            `json:"reserve_price"`
	SellerId     string            `json:"seller_id"`
	SmallImages  SmallImages_model `json:"small_images"`
	Title        string            `json:"title"`
	UserType     string            `json:"user_type"`
	Volume       string            `json:"volume"`
	ZkFinalPrice string            `json:"zk_final_price"`
}
type SmallImages_model struct {
	String []string `json:"string"`
}

/*-------------------淘宝客商品查询结束-----------------------------------------------------------------------------------------------------------------*/
/*------------------- 淘宝客-公用-商品关联推荐-------------------*/

// 淘宝客-公用-商品关联推荐
type Tbk_item_recommend_get_model struct {
	Fields   string `json:"fields"`   //必传项(*)需返回的字段列表
	NumIid   string `json:"num_iid"`  //必传项(*)商品Id表
	Count    string `json:"count"`    //返回数量，默认20，最大值40
	Platform string `json:"platform"` //链接形式：1：PC，2：无线，默认：１
}

//淘宝客-公用-商品关联推荐请求返回淘宝客商品
type Tbk_item_recommend_get_return_model struct {
	//淘宝客商品
	Results []Results_model2 `json:"results"`
}

//淘宝客商品实体
type Results_model2 struct {
	NumIid       string   `json:"num_iid"`        //商品ID
	Title        string   `json:"title"`          //商品标题
	PictUrl      string   `json:"pict_url"`       //商品主图
	SmallImages  []string `json:"small_images"`   //商品小图列表
	ReservePrice string   `json:"reserve_price"`  //商品一口价格
	ZkFinalPrice string   `json:"zk_final_price"` //商品折扣价格
	UserType     string   `json:"user_type"`      //卖家类型，0表示集市，1表示商城
	Provcity     string   `json:"provcity"`       //宝贝所在地
	ItemUrl      string   `json:"item_url"`       //卖家昵称
	Nick         string   `json:"nick"`           //商品地址
	SellerId     string   `json:"seller_id"`      //卖家id
	Volume       string   `json:"volume"`         //30天销量
}

/*------------------- 阿里妈妈推广券-------------------*/
//淘宝客-好券清单API【导购】
type Taobao_tbk_dg_item_coupon_get_model struct {
	AdzoneId int    `json:"adzone_id"` //必传项(*)mm_xxx_xxx_xxx的第三位
	Platform int    `json:"platform"`  //1：PC，2：无线，默认：1
	Cat      string `json:"cat"`       //后台类目ID，用,分割，最大10个，该ID可以通过taobao.itemcats.get接口获取到
	PageSize int    `json:"page_size"` //页大小，默认20，1~100
	Q        string `json:"q"`         //查询词
	PageNo   int    `json:"page_no"`   //第几页，默认：1（当后台类目和查询词均不指定的时候，最多出10000个结果，即page_no*page_size不能超过10000；当指定类目或关键词的时候，则最多出100个结果）
}

//淘宝客-好券清单API【导购】返回
type Taobao_tbk_dg_item_coupon_get_return_model struct {
	TbkDgItemCouponGetResponse Tbk_dg_item_coupon_get_response_model `json:"tbk_dg_item_coupon_get_response"`
}

//淘宝客-好券清单API【导购】返回实体
type Tbk_dg_item_coupon_get_response_model struct {
	Results      Results_model3 `json:"results"`       //TbkCoupon
	TotalResults int            `json:"total_results"` //总请求数
	RequestId    string         `json:"request_id"`    //请求id
}

//好券清单返回results实体
type Results_model3 struct {
	//fmt.Println(_req)
	TbkCoupon []TbkCoupon_model `json:"tbk_coupon"` //信息列表
}

//好券清单返回results实体
type TbkCoupon_model struct {
	SmallImages       SmallImages_model `json:"small_images"`        //商品小图列表
	ShopTitle         string            `json:"shop_title"`          //店铺名称
	UserType          int               `json:"user_type"`           //卖家类型，0表示集市，1表示商城
	ZkFinalPrice      string            `json:"zk_final_price"`      //折扣价
	Title             string            `json:"title"`               //商品标题
	Nick              string            `json:"nick"`                //卖家昵称
	SellerId          int               `json:"seller_id"`           //卖家id
	Volume            int               `json:"volume"`              //30天销量
	PictUrl           string            `json:"pict_url"`            //商品主图
	ItemUrl           string            `json:"item_url"`            //商品详情页链接地址
	CouponTotalCount  int               `json:"coupon_total_count"`  //优惠券总量
	CommissionRate    string            `json:"commission_rate"`     //佣金比率(%)
	CouponInfo        string            `json:"coupon_info"`         //优惠券面额
	Category          int               `json:"category"`            //后台一级类目
	NumIid            int               `json:"num_iid"`             //itemId
	CouponRemainCount int               `json:"coupon_remain_count"` //优惠券剩余量
	CouponStartTime   string            `json:"coupon_start_time"`   //优惠券开始时间
	CouponEndTime     string            `json:"coupon_end_time"`     //优惠券结束时间
	CouponClickUrl    string            `json:"coupon_click_url"`    //商品优惠券推广链接
	ItemDescription   string            `json:"item_description"`    //宝贝描述（推荐理由）
}

///////////////--------------------------淘宝客-推广者-物料精选2019-09-20之后采用接口------------------------------------//////////////////////
/*
淘宝客-推广者-物料精选(2019-09-20之后全新采用)
*/
type Taobao_tbk_dg_optimus_materia_get_model struct {
	PageSize      int    `json:"page_size"`      //页大小，默认20，1~100
	AdzoneId      int    `json:"adzone_id"`      //必传项(*)mm_xxx_xxx_xxx的第三位
	PageNo        int    `json:"page_no"`        //第几页，默认：1（当后台类目和查询词均不指定的时候，最多出10000个结果，即page_no*page_size不能超过10000；当指定类目或关键词的时候，则最多出100个结果）
	MaterialId    int    `json:"material_id"`    //官方的物料Id(详细物料id见：https://tbk.bbs.taobao.com/detail.html?appId=45301&postId=8576096)
	DeviceValue   string `json:"device_value"`   //智能匹配-设备号加密后的值（MD5加密需32位小写）
	DeviceEncrypt string `json:"device_encrypt"` //智能匹配-设备号加密类型：MD5
	DeviceType    string `json:"device_type"`    //智能匹配-设备号类型：IMEI，或者IDFA，或者UTDID（UTDID不支持MD5加密）
	ContentId     int    `json:"content_id"`     //内容专用-内容详情ID
	ContentSource string `json:"content_source"` //内容专用-内容渠道信息
	ItemId        int    `json:"item_id"`        //商品ID，用于相似商品推荐
}

//淘宝客-推广者-物料精选返回对象
type Taobao_tbk_dg_optimus_materia_get_return_model struct {
	TbkDgOptimusMaterialResponse TbkDgOptimusMaterialResponse_Model `json:"tbk_dg_optimus_material_response"`
}

//返回对象
type TbkDgOptimusMaterialResponse_Model struct {
	IsDefault  string           `json:"is_default"`
	ResultList ResultList_Model `json:"result_list"`
	RequestID  string           `json:"request_id"`
}

/*
商品对象数组
*/
type ResultList_Model struct {
	MapData []MapData_Model `json:"map_data"`
}

/*
商品对象
*/
type MapData_Model struct {
	CategoryID           int               `json:"category_id"`
	ClickURL             string            `json:"click_url"`
	CommissionRate       string            `json:"commission_rate"`
	CouponAmount         int               `json:"coupon_amount"`
	CouponClickURL       string            `json:"coupon_click_url"`
	CouponEndTime        string            `json:"coupon_end_time"`
	CouponRemainCount    int               `json:"coupon_remain_count"`
	CouponShareURL       string            `json:"coupon_share_url"`
	CouponStartFee       string            `json:"coupon_start_fee"`
	CouponStartTime      string            `json:"coupon_start_time"`
	CouponTotalCount     int               `json:"coupon_total_count"`
	ItemDescription      string            `json:"item_description"`
	ItemID               int64             `json:"item_id"`
	LevelOneCategoryID   int               `json:"level_one_category_id"`
	LevelOneCategoryName string            `json:"level_one_category_name"`
	Nick                 string            `json:"nick"`
	PictURL              string            `json:"pict_url"`
	SellerID             interface{}       `json:"seller_id"`
	ShopTitle            string            `json:"shop_title"`
	SmallImages          SmallImages_Model `json:"small_images"`
	Title                string            `json:"title"`
	UserType             int               `json:"user_type"`
	Volume               int               `json:"volume"`
	WhiteImage           string            `json:"white_image"`
	ZkFinalPrice         string            `json:"zk_final_price"`
}

//图片列表
type SmallImages_Model struct {
	String []string `json:"string"`
}

///////////////-------------------------------------淘宝客商品搜索接口------------------------------------//////////////////////////////
/*
搜索实体
*/
type Taobao_tbk_search_get_model struct {
	PageSize int    `json:"page_size"` //页大小，默认20，1~100
	PageNo   int    `json:"page_no"`   //第几页，默认：1（当后台类目和查询词均不指定的时候，最多出10000个结果，即page_no*page_size不能超过10000；当指定类目或关键词的时候，则最多出100个结果）
	Vekey    string `json:"vekey"`     //维易淘客接口获取
	Para     string `json:"para"`      //搜索关键字
	Coupon   int    `json:"coupon"`    //当不传递coupon参数或coupon=0时，默认搜索包含无券产品，当传coupon=1时则搜索有券产品。
	IsTmall  int    `json:"is_tmall"`  //是否为天猫商品？设为1表示商品是天猫商城商品，不设置或0表示不限制
	Cat      int    `json:"cat"`       //淘宝分类类目ID，用,分割，最大10个，该ID可以通过taobao.itemcats.get接口获取到。当传递了cat参数时，可以不必使用para参数 例如http://apis.vephp.com/super?vekey=xxxxx&cat=50011129
	Sort     int    `json:"sort"`      //排序，默认 total_sales_des（销量降序），可选值如下：tk_rate_des（淘客佣金比率降序）, tk_rate_asc（淘客佣金比率升序）,  total_sales_des（销量降序）, total_sales_asc（销量升序）,  tk_total_sales_des（累计推广量降序）, tk_total_sales_asc（累计推广量升序）,   tk_total_commi_des（总支出佣金降序）, tk_total_commi_asc（总支出佣金升序）, price_des（价格降序）,price_asc（价格升序）;
}

/*
搜索返回实体
*/
type Taobao_tbk_search_get_return_model struct {
	Error        string             `json:"error"`
	Msg          string             `json:"msg"`
	SearchType   string             `json:"search_type"`
	IsSimilar    string             `json:"is_similar"`
	TotalResults int                `json:"total_results"`
	ResultList   []SearchResultList `json:"result_list"`
}

/*
搜索列表对象
*/
type SearchResultList struct {
	CategoryID           int         `json:"category_id"`
	CategoryName         string      `json:"category_name"`
	CommissionRate       string      `json:"commission_rate"`
	CouponAmount         string      `json:"coupon_amount"`
	CouponEndTime        string      `json:"coupon_end_time"`
	CouponID             string      `json:"coupon_id"`
	CouponInfo           string      `json:"coupon_info"`
	CouponRemainCount    int         `json:"coupon_remain_count"`
	CouponStartFee       string      `json:"coupon_start_fee"`
	CouponStartTime      string      `json:"coupon_start_time"`
	CouponTotalCount     int         `json:"coupon_total_count"`
	ItemDescription      string      `json:"item_description"`
	ItemID               int64       `json:"item_id"`
	ItemURL              string      `json:"item_url"`
	LevelOneCategoryID   int         `json:"level_one_category_id"`
	LevelOneCategoryName string      `json:"level_one_category_name"`
	Nick                 string      `json:"nick"`
	NumIid               int64       `json:"num_iid"`
	PictURL              string      `json:"pict_url"`
	Provcity             string      `json:"provcity"`
	RealPostFee          string      `json:"real_post_fee"`
	ReservePrice         string      `json:"reserve_price"`
	SellerID             interface{} `json:"seller_id"`
	ShopDsr              int         `json:"shop_dsr"`
	ShopTitle            string      `json:"shop_title"`
	ShortTitle           string      `json:"short_title"`
	SmallImages          []string    `json:"small_images"`
	Title                string      `json:"title"`
	TkTotalCommi         string      `json:"tk_total_commi"`
	TkTotalSales         string      `json:"tk_total_sales"`
	UserType             int         `json:"user_type"`
	Volume               int         `json:"volume"`
	WhiteImage           string      `json:"white_image"`
	ZkFinalPrice         string      `json:"zk_final_price"`
}
