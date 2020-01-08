package jd_model

/**
 * 京东实体对象
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
