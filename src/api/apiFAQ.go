package api

import (
	"platformApi/src/model"

)

// FAQ相关操作
type ApiFaqController struct {
	BaseController
}
//func (this *ApiFaqController) Prepare() {
//	lang := this.Ctx.Input.Header("language")
//	beego.Debug("lang =  ", lang)
//	if lang == "zh" {
//		this.Lang = "zh-CN"
//	} else if lang == "tw" {
//		this.Lang = "zh-TW"
//	} else {
//		this.Lang = "en-US"
//	}
//}
// @Title queryFaqList
// @Description 查询faq列表信息
// @Success 200 {object} model.FaqListResult
// @router /queryFaqList [get]
func (o *ApiFaqController) QueryFaqList() {
	var result model.FaqListResult
	result.Status = true
	result.Data = model.QueryFaqs(o.Ctx.Input.Header("language"))

	o.Data["json"] = result
	o.ServeJSON()
}

// @Title queryFaqAllList
// @Description 查询所有的faq信息
// @Success 200 {object} model.FaqListResult
// @router /queryFaqAllList [get]
func (o *ApiFaqController) QueryFaqAllList() {
	var result model.FaqListResult
	result.Status = true
	result.Data = model.QueryAllFaqs()

	o.Data["json"] = result
	o.ServeJSON()
}

// @Title queryFaqInfo
// @Description 查询FAQ的详情
// @Param	id		path 	string	true		"faq记录ID"
// @Success 200 {object} model.FaqResult
// @Failure 403 :faq记录ID不能为空
// @router /queryFaqInfo/:id [get]
func (o *ApiFaqController) QueryFaqInfo() {
	id := o.Ctx.Input.Param(":id")
	var result model.FaqResult
	if id == "" {
		//result.Msg = "FAQ记录ID不能为空！"
		result.Msg = o.Tr("ERROR_FAQ.recordIDIsEmpty")
	} else {
		result.Status = true
		result.Data = model.GetFaqInfo(id)
	}
	o.Data["json"] = result
	o.ServeJSON()
}

// @Title queryTechnologyList
// @Description 技术问题搜索
// @Param	keyWord		formData 	string	true 	"搜索关键字"
// @Success 200 {object} model.TechnologyListResult
// @Failure 403 :搜索关键字不能为空
// @router /queryTechnologyList [post]
func (o *ApiFaqController) QueryTechnologyList() {
	keyWord := o.GetString("keyWord")
	var result model.TechnologyListResult
	if keyWord == "" {
		//result.Msg = "搜索关键字不能为空！"
		result.Msg = o.Tr("ERROR_FAQ.searchKeywordIsEmpty")
		o.Data["json"] = result
		o.ServeJSON()
	}
	result.Status = true
	result.Data = model.SearchTechnologys(keyWord)
	o.Data["json"] = result
	o.ServeJSON()
}

/*
// @Title queryTechnologyInfo
// @Description 查询内容详情
// @Param	id		path 	string	true		"记录ID"
// @Param	ctype	path 	string	true		"内容类型　1:api文档 2:技术教程 3:FAQ"
// @Success 200 {object} model.FaqResult
// @Failure 403 :faq记录ID不能为空
// @router /queryTechnologyInfo/:id/:ctype [get]
func (o *ApiFaqController) QueryTechnologyInfo() {
	id := o.Ctx.Input.Param(":id")
	ctype := o.Ctx.Input.Param(":ctype")

	var result model.FaqResult
	if id == "" {
		result.Msg = "FAQ记录ID不能为空！"
	} else {
		result.Status = true
		result.Data = model.GetFaqInfo(id)
	}
	o.Data["json"] = result
	o.ServeJSON()
}*/
