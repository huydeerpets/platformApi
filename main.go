package main

import (
	"encoding/json"
	"fmt"
	_ "platformApi/routers"
	"platformApi/src/common"
	"platformApi/src/cors"
	"platformApi/src/model"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/orm"
	"github.com/beego/i18n"
	"os"
)

//初始化
func init() {
	dbhost := beego.AppConfig.String("dbhost")
	dbport := beego.AppConfig.String("dbport")
	dbuser := beego.AppConfig.String("dbuser")
	dbpassword := beego.AppConfig.String("dbpassword")
	db := beego.AppConfig.String("db")

	//让beego也采用+8时区的时间
	//Beego的ORM插入Mysql后，时区不一致的解决方案
	//1、orm.RegisterDataBase("default", "mysql", "root:LPET6Plus@tcp(127.0.0.1:18283)/lpet6plusdb?charset=utf8&loc=Local")
	//2、orm.RegisterDataBase("default", "mysql", "db_test:dbtestqwe321@tcp(127.0.0.1:3306)/db_test?charset=utf8&loc=Asia%2FShanghai")
	//orm.DefaultTimeLoc, _ = time.LoadLocation("Asia/Shanghai")
	//注册mysql Driver
	orm.RegisterDriver("mysql", orm.DRMySQL)
	//构造conn连接
	//用户名:密码@tcp(url地址)/数据库
	conn := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + db + "?charset=utf8&parseTime=true&loc=Local"
	//注册数据库连接
	orm.RegisterDataBase("default", "mysql", conn)

	fmt.Printf("数据库连接成功！%s\n", conn)

}

func main() {
	//简单的设置 Debug 为 true 打印查询的语句,可能存在性能问题，不建议使用在产品模式
	orm.Debug = true
	o := orm.NewOrm()
	o.Using("default") // 默认使用 default，你可以指定为其他数据库
	//启用Session
	beego.BConfig.WebConfig.Session.SessionOn = true

	beego.BConfig.WebConfig.Session.SessionName = "platformApisessionID"

	//加载不同语言版本的提示信息
	if err := i18n.SetMessage("zh-CN", "conf/locale_zh-CN.ini"); err != nil {
		beego.Debug("i18n.SetMessage conf/locale_zh-CN.ini err = ", err)
		os.Exit(1)
	}
	if err := i18n.SetMessage("zh-TW", "conf/locale_zh-TW.ini"); err != nil {
		beego.Debug("i18n.SetMessage conf/locale_zh-TW.ini err = ", err)
		os.Exit(1)
	}
	if err := i18n.SetMessage("en-US", "conf/locale_en-US.ini"); err != nil {
		beego.Debug("i18n.SetMessage conf/locale_en-US.ini err = ", err)
		os.Exit(1)
	}

	var FilterUser = func(ctx *context.Context) {
		noValidateMap := map[string]string{
			"/Home/Login":     "GET/POST",
			"/Home/InitReg":   "GET/POST",
			"/Home/InitPass":  "GET/POST",
			"/Home/Reg":       "POST",
			"/Home/UpdatePwd": "POST",
		}
		fmt.Println("****", ctx.Request.RequestURI, ctx.Request.Method)
		requireValidate := true
		for k, v := range noValidateMap {
			result := strings.Contains(v, ctx.Request.Method)
			if ctx.Request.RequestURI == k && result {
				requireValidate = false
				break
			}
		}
		if requireValidate {
			u, ok := ctx.Input.Session(common.USER_INFO).(*model.User)
			if !ok {
				ctx.Redirect(302, "/")
			}
			//必须是管理员才能操作
			if strings.Contains(ctx.Request.RequestURI, "/User/DeleteUser") || strings.Contains(ctx.Request.RequestURI, "/User/AddUser") {
				if u.UserName != "admin" {
					result := new(model.Result)
					result.Status = 1
					result.Msg = "无权限进行此操作!"
					jsonBytes, _ := json.Marshal(result)
					ctx.ResponseWriter.Write(jsonBytes)
					return
				}
			}
		}

	}
	beego.InsertFilter("/Home/*", beego.BeforeRouter, FilterUser)
	beego.InsertFilter("/User/*", beego.BeforeRouter, FilterUser)
	//beego.InsertFilter("/ApiCategory/*", beego.BeforeRouter, FilterUser)
	//beego.InsertFilter("/DocCategory/*", beego.BeforeRouter, FilterUser)
	//beego.InsertFilter("/BaseData/*", beego.BeforeRouter, FilterUser)
	//beego.InsertFilter("/Question/*", beego.BeforeRouter, FilterUser)
	//beego.InsertFilter("/FAQ/*", beego.BeforeRouter, FilterUser)

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	beego.Run()
}
