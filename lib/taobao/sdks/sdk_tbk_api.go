package taobao_sdk

import (
	"encoding/json"
	"fmt"
	"github.com/qnsoft/mall_online/lib/taobao/models"
	"strings"
	"github.com/qnsoft/web_api/utils/ErrorHelper"
	php2go "github.com/qnsoft/web_api/utils/Php2go"
	"github.com/qnsoft/web_api/utils/WebHelper"
)

/*
 post提交
 _map 提交参数
*/
func Self_Post(_map map[string]interface{}) string {
	_HeaderData := map[string]string{"Content-Type": "application/x-www-form-urlencoded;charset=utf-8"}
	_http_url := taobao_model.App_api_url
	_req := WebHelper.HttpPost(_http_url, _HeaderData, _map)
	return _req
}

/*
淘宝客基础接口
*/
//获取淘宝客商品查询
func Taobao_tbk_item_get(_model taobao_model.Tbk_item_get_model) taobao_model.Tbk_item_get_return_model {
	var _rt taobao_model.Tbk_item_get_return_model
	_Params := taobao_model.GetBaseParam()
	/*-----------以上是基础参数-------------*/
	_Params["method"] = "taobao.tbk.item.get"
	_Params["fields"] = _model.Fields
	if len(_model.Q) > 0 {
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
	}
	_Sign := php2go.Rtrim(taobao_model.SignParameters(_Params, false)) //参数排序
	_SignParameters := php2go.Md5(_Sign)
	_Params["sign"] = strings.ToUpper(_SignParameters)
	_req := Self_Post(_Params)
	err := json.Unmarshal([]byte(_req), &_rt)
	ErrorHelper.CheckErr(err)
	return _rt
}

//淘宝客-公用-商品关联推荐
func Taobao_tbk_item_recommend_get(_model taobao_model.Tbk_item_recommend_get_model) taobao_model.Tbk_item_recommend_get_return_model {
	var _rt taobao_model.Tbk_item_recommend_get_return_model
	_Params := taobao_model.GetBaseParam()
	/*-----------以上是基础参数-------------*/
	_Params["method"] = "taobao.tbk.item.recommend.get"
	_Params["fields"] = "num_iid,title,pict_url,small_images,reserve_price,zk_final_price,user_type,provcity,item_url"
	_Params["num_iid"] = "527589948874"
	_Params["count"] = "20"
	_Params["platform"] = "1"
	_Sign := php2go.Rtrim(taobao_model.SignParameters(_Params, false)) //参数排序
	_SignParameters := php2go.Md5(_Sign)
	_Params["sign"] = strings.ToUpper(_SignParameters)
	_req := Self_Post(_Params)
	fmt.Println(_req)
	err := json.Unmarshal([]byte(_req), &_rt)
	ErrorHelper.CheckErr(err)
	return _rt
}

//淘宝客-公用-淘宝客商品详情查询(简版)
func Taobao_tbk_item_info_get() string {
	_Params := taobao_model.GetBaseParam()
	/*-----------以上是基础参数-------------*/
	_Params["method"] = "taobao.tbk.item.info.get"
	_Params["num_iids"] = "123,456"
	// _Params["platform"] = "1"
	// _Params["ip"] = "11.22.33.43"
	_Sign := php2go.Rtrim(taobao_model.SignParameters(_Params, false)) //参数排序
	_SignParameters := php2go.Md5(_Sign)
	_Params["sign"] = strings.ToUpper(_SignParameters)
	_req := Self_Post(_Params)
	return _req
}

//淘宝客-推广者-店铺搜索
func Taobao_tbk_shop_get() string {
	/*-----------基础参数-------------*/
	_Params := taobao_model.GetBaseParam()
	/*-----------方法名-------------*/
	_Params["method"] = "taobao.tbk.shop.get"
	/*-----------请求参数-------------*/
	_Params["fields"] = "user_id,shop_title,shop_type,seller_nick,pict_url,shop_url"
	_Params["q"] = "女装"
	_Params["sort"] = "commission_rate_des"
	// _Params["is_tmall"] = "false"
	// _Params["start_credit"] = "1"
	// _Params["end_credit"] = "20"
	// _Params["start_commission_rate"] = "2000"
	// _Params["end_commission_rate"] = "123"
	// _Params["start_total_action"] = "1"
	// _Params["end_total_action"] = "100"
	// _Params["start_auction_count"] = "123"
	// _Params["end_auction_count"] = "200"
	_Params["platform"] = "1"
	_Params["page_no"] = "1"
	_Params["page_size"] = "20"
	/*-----------sign组合-------------*/
	_Sign := php2go.Rtrim(taobao_model.SignParameters(_Params, false)) //参数排序
	_SignParameters := php2go.Md5(_Sign)
	_Params["sign"] = strings.ToUpper(_SignParameters)
	_req := Self_Post(_Params)
	return _req
}

//淘宝客-公用-店铺关联推荐
func Taobao_tbk_shop_recommend_get() string {
	/*-----------基础参数-------------*/
	_Params := taobao_model.GetBaseParam()
	/*-----------方法名-------------*/
	_Params["method"] = "taobao.tbk.shop.recommend.get"
	/*-----------请求参数-------------*/
	_Params["fields"] = "user_id,shop_title,shop_type,seller_nick,pict_url,shop_url" //需返回的字段列表
	_Params["user_id"] = "3929512070"                                                //卖家Id
	_Params["count"] = "20"                                                          //返回数量，默认20，最大值40
	_Params["platform"] = "1"                                                        //链接形式：1：PC，2：无线，默认：１
	/*-----------sign组合-------------*/
	_Sign := php2go.Rtrim(taobao_model.SignParameters(_Params, false)) //参数排序
	_SignParameters := php2go.Md5(_Sign)
	_Params["sign"] = strings.ToUpper(_SignParameters)
	_req := Self_Post(_Params)
	return _req
}

//淘宝客-推广者-选品库宝贝信息
func Taobao_tbk_uatm_favorites_item_get() string {
	/*-----------基础参数-------------*/
	_Params := taobao_model.GetBaseParam()
	/*-----------方法名-------------*/
	_Params["method"] = "taobao.tbk.uatm.favorites.item.get"
	/*-----------请求参数-------------*/
	_Params["platform"] = "1"                                                                                                                                                                                                  //链接形式：1：PC，2：无线，默认：１
	_Params["page_size"] = "20"                                                                                                                                                                                                //页大小，默认20，1~100
	_Params["adzone_id"] = "34567"                                                                                                                                                                                             //必传项(*)推广位id，需要在淘宝联盟后台创建；且属于appkey备案的媒体id（siteid），如何获取adzoneid，请参考，http://club.alimama.com/read-htm-tid-6333967.html?spm=0.0.0.0.msZnx5
	_Params["unid"] = "3456"                                                                                                                                                                                                   //自定义输入串，英文和数字组成，长度不能大于12个字符，区分不同的推广渠道
	_Params["favorites_id"] = "10010"                                                                                                                                                                                          //必传项(*)选品库的id
	_Params["page_no"] = "2"                                                                                                                                                                                                   //第几页，默认：1，从1开始计数
	_Params["fields"] = "num_iid,title,pict_url,small_images,reserve_price,zk_final_price,user_type,provcity,item_url,seller_id,volume,nick,shop_title,zk_final_price_wap,event_start_time,event_end_time,tk_rate,status,type" //必传项(*)需要输出则字段列表，逗号分隔
	/*-----------sign组合-------------*/
	_Sign := php2go.Rtrim(taobao_model.SignParameters(_Params, false)) //参数排序
	_SignParameters := php2go.Md5(_Sign)
	_Params["sign"] = strings.ToUpper(_SignParameters)
	_req := Self_Post(_Params)
	return _req
}

//淘宝客-推广者-选品库宝贝列表
func Taobao_tbk_uatm_favorites_get() string {
	/*-----------基础参数-------------*/
	_Params := taobao_model.GetBaseParam()
	/*-----------方法名-------------*/
	_Params["method"] = "taobao.tbk.uatm.favorites.get"
	/*-----------请求参数-------------*/
	_Params["page_no"] = "1"                                //第几页，从1开始计数
	_Params["page_size"] = "20"                             //默认20，页大小，即每一页的活动个数
	_Params["fields"] = "favorites_title,favorites_id,type" //必传项(*)需要返回的字段列表，不能为空，字段名之间使用逗号分隔
	_Params["type"] = "1"                                   //默认值-1；选品库类型，1：普通选品组，2：高佣选品组，-1，同时输出所有类型的选品组
	/*-----------sign组合-------------*/
	_Sign := php2go.Rtrim(taobao_model.SignParameters(_Params, false)) //参数排序
	_SignParameters := php2go.Md5(_Sign)
	_Params["sign"] = strings.ToUpper(_SignParameters)
	_req := Self_Post(_Params)
	return _req
}

//淘抢购api
func Taobao_tbk_ju_tqg_get() string {
	/*-----------基础参数-------------*/
	_Params := taobao_model.GetBaseParam()
	/*-----------方法名-------------*/
	_Params["method"] = "taobao.tbk.ju.tqg.get"
	/*-----------请求参数-------------*/
	_Params["adzone_id"] = "123"                                                                                                       //必传项(*)推广位id（推广位申请方式：http://club.alimama.com/read.php?spm=0.0.0.0.npQdST&tid=6306396&ds=1&page=1&toread=1）
	_Params["fields"] = "click_url,pic_url,reserve_price,zk_final_price,total_amount,sold_num,title,category_name,start_time,end_time" //必传项(*)需返回的字段列表
	_Params["start_time"] = "2016-08-09 09:00:00"                                                                                      //必传项(*)最早开团时间
	_Params["end_time"] = "2016-08-09 16:00:00"                                                                                        //必传项(*)最晚开团时间
	_Params["page_no"] = "1"                                                                                                           //必传项(*)第几页，默认1，1~100
	_Params["page_size"] = "40"                                                                                                        //必传项(*)页大小，默认40，1~40
	/*-----------sign组合-------------*/
	_Sign := php2go.Rtrim(taobao_model.SignParameters(_Params, false)) //参数排序
	_SignParameters := php2go.Md5(_Sign)
	_Params["sign"] = strings.ToUpper(_SignParameters)
	_req := Self_Post(_Params)
	return _req
}

//淘宝客-公用-链接解析出商品id
func Taobao_tbk_item_click_extract() string {
	/*-----------基础参数-------------*/
	_Params := taobao_model.GetBaseParam()
	/*-----------方法名-------------*/
	_Params["method"] = "taobao.tbk.item.click.extract"
	/*-----------请求参数-------------*/
	_Params["click_url"] = "https://s.click.taobao.com/***" //必传项(*)长链接或短链接
	/*-----------sign组合-------------*/
	_Sign := php2go.Rtrim(taobao_model.SignParameters(_Params, false)) //参数排序
	_SignParameters := php2go.Md5(_Sign)
	_Params["sign"] = strings.ToUpper(_SignParameters)
	_req := Self_Post(_Params)
	return _req
}

//淘宝客-商品猜你喜欢
func Taobao_tbk_item_guess_like() string {
	/*-----------基础参数-------------*/
	_Params := taobao_model.GetBaseParam()
	/*-----------方法名-------------*/
	_Params["method"] = "taobao.tbk.item.guess.like"
	/*-----------请求参数-------------*/
	_Params["adzone_id"] = "123"                             //必传项(*)广告位ID
	_Params["user_nick"] = "abc"                             //用户昵称，from cookie : _nk_或者tracknick ; from百川sdk : nick
	_Params["user_id"] = "123456"                            //用户数字ID，from cookie : unb
	_Params["os"] = "ios"                                    //必传项(*)系统类型，ios, android, other
	_Params["idfa"] = "65A509BA-227C-49AC-91EC-DE6817E63B10" //ios广告跟踪id
	_Params["imei"] = "641221321098757"                      //android设备imei
	_Params["imei_md5"] = "115d1f360c48b490c3f02fc3e7111111" //android设备imeiMD5值，32位小写
	_Params["ip"] = "106.11.34.15"                           //必传项(*)客户端ip
	_Params["ua"] = "Mozilla/5.0"                            //必传项(*)userAgent
	_Params["apnm"] = "com.xxx"                              //应用包名
	_Params["net"] = "wifi"                                  //必传项(*)联网方式，wifi, cell, unknown
	_Params["mn"] = "iPhone7%2C2"                            //设备型号
	_Params["page_no"] = "1"                                 //第几页
	_Params["page_size"] = "20"                              //页大小
	/*-----------sign组合-------------*/
	_Sign := php2go.Rtrim(taobao_model.SignParameters(_Params, false)) //参数排序
	_SignParameters := php2go.Md5(_Sign)
	_Params["sign"] = strings.ToUpper(_SignParameters)
	_req := Self_Post(_Params)
	return _req
}

//淘宝客-好券清单API【导购】
func Taobao_tbk_dg_item_coupon_get(_model taobao_model.Taobao_tbk_dg_item_coupon_get_model) taobao_model.Taobao_tbk_dg_item_coupon_get_return_model {
	var _rt taobao_model.Taobao_tbk_dg_item_coupon_get_return_model
	/*-----------基础参数-------------*/
	_Params := taobao_model.GetBaseParam()
	/*-----------方法名-------------*/
	_Params["method"] = "taobao.tbk.dg.item.coupon.get"
	/*-----------请求参数-------------*/
	//if len(_model.AdzoneId) > 0 {
	_Params["adzone_id"] = _model.AdzoneId //必传项(*)mm_xxx_xxx_xxx的第三位
	//}
	if _model.Platform > 0 {
		_Params["platform"] = _model.Platform //1：PC，2：无线，默认：1
	}
	if len(_model.Cat) > 0 {
		_Params["cat"] = _model.Cat //后台类目ID，用,分割，最大10个，该ID可以通过taobao.itemcats.get接口获取到
	}
	if _model.PageSize > 0 {
		_Params["page_size"] = _model.PageSize //页大小，默认20，1~100
	}
	if len(_model.Q) > 0 {
		_Params["q"] = _model.Q //查询词
	}
	if _model.PageNo > 0 {
		_Params["page_no"] = _model.PageNo //第几页，默认：1（当后台类目和查询词均不指定的时候，最多出10000个结果，即page_no*page_size不能超过10000；当指定类目或关键词的时候，则最多出100个结果）
	}
	/*-----------sign组合-------------*/
	_Sign := php2go.Rtrim(taobao_model.SignParameters(_Params, false)) //参数排序
	_SignParameters := php2go.Md5(_Sign)
	_Params["sign"] = strings.ToUpper(_SignParameters)
	_req := Self_Post(_Params)
	fmt.Println(_req)
	err := json.Unmarshal([]byte(_req), &_rt)
	ErrorHelper.CheckErr(err)
	return _rt
}

// 淘宝客-公用-阿里妈妈推广券详情查询
func Taobao_tbk_coupon_get() string {
	/*-----------基础参数-------------*/
	_Params := taobao_model.GetBaseParam()
	/*-----------方法名-------------*/
	_Params["method"] = "taobao.tbk.coupon.get"
	/*-----------请求参数-------------*/
	_Params["me"] = "nfr%2BYTo2k1PX18gaNN%2BIPkIG2PadNYbBnwEsv6mRavWieOoOE3L9OdmbDSSyHbGxBAXjHpLKvZbL1320ML%2BCF5FRtW7N7yJ056Lgym4X01A%3D" //带券ID与商品ID的加密串
	_Params["item_id"] = "123"                                                                                                             //商品ID
	_Params["activity_id"] = "sdfwe3eefsdf"                                                                                                //券ID

	/*-----------sign组合-------------*/
	_Sign := php2go.Rtrim(taobao_model.SignParameters(_Params, false)) //参数排序
	_SignParameters := php2go.Md5(_Sign)
	_Params["sign"] = strings.ToUpper(_SignParameters)
	_req := Self_Post(_Params)
	return _req
}

// 淘宝客-淘宝客-公用-淘口令生成
func Taobao_tbk_tpwd_create() string {
	/*-----------基础参数-------------*/
	_Params := taobao_model.GetBaseParam()
	/*-----------方法名-------------*/
	_Params["method"] = "taobao.tbk.tpwd.create"
	/*-----------请求参数-------------*/
	_Params["user_id"] = "123"                    //生成口令的淘宝用户ID
	_Params["text"] = "长度大于5个字符"                  //必传项(*)口令弹框内容
	_Params["url"] = "https://uland.taobao.com/"  //必传项(*)口令跳转目标页
	_Params["logo"] = "https://uland.taobao.com/" //口令弹框logoURL
	_Params["ext"] = "{}"                         //扩展字段JSON格式
	/*-----------sign组合-------------*/
	_Sign := php2go.Rtrim(taobao_model.SignParameters(_Params, false)) //参数排序
	_SignParameters := php2go.Md5(_Sign)
	_Params["sign"] = strings.ToUpper(_SignParameters)
	_req := Self_Post(_Params)
	return _req
}

// 淘宝客-推广者-图文内容输出
func Taobao_tbk_content_get() string {
	/*-----------基础参数-------------*/
	_Params := taobao_model.GetBaseParam()
	/*-----------方法名-------------*/
	_Params["method"] = "taobao.tbk.content.get"
	/*-----------请求参数-------------*/
	_Params["adzone_id"] = "123"                  //必传项(*)推广位
	_Params["type"] = "1"                         //内容类型，1:图文、2: 图集、3: 短视频
	_Params["before_timestamp"] = "1491454244598" //表示取这个时间点以前的数据，以毫秒为单位（出参中的last_timestamp是指本次返回的内容中最早一条的数据，下一次作为before_timestamp传过来，即可实现翻页）
	_Params["count"] = "10"                       //表示期望获取条数
	_Params["cid"] = "2"                          //类目
	_Params["image_width"] = "300"                //图片宽度
	_Params["image_height"] = "300"               //图片高度
	_Params["content_set"] = "1"                  //默认可不传,内容库类型:1 优质内容
	/*-----------sign组合-------------*/
	_Sign := php2go.Rtrim(taobao_model.SignParameters(_Params, false)) //参数排序
	_SignParameters := php2go.Md5(_Sign)
	_Params["sign"] = strings.ToUpper(_SignParameters)
	_req := Self_Post(_Params)
	return _req
}

//淘宝客-推广者-新用户订单明细查询
func Taobao_tbk_dg_newuser_order_get() string {
	/*-----------基础参数-------------*/
	_Params := taobao_model.GetBaseParam()
	/*-----------方法名-------------*/
	_Params["method"] = "taobao.tbk.dg.newuser.order.get"
	/*-----------请求参数-------------*/
	_Params["page_size"] = "20"                   //页大小，默认20，1~100
	_Params["adzone_id"] = "123"                  //mm_xxx_xxx_xxx的第三位
	_Params["page_no"] = "1"                      //页码，默认1
	_Params["start_time"] = "2018-01-24 00:34:05" //开始时间，当活动为淘宝活动，表示最早注册时间；当活动为支付宝活动，表示最早绑定时间；当活动为天猫活动，表示最早领取红包时间
	_Params["end_time"] = "2018-01-24 00:34:05"   //结束时间，当活动为淘宝活动，表示最晚结束时间；当活动为支付宝活动，表示最晚绑定时间；当活动为天猫活动，表示最晚领取红包的时间
	_Params["activity_id"] = "119013_2"           //活动id， 活动名称与活动ID列表，请参见https://tbk.bbs.taobao.com/detail.html?appId=45301&postId=8599277
	/*-----------sign组合-------------*/
	_Sign := php2go.Rtrim(taobao_model.SignParameters(_Params, false)) //参数排序
	_SignParameters := php2go.Md5(_Sign)
	_Params["sign"] = strings.ToUpper(_SignParameters)
	_req := Self_Post(_Params)
	return _req
}

//淘宝客-服务商-新用户订单明细查询
func Taobao_tbk_sc_newuser_order_get() string {
	/*-----------基础参数-------------*/
	_Params := taobao_model.GetBaseParam()
	/*-----------方法名-------------*/
	_Params["method"] = "taobao.tbk.sc.newuser.order.get"
	/*-----------请求参数-------------*/
	_Params["page_size"] = "20"                   //页大小，默认20，1~100
	_Params["adzone_id"] = "123"                  //mm_xxx_xxx_xxx的第三位
	_Params["page_no"] = "1"                      //页码，默认1
	_Params["site_id"] = "1"                      //mm_xxx_xxx_xxx的第二位
	_Params["activity_id"] = "119013_2"           //活动id， 现有活动id包括： 2月手淘拉新：119013_2 3月手淘拉新：119013_3 4月手淘拉新：119013_4 拉手机支付宝新用户_赚赏金：200000_5
	_Params["end_time"] = "2018-01-24 00:34:05"   //结束时间，当活动为淘宝活动，表示最晚结束时间；当活动为支付宝活动，表示最晚绑定时间；当活动为天猫活动，表示最晚领取红包的时间
	_Params["start_time"] = "2018-01-24 00:34:05" //开始时间，当活动为淘宝活动，表示最早注册时间；当活动为支付宝活动，表示最早绑定时间；当活动为天猫活动，表示最早领取红包时间

	/*-----------sign组合-------------*/
	_Sign := php2go.Rtrim(taobao_model.SignParameters(_Params, false)) //参数排序
	_SignParameters := php2go.Md5(_Sign)
	_Params["sign"] = strings.ToUpper(_SignParameters)
	_req := Self_Post(_Params)
	return _req
}

//淘宝客-推广者-物料精选
func Taobao_tbk_dg_optimus_material_get(_model taobao_model.Taobao_tbk_dg_optimus_materia_get_model) taobao_model.Taobao_tbk_dg_optimus_materia_get_return_model {
	var _rt taobao_model.Taobao_tbk_dg_optimus_materia_get_return_model
	/*-----------基础参数-------------*/
	_Params := taobao_model.GetBaseParam()
	/*-----------方法名-------------*/
	_Params["method"] = "taobao.tbk.dg.optimus.material"
	/*-----------请求参数-------------*/
	_Params["adzone_id"] = _model.AdzoneId //必传项(*)mm_xxx_xxx_xxx的第三位
	if _model.PageSize > 0 {
		_Params["page_size"] = _model.PageSize //页大小，默认20，1~100
	}
	if _model.PageNo > 0 {
		_Params["page_no"] = _model.PageNo //第几页，默认：1（当后台类目和查询词均不指定的时候，最多出10000个结果，即page_no*page_size不能超过10000；当指定类目或关键词的时候，则最多出100个结果）
	}
	if _model.MaterialId > 0 {
		_Params["material_id"] = _model.MaterialId //官方的物料Id(详细物料id见：https://tbk.bbs.taobao.com/detail.html?appId=45301&postId=8576096)
	}
	if len(_model.DeviceValue) > 0 {
		_Params["device_value"] = _model.DeviceValue //智能匹配-设备号加密后的值（MD5加密需32位小写）
	}
	if len(_model.DeviceEncrypt) > 0 {
		_Params["device_encrypt"] = _model.DeviceEncrypt //智能匹配-设备号加密类型：MD5
	}
	if len(_model.DeviceType) > 0 {
		_Params["device_type"] = _model.DeviceType //智能匹配-设备号类型：IMEI，或者IDFA，或者UTDID（UTDID不支持MD5加密）
	}
	if _model.ContentId > 0 {
		_Params["content_id"] = _model.ContentId //内容专用-内容详情ID
	}
	if len(_model.ContentSource) > 0 {
		_Params["content_source"] = _model.ContentSource //内容专用-内容渠道信息
	}
	if _model.ItemId > 0 {
		_Params["item_id"] = _model.ItemId //商品ID，用于相似商品推荐
	}
	/*-----------sign组合-------------*/
	_Sign := php2go.Rtrim(taobao_model.SignParameters(_Params, false)) //参数排序
	_SignParameters := php2go.Md5(_Sign)
	_Params["sign"] = strings.ToUpper(_SignParameters)
	_req := Self_Post(_Params)
	err := json.Unmarshal([]byte(_req), &_rt)
	ErrorHelper.CheckErr(err)
	return _rt
}

//淘宝客-推广者-物料搜索
func Taobao_tbk_dg_material_optional() string {
	/*-----------基础参数-------------*/
	_Params := taobao_model.GetBaseParam()
	/*-----------方法名-------------*/
	_Params["method"] = "taobao.tbk.dg.material.optional"
	/*-----------请求参数-------------*/
	_Params["start_dsr"] = "10"             //商品筛选(特定媒体支持)-店铺dsr评分。筛选大于等于当前设置的店铺dsr评分的商品0-50000之间
	_Params["page_size"] = "20"             //页大小，默认20，1~100
	_Params["page_no"] = "1"                //第几页，默认：１
	_Params["platform"] = "1"               //链接形式：1：PC，2：无线，默认：１
	_Params["end_tk_rate"] = "1234"         //商品筛选-淘客佣金比率上限。如：1234表示12.34%
	_Params["start_tk_rate"] = "1234"       //商品筛选-淘客佣金比率下限。如：1234表示12.34%
	_Params["end_price"] = "10"             //商品筛选-折扣价范围上限。单位：元
	_Params["start_price"] = "10"           //商品筛选-折扣价范围下限。单位：元
	_Params["is_overseas"] = "false"        //商品筛选-是否海外商品。true表示属于海外商品，false或不设置表示不限
	_Params["is_tmall"] = "false"           //商品筛选-是否天猫商品。true表示属于天猫商品，false或不设置表示不限
	_Params["sort"] = "tk_rate_des"         //排序_des（降序），排序_asc（升序），销量（total_sales），淘客佣金比率（tk_rate）， 累计推广量（tk_total_sales），总支出佣金（tk_total_commi），价格（price）
	_Params["itemloc"] = "杭州"               //商品筛选-所在地
	_Params["cat"] = "16,18"                //商品筛选-后台类目ID。用,分割，最大10个，该ID可以通过taobao.itemcats.get接口获取到
	_Params["q"] = "女装"                     //商品筛选-查询词
	_Params["material_id"] = "2836"         //官方的物料Id(详细物料id见：https://tbk.bbs.taobao.com/detail.html?appId=45301&postId=8576096)，不传时默认为2836
	_Params["has_coupon"] = "false"         //优惠券筛选-是否有优惠券。true表示该商品有优惠券，false或不设置表示不限
	_Params["ip"] = "13.2.33.4"             //ip参数影响邮费获取，如果不传或者传入不准确，邮费无法精准提供
	_Params["adzone_id"] = "12345678"       //必传项(*)mm_xxx_xxx_12345678三段式的最后一段数字
	_Params["need_free_shipment"] = "true"  //商品筛选-是否包邮。true表示包邮，false或不设置表示不限
	_Params["need_prepay"] = "true"         //商品筛选-是否加入消费者保障。true表示加入，false或不设置表示不限
	_Params["include_pay_rate_30"] = "true" //商品筛选(特定媒体支持)-成交转化是否高于行业均值。True表示大于等于，false或不设置表示不限
	_Params["include_good_rate"] = "true"   //商品筛选-好评率是否高于行业均值。True表示大于等于，false或不设置表示不限
	_Params["include_rfd_rate"] = "true"    //商品筛选(特定媒体支持)-退款率是否低于行业均值。True表示大于等于，false或不设置表示不限
	_Params["npx_level"] = "2"              //商品筛选-牛皮癣程度。取值：1不限，2无，3轻微
	_Params["end_ka_tk_rate"] = "1234"      //商品筛选-KA媒体淘客佣金比率上限。如：1234表示12.34%
	_Params["start_ka_tk_rate"] = "1234"    //商品筛选-KA媒体淘客佣金比率下限。如：1234表示12.34%
	_Params["device_encrypt"] = "MD5"       //智能匹配-设备号加密类型：MD5
	_Params["device_value"] = "xxx"         //智能匹配-设备号加密后的值（MD5加密需32位小写）
	_Params["device_type"] = "IMEI"         //智能匹配-设备号类型：IMEI，或者IDFA，或者UTDID（UTDID不支持MD5加密）
	/*-----------sign组合-------------*/
	_Sign := php2go.Rtrim(taobao_model.SignParameters(_Params, false)) //参数排序
	_SignParameters := php2go.Md5(_Sign)
	_Params["sign"] = strings.ToUpper(_SignParameters)
	_req := Self_Post(_Params)
	return _req
}

//淘宝客-推广者-拉新活动对应数据查询
func Taobao_tbk_dg_newuser_order_sum() string {
	/*-----------基础参数-------------*/
	_Params := taobao_model.GetBaseParam()
	/*-----------方法名-------------*/
	_Params["method"] = "taobao.tbk.dg.newuser.order.sum"
	/*-----------请求参数-------------*/
	_Params["page_size"] = "20"        //必传项(*)页大小，默认20，1~100
	_Params["adzone_id"] = "123"       //mm_xxx_xxx_xxx的第三位
	_Params["page_no"] = "1"           //必传项(*)页码，默认1
	_Params["site_id"] = "123"         //mm_xxx_xxx_xxx的第二位
	_Params["activity_id"] = "sxxx"    //必传项(*)活动id， 活动名称与活动ID列表，请参见https://tbk.bbs.taobao.com/detail.html?appId=45301&postId=8599277
	_Params["settle_month"] = "201807" //结算月份
	/*-----------sign组合-------------*/
	_Sign := php2go.Rtrim(taobao_model.SignParameters(_Params, false)) //参数排序
	_SignParameters := php2go.Md5(_Sign)
	_Params["sign"] = strings.ToUpper(_SignParameters)
	_req := Self_Post(_Params)
	return _req
}

//淘宝客-服务商-拉新活动对应数据查询
func Taobao_tbk_sc_newuser_order_sum() string {
	/*-----------基础参数-------------*/
	_Params := taobao_model.GetBaseParam()
	/*-----------方法名-------------*/
	_Params["method"] = "taobao.tbk.sc.newuser.order.sum"
	/*-----------请求参数-------------*/
	_Params["page_size"] = "20"        //必传项(*)页大小，默认20，1~100
	_Params["adzone_id"] = "123"       //mm_xxx_xxx_xxx的第三位
	_Params["page_no"] = "1"           //必传项(*)页码，默认1
	_Params["site_id"] = "123"         //mm_xxx_xxx_xxx的第二位
	_Params["activity_id"] = "sxxx"    //必传项(*)活动id， 活动名称与活动ID列表，请参见https://tbk.bbs.taobao.com/detail.html?appId=45301&postId=8599277
	_Params["settle_month"] = "201807" //结算月份，需按照活动的结算月份传入具体的值：201807
	/*-----------sign组合-------------*/
	_Sign := php2go.Rtrim(taobao_model.SignParameters(_Params, false)) //参数排序
	_SignParameters := php2go.Md5(_Sign)
	_Params["sign"] = strings.ToUpper(_SignParameters)
	_req := Self_Post(_Params)
	return _req
}

//淘宝客-推广者-图文内容效果数据
func Taobao_tbk_content_effect_get() string {
	/*-----------基础参数-------------*/
	_Params := taobao_model.GetBaseParam()
	/*-----------方法名-------------*/
	_Params["method"] = "taobao.tbk.content.effect.get"
	/*-----------请求参数-------------*/
	_Params["option"] = "{'time':'2018-04-02','page_no':'1','page_size':'50','pid':'mm_1_1_1','content_id':'1234'}"

	/*-----------sign组合-------------*/
	_Sign := php2go.Rtrim(taobao_model.SignParameters(_Params, false)) //参数排序
	_SignParameters := php2go.Md5(_Sign)
	_Params["sign"] = strings.ToUpper(_SignParameters)
	_req := Self_Post(_Params)
	return _req
}

// 淘宝客-服务商-物料精选
func Taobao_tbk_sc_optimus_material() string {
	/*-----------基础参数-------------*/
	_Params := taobao_model.GetBaseParam()
	/*-----------方法名-------------*/
	_Params["method"] = "taobao.tbk.sc.optimus.material"
	/*-----------请求参数-------------*/
	_Params["page_size"] = "20"       //页大小，默认20，1~100
	_Params["adzone_id"] = "123"      //必传项(*)//mm_xxx_xxx_xxx的第三位
	_Params["page_no"] = "1"          //第几页，默认：1
	_Params["material_id"] = "123"    //官方的物料Id(详细物料id见：https://tbk.bbs.taobao.com/detail.html?appId=45301&postId=8576096)
	_Params["site_id"] = "111"        //必传项(*)mm_xxx_xxx_xxx的第二位
	_Params["device_type"] = "IMEI"   //智能匹配-设备号类型：IMEI，或者IDFA，或者UTDID（UTDID不支持MD5加密）
	_Params["device_encrypt"] = "MD5" //智能匹配-设备号加密类型：MD5
	_Params["device_value"] = "xxx"   //智能匹配-设备号加密后的值（MD5加密需32位小写）
	_Params["content_id"] = "323"     //内容专用-内容详情ID
	_Params["content_source"] = "xxx" //内容专用-内容渠道信息
	_Params["item_id"] = "33243"      //商品ID，用于相似商品推荐
	/*-----------sign组合-------------*/
	_Sign := php2go.Rtrim(taobao_model.SignParameters(_Params, false)) //参数排序
	_SignParameters := php2go.Md5(_Sign)
	_Params["sign"] = strings.ToUpper(_SignParameters)
	_req := Self_Post(_Params)
	return _req
}

// 淘宝客-淘礼金创建
func Taobao_tbk_dg_vegas_tlj_create() string {
	/*-----------基础参数-------------*/
	_Params := taobao_model.GetBaseParam()
	/*-----------方法名-------------*/
	_Params["method"] = "taobao.tbk.dg.vegas.tlj.create"
	/*-----------请求参数-------------*/
	_Params["campaign_type"] = "定向：DX；鹊桥：LINK_EVENT；营销：MKT" //CPS佣金计划类型
	_Params["adzone_id"] = "1234567"                        //必传项(*)妈妈广告位Id
	_Params["item_id"] = "1234567"                          //必传项(*)宝贝id
	_Params["total_num"] = "10"                             //必传项(*)淘礼金总个数
	_Params["name"] = "淘礼金来啦"                               //必传项(*)淘礼金名称，最大10个字符
	_Params["user_total_win_num_limit"] = "1"               //必传项(*)单用户累计中奖次数上限
	_Params["security_switch"] = "启动安全：true；不启用安全：false"    //必传项(*)安全开关
	_Params["per_face"] = "10"                              //必传项(*)单个淘礼金面额，支持两位小数，单位元
	_Params["send_start_time"] = "2018-09-01 00:00:00"      //必传项(*)发放开始时间
	_Params["send_end_time"] = "2018-09-01 00:00:00"        //发放截止时间
	_Params["use_end_time"] = "5"                           //使用结束日期。如果是结束时间模式为相对时间，时间格式为1-7直接的整数, 例如，1（相对领取时间1天）； 如果是绝对时间，格式为yyyy-MM-dd，例如，2019-01-29，表示到2019-01-29 23:59:59结束
	_Params["use_end_time_mode"] = "1"                      //结束日期的模式,1:相对时间，2:绝对时间
	_Params["use_start_time"] = "2019-01-29"                //使用开始日期。相对时间，无需填写，以用户领取时间作为使用开始时间。绝对时间，格式 yyyy-MM-dd，例如，2019-01-29，表示从2019-01-29 00:00:00 开始
	/*-----------sign组合-------------*/
	_Sign := php2go.Rtrim(taobao_model.SignParameters(_Params, false)) //参数排序
	_SignParameters := php2go.Md5(_Sign)
	_Params["sign"] = strings.ToUpper(_SignParameters)
	_req := Self_Post(_Params)
	return _req
}

//淘宝客-推广者-官方活动转链
func Taobao_tbk_activitylink_get() string {
	/*-----------基础参数-------------*/
	_Params := taobao_model.GetBaseParam()
	/*-----------方法名-------------*/
	_Params["method"] = "taobao.tbk.activitylink.get"
	/*-----------请求参数-------------*/
	_Params["platform"] = "1"                  //	1：PC，2：无线，默认：１
	_Params["union_id"] = "demo"               //	自定义输入串，英文和数字组成，长度不能大于12个字符，区分不同的推广渠道
	_Params["adzone_id"] = "123"               //必传项(*)推广位id，mm_xx_xx_xx pid三段式中的第三段。adzone_id需属于appKey拥有者
	_Params["promotion_scene_id"] = "12345678" //必传项(*)官方活动ID，从官方活动页获取
	_Params["sub_pid"] = "mm_123_123_123"      //媒体平台下达人的淘客pid
	_Params["relation_id"] = "23"              //渠道关系ID，仅适用于渠道推广场景
	/*-----------sign组合-------------*/
	_Sign := php2go.Rtrim(taobao_model.SignParameters(_Params, false)) //参数排序
	_SignParameters := php2go.Md5(_Sign)
	_Params["sign"] = strings.ToUpper(_SignParameters)
	_req := Self_Post(_Params)
	return _req
}

//淘宝客-服务商-官方活动转链
func Taobao_tbk_sc_activitylink_toolget() string {
	/*-----------基础参数-------------*/
	_Params := taobao_model.GetBaseParam()
	/*-----------方法名-------------*/
	_Params["method"] = "taobao.tbk.sc.activitylink.toolget"
	/*-----------请求参数-------------*/
	_Params["adzone_id"] = "123"               //必传项(*)推广位id，mm_xx_xx_xx pid三段式中的第三段
	_Params["site_id"] = "456"                 //必传项(*)推广位id，mm_xx_xx_xx pid三段式中的第二段
	_Params["platform"] = "1"                  //1：PC，2：无线，默认：１
	_Params["union_id"] = "demo"               //自定义输入串，英文和数字组成，长度不能大于12个字符，区分不同的推广渠道
	_Params["relation_id"] = "23"              //渠道关系ID，仅适用于渠道推广场景
	_Params["promotion_scene_id"] = "12345678" //必传项(*)官方活动ID，从官方活动页获取
	/*-----------sign组合-------------*/
	_Sign := php2go.Rtrim(taobao_model.SignParameters(_Params, false)) //参数排序
	_SignParameters := php2go.Md5(_Sign)
	_Params["sign"] = strings.ToUpper(_SignParameters)
	_req := Self_Post(_Params)
	return _req
}

//淘宝客-推广者-处罚订单查询
func Taobao_tbk_dg_punish_order_get() string {
	/*-----------基础参数-------------*/
	_Params := taobao_model.GetBaseParam()
	/*-----------方法名-------------*/
	_Params["method"] = "taobao.tbk.dg.punish.order.get"
	/*-----------请求参数-------------*/
	_Params["af_order_option"] = "{'span':'1200','relation_id':'2222','tb_trade_id':'258897956183171983','tb_trade_parent_id':'258897956183171983','page_size':'1','page_no':'10','start_time':'2018-11-11 00:01:01','special_id':'23132','violation_type':'1','punish_status':'1'}"
	/*-----------sign组合-------------*/
	_Sign := php2go.Rtrim(taobao_model.SignParameters(_Params, false)) //参数排序
	_SignParameters := php2go.Md5(_Sign)
	_Params["sign"] = strings.ToUpper(_SignParameters)
	_req := Self_Post(_Params)
	return _req
}

//淘宝客-推广者-淘礼金发放及使用报表
func Taobao_tbk_dg_vegas_tlj_instance_report() string {
	/*-----------基础参数-------------*/
	_Params := taobao_model.GetBaseParam()
	/*-----------方法名-------------*/
	_Params["method"] = "taobao.tbk.dg.vegas.tlj.instance.report"
	/*-----------请求参数-------------*/
	_Params["rights_id"] = "UZvYlKXRdTIBf%2B%2F%2FIV9MGw%3D%3D" //必传项(*)实例ID
	/*-----------sign组合-------------*/
	_Sign := php2go.Rtrim(taobao_model.SignParameters(_Params, false)) //参数排序
	_SignParameters := php2go.Md5(_Sign)
	_Params["sign"] = strings.ToUpper(_SignParameters)
	_req := Self_Post(_Params)
	return _req
}

/*------------------- 工具函数处理----------------- */
