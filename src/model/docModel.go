package model

import (
	"bytes"
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

// doc分类
type DocCategory struct {
	Id            int64  `orm:"auto"`
	Name          string `orm:"size(100)" description:"节点名称"`
	Pid           int64  `description:"父ID"`
	Depth         int64  `description:"节点深度"`
	FullPath      string `orm:"size(255)" description:"节点全路径"`
	Status        int64
	OrderNum      int64
	CreateTime    time.Time `orm:"auto_now_add;type(datetime);column(create_time)"`
	CreateTimeFmt string    `orm:"-"`
}

// doc内容
type DocContent struct {
	Id            int64     `orm:"auto" json:"id"`
	CateId        int64     `orm:"size(11)" json:"cateId" description:"分类ID"`
	Title         string    `orm:"size(80)" json:"title" description:"标题"`
	Summary       string    `orm:"size(2000)" json:"summary" description:"简介"`
	Content       string    `orm:"size(4000)" json:"content" description:"内容"`
	AttachAddr    string    `orm:"size(200)" json:"attachAddr" description:"附件地址"`
	CreateTime    time.Time `orm:"auto_now_add;type(datetime);" json:"createTime"`
	UpdateTime    time.Time `orm:"auto_now;type(datetime);" json:"updateTime"` //对于批量的 update 此设置是不生效的
	Status        int64     `orm:"size(11)" json:"status" description:"状态　1:有效果 2:无效"`
	CreateTimeFmt string    `orm:"-"`
	UpdateTimeFmt string    `orm:"-"`
}

// doc主题信息
type DocThemeItem struct {
	Id      int64  `json:"id"  description:"记录ID"`
	Title   string `json:"title" description:"标题"`
	Summary string `json:"summary" description:"简介"`
}

func (cate *DocCategory) TableName() string {
	return "t_doc_category"
}

func (doc *DocContent) TableName() string {
	return "t_doc_content"
}

func GetDocCategoryList(langType string, isAsc bool) []DocCategory {
	o := orm.NewOrm()
	var cates []DocCategory
	sql := "select id,name,pid,depth,full_path,status,CONCAT(full_path,',',id) as bpath from t_doc_category where status=1 and lang_type=? order by pid,order_num desc"
	if isAsc {
		sql = "select id,name,pid,depth,full_path,status,CONCAT(full_path,',',id) as bpath from t_doc_category where status=1 and lang_type=? order by pid,order_num asc"
	}
	num, err := o.Raw(sql, langType).QueryRows(&cates)
	fmt.Println("category err: ", err)
	if err == nil {
		fmt.Println("category nums: ", num)
	}
	return cates
}

func GetDocCateList() ([]DocCategory, error) {
	o := orm.NewOrm()
	qs := o.QueryTable("t_doc_category")
	qs = qs.Filter("Depth", 2).Filter("Status", 1)
	var cates []DocCategory
	cnt, err := qs.All(&cates)

	if err == nil {
		fmt.Println("count", cnt)
	}
	return cates, err
}

/**
 * 通过传入要查找的pid来递归查找他的下级
 * @param  []Category  cates    数组代替数据库中的数据
 * @param  integer pid     父id
 * @param  []DropItem   result 结果数组,&保证变量常驻
 * @param  int depth  输出的分隔符--,无实际意义
 * @return 树状结构数组
 */
func GetDocList(cates []DocCategory, pid int64, result []DropItem, depth int) []DropItem {
	depth = depth + 2
	for _, val := range cates {
		if pid == val.Pid {
			//str := fmt.Sprintf("%s%s", strings.Repeat("--", depth), val.Name)
			str := fmt.Sprintf("%s%s", "", val.Name)
			var dropItem DropItem
			dropItem.Id = val.Id
			dropItem.Title = str
			dropItem.Depth = depth
			result = append(result, dropItem)
			result = GetDocList(cates, val.Id, result, depth)
		}
	}
	return result
}

func GetDocCateInfo(id int64) *DocCategory {
	o := orm.NewOrm()
	cate := new(DocCategory)
	qs := o.QueryTable("t_doc_category")
	qs = qs.Filter("Id", id)
	qs.One(cate)

	return cate
}

func GetDocContentInfo(cateID int64) *DocContent {
	o := orm.NewOrm()
	var doc DocContent
	err := o.Raw("SELECT * FROM t_doc_content WHERE status=1 AND cate_id=?", cateID).QueryRow(&doc)
	if err != nil {
		return nil
	}
	return &doc
}

/*func QueryDocThemeList() []DocThemeItem {
	sql := `SELECT b.cate_id as id,b.title,b.summary from t_doc_category a 
	LEFT JOIN t_doc_content b ON b.cate_id=a.id
	WHERE a.id IN(67,64,63,9) ORDER BY a.order_num ASC`
	var docs []DocThemeItem
	o := orm.NewOrm()
	o.Raw(sql).QueryRows(&docs)

	return docs
}*/
func QueryDocThemeList(lang string) []DocThemeItem {
	sqlEnglish :=`SELECT b.cate_id as id,b.title,b.summary from t_doc_category a 
	LEFT JOIN t_doc_content b ON b.cate_id=a.id
	WHERE a.id IN(85,82,81,76) ORDER BY a.order_num ASC`
	sqlChinese :=`SELECT b.cate_id as id,b.title,b.summary from t_doc_category a 
	LEFT JOIN t_doc_content b ON b.cate_id=a.id
	WHERE a.id IN(67,64,63,9) ORDER BY a.order_num ASC`
	var sql string
	if lang == "zh"{
		sql = sqlChinese
	}else{
		sql = sqlEnglish
	}
	var docs []DocThemeItem
	o := orm.NewOrm()
	o.Raw(sql).QueryRows(&docs)

	return docs
}
func QueryDocAllThemeList() []DocThemeItem {
	sql := `SELECT a.id,b.title,b.summary from t_doc_category a 
		LEFT JOIN t_doc_content b ON b.cate_id=a.id
		WHERE a.status=1 AND a.pid=3 ORDER BY a.order_num ASC`
	var docs []DocThemeItem
	o := orm.NewOrm()
	o.Raw(sql).QueryRows(&docs)

	return docs
}

//保存分类信息
func SaveDocCategroy(cate *DocCategory) (int64, error) {
	o := orm.NewOrm()

	cate.Status = 1
	_, err := o.Insert(cate)

	if err != nil {
		return 0, err
	}

	return cate.Id, nil
}

//保存doc内容信息
func SaveDocInfo(doc DocContent) (int64, error) {
	o := orm.NewOrm()
	_, err := o.Insert(&doc)

	return 0, err
}

//更新文档内容信息
func UpdateDocContentInfo(id int64, name, summary, content, attachAddr string) error {
	o := orm.NewOrm()

	var cate DocCategory
	cate.Id = id
	cate.Name = name
	o.Begin()
	_, err := o.Update(&cate, "Name")
	if err != nil {
		o.Rollback()
		return err
	}
	_, err = o.Raw("UPDATE t_doc_content SET title=?,summary=?,content=?,attach_addr=? WHERE cate_id=?", name, summary, content, attachAddr, id).Exec()
	if err != nil {
		o.Rollback()
		return err
	}
	o.Commit()

	return err
}

//更新分类信息
func UpdateDocCategroy(id int64, orderNum int64, name string) error {
	o := orm.NewOrm()

	var cate DocCategory
	cate.Id = id
	cate.Name = name
	cate.OrderNum = orderNum

	//Update 默认更新所有的字段，可以更新指定的字段：
	_, err := o.Update(&cate, "Name", "OrderNum")

	return err
}

//删除分类信息
func DeleteDocCategroy(id, pid, orderNum, nodeLevel int64) (bool, error) {
	var buf bytes.Buffer
	buf.WriteString("UPDATE t_doc_category SET status=2")
	buf.WriteString(" WHERE ID =?")
	o := orm.NewOrm()
	o.Begin() //开事物
	_, err := o.Raw(buf.String(), id).Exec()
	if err != nil {
		o.Rollback()
		return false, err
	}
	var buf2 bytes.Buffer
	buf2.WriteString("UPDATE t_doc_category SET order_num=order_num-1")
	buf2.WriteString(" WHERE pid = ? AND  order_num>?")
	_, err2 := o.Raw(buf2.String(), pid, orderNum).Exec()
	if err2 != nil {
		o.Rollback()
		return false, err2
	}

	if nodeLevel == 4 {
		delSqls := []string{"UPDATE t_doc_content SET status=2 WHERE cate_id=?"}
		for _, sql := range delSqls {
			_, err = o.Raw(sql, id).Exec()
			if err != nil {
				o.Rollback()
				return false, err
			}
		}
	}

	o.Commit() //所有事物统一提交
	return true, nil
}

//初始化模型
func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(DocCategory), new(DocContent))
}
