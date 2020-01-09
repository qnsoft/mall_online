package taobao

import (
	"fmt"
	"github.com/qnsoft/web_api/controllers/Token"
	"github.com/qnsoft/utils/ErrorHelper"
	"github.com/qnsoft/utils/WebHelper"
)

const (
	//授权口令
	vekey = "V00002504Y26508322"
)

/**
*维易淘客联盟 接口 http://wsd.591hufu.com/doc/taokequdaojiekou
 */
type Wytklm_Controller struct {
	Token.BaseController
}

/*
获取淘宝客邀请url地址
*/
func (this *Wytklm_Controller) Get_auth_url() {
	_state := this.GetString("user_id")
	publishersave_url := "http://api.vephp.com/auth?vekey=" + vekey + "&state=" + _state + "&view=wap&recall_url=https://api1.xhdncppf.com/api/mallother/taobao/Callback_other?member_id=" + _state
	_req := Self_Get(publishersave_url, nil)
	fmt.Println(_req)
	this.Data["json"] = map[string]interface{}{"code": 200, "url": _req}
	this.ServeJSON()
}

/*
三方获取淘宝客邀请码
*/
// func init() {
// 	var _rt map[string]string
// 	invitecode_url := "http://api.vephp.com/invitecode?vekey=" + vekey
// 	_req := Self_Get(invitecode_url, nil)
// 	fmt.Println(_req)
// 	err := json.Unmarshal([]byte(_req), &_rt)
// 	ErrorHelper.CheckErr(err)
// 	Cache_bm.Put("inviter_code", _rt["inviter_code"], 1000*time.Second)
// 	//return _rt["inviter_code"]
// }

/**
*维易淘客联盟 万能高佣转链接口(这里选的是id转链接) http://api.vephp.com/hcapi?vekey=V00002504Y26508322&para=549517028962&relation_id=2224886364
 */
func (this *Wytklm_Controller) Get_directhc() {
	_num_iid := this.GetString("num_iid")
	_relation_id := this.GetString("relation_id")
	publishersave_url := "http://api.vephp.com/hcid?vekey=" + vekey + "&pid=mm_16026811_597150311_109130300047&para=" + _num_iid + "&relationId=" + _relation_id
	fmt.Println(publishersave_url)
	ErrorHelper.LogInfo("高佣转接口领券地址", publishersave_url)
	_req := Self_Get(publishersave_url, nil)
	fmt.Println(_req)
	this.Data["json"] = map[string]interface{}{"code": 200, "url": _req}
	this.ServeJSON()
}

/*
 post提交
 _map 提交参数
*/
func Self_Post(_http_url string, _map map[string]interface{}) string {
	_HeaderData := map[string]string{"Content-Type": "application/x-www-form-urlencoded;charset=utf-8"}
	_req := WebHelper.HttpPost(_http_url, _HeaderData, _map)
	return _req
}

/*
 get提交
 _map 提交参数
*/
func Self_Get(_http_url string, _map map[string]interface{}) string {
	_HeaderData := map[string]string{"Content-Type": "application/json"}
	_req := WebHelper.HttpGet(_http_url, _HeaderData, _map)
	return _req
}
