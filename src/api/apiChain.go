package api

import (
	"encoding/hex"
	"encoding/json"
	"platformApi/src/common"
	"platformApi/src/model"
	"platformApi/src/service"

)

// 链相关操作
type ChainController struct {
	BaseController
}

//func (this *ChainController) Prepare() {
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
// @Title 链申请
// @Description 链申请
// @Param	name		    formData string	true  "链名称"
// @Param	en_short		formData string	true  "链英文简写名称"
// @Param	remark		    formData string	true  "链简介"
// @Param	contact_name	formData string	true  "联系人姓名"
// @Param	contact_tel	formData string	true  "联系人电话"
// @Param	e_mail			formData string	true  "联系人邮箱"
// @Success 200 {object} model.JsonResult
// @router /chainApply [post]
func (o *ChainController) ChainApply() {
	name := o.GetString("name", "")
	en_short := o.GetString("en_short", "")
	remark := o.GetString("remark", "")
	contact_name := o.GetString("contact_name", "")
	contact_tel := o.GetString("contact_tel", "")
	e_mail := o.GetString("e_mail", "")

	obj := make(map[string]string)
	obj["name"] = name
	obj["en_short"] = en_short
	obj["remark"] = remark
	obj["contact_name"] = contact_name
	obj["contact_tel"] = contact_tel
	obj["e_mail"] = e_mail

	bytes, _ := json.Marshal(obj)
	hexStr := hex.EncodeToString(bytes)

	service.ChainApply(hexStr, o.Ctx.Input.Header("language"), func(result model.JsonResult) {
		o.Data["json"] = result
		o.ServeJSON()
	})
}

// @Title 链索索
// @Description 链索索
// @Param	keyWord		    formData string	false  "搜索关键字"
// @Param	page		formData 	int	false		"第几页"
// @Param	pageSize	formData 	int	false		"每页多少条记录""
// @Success 200 {object} model.ChainApplyListResult
// @Failure 403 :搜索关键字不能为空
// @router /chainSearch [post]
func (o *ChainController) ChainSearch() {
	keyWord := o.GetString("keyWord", "")
	page, _ := o.GetInt("page", 1)
	pageSize, _ := o.GetInt("pageSize", 20)
	resultMap := model.SearchChainApplys(pageSize, page, keyWord)
	if resultMap["data"] != nil {
		if infos, ok := resultMap["data"].([]model.ChainApplyInfo); ok {
			for i := range infos {
				infos[i].EMail = common.HideEmail(infos[i].EMail)
				infos[i].ContactTel = common.HideTel(infos[i].ContactTel)
			}
			resultMap["data"] = infos
		}
	}
	resultMap["status"] = true
	o.Data["json"] = resultMap
	o.ServeJSON()
}

// @Title 链申请信息详情
// @Description 链申请信息详情
// @Param	id		path 	string	true		"主键ID"
// @Success 200 {object} model.ChainApplyResult
// @Failure 403 :主键ID不能为空
// @router /chainApplyInfo/:id [get]
func (o *ChainController) ChainApplyInfo() {
	id := o.Ctx.Input.Param(":id")
	var result model.ChainApplyResult
	if id == "" {
		result.Status = false
		//result.Msg = "主键ID不能为空！"
		result.Msg = o.Tr("ERROR_CHAIN.primaryKeyIDIsEmpty")
		o.Data["json"] = result
		o.ServeJSON()
	}
	service.ChainApplyInfo(id, o.Ctx.Input.Header("language"), func(ret model.ChainApplyListResult) {
		result.Status = ret.Status
		if len(ret.Data) > 0 {
			result.Data = &ret.Data[0]
			result.Data.EMail = common.HideEmail(result.Data.EMail)
			result.Data.ContactTel = common.HideTel(result.Data.ContactTel)
		} else {
			result.Status = false
			//result.Msg = "未找到相应的记录!"
			result.Msg = o.Tr("ERROR_CHAIN.noRecord")
		}
		o.Data["json"] = result
		o.ServeJSON()
	})
}

// @Title 发行token所需的平台币标准
// @Description 发行token所需的平台币标准
// @Success 200 {object} model.JsonResult
// @router /queryPublishTokenRequireNum [get]
func (o *ChainController) QueryPublishTokenRequireNum() {
	service.QueryTokenRequireNum(o.Ctx.Input.Header("language"), func(ret model.JsonResult) {
		o.Data["json"] = ret
		o.ServeJSON()
	})
}

// @Title 发布合约所需的平台币标准
// @Description 发布合约所需的平台币标准
// @Success 200 {object} model.JsonResult
// @router /queryPublishCCRequireNum [get]
func (o *ChainController) QueryPublishCCRequireNum() {
	service.QueryCCRequireNum(o.Ctx.Input.Header("language"), func(ret model.JsonResult) {
		o.Data["json"] = ret
		o.ServeJSON()
	})
}

// @Title 手续费返还规则
// @Description 手续费返还规则
// @Success 200 {object} model.ReturnGasConfigResponse
// @router /queryReturnGasConfig [get]
func (o *ChainController) QueryReturnGasConfig() {
	service.QueryReturnGasConfig(o.Ctx.Input.Header("language"), func(ret model.ReturnGasConfigResponse) {
		o.Data["json"] = ret
		o.ServeJSON()
	})
}
