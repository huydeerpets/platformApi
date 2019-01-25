package api

import (
	"fmt"
	"html"
	"platformApi/src/model"
	"strings"

)

// 问题反馈相关操作
type ApiQuestionController struct {
	BaseController
}
//func (this *ApiQuestionController) Prepare() {
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
// @Title 问题反馈
// @Description 问题反馈
// @Param	email			formData  string	true		"邮箱"
// @Param	title			formData  string	true		"标题"
// @Param	question		formData  string	true		"问题简介"
// @Param	baseUrl			formData  string	false		"附件根url"
// @Param	attachName		formData  string	false		"附件名称，多个以逗号分隔"
// @Success 200 {object} model.JsonResult
// @router /addQuestion [post]
func (o *ApiQuestionController) Post() {
	var result model.JsonResult

	email := o.GetString("email")
	title := o.GetString("title")
	question := o.GetString("question")
	baseURL := o.GetString("baseUrl")
	attachName := o.GetString("attachName")

	if email == "" {
		result.Status = false
		//result.Msg = "邮箱不能为空！"
		result.Msg = o.Tr("ERROR_QUESTION.emailIsEmpty")
		o.Data["json"] = result
		o.ServeJSON()
		return
	}

	if title == "" {
		result.Status = false
		//result.Msg = "标题不能为空！"
		result.Msg = o.Tr("ERROR_QUESTION.titleIsBlank")
		o.Data["json"] = result
		o.ServeJSON()
		return
	}

	if question == "" {
		result.Status = false
		//result.Msg = "问题简介不能为空！"
		result.Msg = o.Tr("ERROR_QUESTION.problemProfileIsEmpty")
		o.Data["json"] = result
		o.ServeJSON()
		return
	}
	strArr := strings.Split(attachName, ",")
	if len(strArr) > 5 {
		result.Status = false
		//result.Msg = "最多允许上传5个文件！"
		result.Msg = o.Tr("ERROR_QUESTION.fileUploadLimit")
		o.Data["json"] = result
		o.ServeJSON()
		return
	}
	for _, v := range strArr {
		if strings.Index(v, ".jpg") == -1 &&
			strings.Index(v, ".png") == -1 &&
			strings.Index(v, ".gif") == -1 {
			result.Status = false
			//result.Msg = "上传文件格式只能是png, jpg, gif！"
			result.Msg = o.Tr("ERROR_QUESTION.fileFormatLimit")
			o.Data["json"] = result
			o.ServeJSON()
			return
		}
	}
	question = html.EscapeString(question)

	var quest model.Question
	quest.Email = email
	quest.Title = title
	quest.Question = question
	quest.BaseUrl = baseURL
	quest.AttachName = attachName

	err := model.AddQuestion(&quest)
	if err != nil {
		result.Status = false
		result.Msg = err.Error()
	} else {
		result.Status = true
		result.Data = fmt.Sprintf("%d", quest.Id)
	}
	o.Data["json"] = result
	o.ServeJSON()
}
