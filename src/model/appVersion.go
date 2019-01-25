package model

import (
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

//应用版本
type AppVersion struct {
	Id            int64     `orm:"auto" from:"id"`
	AppName       string    `orm:"size(80)" valid:"Required" form:"appName"`
	Version       string    `orm:"size(20)" valid:"Required" form:"version"`
	PlatType      int       `orm:"size(10)" valid:"Required" form:"platType"`
	UpgradeType   int       `orm:"size(10)" valid:"Required" form:"upgradeType"`
	IsCurrent     int       `orm:"size(10)" valid:"Required" form:"isCurrent"`
	AppAddr       string    `orm:"size(100)" form:"appAddr"`
	AppDesc       string    `orm:"size(100)" form:"appDesc"`
	CreateTime    time.Time `orm:"auto_now_add;type(datetime)"`
	CreateTimeFmt string    `orm:"-"` //创建时间
}

type Version struct {
	Id            int64     `json:"id" description:"记录ID"`
	AppName       string    `json:"appName" description:"app名称"`
	Version       string    `json:"version" description:"版本号"`
	PlatType      int       `json:"platType" description:"平台类型 1、andorid 2、Ios"`
	UpgradeType   int       `json:"upgradeType" description:"升级类型　1 、强制升级 2、可忽略"`
	IsCurrent     int       `json:"isCurrent" description:"是否是当前版本 1、表示是 2、表示不是"`
	AppAddr       string    `json:"appAddr" description:"下载地址"`
	Size          string    `orm:"-" json:"size" description:"文件大小"`
	AppDesc       string    `json:"appDesc" description:"app描述"`
	CreateTime    time.Time `json:"createTime" description:"创建时间"`
	CreateTimeFmt string    `json:"createTimeFmt"`
}

func (u *AppVersion) TableName() string {
	return "t_app_version"
}

//根据应用名称、平台类型查询当前使用的版本
func GetCurrentVersion(appName string, platType int64) (*AppVersion, error) {
	o := orm.NewOrm()
	version := new(AppVersion)
	err := o.Raw("SELECT * FROM t_app_version WHERE is_current=1 AND app_name = ? AND plat_type=?", appName, platType).QueryRow(version)
	if err != nil {
		return nil, err
	}
	return version, nil
}

//初始化模型
func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(AppVersion))
}
