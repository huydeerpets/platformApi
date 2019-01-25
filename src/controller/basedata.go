package controller

import (
	"fmt"
	"platformApi/src/model"
	"strconv"

	"github.com/astaxie/beego"
)

type BaseDataController struct {
	beego.Controller
}

// @router /BaseData/InitList/ [get]
func (c *BaseDataController) InitList() {
	c.TplName = "basedatalist.html"
}

// @router /BaseData/List/ [post]
func (c *BaseDataController) List() {

	dataType := c.GetString("dataType")

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
	resultMap := model.SearchBaseData(rowCount, pageNo, dataType, keyWord, langType)
	c.Data["json"] = map[string]interface{}{"rows": resultMap["data"], "rowCount": rowCount, "current": pageNo, "total": resultMap["total"]}

	c.ServeJSON()
}

// @router /BaseData/InitAdd/ [get]
func (c *BaseDataController) InitAdd() {
	id := c.GetString("id")
	langType := c.GetString("langType")
	if id != "" {
		baseData, _ := model.GetBaseDataInfo(id)
		if baseData != nil {
			c.Data["baseData"] = baseData
		}
	}
	c.Data["langType"] = langType
	c.TplName = "basedataadd.html"
}

// @router /BaseData/AddBaseData [post]
func (c *BaseDataController) AddBaseData() {
	//自动解析绑定到对象中,ParseForm 不支持解析raw data,必须是表单form提交
	baseData := model.BaseData{}
	result := new(model.Result)
	var err error
	err = c.ParseForm(&baseData)
	idStr := c.GetString("id")
	if err == nil {
		if idStr != "" {
			id, _ := strconv.ParseInt(idStr, 10, 64)
			baseData.Id = id
			err = model.UpdateBaseData(&baseData)
		} else {
			err = model.AddBaseData(&baseData)
		}

		if err != nil {
			result.Status = 1
			result.Msg = err.Error()
		} else {
			result.Status = 0
			result.Data = map[string]int64{"Id": baseData.Id}
		}
	} else {
		result.Status = 1
		result.Msg = fmt.Sprintf("操作失败,%s", err.Error())
	}
	c.Data["json"] = result
	c.ServeJSON()
}

// @router /BaseData/DeleteBaseData/:id [get]
func (c *BaseDataController) DeleteBaseData() {
	result := new(model.Result)
	id, _ := strconv.ParseInt(c.Ctx.Input.Param(":id"), 10, 64)
	err := model.DeleteBaseData(id)

	if err != nil {
		result.Status = 1
		result.Msg = fmt.Sprintf("删除失败,%s", err.Error())
	} else {
		result.Status = 0
	}

	c.Data["json"] = result
	c.ServeJSON()
}

// @router /BaseData/DisabledBaseData/:id [get]
func (c *BaseDataController) DisabledBaseData() {
	result := new(model.Result)
	id, _ := strconv.ParseInt(c.Ctx.Input.Param(":id"), 10, 64)
	err := model.DisabledBaseData(id)

	if err != nil {
		result.Status = 1
		result.Msg = fmt.Sprintf("禁用失败,%s", err.Error())
	} else {
		result.Status = 0
	}

	c.Data["json"] = result
	c.ServeJSON()
}

// @router /BaseData/EnabledBaseData/:id [get]
func (c *BaseDataController) EnabledBaseData() {
	result := new(model.Result)
	id, _ := strconv.ParseInt(c.Ctx.Input.Param(":id"), 10, 64)
	err := model.EnabledBaseData(id)

	if err != nil {
		result.Status = 1
		result.Msg = fmt.Sprintf("启用失败,%s", err.Error())
	} else {
		result.Status = 0
	}

	c.Data["json"] = result
	c.ServeJSON()
}

// @router /BaseData/GetBaseDataByType/:dataType/:langType [get]
func (c *BaseDataController) GetBaseDataByType() {
	result := new(model.Result)
	dataType := c.Ctx.Input.Param(":dataType")
	langType := c.Ctx.Input.Param(":langType")
	res, err := model.QueryBaseDataByType(dataType, langType)

	if err != nil {
		result.Status = 1
		result.Msg = fmt.Sprintf("查询失败,%s", err.Error())
	} else {
		result.Status = 0
		result.Data = res
	}

	c.Data["json"] = result
	c.ServeJSON()
}

// @router /BaseData/GetBaseDataByTypeExt/:dataType/:langType [get]
func (c *BaseDataController) GetBaseDataByTypeExt() {
	result := new(model.Result)
	dataType := c.Ctx.Input.Param(":dataType")
	langType := c.Ctx.Input.Param(":langType")
	res, err := model.QueryBaseDataByTypeExt(dataType, langType)

	if err != nil {
		result.Status = 1
		result.Msg = fmt.Sprintf("查询失败,%s", err.Error())
	} else {
		result.Status = 0
		result.Data = res
	}

	c.Data["json"] = result
	c.ServeJSON()
}
