package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["platformApi/src/api:ApiDocController"] = append(beego.GlobalControllerRouter["platformApi/src/api:ApiDocController"],
        beego.ControllerComments{
            Method: "QueryDocAllTheme",
            Router: `/queryDocAllTheme`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["platformApi/src/api:ApiDocController"] = append(beego.GlobalControllerRouter["platformApi/src/api:ApiDocController"],
        beego.ControllerComments{
            Method: "QueryDocInfo",
            Router: `/queryDocInfo/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["platformApi/src/api:ApiDocController"] = append(beego.GlobalControllerRouter["platformApi/src/api:ApiDocController"],
        beego.ControllerComments{
            Method: "QueryDocTheme",
            Router: `/queryDocTheme`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["platformApi/src/api:ApiFaqController"] = append(beego.GlobalControllerRouter["platformApi/src/api:ApiFaqController"],
        beego.ControllerComments{
            Method: "QueryFaqAllList",
            Router: `/queryFaqAllList`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["platformApi/src/api:ApiFaqController"] = append(beego.GlobalControllerRouter["platformApi/src/api:ApiFaqController"],
        beego.ControllerComments{
            Method: "QueryFaqInfo",
            Router: `/queryFaqInfo/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["platformApi/src/api:ApiFaqController"] = append(beego.GlobalControllerRouter["platformApi/src/api:ApiFaqController"],
        beego.ControllerComments{
            Method: "QueryFaqList",
            Router: `/queryFaqList`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["platformApi/src/api:ApiFaqController"] = append(beego.GlobalControllerRouter["platformApi/src/api:ApiFaqController"],
        beego.ControllerComments{
            Method: "QueryTechnologyList",
            Router: `/queryTechnologyList`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["platformApi/src/api:ApiQueryController"] = append(beego.GlobalControllerRouter["platformApi/src/api:ApiQueryController"],
        beego.ControllerComments{
            Method: "QueryApiInfo",
            Router: `/queryApiInfo/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["platformApi/src/api:ApiQueryController"] = append(beego.GlobalControllerRouter["platformApi/src/api:ApiQueryController"],
        beego.ControllerComments{
            Method: "QueryApiTheme",
            Router: `/queryApiTheme/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["platformApi/src/api:ApiQuestionController"] = append(beego.GlobalControllerRouter["platformApi/src/api:ApiQuestionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/addQuestion`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["platformApi/src/api:ApiSignController"] = append(beego.GlobalControllerRouter["platformApi/src/api:ApiSignController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/addSignData`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["platformApi/src/api:ApiSignController"] = append(beego.GlobalControllerRouter["platformApi/src/api:ApiSignController"],
        beego.ControllerComments{
            Method: "GetSignStatus",
            Router: `/getSignStatus/:qrCode`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["platformApi/src/api:ApiSignController"] = append(beego.GlobalControllerRouter["platformApi/src/api:ApiSignController"],
        beego.ControllerComments{
            Method: "QuerySignInfo",
            Router: `/querySignInfo/:qrCode`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["platformApi/src/api:ApiVersionController"] = append(beego.GlobalControllerRouter["platformApi/src/api:ApiVersionController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/getVersionInfo/:platType`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["platformApi/src/api:ChainController"] = append(beego.GlobalControllerRouter["platformApi/src/api:ChainController"],
        beego.ControllerComments{
            Method: "ChainApply",
            Router: `/chainApply`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["platformApi/src/api:ChainController"] = append(beego.GlobalControllerRouter["platformApi/src/api:ChainController"],
        beego.ControllerComments{
            Method: "ChainApplyInfo",
            Router: `/chainApplyInfo/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["platformApi/src/api:ChainController"] = append(beego.GlobalControllerRouter["platformApi/src/api:ChainController"],
        beego.ControllerComments{
            Method: "ChainSearch",
            Router: `/chainSearch`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["platformApi/src/api:ChainController"] = append(beego.GlobalControllerRouter["platformApi/src/api:ChainController"],
        beego.ControllerComments{
            Method: "QueryPublishCCRequireNum",
            Router: `/queryPublishCCRequireNum`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["platformApi/src/api:ChainController"] = append(beego.GlobalControllerRouter["platformApi/src/api:ChainController"],
        beego.ControllerComments{
            Method: "QueryPublishTokenRequireNum",
            Router: `/queryPublishTokenRequireNum`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["platformApi/src/api:ChainController"] = append(beego.GlobalControllerRouter["platformApi/src/api:ChainController"],
        beego.ControllerComments{
            Method: "QueryReturnGasConfig",
            Router: `/queryReturnGasConfig`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["platformApi/src/api:ContractController"] = append(beego.GlobalControllerRouter["platformApi/src/api:ContractController"],
        beego.ControllerComments{
            Method: "QueryContractInfo",
            Router: `/queryContractInfo/:address/:version`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["platformApi/src/api:ContractController"] = append(beego.GlobalControllerRouter["platformApi/src/api:ContractController"],
        beego.ControllerComments{
            Method: "QueryContractList",
            Router: `/queryContractList/:address`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["platformApi/src/api:TokenController"] = append(beego.GlobalControllerRouter["platformApi/src/api:TokenController"],
        beego.ControllerComments{
            Method: "IsFirstSetMasterAndManager",
            Router: `/isFirstSetMasterAndManager/:tokenID`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["platformApi/src/api:TokenController"] = append(beego.GlobalControllerRouter["platformApi/src/api:TokenController"],
        beego.ControllerComments{
            Method: "QueryAllTokens",
            Router: `/queryAllTokens`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["platformApi/src/api:TokenController"] = append(beego.GlobalControllerRouter["platformApi/src/api:TokenController"],
        beego.ControllerComments{
            Method: "QueryTokenInfo",
            Router: `/queryTokenInfo/:tokenID`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["platformApi/src/api:TokenController"] = append(beego.GlobalControllerRouter["platformApi/src/api:TokenController"],
        beego.ControllerComments{
            Method: "QueryTokenList",
            Router: `/queryTokenList`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
