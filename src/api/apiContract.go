package api

import (
	"platformApi/src/common"
	"platformApi/src/model"
	"platformApi/src/service"

)

// 合约相关操作
type ContractController struct {
	BaseController
}
//func (this *ContractController) Prepare() {
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
// @Title queryContractList
// @Description 查询某个地址下的合约信息列表
// @Param	address		path 	string	true		"钱包地址或合约地址"
// @Success 200 {object} model.ContractListResult
// @Failure 403 :钱包地址和合约地址必须二选其一
// @router /queryContractList/:address [get]
func (o *ContractController) QueryContractList() {
	address := o.Ctx.Input.Param(":address")
	var result model.ContractListResult
	if address == "" {
		result.Status = false
		//result.Msg = "钱包地址或合约地址不能为空！"
		result.Msg = o.Tr("ERROR_CONTRACT.walletAddressOrContractAddressIsEmpty")
		o.Data["json"] = result
		o.ServeJSON()
	}
	service.QueryContractList(address, o.Ctx.Input.Header("language"), func(ret model.ContractListResult) {
		if len(ret.Data) > 0 {
			for i := range ret.Data {
				ret.Data[i].CreateTime = common.FormatDate(ret.Data[i].CreateTime)
				ret.Data[i].UpdateTime = common.FormatDate(ret.Data[i].UpdateTime)
			}
		}
		o.Data["json"] = ret
		o.ServeJSON()
	})
}

// @Title queryContractInfo
// @Description 查询合约的详情
// @Param	address		path 	string	true		"合约地址"
// @Param	version		path 	string	false		"合约版本号"
// @Success 200 {object} model.ContractResult
// @Failure 403 :合约地址不能为空
// @router /queryContractInfo/:address/:version [get]
func (o *ContractController) QueryContractInfo() {
	address := o.Ctx.Input.Param(":address")
	version := o.Ctx.Input.Param(":version")
	var result model.ContractResult
	if address == "" {
		//result.Msg = "合约地址不能为空！"
		result.Msg = o.Tr("ERROR_CONTRACT.contractAddressIsEmpty")
		o.Data["json"] = result
		o.ServeJSON()
	}
	if version == "" {
		//result.Msg = "合约版本号不能为空！"
		result.Msg = o.Tr("ERROR_CONTRACT.contractVersionNumberIsEmpty")
		o.Data["json"] = result
		o.ServeJSON()
	}
	service.QueryContractInfo(address, o.Ctx.Input.Header("language"), version, func(ret model.ContractResult) {
		o.Data["json"] = ret
		o.ServeJSON()
	})
}
