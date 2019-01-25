package controller

import (
	"bytes"
	"fmt"
	"platformApi/src/model"

	"strconv"

	"github.com/astaxie/beego"
)

type DocCategoryController struct {
	beego.Controller
}

func (c *DocCategoryController) InitDocCategory() {
	paramID := c.Ctx.Input.Params()["0"]
	subCount := c.Ctx.Input.Params()["1"]
	nodeLevel := c.Ctx.Input.Params()["2"]
	id, _ := strconv.ParseInt(paramID, 10, 64)
	cate := model.GetDocCateInfo(id)
	if nodeLevel == "4" {
		docCont := model.GetDocContentInfo(id)
		c.Data["docCont"] = docCont
	}

	c.Data["cate"] = cate
	c.Data["subCount"] = subCount
	c.Data["nodeLevel"] = nodeLevel
	if nodeLevel == "3" || nodeLevel == "4" {
		c.TplName = "doc_info_add.html"
	} else {
		c.TplName = "doc_category_add.html"
	}
}

func (c *DocCategoryController) GetDocCateList() {
	cates, err := model.GetDocCateList()
	result := new(model.Result)
	if err != nil {
		result.Status = 1
		result.Msg = "获取分类信息失败！"
	} else {
		result.Status = 0
		result.Data = cates
	}
	c.Data["json"] = result
	c.ServeJSON()
}

/**
获取分级的下拉列表数据
*/
func (c *DocCategoryController) GetDocCateDropList() {
	cates := model.GetDocCategoryList("1", true)

	arr := make([]model.DropItem, 0)

	arr = model.GetDocList(cates, 1, arr, 0)

	result := new(model.Result)
	result.Status = 0
	result.Data = arr
	c.Data["json"] = result
	c.ServeJSON()
}

func (c *DocCategoryController) GetDocCateInfo() {
	paramID := c.Ctx.Input.Params()["0"]
	id, _ := strconv.ParseInt(paramID, 10, 64)
	cate := model.GetDocCateInfo(id)

	//b, _ := json.Marshal(cate)
	//fmt.Printf("***********%#v", string(b))

	result := new(model.Result)
	result.Status = 0
	result.Data = cate
	c.Data["json"] = result
	c.ServeJSON()
}

func (c *DocCategoryController) SaveDocCateInfo() {
	pid, _ := strconv.ParseInt(c.GetString("pid"), 10, 64)
	full_path := c.GetString("full_path")
	depth, _ := strconv.ParseInt(c.GetString("depth"), 10, 64)
	name := c.GetString("name")
	subCount, _ := strconv.ParseInt(c.GetString("subCount"), 10, 64)

	newPath := fmt.Sprintf("%s,%d", full_path, pid)

	var cate model.DocCategory
	cate.Pid = pid
	cate.Name = name
	cate.Depth = depth + 1
	cate.FullPath = newPath
	cate.OrderNum = subCount

	v, err := model.SaveDocCategroy(&cate)

	result := new(model.Result)
	if err == nil {
		result.Status = 0
		result.Data = map[string]int64{"Id": v}
	} else {
		result.Status = 1
		result.Msg = fmt.Sprintf("保存分类信息失败,%s", err.Error())
	}
	c.Data["json"] = result
	c.ServeJSON()
}

func (c *DocCategoryController) UpdateDocCateInfo() {
	paramID := c.GetString("id")
	id, _ := strconv.ParseInt(paramID, 10, 64)
	name := c.GetString("name")
	orderNum, _ := strconv.ParseInt(c.GetString("orderNum"), 10, 64)
	err := model.UpdateDocCategroy(id, orderNum, name)

	result := new(model.Result)
	if err == nil {
		result.Status = 0
	} else {
		result.Status = 1
		result.Msg = "更新分类信息失败！"
	}
	c.Data["json"] = result
	c.ServeJSON()
}

func (c *DocCategoryController) DelDocCateInfo() {
	id, _ := strconv.ParseInt(c.GetString("id"), 10, 64)
	pid, _ := strconv.ParseInt(c.GetString("pid"), 10, 64)
	orderNum, _ := strconv.ParseInt(c.GetString("orderNum"), 10, 64)
	nodeLevel, _ := strconv.ParseInt(c.GetString("nodeLevel"), 10, 64)

	_, err := model.DeleteDocCategroy(id, pid, orderNum, nodeLevel)

	result := new(model.Result)
	if err == nil {
		result.Status = 0
	} else {
		result.Status = 1
		result.Msg = "删除信息失败！"
	}
	c.Data["json"] = result
	c.ServeJSON()
}

func (c *DocCategoryController) SaveDocInfo() {
	pid, _ := strconv.ParseInt(c.GetString("pid"), 10, 64)
	full_path := c.GetString("full_path")
	depth, _ := strconv.ParseInt(c.GetString("depth"), 10, 64)
	name := c.GetString("name")
	subCount, _ := strconv.ParseInt(c.GetString("subCount"), 10, 64)

	newPath := fmt.Sprintf("%s,%d", full_path, pid)

	var cate model.DocCategory
	cate.Pid = pid
	cate.Name = name
	cate.Depth = depth + 1
	cate.FullPath = newPath
	cate.OrderNum = subCount

	v, err := model.SaveDocCategroy(&cate)

	result := new(model.Result)
	if err == nil {
		var docCont model.DocContent
		docCont.CateId = v
		docCont.Title = name
		docCont.Summary = c.GetString("summary")
		docCont.Content = c.GetString("content")
		docCont.AttachAddr = c.GetString("attachAddr")
		docCont.Status = 1
		_, err = model.SaveDocInfo(docCont)
	}

	if err == nil {
		result.Status = 0
		result.Data = map[string]int64{"Id": v}
	} else {
		result.Status = 1
		result.Msg = fmt.Sprintf("保存信息失败,%s", err.Error())
	}
	c.Data["json"] = result
	c.ServeJSON()
}

func (c *DocCategoryController) UpdateDocInfo() {
	cateID, _ := strconv.ParseInt(c.GetString("pid"), 10, 64)
	name := c.GetString("name")
	summary := c.GetString("summary")
	content := c.GetString("content")
	attachAddr := c.GetString("attachAddr")

	err := model.UpdateDocContentInfo(cateID, name, summary, content, attachAddr)

	result := new(model.Result)
	if err == nil {
		result.Status = 0
	} else {
		result.Status = 1
		result.Msg = fmt.Sprintf("保存信息失败,%s", err.Error())
	}
	c.Data["json"] = result
	c.ServeJSON()
}

//获取分类信息
func (c *DocCategoryController) List() {
	langType := c.GetString("langType")
	if langType == "" {
		langType = "1"
	}
	cates := model.GetDocCategoryList(langType, false)

	var buffer bytes.Buffer //Buffer是一个实现了读写方法的可变大小的字节缓冲
	for _, m := range cates {
		/*
		   func (b *Buffer) WriteString(s string) (n int, err error)
		   Write将s的内容写入缓冲中，如必要会增加缓冲容量。返回值n为len(p)，err总是nil。如果缓冲变得太大，Write会采用错误值ErrTooLarge引发panic。
		*/
		str := fmt.Sprintf("tree.insertNewChild('%d',%d,'%s',0,0,0,0,'SELECT,CALL,TOP,CHILD,CHECKED');", m.Pid, m.Id, m.Name)
		buffer.WriteString(str)
	}
	c.Data["tree"] = buffer.String()
	c.Data["Cates"] = cates
	c.Data["langType"] = langType
	c.TplName = "doc_category.html"
}
