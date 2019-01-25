package model

import (
	"bytes"
	"fmt"
	"platformApi/src/common"
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

// faq
type FaqModel struct {
	Id            int64     `orm:"auto" json:"id"`
	Title         string    `orm:"size(80)" json:"title" description:"标题"`
	Content       string    `orm:"size(4000)" json:"content" description:"内容"`
	CreateTime    time.Time `orm:"auto_now_add;type(datetime);" json:"createTime"`
	UpdateTime    time.Time `orm:"auto_now;type(datetime);" json:"updateTime"` //对于批量的 update 此设置是不生效的
	Status        int64     `orm:"size(11)" json:"status" description:"状态　1:有效果 2:无效"`
	LangType      int       `orm:"size(2)" json:"langType" description:"语言　1:中文 2:英文"`
	CreateTimeFmt string    `orm:"-" json:"createTimeFmt"`
	UpdateTimeFmt string    `orm:"-" json:"updateTimeFmt"`
}

//技术问题搜索结果
type Technology struct {
	Id     int64  `json:"id"`
	Pid    int64  `json:"pid" description:"父ID"`
	Title  string `json:"title" description:"标题"`
	Remark string `json:"remark" description:"描述"`
	Ctype  int    `json:"ctype" description:"内容类型　1:api文档 2:技术教程 3:FAQ"`
}

func (faq *FaqModel) TableName() string {
	return "t_platform_faq"
}

func GetFaqInfo(id string) *FaqModel {
	o := orm.NewOrm()
	faq := new(FaqModel)
	qs := o.QueryTable("t_platform_faq")
	qs = qs.Filter("Id", id)
	qs.One(faq)

	if faq != nil && faq.Title != "" {
		faq.CreateTimeFmt = common.FormatTime(faq.CreateTime)
		faq.UpdateTimeFmt = common.FormatTime(faq.UpdateTime)
	} else {
		faq = nil
	}
	return faq
}

//保存faq信息
func SaveFaqInfo(faq *FaqModel) (int64, error) {
	o := orm.NewOrm()

	faq.Status = 1
	_, err := o.Insert(faq)

	if err != nil {
		return 0, err
	}

	return faq.Id, nil
}

//更新分类信息
func UpdateFaqInfo(id int64, title, content string) error {
	o := orm.NewOrm()

	var faq FaqModel
	faq.Id = id
	faq.Title = title
	faq.Content = content

	_, err := o.Update(&faq, "Title", "Content")

	return err
}

//删除faq信息
func DeleteFaqInfo(id int64) (bool, error) {
	var buf bytes.Buffer
	buf.WriteString("DELETE FROM t_platform_faq")
	buf.WriteString(" WHERE ID =?")
	o := orm.NewOrm()
	_, err := o.Raw(buf.String(), id).Exec()
	if err != nil {
		return false, err
	}
	return true, nil
}

//更新faq状态
func UpdateFAQStatus(id int64) error {
	o := orm.NewOrm()
	faq := new(FaqModel)
	err := o.Raw("SELECT * FROM t_platform_faq WHERE id=?", id).QueryRow(faq)
	if err != nil {
		return err
	}
	if faq.Status == 1 {
		faq.Status = 2
	} else {
		faq.Status = 1
	}
	_, err = o.Update(faq, "Status", "UpdateTime")

	return err
}

/*func QueryFaqs() []FaqModel {
	o := orm.NewOrm()
	var faqs []FaqModel
	o.Raw("SELECT id,title,create_time,update_time,status FROM t_platform_faq WHERE status=1").QueryRows(&faqs)
	for i, v := range faqs {
		v.CreateTimeFmt = common.FormatTime(v.CreateTime)
		v.UpdateTimeFmt = common.FormatTime(v.UpdateTime)
		faqs[i] = v
	}
	return faqs
}*/
func QueryFaqs(lang string) []FaqModel {
	o := orm.NewOrm()
	var faqs []FaqModel
	sqlEnglish :=`SELECT id,title,create_time,update_time,status FROM t_platform_faq WHERE status=1 AND lang_type = 2`
	sqlChinese :=`SELECT id,title,create_time,update_time,status FROM t_platform_faq WHERE status=1 AND lang_type = 1`
	var sql string
	if lang == "zh"{
		sql = sqlChinese
	}else{
		sql = sqlEnglish
	}
	o.Raw(sql).QueryRows(&faqs)
	for i, v := range faqs {
		v.CreateTimeFmt = common.FormatTime(v.CreateTime)
		v.UpdateTimeFmt = common.FormatTime(v.UpdateTime)
		faqs[i] = v
	}
	return faqs
}
func QueryAllFaqs() []FaqModel {
	o := orm.NewOrm()
	var faqs []FaqModel
	o.Raw("SELECT * FROM t_platform_faq WHERE status=1").QueryRows(&faqs)
	for i, v := range faqs {
		v.CreateTimeFmt = common.FormatTime(v.CreateTime)
		v.UpdateTimeFmt = common.FormatTime(v.UpdateTime)
		faqs[i] = v
	}
	return faqs
}

//技术问题搜索
func SearchTechnologys(keyWord string) []Technology {
	o := orm.NewOrm()
	var arr []Technology
	var sql = `SELECT * FROM (
				SELECT id,pid,name as title,remark,1 as ctype FROM t_api_category WHERE depth=3 and status=1
				union
				SELECT cate_id as id,0 as pid,title,summary as remark,2 as ctype FROM t_doc_content WHERE status=1
				union
				SELECT id,0 as pid,title,'具体详细内容请点击此条目...' as remark,3 as ctype from t_platform_faq WHERE status=1) a WHERE a.title LIKE '%s' OR a.remark LIKE '%s' ORDER BY a.ctype ASC`

	sql = fmt.Sprintf(sql, "%"+keyWord+"%", "%"+keyWord+"%")
	o.Raw(sql).QueryRows(&arr)

	return arr
}

func SearchFAQS(pageSize, pageNo int, search, langType string) map[string]interface{} {
	o := orm.NewOrm()
	qs := o.QueryTable("t_platform_faq").Filter("LangType", langType)

	if search != "" {
		cond := orm.NewCondition()
		cond = cond.And("Title__icontains", search).Or("Content__icontains", search)
		qs = qs.SetCond(cond)
	}

	resultMap := make(map[string]interface{})
	cnt, err := qs.Count()

	if err == nil && cnt > 0 {
		var faqs []FaqModel
		_, err := qs.Limit(pageSize, (pageNo-1)*pageSize).OrderBy("-UpdateTime").All(&faqs)

		if err == nil {
			resultMap["total"] = cnt
			for i, v := range faqs {
				v.CreateTimeFmt = common.FormatTime(v.CreateTime)
				v.UpdateTimeFmt = common.FormatTime(v.UpdateTime)
				faqs[i] = v
			}
			resultMap["data"] = faqs
			return resultMap
		}
	}

	resultMap["total"] = 0
	resultMap["data"] = nil

	return resultMap
}

//获取内容详情
func GetTechnologyInfo(id, ctype string) *FaqModel {
	o := orm.NewOrm()
	faq := new(FaqModel)
	qs := o.QueryTable("t_platform_faq")
	qs = qs.Filter("Id", id)
	qs.One(faq)

	if faq != nil && faq.Title != "" {
		faq.CreateTimeFmt = common.FormatTime(faq.CreateTime)
		faq.UpdateTimeFmt = common.FormatTime(faq.UpdateTime)
	} else {
		faq = nil
	}
	return faq
}

//初始化模型
func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(FaqModel))
}
