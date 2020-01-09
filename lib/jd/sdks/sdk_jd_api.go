package jd_sdk

import (
	"encoding/json"
	"fmt"
	"github.com/qnsoft/mall_online/lib/jd/models"
	"strings"
	"github.com/qnsoft/utils/ErrorHelper"
	"github.com/qnsoft/utils/WebHelper"
	"github.com/qnsoft/utils/php2go"
)

/*
 post提交
 _map 提交参数
*/
func Self_Post(_map map[string]interface{}) string {
	_HeaderData := map[string]string{"Content-Type": "application/x-www-form-urlencoded;charset=utf-8"}
	_http_url := jd_model.App_api_url
	_req := WebHelper.HttpPost(_http_url, _HeaderData, _map)
	return _req
}

/*
 get提交
 _map 提交参数
*/
func Self_Get(_map map[string]interface{}) string {
	_HeaderData := map[string]string{"Content-Type": "application/x-www-form-urlencoded;charset=utf-8"}
	_http_url := jd_model.App_api_url
	_req := WebHelper.HttpGet(_http_url, _HeaderData, _map)
	return _req
}

/*
测试接口
*/
//根据商品的父类目id查询子类目id信息，通常用获取各级类目对应关系，以便将推广商品归类。业务参数parentId、grade都输入0可查询所有一级类目ID，之后再用其作为parentId查询其子类目。
func Jd_union_open_category_goods_get() interface{} {
	var _rt interface{}
	_Params := jd_model.GetBaseParam()
	/*-----------以上是基础参数-------------*/
	_Params["method"] = "jd.union.open.category.goods.get"
	_Params["360buy_param_json"] = "{'req':{'parentId':1342,'grade':2}}"
	/* if len(_model.Q) > 0 {
		_Params["q"] = _model.Q
	}
	if len(_model.Cat) > 0 {
		_Params["cat"] = _model.Cat
	}
	if len(_model.Itemloc) > 0 {
		_Params["itemloc"] = _model.Itemloc
	}
	if len(_model.Sort) > 0 {
		_Params["sort"] = _model.Sort
	}
	if len(_model.IsTmall) > 0 {
		_Params["is_tmall"] = _model.IsTmall
	}
	if len(_model.IsOverseas) > 0 {
		_Params["is_overseas"] = _model.IsOverseas
	}
	if len(_model.StartPrice) > 0 {
		_Params["start_price"] = _model.StartPrice
	}
	if len(_model.EndPrice) > 0 {
		_Params["end_price"] = _model.EndPrice
	}
	if len(_model.StartTkRate) > 0 {
		_Params["start_tk_rate"] = _model.StartTkRate
	}
	if len(_model.EndTkRate) > 0 {
		_Params["end_tk_rate"] = _model.EndTkRate
	}
	if len(_model.Platform) > 0 {
		_Params["platform"] = _model.Platform
	}
	if len(_model.PageNo) > 0 {
		_Params["page_no"] = _model.PageNo
	}
	if len(_model.PageSize) > 0 {
		_Params["page_size"] = _model.PageSize
	} */
	_Sign := php2go.Rtrim(jd_model.SignParameters(_Params, false)) //参数排序
	_SignParameters := php2go.Md5(_Sign)
	_Params["sign"] = strings.ToUpper(_SignParameters)
	_req := Self_Post(_Params)
	fmt.Println(_req)
	err := json.Unmarshal([]byte(_req), &_rt)
	ErrorHelper.CheckErr(err)
	return _rt
}
