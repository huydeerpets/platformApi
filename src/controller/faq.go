package controller

import (
	"fmt"
	"platformApi/src/model"
	"strconv"

	"github.com/astaxie/beego"
)

type FaqController struct {
	beego.Controller
}

// @router /FAQ/InitList/ [get]
func (c *FaqController) InitList() {
	c.TplName = "faq_list.html"
}

// @router /FAQ/List/ [post]
func (c *FaqController) List() {
	keyWord := c.GetString("keyWord")

	pageNo, _ := c.GetInt("current")

	rowCount, _ := c.GetInt("rowCount")

	langType := c.GetString("langType")
	if langType == "" {
		langType = "1"
	}

	if pageNo == 0 {
		pageNo = 1
	}
	resultMap := model.SearchFAQS(rowCount, pageNo, keyWord, langType)
	c.Data["json"] = map[string]interface{}{"rows": resultMap["data"], "rowCount": rowCount, "current": pageNo, "total": resultMap["total"]}

	c.ServeJSON()
}

// @router /FAQ/InitAdd/ [get]
func (c *FaqController) InitAdd() {
	id := c.GetString("id")
	if id != "" {
		faq := model.GetFaqInfo(id)
		if faq != nil {
			c.Data["faq"] = faq
		}
	}
	c.TplName = "faq_add.html"
}

// @router /FAQ/SaveFAQ [post]
func (c *FaqController) SaveFAQ() {
	//自动解析绑定到对象中,ParseForm 不支持解析raw data,必须是表单form提交
	faq := model.FaqModel{}
	result := new(model.Result)
	var err error
	title := c.GetString("title")
	content := c.GetString("content")
	idStr := c.GetString("id")
	faq.Title = title
	faq.Content = content
	var id int64
	if err == nil {
		if idStr != "" {
			id, _ = strconv.ParseInt(idStr, 10, 64)
			faq.Id = id
			err = model.UpdateFaqInfo(id, faq.Title, faq.Content)
		} else {
			id, err = model.SaveFaqInfo(&faq)
		}

		if err != nil {
			result.Status = 1
			result.Msg = err.Error()
		} else {
			result.Status = 0
			result.Data = map[string]int64{"Id": id}
		}
	} else {
		result.Status = 1
		result.Msg = fmt.Sprintf("操作失败,%s", err.Error())
	}
	c.Data["json"] = result
	c.ServeJSON()
}

// @router /FAQ/DeleteFAQ/:id [get]
func (c *FaqController) DeleteFAQ() {
	result := new(model.Result)
	id, _ := strconv.ParseInt(c.Ctx.Input.Param(":id"), 10, 64)
	_, err := model.DeleteFaqInfo(id)

	if err != nil {
		result.Status = 1
		result.Msg = fmt.Sprintf("删除失败,%s", err.Error())
	} else {
		result.Status = 0
	}

	c.Data["json"] = result
	c.ServeJSON()
}

// @router /FAQ/UpdateFAQStatus/:id [get]
func (c *FaqController) UpdateFAQStatus() {
	result := new(model.Result)
	id, _ := strconv.ParseInt(c.Ctx.Input.Param(":id"), 10, 64)
	err := model.UpdateFAQStatus(id)

	if err != nil {
		result.Status = 1
		result.Msg = fmt.Sprintf("操作失败,%s", err.Error())
	} else {
		result.Status = 0
	}

	c.Data["json"] = result
	c.ServeJSON()
}
