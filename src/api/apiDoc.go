package api

import (
	"platformApi/src/model"
	"strconv"

)

// 相关文档查询接口
type ApiDocController struct {
	BaseController
}

// @Title 获取新手指导主题信息
// @Description 获取新手指导主题信息
// @Success 200 {object} model.DocThemeListResult
// @router /queryDocTheme [get]
func (o *ApiDocController) QueryDocTheme() {
	var result model.DocThemeListResult

	data := model.QueryDocThemeList(o.Ctx.Input.Header("language"))
	result.Status = true
	result.Data = data

	o.Data["json"] = result
	o.ServeJSON()
}

// @Title 获取新手指导所有主题信息
// @Description 获取新手指导所有主题信息
// @Success 200 {object} model.DocThemeListResult
// @router /queryDocAllTheme [get]
func (o *ApiDocController) QueryDocAllTheme() {
	var result model.DocThemeListResult

	data := model.QueryDocAllThemeList()
	result.Status = true
	result.Data = data

	o.Data["json"] = result
	o.ServeJSON()
}

// @Title 获取文档详情
// @Description 获取文档详情
// @Param	id		path 	string	true		"文档记录ID"
// @Success 200 {object} model.DocInfoResult
// @router /queryDocInfo/:id [get]
func (o *ApiDocController) QueryDocInfo() {
	var result model.DocInfoResult
	idStr := o.Ctx.Input.Param(":id")
	if idStr == "undefined" {
		idStr = "0"
	}
	id, _ := strconv.ParseInt(idStr, 10, 64)
	data := model.GetDocContentInfo(id)
	result.Status = true
	result.Data = data

	o.Data["json"] = result
	o.ServeJSON()
}
