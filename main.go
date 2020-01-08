package main

import (

	// "qnsoft/mall_online/lib/taobao/models"
	"fmt"
	//"qnsoft/mall_online/lib/taobao/sdks"

	"qnsoft/mall_online/lib/jd/sdks"
	_ "qnsoft/mall_online/routers"

	"github.com/astaxie/beego"
)

func main() {
	//fmt.Println(123, taobao.App_key)
	//获取淘宝系统时间
	//taobao_sdk.Taobao_time_get()
	//taobao.Taobao_product_get()
	//str_json := taobao_sdk.Taobao_tbk_item_get()
	// if strings.Index(str_json, "tbk_item_get_response") >= 0 {
	// 	var _rt_model taobao_model.Tbk_item_get_response
	// 	err := json.Unmarshal([]byte(str_json), &_rt_model)
	// 	ErrorHelper.CheckErr(err)
	// 	fmt.Println(_rt_model)
	// }
	//淘宝客-公用-商品关联推荐
	//str_json := taobao_sdk.Taobao_tbk_item_recommend_get()
	//淘宝客-公用-淘宝客商品详情查询(简版)
	//str_json := taobao_sdk.Taobao_tbk_item_info_get
	//淘宝客-推广者-店铺搜索
	//str_json := taobao_sdk.Taobao_tbk_shop_get()
	//淘宝客-公用-店铺关联推荐
	//str_json := taobao_sdk.Taobao_tbk_shop_recommend_get()
	//淘宝客-推广者-选品库宝贝信息
	//str_json := taobao_sdk.Taobao_tbk_uatm_favorites_item_get()
	//淘宝客-推广者-选品库宝贝列表
	//str_json := taobao_sdk.Taobao_tbk_uatm_favorites_get()
	//淘抢购api
	//str_json := taobao_sdk.Taobao_tbk_ju_tqg_get()
	//淘宝客-公用-链接解析出商品id
	//str_json := taobao_sdk.Taobao_tbk_item_click_extract()
	//淘宝客-商品猜你喜欢
	//str_json := taobao_sdk.Taobao_tbk_item_guess_like()
	//淘宝客-好券清单API【导购】
	//str_json := taobao_sdk.Taobao_tbk_dg_item_coupon_get()
	// 淘宝客-公用-阿里妈妈推广券详情查询
	//str_json := taobao_sdk.Taobao_tbk_coupon_get()
	// 淘宝客-淘宝客-公用-淘口令生成
	//str_json := taobao_sdk.Taobao_tbk_tpwd_create()
	//淘宝客-推广者-图文内容输出
	//str_json := taobao_sdk.Taobao_tbk_content_get()
	// 淘宝客-推广者-新用户订单明细查询
	//str_json := taobao_sdk.Taobao_tbk_dg_newuser_order_get()
	// 淘宝客-服务商-新用户订单明细查询
	//str_json := taobao_sdk.Taobao_tbk_sc_newuser_order_get()
	// 淘宝客-推广者-物料精选
	//str_json := taobao_sdk.Taobao_tbk_dg_optimus_material()
	//淘宝客-推广者-物料搜索
	//str_json := taobao_sdk.Taobao_tbk_dg_material_optional()
	//淘宝客-推广者-拉新活动对应数据查询
	//str_json := taobao_sdk.Taobao_tbk_dg_newuser_order_sum()
	//淘宝客-服务商-拉新活动对应数据查询
	//str_json := taobao_sdk.Taobao_tbk_sc_newuser_order_sum()
	//淘宝客-推广者-图文内容效果数据
	//str_json := taobao_sdk.Taobao_tbk_content_effect_get()
	// 淘宝客-服务商-物料精选
	//str_json := taobao_sdk.Taobao_tbk_sc_optimus_material()
	// 淘宝客-淘礼金创建
	//str_json := taobao_sdk.Taobao_tbk_dg_vegas_tlj_create()
	//淘宝客-推广者-官方活动转链
	//str_json := taobao_sdk.Taobao_tbk_activitylink_get()
	//淘宝客-服务商-官方活动转链
	//str_json := taobao_sdk.Taobao_tbk_sc_activitylink_toolget()
	//淘宝客-推广者-处罚订单查询
	//str_json := taobao_sdk.Taobao_tbk_dg_punish_order_get()
	//淘宝客-推广者-淘礼金发放及使用报表
	//str_json := taobao_sdk.Taobao_tbk_dg_vegas_tlj_instance_report()
	//fmt.Println(str_json)
	/* ---------------以下是正式对象接口------------------- */
	//--------------淘宝客-获取商品查询---------------------------------------------------------------------//
	/* 	_model := taobao_model.Tbk_item_get_model{}
	   	_model.Fields = "num_iid,title,pict_url,small_images,reserve_price,zk_final_price,user_type,provcity,item_url,seller_id,volume,nick"
	   	_model.Q = "女装"

	   	_model.Cat = "16,18"
	   	_model.Itemloc = "郑州"
	   	_model.Sort = "tk_rate_des"
	   	_model.IsTmall = "true"
	   	// _model.IsOverseas = "false"
	   	// _model.StartPrice = "10"
	   	// _model.EndPrice = "10"
	   	// _model.StartTkRate = "123"
	   	// _model.EndTkRate = "123"
	   	_model.Platform = "1"
	   	_model.PageNo = "1"
	   	_model.PageSize = "20"
	   	str_json1 := taobao_sdk.Taobao_tbk_item_get(_model)
	   	fmt.Printf("%#v", str_json1.TbkItemGetResponse.Results.NtbkItem) */
	//--------------淘宝客-公用-商品关联推荐---------------------------------------------------------------------//
	// _model := taobao_model.Tbk_item_recommend_get_model{}
	// str_json2 := taobao_sdk.Taobao_tbk_item_recommend_get(_model)
	// fmt.Printf("%#v", str_json2)

	// str_json2 := taobao_sdk.Taobao_product_get()
	// fmt.Println(str_json2)
	str_json2 := jd_sdk.Jd_union_open_category_goods_get()
	fmt.Println(str_json2)

	//taobao.Auth() //获取session
	beego.Run()
}
