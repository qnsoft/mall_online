package taobao_model

import (
	"fmt"
	"sort"
	"strconv"
	"time"
	"github.com/qnsoft/utils/DateHelper"
	php2go "github.com/qnsoft/utils/Php2go"
)

const (
	App_key    string = "27675801"                         //27675801 27730028
	App_secret string = "1141167f91a1861c73a2cb2df014d9eb" //1141167f91a1861c73a2cb2df014d9eb  1ba3669aadc430adfa51aaca10fbf1c8
	//App_sessionkey string = "test"
	App_api_url string = "http://gw.api.taobao.com/router/rest"
	//App_api_url string = "https://eco.taobao.com/router/rest"
)

/*
获取基础参数
*/
func GetBaseParam() map[string]interface{} {
	_rt := make(map[string]interface{})
	_rt["timestamp"] = date.FormatDate(time.Now(), "yyyy-MM-dd HH:mm:ss")
	_rt["format"] = "json"
	_rt["v"] = "2.0"
	_rt["sign_method"] = "md5"
	_rt["app_key"] = App_key
	return _rt
}

/*
参数集排序Sign签名 根据key排序
mp参数列表
is_url是否URLEncode加密
*/
func SignParameters(mp map[string]interface{}, is_url bool) string {
	_str := ""
	var newMp = make([]string, 0)
	for k, _ := range mp {
		newMp = append(newMp, k)
	}
	sort.Strings(newMp)
	for _, v := range newMp {
		_srt := ""
		switch _value := mp[v].(type) {
		case string:
			_srt = _value
		case int:
			_srt = strconv.Itoa(_value)
		}
		if len(_srt) > 0 {
			if is_url {
				_str += v + php2go.URLEncode(_srt)
			} else {
				_str += v + _srt
			}
		}
	}
	_str = fmt.Sprintf("%s%s%s", App_secret, _str, App_secret)
	return _str
}

/*
参数集排序Sign签名 根据key排序
*/
func SignParametersMap(mp map[string]string, is_url bool) map[string]string {
	_map := make(map[string]string)
	var newMp = make([]string, 0)
	for k, _ := range mp {
		newMp = append(newMp, k)
	}
	sort.Strings(newMp)
	for _, v := range newMp {
		if is_url {
			_map[v] = php2go.URLEncode(mp[v])
		} else {
			_map[v] = mp[v]
		}
	}
	return _map
}
