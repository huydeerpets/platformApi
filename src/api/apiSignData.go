package api

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"platformApi/src/common"
	"platformApi/src/model"
	"platformApi/src/service"
	"strconv"
	"time"
)

// 签名相关操作
type ApiSignController struct {
	BaseController
}

//func (this *ApiSignController) Prepare() {
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
// @Title Create
// @Description 生成二维码的唯一标识
// @Param	signType		formData int	true  "签名类型 1、发布合约 2、实例化合约 3、升级合约 4、发行token 5、设置master、manager 6、manager签名确认 7、添加manager　8、替换manager 9 、删除manager 10、设置manager操作确认的阀值 11、设置发行token所需要的平台币 12、设置发布合约所需要的平台币 13、设置手续费返还规则 14、删除合约 15、设置master操作所需的确认数阀值 16、设置token图标"
// @Param	originData		formData string	true  "需要签名的原数据，上传附件时，直接传递到OSS上，将返回的文件地址与其它信息合并生成原始数据"
// @Success 200 {object} model.JsonResult
// @Failure 403 只需要传递签名类型、待签名的原始数据即可
// @router /addSignData [post]
func (o *ApiSignController) Post() {
	var result model.JsonResult
	signType, _ := o.GetInt("signType", -1)
	originData := o.GetString("originData", "")
	if signType == -1 {
		result.Status = false
		//result.Msg = "签名类型不能为空！"
		result.Msg = o.Tr("ERROR_SIGN.signatureTypeIsEmpty")
		o.Data["json"] = result
		o.ServeJSON()
		return
	}
	if originData == "" {
		result.Status = false
		//result.Msg = "待签名的原始数据不能为空！"
		result.Msg = o.Tr("ERROR_SIGN.originalDataIsEmpty")
		o.Data["json"] = result
		o.ServeJSON()
		return
	}
	var sign model.SignData
	sign.QrCode = common.PLATSIGN + common.GetUUID()
	sign.SignType = signType
	sign.Status = 1
	timeUnix := time.Now().Unix() + 600
	sign.ValidTime = timeUnix

	hexBytes, _ := hex.DecodeString(originData)
	originMap := make(map[string]interface{})
	err := json.Unmarshal(hexBytes, &originMap)
	if err != nil {
		result.Status = false
		//result.Msg = "16进制数据转json出错！"
		result.Msg = o.Tr("ERROR_SIGN.jsonDeserialization")
		o.Data["json"] = result
		o.ServeJSON()
		return
	}
	originMap["funcName"] = common.MultiSignMethod[signType]
	if sign.SignType == 4 {
		originMap["enableNumber"] = fmt.Sprintf("%v", originMap["totalNumber"])
	} else if sign.SignType == 10 || sign.SignType == 15 {
		num, _ := strconv.Atoi(fmt.Sprintf("%v", originMap["threshold"]))
		originMap["threshold"] = num
	}
	//主网参数
	if sign.SignType == 11 || sign.SignType == 12 || sign.SignType == 13 {
		token := service.QueryMasterTokenInfo(o.Ctx.Input.Header("language"))
		if token == nil {
			result.Status = false
			//result.Msg = "获取主币信息出错！"
			result.Msg = o.Tr("ERROR_SIGN.getTheMainCurrencyInformation")
			o.Data["json"] = result
			o.ServeJSON()
			return
		}
		originMap["tokenID"] = token.TokenID
		if sign.SignType == 11 || sign.SignType == 12 {
			originMap["number"] = fmt.Sprintf("%v", originMap["number"])
		}
	}
	jsonBytes, _ := json.Marshal(originMap)
	sign.OriginData = hex.EncodeToString(jsonBytes)
	sign.LangType = o.Lang

	err = model.AddSignData(&sign)
	if err != nil {
		result.Status = false
		result.Msg = err.Error()
	} else {
		result.Status = true
		result.Data = sign.QrCode
	}
	o.Data["json"] = result
	o.ServeJSON()
}

// @Title querySignInfo
// @Description 获取签名信息
// @Param	qrCode		path 	string	true		"二维码标识 9f9acb8f749b4f32b2e9dcab30601560"
// @Success 200 {object} model.SignInfoResult
// @Failure 403 :二维码标识不能为空
// @router /querySignInfo/:qrCode [get]
func (o *ApiSignController) QuerySignInfo() {
	qrCode := o.Ctx.Input.Param(":qrCode")
	var result model.SignInfoResult
	if qrCode != "" {
		sign, err := model.GetSingInfo(qrCode)
		if err != nil || sign == nil {
			result.Status = false
			//result.Msg = "未找到签名信息！"
			result.Msg = o.Tr("ERROR_SIGN.signatureInformationNotFound")
		} else {
			result.Status = true
			result.Data = sign
		}
	} else {
		result.Status = false
		//result.Msg = "二维码标识不能为空！"
		result.Msg = o.Tr("ERROR_SIGN.QRCodeIDIsEmpty")
	}
	o.Data["json"] = result
	o.ServeJSON()
}

// @Title Query
// @Description 查询签名是否成功
// @Param	qrCode		path 	string	true		"二维码标识"
// @Success 200 {object} model.CCInvokeResult
// @Failure 403 :二维码标识不能为空
// @router /getSignStatus/:qrCode [get]
func (o *ApiSignController) GetSignStatus() {
	qrCode := o.Ctx.Input.Param(":qrCode")
	var result model.CCInvokeResult
	result.Status = false
	if qrCode != "" {
		sign, err := model.GetSingInfo(qrCode)
		if err != nil || sign == nil {
			//result.Msg = "未找到签名确认信息！"
			result.Msg = o.Tr("ERROR_SIGN.signatureConfirmationNotFound")
		} else {
			if sign.Status == 3 {
				json.Unmarshal([]byte(sign.RespResult), &result)
			} else {
				//result.Msg = "请重新扫描二维码进行确认操作！"
				result.Msg = o.Tr("ERROR_SIGN.QRCodeScanning")
			}
		}
	} else {
		//result.Msg = "二维码标识不能为空！"
		result.Msg = o.Tr("ERROR_SIGN.QRCodeIDIsEmpty")
	}
	o.Data["json"] = result
	o.ServeJSON()
}
