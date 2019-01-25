package service

import (
	"encoding/json"
	"fmt"
	"net/url"
	"platformApi/src/common"
	"platformApi/src/model"

	"github.com/astaxie/beego"
)

type Result struct {
	Status string `json:"status"`
	Msg    string `json:"msg"`
	Data   string `json:"data"`
}

//根据地址查询合约列表
func QueryContractList(address string, lang string, callback func(model.ContractListResult)) {
	chainurl := beego.AppConfig.String("chainurl") + "/udo_cc_query"
	data := url.Values{"address": {address}}
	common.UDOHttpPost(chainurl, lang, data, func(result interface{}, err error) {
		if err == nil && result != nil {
			if data, ok := result.([]byte); ok {
				var jsonResult model.ContractListResult
				err := json.Unmarshal(data, &jsonResult)
				if err != nil {
					jsonResult.Status = false
					jsonResult.Msg = fmt.Sprintf("json解析出错,%s！", string(data))
				}
				callback(jsonResult)
			}
		}
	})
}

//根据合约地址获取合约信息
func QueryContractInfo(address, lang, version string, callback func(model.ContractResult)) {
	chainurl := fmt.Sprintf(beego.AppConfig.String("chainurl")+"/udo_cc_info/%s/%s", address, version)
	common.UDOHttpGet(chainurl, lang, func(result interface{}, err error) {
		if err == nil && result != nil {
			if data, ok := result.([]byte); ok {
				var jsonResult model.ContractResult
				err := json.Unmarshal(data, &jsonResult)
				if err != nil {
					jsonResult.Status = false
					jsonResult.Msg = fmt.Sprintf("json解析出错,%s！", string(data))
				}
				callback(jsonResult)
			}
		}
	})
}

//根据tokenID获取token信息
func QueryTokenInfo(tokenID string, lang string, callback func(model.TokenResult)) {
	chainurl := fmt.Sprintf(beego.AppConfig.String("chainurl")+"/queryTokenInfo/%s", tokenID)
	common.UDOHttpGet(chainurl, lang, func(result interface{}, err error) {
		if err == nil && result != nil {
			if data, ok := result.([]byte); ok {
				var jsonResult model.TokenResult
				err := json.Unmarshal(data, &jsonResult)
				if err != nil {
					jsonResult.Status = false
					jsonResult.Msg = fmt.Sprintf("json解析出错,%s！", string(data))
				}
				if jsonResult.Status == true {
					token, err := model.QueryTokenInfo(tokenID)
					if err != nil {
						jsonResult.Status = false
						jsonResult.Msg = "未找到相应的token信息！"
					} else {
						jsonResult.Data.IconUrl = token.IconUrl
						bigTotal := common.HexToBigInt(jsonResult.Data.TotalNumber)
						bigRest := common.HexToBigInt(jsonResult.Data.RestNumber)
						digits := common.GetNumberWithDigits(jsonResult.Data.DecimalUnits)

						jsonResult.Data.TotalNumber = common.BigIntDiv(bigTotal, digits)
						jsonResult.Data.RestNumber = common.BigIntDiv(bigRest, digits)
					}
				}
				callback(jsonResult)
			}
		}
	})
}

//获取某个token下的manager
func QueryManagerList(tokenID string, lang string, callback func(model.CommonResult)) {
	chainurl := fmt.Sprintf(beego.AppConfig.String("chainurl")+"/queryManagerList/%s", tokenID)
	common.UDOHttpGet(chainurl, lang, func(result interface{}, err error) {
		if err == nil && result != nil {
			if data, ok := result.([]byte); ok {
				var jsonResult model.CommonResult
				err := json.Unmarshal(data, &jsonResult)
				if err != nil {
					jsonResult.Status = false
					jsonResult.Msg = fmt.Sprintf("json解析出错,%s！", string(data))
				}
				callback(jsonResult)
			}
		}
	})
}

//获取某个token下的多重签名信息
func QuerySignInfoByToken(tokenID string,lang string) model.SignThemeResult {
	chainurl := fmt.Sprintf(beego.AppConfig.String("chainurl")+"/querySignInfoByToken/%s", tokenID)
	bytes, _ := common.UDOSyncHttpGet(chainurl,lang)
	var jsonResult model.SignThemeResult
	err := json.Unmarshal(bytes, &jsonResult)
	if err != nil {
		jsonResult.Status = false
		jsonResult.Msg = fmt.Sprintf("json解析出错,%s！", string(bytes))
	}
	return jsonResult
}

//根据tokenID获取manager的个数
func ManagerCount(tokenID string, lang string, callback func(model.ManagerCountResult)) {
	chainurl := fmt.Sprintf(beego.AppConfig.String("chainurl")+"/managersCount/%s", tokenID)
	common.UDOHttpGet(chainurl, lang, func(result interface{}, err error) {
		if err == nil && result != nil {
			if data, ok := result.([]byte); ok {
				var jsonResult model.ManagerCountResult
				err := json.Unmarshal(data, &jsonResult)
				if err != nil {
					jsonResult.Status = false
					jsonResult.Msg = fmt.Sprintf("json解析出错,%s！", string(data))
				}
				callback(jsonResult)
			}
		}
	})
}

//链申请
func ChainApply(hexStr string, lang string, callback func(model.JsonResult)) {
	chainurl := beego.AppConfig.String("chainurl") + "/udo_chainEnter"
	data := url.Values{"rawData": {hexStr}}
	common.UDOHttpPost(chainurl, lang, data, func(result interface{}, err error) {

		if err == nil && result != nil {
			if data, ok := result.([]byte); ok {
				var jsonResult model.JsonResult
				type Result struct {
					Status bool
					Msg    string
					Data   int64
				}
				var ret Result
				err := json.Unmarshal(data, &ret)
				if err != nil {
					jsonResult.Status = false
					jsonResult.Msg = fmt.Sprintf("json解析出错,%s！", string(data))
				} else {
					jsonResult.Status = ret.Status
					jsonResult.Msg = ret.Msg
					jsonResult.Data = fmt.Sprintf("%d", ret.Data)
				}

				callback(jsonResult)
			}
		}
	})
}

//链搜索
func ChainSearch(keyWord string, lang string, callback func(model.ChainApplyListResult)) {
	chainurl := beego.AppConfig.String("chainurl") + "/udo_chainEnterSearch"
	data := url.Values{"keyWord": {keyWord}}
	common.UDOHttpPost(chainurl, lang, data, func(result interface{}, err error) {

		if err == nil && result != nil {
			if data, ok := result.([]byte); ok {
				var jsonResult model.ChainApplyListResult
				err := json.Unmarshal(data, &jsonResult)
				if err != nil {
					jsonResult.Status = false
					jsonResult.Msg = fmt.Sprintf("json解析出错,%s！", string(data))
				}
				if jsonResult.Status == true {
					for i := range jsonResult.Data {
						jsonResult.Data[i].CreateTimeFmt = common.FormatTime(jsonResult.Data[i].CreateTime)
					}
				}
				callback(jsonResult)
			}
		}
	})
}

//获取链申请信息
func ChainApplyInfo(id string, lang string, callback func(model.ChainApplyListResult)) {
	chainurl := fmt.Sprintf(beego.AppConfig.String("chainurl")+"/udo_chainEnterInfo/%s", id)
	common.UDOHttpGet(chainurl, lang, func(result interface{}, err error) {
		if err == nil && result != nil {
			if data, ok := result.([]byte); ok {
				var jsonResult model.ChainApplyListResult
				err := json.Unmarshal(data, &jsonResult)
				if err != nil {
					jsonResult.Status = false
					jsonResult.Msg = fmt.Sprintf("json解析出错,%s！", string(data))
				}
				if jsonResult.Status == true {
					for i := range jsonResult.Data {
						jsonResult.Data[i].CreateTimeFmt = common.FormatTime(jsonResult.Data[i].CreateTime)
					}
				}
				callback(jsonResult)
			}
		}
	})
}

//获取发行token所需要的平台币标准
func QueryTokenRequireNum(lang string, callback func(model.JsonResult)) {
	chainurl := beego.AppConfig.String("chainurl") + "/udo_queryPublishTokenRequireNum"
	common.UDOHttpGet(chainurl, lang, func(result interface{}, err error) {
		if err == nil && result != nil {
			if data, ok := result.([]byte); ok {
				var jsonResult model.JsonResult
				err := json.Unmarshal(data, &jsonResult)
				if err != nil {
					jsonResult.Status = false
					jsonResult.Msg = fmt.Sprintf("json解析出错,%s！", string(data))
				}
				callback(jsonResult)
			}
		}
	})
}

//获取发布合约所需要的平台币标准
func QueryCCRequireNum(lang string, callback func(model.JsonResult)) {
	chainurl := beego.AppConfig.String("chainurl") + "/udo_queryPublishCCRequireNum"
	common.UDOHttpGet(chainurl, lang, func(result interface{}, err error) {
		if err == nil && result != nil {
			if data, ok := result.([]byte); ok {
				var jsonResult model.JsonResult
				err := json.Unmarshal(data, &jsonResult)
				if err != nil {
					jsonResult.Status = false
					jsonResult.Msg = fmt.Sprintf("json解析出错,%s！", string(data))
				}
				callback(jsonResult)
			}
		}
	})
}

//获取手续费返还规则
func QueryReturnGasConfig(lang string, callback func(model.ReturnGasConfigResponse)) {
	chainurl := beego.AppConfig.String("chainurl") + "/udo_queryReturnGasConfig"
	common.UDOHttpGet(chainurl, lang, func(result interface{}, err error) {
		if err == nil && result != nil {
			if data, ok := result.([]byte); ok {
				var jsonResult model.ReturnGasConfigResponse
				err := json.Unmarshal(data, &jsonResult)
				if err != nil {
					jsonResult.Status = false
					jsonResult.Msg = fmt.Sprintf("json解析出错,%s！", string(data))
				}
				callback(jsonResult)
			}
		}
	})
}

//查询主币信息
func QueryMasterTokenInfo(lang string) *model.Token {
	chainurl := beego.AppConfig.String("chainurl") + "/queryMasterTokenInfo"
	bytes, err := common.UDOSyncHttpGet(chainurl, lang)
	if err != nil || len(bytes) == 0 {
		return nil
	}
	type TempResult struct {
		Status bool        `json:"status"`
		Msg    string      `json:"msg"`
		Data   model.Token `json:"data"`
	}
	var result TempResult
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		return nil
	}
	return &result.Data
}
