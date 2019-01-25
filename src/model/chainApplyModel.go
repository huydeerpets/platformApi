package model

import (
	"platformApi/src/common"
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

//链申请信息
type ChainApplyInfo struct {
	Id            int64     `orm:"auto" json:"id" description:"记录ID"`
	Name          string    `orm:"size(80)" json:"name" description:"链中文名称"`
	EnShort       string    `orm:"size(40)" json:"en_short" description:"链英文简称"`
	EMail         string    `orm:"size(40)" json:"e_mail" description:"邮箱"`
	ContactName   string    `orm:"size(40)" json:"contact_name" description:"联系人"`
	ContactTel    string    `orm:"size(20)" json:"contact_tel" description:"联系电话"`
	Remark        string    `orm:"size(200)" json:"remark" description:"备注"`
	Status        int       `orm:"size(2)" json:"status" description:"状态 1:线下沟通 2:入驻成功　3:入驻失败"`
	CreateTime    time.Time `orm:"auto_now_add;type(datetime)" json:"create_time" description:"创建日期"`
	CreateTimeFmt string    `orm:"-" json:"createTimeFmt" description:"格式化日期输出"`
}

func (c *ChainApplyInfo) TableName() string {
	return "t_chain_enter"
}

//修改
func UpdateChainApplyStatus(id, status string) error {
	o := orm.NewOrm()
	_, err := o.Raw("UPDATE t_chain_enter SET status=? WHERE id=?", status, id).Exec()

	return err
}

//查询详情
func QueryChainApplyInfo(id string) (*ChainApplyInfo, error) {
	o := orm.NewOrm()
	applyInfo := new(ChainApplyInfo)
	err := o.Raw("SELECT * FROM t_chain_enter WHERE id = ?", id).QueryRow(applyInfo)
	if err != nil {
		return nil, err
	}
	applyInfo.CreateTimeFmt = common.FormatTime(applyInfo.CreateTime)

	return applyInfo, nil
}

//搜索相关申请信息
func SearchChainApplys(pageSize, pageNo int, search string) map[string]interface{} {
	o := orm.NewOrm()
	qs := o.QueryTable("t_chain_enter")

	if search != "" {
		cond := orm.NewCondition()
		cond = cond.And("Name__icontains", search).Or("EnShort__icontains", search).Or("ContactName__icontains", search).Or("ContactTel__icontains", search)
		qs = qs.SetCond(cond)
	}

	resultMap := make(map[string]interface{})
	cnt, err := qs.Count()

	if err == nil && cnt > 0 {
		var applys []ChainApplyInfo
		_, err := qs.Limit(pageSize, (pageNo-1)*pageSize).OrderBy("-CreateTime").All(&applys)

		if err == nil {
			resultMap["total"] = cnt
			for i, v := range applys {
				v.CreateTimeFmt = common.FormatTime(v.CreateTime)

				applys[i] = v
			}
			resultMap["data"] = applys
			return resultMap
		}
	}

	resultMap["total"] = 0
	resultMap["data"] = nil

	return resultMap
}

//初始化模型
func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(ChainApplyInfo))
}
