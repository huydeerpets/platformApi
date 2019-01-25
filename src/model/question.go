package model

import (
	"platformApi/src/common"
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

//意见反馈
type Question struct {
	Id            int64     `orm:"auto" json:"id" description:"记录ID"`
	Title         string    `orm:"size(100)" json:"title" description:"标题"`
	Email         string    `orm:"size(100)" json:"email" description:"邮箱"`
	Question      string    `orm:"size(1000)" json:"question" description:"问题描述"`
	BaseUrl       string    `orm:"size(100)" json:"baseUrl" description:"附件根url"`
	AttachName    string    `orm:"size(2000)" json:"attachName" description:"附件名称，文件名加后缀名，多个以逗号分隔"`
	CreateTime    time.Time `orm:"auto_now_add;type(datetime)" json:"createTime" description:"创建日期"`
	CreateTimeFmt string    `orm:"-" json:"createTimeFmt" description:"格式化日期输出"`
}

func (q *Question) TableName() string {
	return "t_platform_question"
}

//添加
func AddQuestion(quest *Question) error {
	o := orm.NewOrm()
	_, err := o.Insert(quest)
	return err
}

//修改
func UpdateQuestion(quest *Question) error {
	o := orm.NewOrm()
	//Update 默认更新所有的字段，可以更新指定的字段：
	var err error
	if quest.BaseUrl == "" || quest.AttachName == "" {
		_, err = o.Update(quest, "Title", "Question")
	} else {
		_, err = o.Update(quest, "Title", "Question", "BaseUrl", "AttachName")
	}

	return err
}

//删除
func DeleteQuestion(id int64) error {
	o := orm.NewOrm()
	quest := Question{Id: id}
	err := o.Read(&quest)
	if err != nil {
		return err
	}
	_, err = o.Delete(&quest)
	if err != nil {
		return err
	}
	return nil
}

//查询详情
func GetQuestionInfo(id string) (*Question, error) {
	o := orm.NewOrm()
	quest := new(Question)
	err := o.Raw("SELECT * FROM t_platform_question WHERE id = ?", id).QueryRow(quest)
	if err != nil {
		return nil, err
	}
	quest.CreateTimeFmt = common.FormatTime(quest.CreateTime)
	return quest, nil
}

//搜索相关问题
func SearchQuestions(pageSize, pageNo int, search string) map[string]interface{} {
	o := orm.NewOrm()
	qs := o.QueryTable("t_platform_question")

	if search != "" {
		cond := orm.NewCondition()
		cond = cond.And("Question__icontains", search).Or("Email__icontains", search).Or("Title__icontains", search)
		qs = qs.SetCond(cond)
	}

	resultMap := make(map[string]interface{})
	cnt, err := qs.Count()

	if err == nil && cnt > 0 {
		var questions []Question
		_, err := qs.Limit(pageSize, (pageNo-1)*pageSize).All(&questions)

		if err == nil {
			resultMap["total"] = cnt
			for i, v := range questions {
				v.CreateTimeFmt = common.FormatTime(v.CreateTime)

				questions[i] = v
			}
			resultMap["data"] = questions
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
	orm.RegisterModel(new(Question))
}
