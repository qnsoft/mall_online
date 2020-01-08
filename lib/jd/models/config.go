package jd_model

import (
	"fmt"
	"sort"
	"strconv"
	"time"
	"github.com/qnsoft/web_api/utils/DateHelper"
	"github.com/qnsoft/web_api/utils/php2go"
)

const (
	App_key    string = "2f6faf492a5597e1415f47c7f107fedd"
	App_secret string = "3305d2a84356460a981dacab01be916e"
	//App_sessionkey string = "test"
	App_api_url string = "https://api.jd.com/routerjson"
	//App_api_url string = "https://eco.taobao.com/router/rest"
)

/*
获取基础参数
*/
func GetBaseParam() map[string]interface{} {
	_rt := make(map[string]interface{})
	_rt["timestamp"] = date.FormatDate(time.Now(), "yyyy-MM-dd HH:mm:ss")
	_rt["format"] = "json"
	_rt["v"] = "1.0"
	_rt["access_token"] = "36e8de7ac5ea43748a9559d5c30ff23agrln"
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
	_str = php2go.Rtrim(_str, "&")
	_str = fmt.Sprintf("%s%s%s", App_secret, _str, App_secret)
	fmt.Println(_str)
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
