package taobao

import (
	"encoding/json"
	"fmt"
	"qnsoft/mall_online/lib/taobao/models"
	"qnsoft/mall_online/lib/taobao/sdks"
	"strconv"
	"zhenfangbian/web_api/controllers/Token"
	"zhenfangbian/web_api/models/shop"
	"zhenfangbian/web_api/utils/DbHelper"
	"zhenfangbian/web_api/utils/ErrorHelper"
	"zhenfangbian/web_api/utils/WebHelper"

	"github.com/astaxie/beego/cache"
)

var Cache_bm cache.Cache

func init() {
	Cache_bm, _ = cache.NewCache("memory", `{"interval":600}`)
}

/**
*淘宝信息
 */
type Mall_other_taobao_Controller struct {
	Token.BaseController
}

/*
获取SessionKey
*/
func Auth() {
	_http_url := "http://container.open.taobao.com/container?appkey=" + taobao_model.App_key
	_HeaderData := map[string]string{"Content-Type": "application/json"}
	_req := WebHelper.HttpGet(_http_url, _HeaderData, nil)
	fmt.Println(_req)
	//return _req
}

/*
接收官方回传获取SessionKey
*/
func (this *Mall_other_taobao_Controller) Callback() {
	_session_key := this.GetString("top_session")
	_user_id := this.GetString("user_id")
	// _session := this.GetString("session")

	fmt.Println(_user_id)
	// fmt.Println(_session)
	// Cache_bm.Put("taobao_session", _session, 100*time.Second)
	this.Data["json"] = map[string]interface{}{"code": 200, "data": "", "session_key": _session_key}
	this.ServeJSON()
}

/*
回传三方接口回传
*/
func (this *Mall_other_taobao_Controller) Callback_other() {
	_session := this.GetString("session")
	_member_id := this.GetString("member_id")
	fmt.Println(_session)
	fmt.Println(_member_id)
	ErrorHelper.LogInfo("返回的session为："+_session, "member_id为："+_member_id)
	_json := Get_publishersave(_session)
	fmt.Println(_json)
	if _json.Error == "0" {
		_model_user := shop.User{UserId: _member_id}
		_ok, _ := DbHelper.MySqlDb().Get(&_model_user)
		ErrorHelper.LogInfo("根据user_id查询对应id为：", _model_user.Id)
		if _ok {
			ErrorHelper.LogInfo("返回的relationId为：", _json.Data.RelationID)
			_model_user.RelationId = strconv.Itoa(_json.Data.RelationID)
			_, err := DbHelper.MySqlDb().Id(_model_user.Id).Cols("relation_id").Update(_model_user)
			ErrorHelper.CheckErr(err)
		}
	}
	this.Data["taobao_json"] = _json
	this.TplName = "taobaoke/ok.html"
}

/*
回传后要执行的绑定事件
*/
func Get_publishersave(_accesstoken string) *Relation {
	var _rt Relation
	_invitecode := "MQTWDD"
	publishersave_url := "http://api.vephp.com/publishersave?vekey=" + vekey + "&inviter_code=" + _invitecode + "&accesstoken=" + _accesstoken
	_req := Self_Get(publishersave_url, nil)
	fmt.Println(_req)
	err := json.Unmarshal([]byte(_req), &_rt)
	ErrorHelper.CheckErr(err)
	return &_rt
}

/*
返回淘宝客信息实体
*/
type Relation struct {
	Error string            `json:"error"`
	Msg   string            `json:"msg"`
	Data  Relation_SubModel `json:"data"`
}

/*
返回淘宝客信息实体详情
*/
type Relation_SubModel struct {
	RelationID  int    `json:"relation_id"`
	AccountName string `json:"account_name"`
	SpecialID   int    `json:"special_id"`
	Desc        string `json:"desc"`
}

/* 获取商品列表 */
func (this *Mall_other_taobao_Controller) Goods() {
	//--------------淘宝客-获取商品查询---------------------------------------------------------------------//
	_model := taobao_model.Tbk_item_get_model{}
	_model.Fields = "num_iid,title,pict_url,small_images,reserve_price,zk_final_price,user_type,provcity,item_url,seller_id,volume,nick"
	_model.Q = "女装"

	_model.Cat = "16,18"
	//_model.Itemloc = "郑州"
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
	fmt.Printf("%#v", str_json1.TbkItemGetResponse.Results.NtbkItem)
	this.Data["json"] = map[string]interface{}{"code": 20000, "data": str_json1}
	this.ServeJSON()
}

/* 导购列表 */
func (this *Mall_other_taobao_Controller) Taobao_YouhuiQuan() {
	_AdzoneId, _ := this.GetInt("AdzoneId", 109130300047) //必传项(*)mm_xxx_xxx_xxx的第三位 109130300047   109204600214
	_Platform, _ := this.GetInt("Platform", 1)            //1：PC，2：无线，默认：1
	_Cat := this.GetString("Cat")                         //后台类目ID，用,分割，最大10个，该ID可以通过taobao.itemcats.get接口获取到
	_PageNo, _ := this.GetInt("PageNo", 1)                //第几页，默认：1（当后台类目和查询词均不指定的时候，最多出10000个结果，即page_no*page_size不能超过10000；当指定类目或关键词的时候，则最多出100个结果）
	_PageSize, _ := this.GetInt("PageSize", 30)           //页大小，默认20，1~100
	_Q := this.GetString("Q")                             //查询词
	/*------------------- 阿里妈妈推广券-------------------*/
	_model := taobao_model.Taobao_tbk_dg_item_coupon_get_model{}
	_model.AdzoneId = _AdzoneId
	_model.Platform = _Platform //满减吗？
	if len(_Cat) > 0 {
		_model.Cat = _Cat
	}
	_model.PageSize = _PageSize
	if len(_Q) > 0 {
		_model.Q = _Q
	}
	_model.PageNo = _PageNo
	return_json := taobao_sdk.Taobao_tbk_dg_item_coupon_get(_model)
	fmt.Printf("%#v", return_json)
	this.Data["json"] = map[string]interface{}{"code": 20000, "data": return_json}
	this.ServeJSON()
}

/* 导购列表新2019-09-20之后*/
func (this *Mall_other_taobao_Controller) Taobao_YouhuiQuan1() {
	_AdzoneId, _ := this.GetInt("AdzoneId", 109130300047) //必传项(*)mm_xxx_xxx_xxx的第三位 109130300047   109204600214
	_PageNo, _ := this.GetInt("PageNo", 1)                //第几页，默认：1（当后台类目和查询词均不指定的时候，最多出10000个结果，即page_no*page_size不能超过10000；当指定类目或关键词的时候，则最多出100个结果）
	_PageSize, _ := this.GetInt("PageSize", 30)           //页大小，默认20，1~100
	_MaterialId, _ := this.GetInt("MaterialId", 13366)    //官方的物料Id
	_model := taobao_model.Taobao_tbk_dg_optimus_materia_get_model{}
	_model.AdzoneId = _AdzoneId
	_model.PageSize = _PageSize
	_model.PageNo = _PageNo
	_model.MaterialId = _MaterialId
	return_json := taobao_sdk.Taobao_tbk_dg_optimus_material_get(_model)
	this.Data["json"] = map[string]interface{}{"code": 20000, "data": return_json}
	this.ServeJSON()
}

/* 导购列表新2019-09-20之后搜索接口--直接用三方维易淘客*/
func (this *Mall_other_taobao_Controller) Taobao_Search() {
	_PageNo, _ := this.GetInt("PageNo", 1)             //第几页，默认：1（当后台类目和查询词均不指定的时候，最多出10000个结果，即page_no*page_size不能超过10000；当指定类目或关键词的时候，则最多出100个结果）
	_PageSize, _ := this.GetInt("PageSize", 30)        //页大小，默认20，1~100
	_Para := this.GetString("Para", "特价")              //商品关键字
	_Cat := this.GetString("Cat")                      //淘宝分类类目ID
	_IsTmall, _ := this.GetInt("IsTmall", 0)           //是否天猫商城商品
	_Coupon, _ := this.GetInt("Coupon", 1)             //是否有优惠券，0没有1有
	_Sort := this.GetString("Sort", "total_sales_des") //排序
	_http_url := "http://api.vephp.com/super?vekey=V00002504Y26508322&para=" + _Para + "&page=" + strconv.Itoa(_PageNo) + "&pagesize=" + strconv.Itoa(_PageSize) + "&cat=" + _Cat + "&is_tmall=" + strconv.Itoa(_IsTmall) + "&sort=" + _Sort + "&coupon=" + strconv.Itoa(_Coupon)
	_req := Self_Get(_http_url, nil)
	//fmt.Println(_req)
	return_json := taobao_model.Taobao_tbk_search_get_return_model{}
	json.Unmarshal([]byte(_req), &return_json)
	this.Data["json"] = map[string]interface{}{"code": 20000, "data": return_json}
	this.ServeJSON()
}
