package api

import (
	"platformApi/src/model"
	"sort"
	"strconv"
)

// API查询相关操作
type ApiQueryController struct {
	BaseController
}

//func (this *ApiQueryController) Prepare() {
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
// @Title API主题查询
// @Description API主题查询
// @Param	id		path 	string	false		"接口ID"
// @Success 200 {object} model.ApiThemeResponse
// @router /queryApiTheme/:id [get]
func (o *ApiQueryController) QueryApiTheme() {
	var result model.ApiThemeResponse
	id := o.Ctx.Input.Param(":id")
	if id == "undefined" {
		id = ""
	}
	data := model.QueryApiTheme(id, o.Ctx.Input.Header("language"))
	result.Status = true
	result.Data = data

	o.Data["json"] = result
	o.ServeJSON()
}

// @Title 查询某接口的详情
// @Description 查询某接口的详情
// @Param	id		path 	string	true		"接口ID"
// @Success 200 {object} model.ApiInfoResponse
// @router /queryApiInfo/:id [get]
func (o *ApiQueryController) QueryApiInfo() {
	var result model.ApiInfoResponse
	paramID := o.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(paramID, 10, 64)
	langType := "1"
	if o.Lang == "en-US" {
		langType = "2"
	}
	baseData, _ := model.QueryBaseDataByTypeExt("request_method,req_arg_type,data_type,comp_lang,bool_type", langType)

	cate := model.GetCategoryInfo(id)
	var mainInfo model.ApiMainInfo
	mainInfo.Id = cate.Id
	mainInfo.Name = cate.Name
	mainInfo.RequestUrl = cate.RequestUrl
	mainInfo.Method = o.getName("request_method", cate.Method, baseData)
	mainInfo.Remark = cate.Remark

	requestParams := model.GetApiRequestInfo(id)
	if len(requestParams) > 0 {
		for i := range requestParams {
			requestParams[i].ArgType = o.getName("req_arg_type", requestParams[i].ArgType, baseData)
			requestParams[i].DataType = o.getName("data_type", requestParams[i].DataType, baseData)
			requestParams[i].IsRequire = o.getName("bool_type", requestParams[i].IsRequire, baseData)
		}
	}
	mainInfo.RequestParams = requestParams

	fieldData := model.GetApiRequestFieldInfo(id)
	if len(fieldData) > 0 {
		for i := range fieldData {
			for j := range fieldData[i].Items {
				fieldData[i].Items[j].DataType = o.getName("data_type", fieldData[i].Items[j].DataType, baseData)
				fieldData[i].Items[j].IsRequire = o.getName("bool_type", fieldData[i].Items[j].IsRequire, baseData)
			}
		}
	}
	mainInfo.RequestParamsMemos = fieldData

	formDatas := model.GetApiFormDataInfo(id)
	if len(formDatas) > 0 {
		for i := range formDatas {
			formDatas[i].DataType = o.getName("data_type", formDatas[i].DataType, baseData)
			formDatas[i].IsRequire = o.getName("bool_type", formDatas[i].IsRequire, baseData)
		}
	}
	mainInfo.FormParams = formDatas

	requestSamples := model.GetRequestSampleInfo(id)
	mainInfo.RequestSamples = requestSamples

	requestCodes := model.GetRequestCodeInfo(id)
	if len(requestCodes) > 0 {
		for i := range requestCodes {
			requestCodes[i].LangType = o.getName("comp_lang", requestCodes[i].LangType, baseData)
		}
	}
	mainInfo.RequestCodes = requestCodes

	respParams := model.GetRespParamInfo(id)
	var respParamResults model.RespParamResults
	if len(respParams) > 0 {
		for r := range respParams {
			respParams[r].DataType = o.getName("data_type", respParams[r].DataType, baseData)
		}
		paramsMap := o.convertData(respParams)

		for id, items := range paramsMap {
			if id == 0 {
				var rt model.RespParamResult
				//rt.Title = "响应结果说明"
				rt.Title = o.Tr("QUERY.responseResultDescription")
				rt.RespParams = items
				respParamResults = append(respParamResults, rt)
			} else {
				parentItem := o.findParamItem(id, respParams)
				var rt model.RespParamResult
				//rt.Title = "【" + parentItem.DataName + "】字段说明"
				rt.Title = "【" + parentItem.DataName + "】" + o.Tr("QUERY.fieldDescription")
				rt.RespParams = items
				respParamResults = append(respParamResults, rt)
			}
		}
	}
	if len(respParamResults) > 0 {
		sort.Sort(respParamResults)
	}
	mainInfo.RespParamResults = respParamResults

	respSamples := model.GetRespSampleInfo(id)
	if len(respSamples) > 0 {
		for i := range respSamples {
			if respSamples[i].RespType == "1" {
				//respSamples[i].RespType = "成功实例"
				respSamples[i].RespType = o.Tr("QUERY.successfulInstance")
			} else if respSamples[i].RespType == "2" {
				//respSamples[i].RespType = "失败实例"
				respSamples[i].RespType = o.Tr("QUERY.failedInstance")
			}
		}
	}
	mainInfo.RespSamples = respSamples

	result.Status = true
	result.Data = mainInfo

	o.Data["json"] = result
	o.ServeJSON()
}

func (o *ApiQueryController) getName(bType, bVal string, baseData map[string]([]map[string]string)) string {
	for key := range baseData {
		if key == bType {
			arr := baseData[key]
			for _, v := range arr {
				if v["data_code"] == bVal {
					return v["data_name"]
				}
			}
		}
	}
	return ""
}

func (o *ApiQueryController) convertData(respParams []model.ResponseParam) map[int64]([]model.ResponseParam) {
	paramsMap := make(map[int64]([]model.ResponseParam), 0)
	for _, v := range respParams {
		paramsMap[v.ParentId] = append(paramsMap[v.ParentId], v)
	}
	return paramsMap
}

func (o *ApiQueryController) findParamItem(id int64, respParams []model.ResponseParam) *model.ResponseParam {
	for _, v := range respParams {
		if v.Id == id {
			return &v
		}
	}
	return nil
}
