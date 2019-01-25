package api

import (
	"fmt"
	"platformApi/src/model"
	"platformApi/src/service"
	"strconv"

)

// token相关操作
type TokenController struct {
	BaseController
}
//func (this *TokenController) Prepare() {
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
// @Title queryTokenInfo
// @Description 查询token详情
// @Param	tokenID		path 	string	true		"token标识"
// @Success 200 {object} model.TokenResult
// @Failure 403 :钱包地址不能为空
// @router /queryTokenInfo/:tokenID [get]
func (o *TokenController) QueryTokenInfo() {
	tokenID := o.Ctx.Input.Param(":tokenID")
	var result model.TokenResult
	if tokenID == "" {
		result.Status = false
		//result.Msg = "token标识不能为空！"
		result.Msg = o.Tr("ERROR_TOKEN.tokenIDIsEmpty")
		o.Data["json"] = result
		o.ServeJSON()
	}
	service.QueryTokenInfo(tokenID,  o.Ctx.Input.Header("language"),func(ret model.TokenResult) {
		if ret.Status == true {
			signThemeResult := service.QuerySignInfoByToken(tokenID ,o.Ctx.Input.Header("language"))
			ret.SignTheme = signThemeResult.Data
		}
		o.Data["json"] = ret
		o.ServeJSON()
	})
}

/*
// @Title UpdateTokenIcon
// @Description 设置token的图标
// @Param	tokenID		formData string	true  "token标识或ID"
// @Param	iconUrl		formData string	true  "图标地址url"
// @Success 200 {object} model.JsonResult
// @Failure 403 token标识和图标不能为空
// @router /updateTokenIcon [post]
func (o *TokenController) UpdateTokenIcon() {
	var result model.JsonResult
	tokenID := o.GetString("tokenID", "")
	iconUrl := o.GetString("iconUrl", "")
	if tokenID == "" {
		result.Status = false
		result.Msg = "token标识或ID不能为空！"
		o.Data["json"] = result
		o.ServeJSON()
		return
	}
	if iconUrl == "" {
		result.Status = false
		result.Msg = "token图标地址不能为空！"
		o.Data["json"] = result
		o.ServeJSON()
		return
	}

	err := model.UpdateTokenIcon(tokenID, iconUrl)
	if err != nil {
		result.Status = false
		result.Msg = err.Error()
	} else {
		result.Status = true
		result.Data = tokenID
	}
	o.Data["json"] = result
	o.ServeJSON()
}*/

// @Title isFirstSetMasterAndManager
// @Description 是否首次设置master和manager
// @Param	tokenID		path 	string	true		"token标识"
// @Success 200 {object} model.JsonResult
// @Failure 403 :token标识不能为空
// @router /isFirstSetMasterAndManager/:tokenID [get]
func (o *TokenController) IsFirstSetMasterAndManager() {
	tokenID := o.Ctx.Input.Param(":tokenID")
	var result model.JsonResult
	if tokenID == "" {
		result.Status = false
		//result.Msg = "token标识不能为空！"
		result.Msg = o.Tr("ERROR_TOKEN.tokenIDIsEmpty")
		o.Data["json"] = result
		o.ServeJSON()
	}
	service.ManagerCount(tokenID, o.Ctx.Input.Header("language"),func(ret model.ManagerCountResult) {
		if ret.Status == true {
			count, _ := strconv.Atoi(ret.Count)
			if count > 0 {
				result.Status = false
				//result.Msg = "已设置！"
				result.Msg = o.Tr("ERROR_TOKEN.masterAndManagerSet")
			} else {
				result.Status = true
				//result.Msg = "未设置！"
				result.Msg = o.Tr("ERROR_TOKEN.masterAndManagerNotSet")
			}
		} else {
			result.Status = false
			//result.Msg = "获取manager个数失败！"
			result.Msg = o.Tr("ERROR_TOKEN.getTheNumberOfManagers")
		}
		o.Data["json"] = result
		o.ServeJSON()
	})
}

// @Title queryAllTokens
// @Description 获取所有发行的token
// @Success 200 {object} model.TokenItemListResult
// @router /queryAllTokens [get]
func (o *TokenController) QueryAllTokens() {
	var result model.TokenItemListResult
	items, err := model.QueryAllTokens()
	if err != nil {
		result.Status = false
		result.Msg = fmt.Sprintf(o.Tr("ERROR_TOKEN.queryError") + "，%s！", err.Error())
	} else {
		result.Status = true
		result.Data = items
	}
	o.Data["json"] = result
	o.ServeJSON()
}

// @Title 搜索token信息
// @Description 搜索token信息
// @Param	keyWord		formData 	string	false	"关键字"
// @Param	page		formData 	int	false		"第几页"
// @Param	pageSize	formData 	int	false		"每页多少条记录""
// @Success 200 {object} model.TokenListResult
// @router /queryTokenList [post]
func (o *TokenController) QueryTokenList() {
	keyWord := o.GetString("keyWord")
	page, _ := o.GetInt("page", 1)
	pageSize, _ := o.GetInt("pageSize", 100)
	resultMap := model.SearchTokens(pageSize, page, keyWord)

	resultMap["status"] = true
	if arr, ok := resultMap["data"].([]model.Token); ok {
		var tokens []model.TokenFlex
		for _, v := range arr {
			var token model.TokenFlex
			token.TokenID = v.TokenID
			token.Name = v.Name
			token.TokenSymbol = v.TokenSymbol
			token.Status = v.Status

			ret := service.QuerySignInfoByToken(v.TokenID,o.Ctx.Input.Header("language"))
			if ret.Status == true {
				managerCount := ret.Data.ManagerCount
				masterThreshold := ret.Data.MasterThreshold
				managerThreshold := ret.Data.ManagerThreshold

				token.ManagerCount = managerCount
				token.MasterThreshold = masterThreshold
				token.ManagerThreshold = managerThreshold
				if managerCount > 0 {
					token.HasMutliSign = true
				}
			}
			tokens = append(tokens, token)
		}
		resultMap["data"] = tokens
		o.Data["json"] = resultMap
		o.ServeJSON()
	} else {
		o.Data["json"] = resultMap
		o.ServeJSON()
	}
}
