package controller

import (
	"platformApi/src/common"

	"github.com/astaxie/beego"
)

type OSSController struct {
	beego.Controller
}

//get请求
func (c *OSSController) GetPolicyToken() {
	common.HandlerRequest(c.Ctx.ResponseWriter, c.Ctx.Request)
}

//post请求
func (c *OSSController) Callback() {
	common.HandlerRequest(c.Ctx.ResponseWriter, c.Ctx.Request)
}
