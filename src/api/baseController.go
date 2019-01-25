package api

import (
	"github.com/astaxie/beego"
	"github.com/beego/i18n"
)

type (
	BaseController struct {
		beego.Controller
		i18n.Locale //处理国际化的包
	}
)

//语言版本判断，便于选择对应的语言信息文件
func (this *BaseController) Prepare() {
	lang := this.Ctx.Input.Header("language")
	beego.Debug("lang =  ", lang)
	if lang == "zh" {
		this.Lang = "zh-CN"
	} else {
		this.Lang = "en-US"
	}
}
