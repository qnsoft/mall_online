package taobao_sdk

import (
	"encoding/json"
	"fmt"
	"github.com/qnsoft/mall_online/lib/taobao/models"
	"strings"
	"github.com/qnsoft/web_api/utils/ErrorHelper"
	"github.com/qnsoft/web_api/utils/WebHelper"
	php2go "github.com/qnsoft/web_api/utils/Php2go"
)

/*
测试接口用来熟悉淘宝api
*/
//获取淘宝当前系统时间
func Taobao_time_get() {
	var _rt taobao.Error_Response 
	_Params := taobao.GetBaseParam()
	/*-----------以上是基础参数-------------*/
	_Params["method"] = "taobao.time.get"
	_Sign := php2go.Rtrim(taobao.SignParameters(_Params, false)) //参数排序
	_SignParameters := php2go.Md5(_Sign)
	_Params["sign"] = strings.ToUpper(_SignParameters)
	_HeaderData := map[string]string{"Content-Type": "application/x-www-form-urlencoded;charset=utf-8"}
	_http_url := taobao.App_api_url
	_req := WebHelper.HttpPost(_http_url, _HeaderData, _Params)
	fmt.Println(_req)
	err := json.Unmarshal([]byte(_req), &_rt)
	ErrorHelper.CheckErr(err)
	fmt.Println(_rt.ErrorResponse.Code)
}

// //获取一个产品的信息
// func Taobao_product_get() {
// 	var _rt taobao.Error_Response
// 	_Params := taobao.GetBaseParam()
// 	/*-----------以上是基础参数-------------*/
// 	_Params["method"] = "taobao.product.get"
// 	_Params["fields"] = "product_id,outer_id"
// 	_Params["product_id"] = "86126527"
// 	_Params["cid"] = "50012286"
// 	_Params["props"] = "10005:10027;10006:29729"
// 	_Sign := php2go.Rtrim(taobao.SignParameters(_Params, false)) //参数排序
// 	_SignParameters := php2go.Md5(_Sign)
// 	_Params["sign"] = strings.ToUpper(_SignParameters)
// 	_HeaderData := map[string]string{"Content-Type": "application/x-www-form-urlencoded;charset=utf-8"}
// 	_http_url := taobao.App_api_url
// 	_req := WebHelper.HttpPost(_http_url, _HeaderData, _Params)
// 	fmt.Println(_req)
// 	err := json.Unmarshal([]byte(_req), &_rt)
// 	ErrorHelper.CheckErr(err)
// 	fmt.Println(_rt.ErrorResponse.Code)
// }
