package routers

import (
	"github.com/qnsoft/mall_online/controllers/mall_other/taobao"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
)

func init() {
	//加入跨域权限
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type", "Token"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type", "Token"},
		AllowCredentials: true}))
	beego.SetStaticPath("/upload", "upload") //重定向静态路径
	beego.SetStaticPath("/views", "views")   //重定向静态路径
	//-------------------------基础接口开始-----------------------------------//
	/* //根目录
	beego.Router("/", &controllers.Default_Controller{}, "get,post:Get")
	//先获取access_token
	beego.Router("/access_token", &Token.AccesstokenController{}, "post:Access_Token")
	//token测试
	beego.Router("/testtoken", &controllers.Default_Controller{}, "get,post:TestToken")
	//上传图片
	beego.Router("/img/upload_pic", &ImageUpload.Image_Uplaod_Controller{}, "post:Uplaod_Pic")
	//获取图片
	beego.Router("/img/info_pic", &ImageUpload.Image_Uplaod_Controller{}, "get:Info_Pic")
	//发送短信验证码
	beego.Router("/sms/MsgSend", &Sms.Sms_Controller{}, "post:MsgSend") */
	//-------------------------基础接口结束-----------------------------------//
	//-------------------------对接接口开始-----------------------------------//
	//阿里商品列表
	//接收官方回传
	beego.Router("/api/mallother/taobao/Callback", &taobao.Mall_other_taobao_Controller{}, "post,get:Callback")
	//接收三方接口回传
	beego.Router("/api/mallother/taobao/Callback_other", &taobao.Mall_other_taobao_Controller{}, "post,get:Callback_other")
	beego.Router("/api/mallother/taobao/goods", &taobao.Mall_other_taobao_Controller{}, "post:Goods")
	//阿里妈妈推广券
	beego.Router("/api/mallother/taobao/Taobao_YouhuiQuan", &taobao.Mall_other_taobao_Controller{}, "post:Taobao_YouhuiQuan")
	///--------------维易淘宝客API---------------------------//
	//维易淘宝客API--获取授权地址url
	beego.Router("/api/mallother/taobao/Get_auth_url?:user_id", &taobao.Wytklm_Controller{}, "get:Get_auth_url")
	beego.Router("/api/mallother/taobao/Get_directhc?:num_iid/:relation_id", &taobao.Wytklm_Controller{}, "get:Get_directhc")
}
