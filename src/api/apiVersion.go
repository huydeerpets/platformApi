package api

import (
	"platformApi/src/common"
	"platformApi/src/model"
	"strconv"

)

// 有关应用版本的相关操作
type ApiVersionController struct {
	BaseController
}
//func (this *ChainController) ApiVersionController() {
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
// @Title Get
// @Description 获取某平台的版本信息
// @Param	platType		path 	string	true		"平台类型(1、Android;2、IOS)"
// @Success 200 {object} model.VersionResult
// @Failure 403 :platType不能为空
// @router /getVersionInfo/:platType [get]
func (o *ApiVersionController) Get() {
	platType := o.Ctx.Input.Param(":platType")
	var result model.VersionResult
	if platType != "" {
		_type, _ := strconv.ParseInt(platType, 10, 64)
		ob, err := model.GetCurrentVersion("wallet", _type)
		if err != nil {
			result.Status = false
			result.Msg = err.Error()
		} else {
			result.Status = true

			var version model.Version
			version.Id = ob.Id
			version.AppName = ob.AppName
			version.Version = ob.Version
			version.AppAddr = ob.AppAddr
			version.AppDesc = ob.AppDesc
			version.CreateTime = ob.CreateTime
			version.UpgradeType = ob.UpgradeType
			version.CreateTimeFmt = common.FormatFullTime(version.CreateTime)
			version.IsCurrent = ob.IsCurrent
			version.PlatType = ob.PlatType

			if version.AppAddr != "" {
				fileSize := common.GetFileSize(version.AppAddr)
				version.Size = common.FormatFileSize(fileSize)
			}

			result.Data = version
		}
	} else {
		result.Status = false
		//result.Msg = "平台类型不能为空！"
		result.Msg = o.Tr("ERROR_VERSION.platformTypeIsEmpty")
	}
	o.Data["json"] = result
	o.ServeJSON()
}
