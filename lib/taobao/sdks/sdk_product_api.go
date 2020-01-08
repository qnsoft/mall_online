package taobao_sdk

import (
	"github.com/qnsoft/mall_online/lib/taobao/models"
	"strings"
	php2go "zhenfangbian/web_api/utils/Php2go"
)

/**
 * 淘宝商品对象
 */

//淘宝-获取一个产品的信息
func Taobao_product_get() string {
	/*-----------基础参数-------------*/
	_Params := taobao_model.GetBaseParam()
	/*-----------方法名-------------*/
	_Params["method"] = "taobao.product.get"
	/*-----------请求参数-------------*/
	_Params["fields"] = "product_id,outer_id"
	//_Params["product_id"] = 596473597392
	//_Params["cid"] = ""
	//_Params["props"] = ""
	/*-----------sign组合-------------*/
	_Sign := php2go.Rtrim(taobao_model.SignParameters(_Params, false)) //参数排序
	_SignParameters := php2go.Md5(_Sign)
	_Params["sign"] = strings.ToUpper(_SignParameters)
	_req := Self_Post(_Params)
	return _req
}
