package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["platformApi/src/controller:BaseDataController"] = append(beego.GlobalControllerRouter["platformApi/src/controller:BaseDataController"],
        beego.ControllerComments{
            Method: "AddBaseData",
            Router: `/BaseData/AddBaseData`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["platformApi/src/controller:BaseDataController"] = append(beego.GlobalControllerRouter["platformApi/src/controller:BaseDataController"],
        beego.ControllerComments{
            Method: "DeleteBaseData",
            Router: `/BaseData/DeleteBaseData/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["platformApi/src/controller:BaseDataController"] = append(beego.GlobalControllerRouter["platformApi/src/controller:BaseDataController"],
        beego.ControllerComments{
            Method: "DisabledBaseData",
            Router: `/BaseData/DisabledBaseData/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["platformApi/src/controller:BaseDataController"] = append(beego.GlobalControllerRouter["platformApi/src/controller:BaseDataController"],
        beego.ControllerComments{
            Method: "EnabledBaseData",
            Router: `/BaseData/EnabledBaseData/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["platformApi/src/controller:BaseDataController"] = append(beego.GlobalControllerRouter["platformApi/src/controller:BaseDataController"],
        beego.ControllerComments{
            Method: "GetBaseDataByType",
            Router: `/BaseData/GetBaseDataByType/:dataType/:langType`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["platformApi/src/controller:BaseDataController"] = append(beego.GlobalControllerRouter["platformApi/src/controller:BaseDataController"],
        beego.ControllerComments{
            Method: "GetBaseDataByTypeExt",
            Router: `/BaseData/GetBaseDataByTypeExt/:dataType/:langType`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["platformApi/src/controller:BaseDataController"] = append(beego.GlobalControllerRouter["platformApi/src/controller:BaseDataController"],
        beego.ControllerComments{
            Method: "InitAdd",
            Router: `/BaseData/InitAdd/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["platformApi/src/controller:BaseDataController"] = append(beego.GlobalControllerRouter["platformApi/src/controller:BaseDataController"],
        beego.ControllerComments{
            Method: "InitList",
            Router: `/BaseData/InitList/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["platformApi/src/controller:BaseDataController"] = append(beego.GlobalControllerRouter["platformApi/src/controller:BaseDataController"],
        beego.ControllerComments{
            Method: "List",
            Router: `/BaseData/List/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["platformApi/src/controller:ChainApplyController"] = append(beego.GlobalControllerRouter["platformApi/src/controller:ChainApplyController"],
        beego.ControllerComments{
            Method: "InitAdd",
            Router: `/ChainApply/InitAdd/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["platformApi/src/controller:ChainApplyController"] = append(beego.GlobalControllerRouter["platformApi/src/controller:ChainApplyController"],
        beego.ControllerComments{
            Method: "InitList",
            Router: `/ChainApply/InitList/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["platformApi/src/controller:ChainApplyController"] = append(beego.GlobalControllerRouter["platformApi/src/controller:ChainApplyController"],
        beego.ControllerComments{
            Method: "List",
            Router: `/ChainApply/List/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["platformApi/src/controller:ChainApplyController"] = append(beego.GlobalControllerRouter["platformApi/src/controller:ChainApplyController"],
        beego.ControllerComments{
            Method: "UpdateChainApplyStatus",
            Router: `/ChainApply/UpdateChainApplyStatus/:id/:status`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["platformApi/src/controller:FaqController"] = append(beego.GlobalControllerRouter["platformApi/src/controller:FaqController"],
        beego.ControllerComments{
            Method: "DeleteFAQ",
            Router: `/FAQ/DeleteFAQ/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["platformApi/src/controller:FaqController"] = append(beego.GlobalControllerRouter["platformApi/src/controller:FaqController"],
        beego.ControllerComments{
            Method: "InitAdd",
            Router: `/FAQ/InitAdd/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["platformApi/src/controller:FaqController"] = append(beego.GlobalControllerRouter["platformApi/src/controller:FaqController"],
        beego.ControllerComments{
            Method: "InitList",
            Router: `/FAQ/InitList/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["platformApi/src/controller:FaqController"] = append(beego.GlobalControllerRouter["platformApi/src/controller:FaqController"],
        beego.ControllerComments{
            Method: "List",
            Router: `/FAQ/List/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["platformApi/src/controller:FaqController"] = append(beego.GlobalControllerRouter["platformApi/src/controller:FaqController"],
        beego.ControllerComments{
            Method: "SaveFAQ",
            Router: `/FAQ/SaveFAQ`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["platformApi/src/controller:FaqController"] = append(beego.GlobalControllerRouter["platformApi/src/controller:FaqController"],
        beego.ControllerComments{
            Method: "UpdateFAQStatus",
            Router: `/FAQ/UpdateFAQStatus/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["platformApi/src/controller:FileController"] = append(beego.GlobalControllerRouter["platformApi/src/controller:FileController"],
        beego.ControllerComments{
            Method: "UploadToken",
            Router: `/File/UploadToken/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["platformApi/src/controller:QuestionController"] = append(beego.GlobalControllerRouter["platformApi/src/controller:QuestionController"],
        beego.ControllerComments{
            Method: "DeleteQuestion",
            Router: `/Question/DeleteQuestion/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["platformApi/src/controller:QuestionController"] = append(beego.GlobalControllerRouter["platformApi/src/controller:QuestionController"],
        beego.ControllerComments{
            Method: "InitAdd",
            Router: `/Question/InitAdd/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["platformApi/src/controller:QuestionController"] = append(beego.GlobalControllerRouter["platformApi/src/controller:QuestionController"],
        beego.ControllerComments{
            Method: "InitList",
            Router: `/Question/InitList/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["platformApi/src/controller:QuestionController"] = append(beego.GlobalControllerRouter["platformApi/src/controller:QuestionController"],
        beego.ControllerComments{
            Method: "List",
            Router: `/Question/List/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["platformApi/src/controller:UserController"] = append(beego.GlobalControllerRouter["platformApi/src/controller:UserController"],
        beego.ControllerComments{
            Method: "AddUser",
            Router: `/User/AddUser`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["platformApi/src/controller:UserController"] = append(beego.GlobalControllerRouter["platformApi/src/controller:UserController"],
        beego.ControllerComments{
            Method: "DeleteUser",
            Router: `/User/DeleteUser/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["platformApi/src/controller:UserController"] = append(beego.GlobalControllerRouter["platformApi/src/controller:UserController"],
        beego.ControllerComments{
            Method: "InitAdd",
            Router: `/User/InitAdd/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["platformApi/src/controller:UserController"] = append(beego.GlobalControllerRouter["platformApi/src/controller:UserController"],
        beego.ControllerComments{
            Method: "InitList",
            Router: `/User/InitList/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["platformApi/src/controller:UserController"] = append(beego.GlobalControllerRouter["platformApi/src/controller:UserController"],
        beego.ControllerComments{
            Method: "List",
            Router: `/User/List/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["platformApi/src/controller:UserController"] = append(beego.GlobalControllerRouter["platformApi/src/controller:UserController"],
        beego.ControllerComments{
            Method: "Qrcode",
            Router: `/User/Qrcode`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
