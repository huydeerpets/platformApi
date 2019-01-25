package controller

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"platformApi/src/model"
	"regexp"

	"strconv"

	"github.com/astaxie/beego"
)

type ApiCategoryController struct {
	beego.Controller
}

func (c *ApiCategoryController) InitCategory() {
	paramID := c.Ctx.Input.Params()["0"]
	subCount := c.Ctx.Input.Params()["1"]
	nodeLevel := c.Ctx.Input.Params()["2"]
	langType := c.Ctx.Input.Params()["3"]
	id, _ := strconv.ParseInt(paramID, 10, 64)
	cate := model.GetCategoryInfo(id)
	c.Data["cate"] = cate
	c.Data["subCount"] = subCount
	c.Data["nodeLevel"] = nodeLevel
	c.Data["langType"] = langType
	if nodeLevel == "3" {
		c.TplName = "addapiinfo.html"
	} else {
		c.TplName = "addapicategory.html"
	}
}

func (c *ApiCategoryController) InitApiInfo() {
	paramID := c.Ctx.Input.Params()["0"]
	nodeLevel := c.Ctx.Input.Params()["1"]
	langType := c.Ctx.Input.Params()["2"]
	id, _ := strconv.ParseInt(paramID, 10, 64)
	cate := model.GetCategoryInfo(id)

	pat := "<br>" //正则
	re, _ := regexp.Compile(pat)
	//将匹配到的部分替换为"##.#"
	cate.Remark = re.ReplaceAllString(cate.Remark, "\n")

	c.Data["cate"] = cate
	c.Data["subCount"] = 0
	c.Data["nodeLevel"] = nodeLevel
	c.Data["langType"] = langType

	c.TplName = "addapiinfo.html"
}

func (c *ApiCategoryController) GetApiRequestInfo() {
	paramID := c.Ctx.Input.Params()["0"]
	id, _ := strconv.ParseInt(paramID, 10, 64)
	requestParams := model.GetApiRequestInfo(id)

	data := make(map[string]interface{})
	data["requestParams"] = requestParams

	formDatas := model.GetApiFormDataInfo(id)
	data["formDatas"] = formDatas

	requestSamples := model.GetRequestSampleInfo(id)
	data["requestSamples"] = requestSamples

	requestCodes := model.GetRequestCodeInfo(id)
	data["requestCodes"] = requestCodes

	respSamples := model.GetRespSampleInfo(id)
	data["respSamples"] = respSamples

	respItems := model.GetRespParamInfo(id)
	data["respItems"] = respItems

	fieldData := model.GetApiRequestFieldInfo(id)
	data["fieldData"] = fieldData

	var result model.Result
	result.Status = 0
	result.Data = data

	c.Data["json"] = result

	c.ServeJSON()
}

func (c *ApiCategoryController) GetCateList() {
	cates, err := model.GetCateList()
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
func (c *ApiCategoryController) GetCateDropList() {
	cates := model.GetCategoryList("1", true)

	arr := make([]model.DropItem, 0)

	arr = model.GetList(cates, 1, arr, 0)

	result := new(model.Result)
	result.Status = 0
	result.Data = arr
	c.Data["json"] = result
	c.ServeJSON()
}

func (c *ApiCategoryController) GetCateInfo() {
	paramID := c.Ctx.Input.Params()["0"]
	id, _ := strconv.ParseInt(paramID, 10, 64)
	cate := model.GetCategoryInfo(id)

	//b, _ := json.Marshal(cate)
	//fmt.Printf("***********%#v", string(b))

	result := new(model.Result)
	result.Status = 0
	result.Data = cate
	c.Data["json"] = result
	c.ServeJSON()
}

func (c *ApiCategoryController) SaveCateInfo() {
	pid, _ := strconv.ParseInt(c.GetString("pid"), 10, 64)
	full_path := c.GetString("full_path")
	depth, _ := strconv.ParseInt(c.GetString("depth"), 10, 64)
	name := c.GetString("name")
	subCount, _ := strconv.ParseInt(c.GetString("subCount"), 10, 64)

	newPath := fmt.Sprintf("%s,%d", full_path, pid)

	var cate model.ApiCategory
	cate.Pid = pid
	cate.Name = name
	cate.Depth = depth + 1
	cate.FullPath = newPath
	cate.OrderNum = subCount

	v, err := model.SaveCategroy(&cate)

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

func (c *ApiCategoryController) SaveApiInfo() {
	hexStr := c.GetString("rawData")
	bytes, err := hex.DecodeString(hexStr)
	result := new(model.Result)
	if err != nil {
		result.Status = 1
		result.Msg = "解析16进制格式数据出错!"
		c.Data["json"] = result
		c.ServeJSON()
		return
	}
	var request model.SaveApiRequest
	err = json.Unmarshal(bytes, &request)
	if err != nil {
		result.Status = 1
		result.Msg = fmt.Sprintf("json解析出错,信息为%s!", err.Error())
		c.Data["json"] = result
		c.ServeJSON()
		return
	}

	pid, _ := strconv.ParseInt(request.Pid, 10, 64)
	fullPath := request.FullPath
	depth, _ := strconv.ParseInt(request.Depth, 10, 64)
	name := request.Name
	subCount, _ := strconv.ParseInt(request.SubCount, 10, 64)

	newPath := fmt.Sprintf("%s,%d", fullPath, pid)

	var cate model.ApiCategory
	cate.Pid = pid
	cate.Name = name
	cate.Depth = depth + 1
	cate.FullPath = newPath
	cate.OrderNum = subCount
	cate.Method = request.Method
	cate.RequestUrl = request.RequestUrl
	cate.Remark = request.Remark

	cateID, err := model.SaveCategroy(&cate)

	if len(request.Params) > 0 {
		for i := range request.Params {
			request.Params[i].CateId = cateID
			request.Params[i].Status = 1
		}
	}

	if len(request.FormDatas) > 0 {
		for i := range request.FormDatas {
			request.FormDatas[i].CateId = cateID
			request.FormDatas[i].Status = 1
		}
	}

	if len(request.ReqSamples) > 0 {
		for i := range request.ReqSamples {
			request.ReqSamples[i].CateId = cateID
		}
	}

	if len(request.ReqCodes) > 0 {
		for i := range request.ReqCodes {
			request.ReqCodes[i].CateId = cateID
		}
	}

	if len(request.RespSamples) > 0 {
		for i := range request.RespSamples {
			request.RespSamples[i].CateId = cateID
		}
	}

	if len(request.RespItems) > 0 {
		for i := range request.RespItems {
			request.RespItems[i].CateId = cateID
		}
	}

	if len(request.FieldData) > 0 {
		for i := range request.FieldData {
			request.FieldData[i].ReqFiled.CateId = cateID
		}
	}

	_, err = model.SaveApiInfo(request)
	if err == nil {
		result.Status = 0
		result.Data = cateID
	} else {
		result.Status = 1
		result.Msg = fmt.Sprintf("保存信息失败,%s", err.Error())
	}
	c.Data["json"] = result
	c.ServeJSON()
}

//更新api信息
func (c *ApiCategoryController) UpdateApiInfo() {
	hexStr := c.GetString("rawData")
	bytes, err := hex.DecodeString(hexStr)
	result := new(model.Result)
	if err != nil {
		result.Status = 1
		result.Msg = "解析16进制格式数据出错!"
		c.Data["json"] = result
		c.ServeJSON()
		return
	}
	var request model.SaveApiRequest
	err = json.Unmarshal(bytes, &request)
	if err != nil {
		result.Status = 1
		result.Msg = fmt.Sprintf("json解析出错,信息为%s!", err.Error())
		c.Data["json"] = result
		c.ServeJSON()
		return
	}

	err = model.UpdateApiInfo(request)
	if err == nil {
		result.Status = 0
		result.Data = "修改成功"
	} else {
		result.Status = 1
		result.Msg = fmt.Sprintf("保存信息失败,%s", err.Error())
	}
	c.Data["json"] = result
	c.ServeJSON()
}

func (c *ApiCategoryController) UpdateCateInfo() {
	paramID := c.GetString("id")
	id, _ := strconv.ParseInt(paramID, 10, 64)
	name := c.GetString("name")
	orderNum, _ := strconv.ParseInt(c.GetString("orderNum"), 10, 64)
	err := model.UpdateCategroy(id, orderNum, name)

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

func (c *ApiCategoryController) DelCateInfo() {
	id, _ := strconv.ParseInt(c.GetString("id"), 10, 64)
	pid, _ := strconv.ParseInt(c.GetString("pid"), 10, 64)
	orderNum, _ := strconv.ParseInt(c.GetString("orderNum"), 10, 64)
	nodeLevel, _ := strconv.ParseInt(c.GetString("nodeLevel"), 10, 64)

	_, err := model.DeleteCategroy(id, pid, orderNum, nodeLevel)

	result := new(model.Result)
	if err == nil {
		result.Status = 0
	} else {
		result.Status = 1
		result.Msg = "删除分类信息失败！"
	}
	c.Data["json"] = result
	c.ServeJSON()
}

//复制api信息
func (c *ApiCategoryController) CopyCateInfo() {
	id, _ := strconv.ParseInt(c.GetString("id"), 10, 64)
	pid, _ := strconv.ParseInt(c.GetString("pid"), 10, 64)
	//orderNum, _ := strconv.ParseInt(c.GetString("orderNum"), 10, 64)
	//nodeLevel, _ := strconv.ParseInt(c.GetString("nodeLevel"), 10, 64)

	newId, err := model.CopyCategroy(id, pid)

	result := new(model.Result)
	if err == nil {
		result.Status = 0
		result.Data = newId
	} else {
		result.Status = 1
		result.Msg = "复制失败！"
	}
	c.Data["json"] = result
	c.ServeJSON()
}

//获取分类信息
func (c *ApiCategoryController) List() {
	langType := c.GetString("langType")
	if langType == "" {
		langType = "1"
	}
	cates := model.GetCategoryList(langType, false)

	var buffer bytes.Buffer //Buffer是一个实现了读写方法的可变大小的字节缓冲
	for _, m := range cates {
		/*
		   func (b *Buffer) WriteString(s string) (n int, err error)
		   Write将s的内容写入缓冲中，如必要会增加缓冲容量。返回值n为len(p)，err总是nil。如果缓冲变得太大，Write会采用错误值ErrTooLarge引发panic。
		*/
		str := fmt.Sprintf(`tree.insertNewChild('%d',%d,"%s",0,0,0,0,'SELECT,CALL,TOP,CHILD,CHECKED');`, m.Pid, m.Id, m.Name)
		buffer.WriteString(str)
	}
	c.Data["tree"] = buffer.String()
	c.Data["Cates"] = cates
	c.Data["langType"] = langType
	c.TplName = "apicategory.html"
}

//登录功能
func (c *ApiCategoryController) Post() {

	c.Data["json"] = map[string]interface{}{"islogin": 1}
	c.ServeJSON()
}
