package controller

import (
	"github.com/astaxie/beego"
)

type FileController struct {
	beego.Controller
}

// @router /File/UploadToken/ [get]
func (c *FileController) UploadToken() {
	//c.Data["json"] = map[string]interface{}{"uptoken": common.GetUploadToken()}
	c.ServeJSON()
}
