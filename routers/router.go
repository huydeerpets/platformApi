// @APIVersion 1.0.0
// @Title platform API
// @Description 平台链官网相关api接口
// @Contact 278985177@qq.com
// @TermsOfServiceUrl http://www.m-chain.com
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"platformApi/src/api"
	"platformApi/src/controller"

	"github.com/astaxie/beego"
	"github.com/dchest/captcha"
)

func init() {
	//默认的
	beego.Router("/", &controller.MainController{})

	//验证码
	beego.Handler("/captcha/*.png", captcha.Server(106, 36))

	//JS分页
	beego.Router("/Home/PageNextData", &controller.YonghuController{})

	//Easyui使用
	beego.Router("/Home/EasyUI", &controller.EasyUIController{})
	beego.Router("/Home/EasyUIData", &controller.EasyUIDataController{})

	//Api接口部分
	beego.Router("/api/Html", &api.ApiController{})
	beego.Router("/api/GetJson", &api.ApiJsonController{})
	beego.Router("/api/GetXml", &api.ApiXMLController{})
	beego.Router("/api/GetJsonp", &api.ApiJsonpController{})
	beego.Router("/api/GetDictionary", &api.ApiDictionaryController{})
	beego.Router("/api/GetParams", &api.ApiParamsController{})

	//session部分
	beego.Router("/Home/Login", &controller.LoginController{})
	beego.Router("/Home/InitReg", &controller.LoginController{}, "get:InitReg")
	beego.Router("/Home/InitPass", &controller.LoginController{}, "get:InitPass")
	beego.Router("/Home/Reg", &controller.LoginController{}, "post:Reg")
	beego.Router("/Home/UpdatePwd", &controller.LoginController{}, "post:UpdatePwd")

	//文件管理
	beego.Include(&controller.FileController{})

	//用户管理
	beego.Include(&controller.UserController{})

	//api分类
	beego.AutoRouter(&controller.ApiCategoryController{})

	//文档管理
	beego.AutoRouter(&controller.DocCategoryController{})

	//GET upload policy token
	beego.AutoRouter(&controller.OSSController{})

	//基础数据管理
	beego.Include(&controller.BaseDataController{})

	//问题反馈
	beego.Include(&controller.QuestionController{})

	//FAQ管理
	beego.Include(&controller.FaqController{})

	//链入驻申请管理
	beego.Include(&controller.ChainApplyController{})

	beego.Router("/Home/Logout", &controller.LogoutController{})
	//布局页面部分
	beego.Router("/Home/Layout", &controller.LayoutController{})

	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/sign",
			beego.NSInclude(
				&api.ApiSignController{},
			),
		),
		beego.NSNamespace("/contract",
			beego.NSInclude(
				&api.ContractController{},
			),
		),
		beego.NSNamespace("/token",
			beego.NSInclude(
				&api.TokenController{},
			),
		),
		beego.NSNamespace("/chain",
			beego.NSInclude(
				&api.ChainController{},
			),
		),
		beego.NSNamespace("/question",
			beego.NSInclude(
				&api.ApiQuestionController{},
			),
		),
		beego.NSNamespace("/api",
			beego.NSInclude(
				&api.ApiQueryController{},
			),
		),
		beego.NSNamespace("/faq",
			beego.NSInclude(
				&api.ApiFaqController{},
			),
		),
		beego.NSNamespace("/doc",
			beego.NSInclude(
				&api.ApiDocController{},
			),
		),
		beego.NSNamespace("/version",
			beego.NSInclude(
				&api.ApiVersionController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
