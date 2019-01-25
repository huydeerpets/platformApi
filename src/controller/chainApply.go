package controller

import (
	"fmt"
	"platformApi/src/model"

	"github.com/astaxie/beego"
)

type ChainApplyController struct {
	beego.Controller
}

// @router /ChainApply/InitList/ [get]
func (c *ChainApplyController) InitList() {
	c.TplName = "chainapply_list.html"
}

// @router /ChainApply/InitAdd/ [get]
func (c *ChainApplyController) InitAdd() {
	id := c.GetString("id")
	if id != "" {
		apply, _ := model.QueryChainApplyInfo(id)
		c.Data["chainApply"] = apply
	}
	c.TplName = "chainapply_info.html"
}

// @router /ChainApply/List/ [post]
func (c *ChainApplyController) List() {

	keyWord := c.GetString("keyWord")

	pageNo, _ := c.GetInt("current")

	rowCount, _ := c.GetInt("rowCount")

	if pageNo == 0 {
		pageNo = 1
	}
	resultMap := model.SearchChainApplys(rowCount, pageNo, keyWord)
	c.Data["json"] = map[string]interface{}{"rows": resultMap["data"], "rowCount": rowCount, "current": pageNo, "total": resultMap["total"]}

	c.ServeJSON()
}

// @router /ChainApply/UpdateChainApplyStatus/:id/:status [get]
func (c *ChainApplyController) UpdateChainApplyStatus() {
	result := new(model.Result)
	id := c.Ctx.Input.Param(":id")
	status := c.Ctx.Input.Param(":status")
	err := model.UpdateChainApplyStatus(id, status)
	if err != nil {
		result.Status = 1
		result.Msg = fmt.Sprintf("设置状态失败,%s", err.Error())
	} else {
		result.Status = 0
	}

	c.Data["json"] = result
	c.ServeJSON()
}
