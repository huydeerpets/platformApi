package model

import (
	"bytes"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

// api分类
type ApiCategory struct {
	Id            int64  `orm:"auto"`
	Name          string `orm:"size(100)" description:"节点名称"`
	Pid           int64  `description:"父ID"`
	Depth         int64  `description:"节点深度"`
	FullPath      string `orm:"size(255)" description:"节点全路径"`
	RequestUrl    string `orm:"size(100)" description:"请求路径"`
	Method        string `orm:"size(10)" description:"请求方法GET、POST"`
	Status        int64
	OrderNum      int64
	Remark        string    `orm:"size(1000)" description:"描述或说明"`
	CreateTime    time.Time `orm:"auto_now_add;type(datetime);column(create_time)"`
	CreateTimeFmt string    `orm:"-"`
}

// 请求参数
type RequestParam struct {
	Id            int64     `orm:"auto" json:"id"`
	CateId        int64     `orm:"size(11)" json:"cateId" description:"API接口ID"`
	ArgName       string    `orm:"size(80)" json:"arg_name" description:"参数名称"`
	ArgType       string    `orm:"size(40)" json:"arg_type" description:"请求参数类型"`
	DataType      string    `orm:"size(40)" json:"data_type" description:"数据类型"`
	IsRequire     string    `orm:"size(1)" json:"is_require" description:"是否必填"`
	Description   string    `orm:"size(100)" json:"description" description:"参数描述"`
	Status        int       `orm:"size(2)" json:"status" description:"状态 1、启用 2、禁用"`
	OrderNum      int64     `orm:"size(11)" json:"order_num" description:"排序号"`
	CreateTime    time.Time `orm:"auto_now_add;type(datetime);" json:"createTime"`
	CreateTimeFmt string    `orm:"-"`
}

// 表单构造参数
type FormData struct {
	Id            int64     `orm:"auto" json:"id"`
	CateId        int64     `orm:"size(11)" json:"cateId" description:"API接口ID"`
	DataName      string    `orm:"size(80)" json:"data_name" description:"名称"`
	DataType      string    `orm:"size(40)" json:"data_type" description:"数据类型"`
	IsRequire     string    `orm:"size(1)" json:"is_require" description:"是否必填"`
	Description   string    `orm:"size(100)" json:"description" description:"参数描述"`
	Status        int       `orm:"size(2)" json:"status" description:"状态 1、启用 2、禁用"`
	OrderNum      int64     `orm:"size(11)" json:"order_num" description:"排序号"`
	CreateTime    time.Time `orm:"auto_now_add;type(datetime);" json:"createTime"`
	CreateTimeFmt string    `orm:"-"`
}

//字段汇总表
type RequestField struct {
	Id            int64     `orm:"auto" json:"id"`
	CateId        int64     `orm:"size(11)" json:"cateId" description:"api接口分类ID"`
	Title         string    `orm:"size(100)" json:"title" description:"标题"`
	CreateTime    time.Time `orm:"auto_now_add;type(datetime);" json:"createTime"`
	CreateTimeFmt string    `orm:"-"`
}

//字段明细说明
type FieldItem struct {
	Id            int64     `orm:"auto" json:"id"`
	FieldId       int64     `orm:"size(11)" json:"fieldId" description:"关联t_request_fields表ID"`
	DataName      string    `orm:"size(80)" json:"data_name" description:"名称"`
	DataType      string    `orm:"size(40)" json:"data_type" description:"数据类型"`
	IsRequire     string    `orm:"size(1)" json:"is_require" description:"是否必填"`
	Description   string    `orm:"size(100)" json:"description" description:"参数描述"`
	Status        int       `orm:"size(2)" json:"status" description:"状态 1、启用 2、禁用"`
	OrderNum      int64     `orm:"size(11)" json:"order_num" description:"排序号"`
	CreateTime    time.Time `orm:"auto_now_add;type(datetime);" json:"createTime"`
	CreateTimeFmt string    `orm:"-"`
}

// 请求实例
type RequestSample struct {
	Id            int64     `orm:"auto" json:"id"`
	CateId        int64     `orm:"size(11)" json:"cateId" description:"API接口ID"`
	Title         string    `orm:"size(80)" json:"title" description:"实例标题"`
	Content       string    `orm:"size(4000)" json:"content" description:"实例内容"`
	CreateTime    time.Time `orm:"auto_now_add;type(datetime);" json:"createTime"`
	CreateTimeFmt string    `orm:"-"`
}

//代码
type RequestCode struct {
	Id            int64     `orm:"auto" json:"id"`
	CateId        int64     `orm:"size(11)" json:"cateId" description:"API接口ID"`
	LangType      string    `orm:"size(40)" json:"langType" description:"语言类型"`
	Content       string    `orm:"size(4000)" json:"content" description:"代码"`
	CreateTime    time.Time `orm:"auto_now_add;type(datetime);" json:"createTime"`
	CreateTimeFmt string    `orm:"-"`
}

//响应实例
type ResponseSample struct {
	Id            int64     `orm:"auto" json:"id"`
	CateId        int64     `orm:"size(11)" json:"cateId" description:"API接口ID"`
	RespType      string    `orm:"size(40)" json:"respType" description:"响应类型"`
	Content       string    `orm:"size(4000)" json:"content" description:"内容"`
	CreateTime    time.Time `orm:"auto_now_add;type(datetime);" json:"createTime"`
	CreateTimeFmt string    `orm:"-"`
}

type ResponseParam struct {
	Id            int64     `orm:"auto" json:"id"`
	ParentId      int64     `orm:"size(20)" json:"parentId" description:"父ID"`
	CateId        int64     `orm:"size(11)" json:"cateId" description:"API接口ID"`
	DataName      string    `orm:"size(80)" json:"data_name" description:"名称"`
	DataType      string    `orm:"size(40)" json:"data_type" description:"数据类型"`
	Description   string    `orm:"size(100)" json:"description" description:"参数描述"`
	Idx           int64     `orm:"size(11)" json:"idx" description:"同一个父ID下索引相同的为一组元素"`
	Status        int       `orm:"size(2)" json:"status" description:"状态 1、启用 2、禁用"`
	OrderNum      int64     `orm:"size(11)" json:"order_num" description:"排序号"`
	CreateTime    time.Time `orm:"auto_now_add;type(datetime);" json:"createTime"`
	CreateTimeFmt string    `orm:"-"`
}

type ResponseParamReq struct {
	Id            int64     `orm:"auto" json:"id"`
	ParentId      string    `orm:"size(20)" json:"parentId" description:"父ID"`
	CateId        int64     `orm:"size(11)" json:"cateId" description:"API接口ID"`
	DataName      string    `orm:"size(80)" json:"data_name" description:"名称"`
	DataType      string    `orm:"size(40)" json:"data_type" description:"数据类型"`
	Description   string    `orm:"size(100)" json:"description" description:"参数描述"`
	Status        int       `orm:"size(2)" json:"status" description:"状态 1、启用 2、禁用"`
	OrderNum      int64     `orm:"size(11)" json:"order_num" description:"排序号"`
	CreateTime    time.Time `orm:"auto_now_add;type(datetime);" json:"createTime"`
	CreateTimeFmt string    `orm:"-"`
}

type RequestFieldData struct {
	ReqFiled RequestField `json:"reqFiled"`
	Items    []FieldItem  `json:"items"`
}

type SaveApiRequest struct {
	Pid                   string             `json:"pid"`
	FullPath              string             `json:"full_path"`
	Depth                 string             `json:"depth"`
	SubCount              string             `json:"subCount"`
	Name                  string             `json:"name"`
	Method                string             `json:"method"`
	RequestUrl            string             `json:"request_url"`
	Remark                string             `json:"remark"`
	DeleteIds             string             `json:"deleteIds"`
	DeleteReqSampleIds    string             `json:"deleteReqSampleIds"`
	DeleteReqCodeIds      string             `json:"deleteReqCodeIds"`
	DeleteFormDataIds     string             `json:"deleteFormDataIds"`
	DeleteReqFieldIds     string             `json:"deleteReqFieldIds"`     //字段说明汇总ID
	DeleteReqFieldItemIds string             `json:"deleteReqFieldItemIds"` //字段说明明细ID
	Params                []RequestParam     `json:"params"`
	FormDatas             []FormData         `json:"formDatas"`
	ReqSamples            []RequestSample    `json:"reqSamples"`
	ReqCodes              []RequestCode      `json:"reqCodes"`
	RespSamples           []ResponseSample   `json:"respSamples"`
	RespItems             []ResponseParamReq `json:"respItems"`
	FieldData             []RequestFieldData `json:"fieldData"`
}

type DropItem struct {
	Id    int64
	Title string
	Depth int
}

type ApiThemeItem struct {
	Id    int64          `json:"id"`
	Pid   int64          `json:"pid"`
	Title string         `json:"title" description:"主题名称"`
	Items []ApiThemeItem `json:"items" description:"主题项"`
}

func (cate *ApiCategory) TableName() string {
	return "t_api_category"
}

func (param *RequestParam) TableName() string {
	return "t_request_params"
}

func (reqField *RequestField) TableName() string {
	return "t_request_fields"
}

func (fieldItem *FieldItem) TableName() string {
	return "t_field_items"
}

func (formData *FormData) TableName() string {
	return "t_form_data"
}

func (sample *RequestSample) TableName() string {
	return "t_request_samples"
}

func (code *RequestCode) TableName() string {
	return "t_request_codes"
}

func (respSample *ResponseSample) TableName() string {
	return "t_response_samples"
}

func (respParam *ResponseParam) TableName() string {
	return "t_response_params"
}

func GetCategoryList(langType string, isAsc bool) []ApiCategory {
	o := orm.NewOrm()
	var cates []ApiCategory
	sql := "select id,name,pid,depth,full_path,status,CONCAT(full_path,',',id) as bpath from t_api_category where lang_type=? order by pid,order_num desc"
	if isAsc {
		sql = "select id,name,pid,depth,full_path,status,CONCAT(full_path,',',id) as bpath from t_api_category where lang_type=? order by pid,order_num asc"
	}
	num, err := o.Raw(sql, langType).QueryRows(&cates)
	fmt.Println("category err: ", err)
	if err == nil {
		fmt.Println("category nums: ", num)
	}
	return cates
}

func GetCateList() ([]ApiCategory, error) {
	o := orm.NewOrm()
	qs := o.QueryTable("t_api_category")
	qs = qs.Filter("Depth", 2)
	var cates []ApiCategory
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
func GetList(cates []ApiCategory, pid int64, result []DropItem, depth int) []DropItem {
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
			result = GetList(cates, val.Id, result, depth)
		}
	}
	return result
}

func GetCategoryInfo(id int64) *ApiCategory {
	o := orm.NewOrm()
	cate := new(ApiCategory)
	qs := o.QueryTable("t_api_category")
	qs = qs.Filter("Id", id)
	qs.One(cate)
	return cate
}

func GetApiRequestInfo(id int64) []RequestParam {
	o := orm.NewOrm()
	var params []RequestParam
	_, err := o.Raw("SELECT * FROM t_request_params WHERE status=1 AND cate_id=? ORDER BY order_num ASC", id).QueryRows(&params)
	if err != nil {
		return nil
	}
	return params
}

func GetApiRequestFieldInfo(id int64) []RequestFieldData {
	o := orm.NewOrm()
	var fields []RequestField
	//DATE_FORMAT(create_time,'%Y-%m-%d %H:%i:%s')
	_, err := o.Raw("SELECT * FROM t_request_fields WHERE cate_id=? ORDER BY id ASC", id).QueryRows(&fields)
	if err != nil {
		return nil
	}
	var fieldData []RequestFieldData
	for _, v := range fields {
		var items []FieldItem
		o.Raw("SELECT * FROM t_field_items WHERE field_id=? ORDER BY order_num ASC", v.Id).QueryRows(&items)
		var fieldDataItem RequestFieldData
		fieldDataItem.ReqFiled = v
		fieldDataItem.Items = items
		fieldData = append(fieldData, fieldDataItem)
	}
	return fieldData
}

func GetApiFormDataInfo(id int64) []FormData {
	o := orm.NewOrm()
	var datas []FormData
	_, err := o.Raw("SELECT * FROM t_form_data WHERE status=1 AND cate_id=? ORDER BY order_num ASC", id).QueryRows(&datas)
	if err != nil {
		return nil
	}
	return datas
}

func GetRequestSampleInfo(cateID int64) []RequestSample {
	o := orm.NewOrm()
	var samples []RequestSample
	_, err := o.Raw("SELECT * FROM t_request_samples WHERE cate_id=? ORDER BY id ASC", cateID).QueryRows(&samples)
	if err != nil {
		return nil
	}
	return samples
}

func GetRespSampleInfo(cateID int64) []ResponseSample {
	o := orm.NewOrm()
	var samples []ResponseSample
	_, err := o.Raw("SELECT * FROM t_response_samples WHERE cate_id=? ORDER BY id ASC", cateID).QueryRows(&samples)
	if err != nil {
		return nil
	}
	return samples
}

func GetRequestCodeInfo(cateID int64) []RequestCode {
	o := orm.NewOrm()
	var codes []RequestCode
	_, err := o.Raw("SELECT * FROM t_request_codes WHERE cate_id=? ORDER BY id ASC", cateID).QueryRows(&codes)
	if err != nil {
		return nil
	}
	return codes
}

func GetRespParamInfo(cateID int64) []ResponseParam {
	o := orm.NewOrm()
	var params []ResponseParam
	_, err := o.Raw("SELECT * FROM t_response_params WHERE cate_id=? ORDER BY parent_id,idx,order_num ASC", cateID).QueryRows(&params)
	if err != nil {
		return nil
	}
	return params
}

//保存分类信息
func SaveCategroy(cate *ApiCategory) (int64, error) {
	o := orm.NewOrm()

	cate.Status = 1
	_, err := o.Insert(cate)

	if err != nil {
		return 0, err
	}

	return cate.Id, nil
}

//保存接口信息
func SaveApiInfo(request SaveApiRequest) (int64, error) {
	o := orm.NewOrm()
	o.InsertMulti(len(request.Params), request.Params)
	o.InsertMulti(len(request.FormDatas), request.FormDatas)
	o.InsertMulti(len(request.ReqSamples), request.ReqSamples)
	o.InsertMulti(len(request.ReqCodes), request.ReqCodes)
	o.InsertMulti(len(request.RespSamples), request.RespSamples)

	//处理请求参数字段说明信息
	if len(request.FieldData) > 0 {
		for i := range request.FieldData {
			dataItem := request.FieldData[i]
			field := dataItem.ReqFiled
			o.Insert(&field)
			if len(dataItem.Items) > 0 {
				for j := range dataItem.Items {
					dataItem.Items[j].FieldId = field.Id
				}
				o.InsertMulti(len(dataItem.Items), dataItem.Items)
			}
		}
	}

	//处理响应结果
	if len(request.RespItems) > 0 {
		saveAllRespParams(o, request.RespItems)
	}

	return 0, nil
}

func saveAllRespParams(o orm.Ormer, items []ResponseParamReq) {
	for _, item := range items {
		if item.ParentId == "respTable" { //第一级
			var param ResponseParam
			param.ParentId = 0
			param.CateId = item.CateId
			param.DataName = item.DataName
			param.DataType = item.DataType
			param.Description = item.Description
			param.Idx = 0
			param.OrderNum = item.OrderNum
			param.Status = 1
			_, err := o.Insert(&param)
			if err == nil {
				//第二级
				if param.DataType == "5" || param.DataType == "6" {
					saveResponseParams(o, items, fmt.Sprintf(item.ParentId+"%d", item.OrderNum), param.Id)
				}

			}
		}
	}
}

func saveResponseParams(o orm.Ormer, items []ResponseParamReq, tableId string, parentID int64) {
	fmt.Printf("******tableId**********" + tableId + "_\n")
	for _, item := range items {
		match, _ := regexp.MatchString(tableId+"_\\d+$", item.ParentId)
		fmt.Printf("******tableId**********" + tableId + "_  match=" + item.ParentId + "\n")
		if match {
			start := len(tableId + "_")
			end := len(item.ParentId)
			idxStr := item.ParentId[start:end]
			idx, _ := strconv.ParseInt(idxStr, 10, 64)
			var param ResponseParam
			param.ParentId = parentID
			param.CateId = item.CateId
			param.DataName = item.DataName
			param.DataType = item.DataType
			param.Description = item.Description
			param.Idx = idx
			param.OrderNum = item.OrderNum
			param.Status = 1
			_, err := o.Insert(&param)
			if err == nil {
				if param.DataType == "5" || param.DataType == "6" {
					saveResponseParams(o, items, fmt.Sprintf(item.ParentId+"%d", item.OrderNum), param.Id)
				}
			}
		}
	}
}

//更新接口信息
func UpdateApiInfo(request SaveApiRequest) error {
	o := orm.NewOrm()
	o.Begin()
	//更新接口主体信息
	_, err := o.Raw("UPDATE t_api_category SET name=?,method=?,request_url=?,remark=? WHERE id=?", request.Name, request.Method, request.RequestUrl, request.Remark, request.Pid).Exec()
	if err != nil {
		o.Rollback()
		return err
	}
	//对要删除的数据时行删除
	oldIds := strings.Split(request.DeleteIds, ",")
	var newIds string
	for _, v := range oldIds {
		if len(v) > 0 && v != "0" {
			if len(newIds) > 0 {
				newIds += ","
			}
			newIds += v
		}
	}
	if len(newIds) > 0 {
		_, err = o.Raw("DELETE FROM t_request_params WHERE id IN(" + newIds + ")").Exec()
		if err != nil {
			o.Rollback()
			return err
		}
	}

	//表单参数要删除的记录
	oldFormDataIds := strings.Split(request.DeleteFormDataIds, ",")
	var newFormDataIds string
	for _, v := range oldFormDataIds {
		if len(v) > 0 && v != "0" {
			if len(newFormDataIds) > 0 {
				newFormDataIds += ","
			}
			newFormDataIds += v
		}
	}
	if len(newFormDataIds) > 0 {
		_, err = o.Raw("DELETE FROM t_form_data WHERE id IN(" + newFormDataIds + ")").Exec()
		if err != nil {
			o.Rollback()
			return err
		}
	}

	oldReqSampleIds := strings.Split(request.DeleteReqSampleIds, ",")
	var newReqSampleIds string
	for _, v := range oldReqSampleIds {
		if len(v) > 0 && v != "0" {
			if len(newReqSampleIds) > 0 {
				newReqSampleIds += ","
			}
			newReqSampleIds += v
		}
	}
	if len(newReqSampleIds) > 0 {
		_, err = o.Raw("DELETE FROM t_request_samples WHERE id IN(" + newReqSampleIds + ")").Exec()
		if err != nil {
			o.Rollback()
			return err
		}
	}

	oldReqCodeIds := strings.Split(request.DeleteReqCodeIds, ",")
	var newReqCodeIds string
	for _, v := range oldReqCodeIds {
		if len(v) > 0 && v != "0" {
			if len(newReqCodeIds) > 0 {
				newReqCodeIds += ","
			}
			newReqCodeIds += v
		}
	}
	if len(newReqCodeIds) > 0 {
		_, err = o.Raw("DELETE FROM t_request_codes WHERE id IN(" + newReqCodeIds + ")").Exec()
		if err != nil {
			o.Rollback()
			return err
		}
	}

	cateID, _ := strconv.ParseInt(request.Pid, 10, 64)
	//更新之前的数据
	var insertObjs []RequestParam
	for _, item := range request.Params {
		if item.Id > 0 {
			_, err = o.Update(&item, "ArgName", "ArgType", "DataType", "IsRequire", "Description", "OrderNum")
		} else {
			item.CateId = cateID
			item.Status = 1
			insertObjs = append(insertObjs, item)
		}
		if err != nil {
			o.Rollback()
			return err
		}
	}

	//更新之前的表单数据
	var insertFormDataObjs []FormData
	for _, item := range request.FormDatas {
		if item.Id > 0 {
			_, err = o.Update(&item, "DataName", "DataType", "IsRequire", "Description", "OrderNum")
		} else {
			item.CateId = cateID
			item.Status = 1
			insertFormDataObjs = append(insertFormDataObjs, item)
		}
		if err != nil {
			o.Rollback()
			return err
		}
	}

	var insertReqSampleObjs []RequestSample
	for _, item := range request.ReqSamples {
		if item.Id > 0 {
			_, err = o.Update(&item, "Title", "Content")
		} else {
			item.CateId = cateID
			insertReqSampleObjs = append(insertReqSampleObjs, item)
		}
		if err != nil {
			o.Rollback()
			return err
		}
	}

	var insertReqCodeObjs []RequestCode
	for _, item := range request.ReqCodes {
		if item.Id > 0 {
			_, err = o.Update(&item, "LangType", "Content")
		} else {
			item.CateId = cateID
			insertReqCodeObjs = append(insertReqCodeObjs, item)
		}
		if err != nil {
			o.Rollback()
			return err
		}
	}

	//更新需要修改的响应实例数据
	var insertRespObjs []ResponseSample
	for _, item := range request.RespSamples {
		if item.Id > 0 {
			_, err = o.Update(&item, "RespType", "Content")
		} else {
			item.CateId = cateID
			insertRespObjs = append(insertRespObjs, item)
		}
		if err != nil {
			o.Rollback()
			return err
		}
	}

	//更新需要修改的响应结果参数,采用先删除后添加的方式
	if len(request.RespItems) > 0 {
		for i := range request.RespItems {
			request.RespItems[i].CateId = cateID
			request.RespItems[i].Status = 1
		}
	}
	_, err = o.Raw("DELETE FROM t_response_params WHERE cate_id=?", cateID).Exec()
	if err != nil {
		o.Rollback()
		return err
	}
	if len(request.RespItems) > 0 {
		saveAllRespParams(o, request.RespItems)
	}
	//添加新的数据
	if len(insertObjs) > 0 {
		_, err = o.InsertMulti(len(insertObjs), insertObjs)
		if err != nil {
			o.Rollback()
			return err
		}
	}

	if len(insertFormDataObjs) > 0 {
		_, err = o.InsertMulti(len(insertFormDataObjs), insertFormDataObjs)
		if err != nil {
			o.Rollback()
			return err
		}
	}

	if len(insertReqSampleObjs) > 0 {
		_, err = o.InsertMulti(len(insertReqSampleObjs), insertReqSampleObjs)
		if err != nil {
			o.Rollback()
			return err
		}
	}

	if len(insertReqCodeObjs) > 0 {
		_, err = o.InsertMulti(len(insertReqCodeObjs), insertReqCodeObjs)
		if err != nil {
			o.Rollback()
			return err
		}
	}

	if len(insertRespObjs) > 0 {
		_, err = o.InsertMulti(len(insertRespObjs), insertRespObjs)
		if err != nil {
			o.Rollback()
			return err
		}
	}

	//处理请求参数说明信息 start
	if len(request.DeleteReqFieldIds) > 1 {
		fieldIDs := strings.Split(request.DeleteReqFieldIds, ",")
		for _, fieldID := range fieldIDs {
			_, err = o.Raw("DELETE FROM t_request_fields WHERE id=?", fieldID).Exec()
			if err != nil {
				o.Rollback()
				return err
			}
			_, err = o.Raw("DELETE FROM t_field_items WHERE field_id=?", fieldID).Exec()
			if err != nil {
				o.Rollback()
				return err
			}
		}
	}

	if len(request.DeleteReqFieldItemIds) > 1 {
		itemIDs := request.DeleteReqFieldItemIds[1 : len(request.DeleteReqFieldItemIds)-1]
		_, err = o.Raw("DELETE FROM t_field_items WHERE id IN(" + itemIDs + ")").Exec()
		if err != nil {
			o.Rollback()
			return err
		}
	}
	if len(request.FieldData) > 0 {
		var insertFieldItems []FieldItem
		for _, dataItem := range request.FieldData {
			fieldItem := dataItem.ReqFiled
			items := dataItem.Items
			if fieldItem.Id > 0 {
				_, err = o.Raw("UPDATE t_request_fields SET title=? WHERE id=?", fieldItem.Title, fieldItem.Id).Exec()
				if err != nil {
					o.Rollback()
					return err
				}
				for _, item := range items {
					if item.Id > 0 { //修改
						_, err = o.Update(&item, "DataName", "DataType", "IsRequire", "Description", "OrderNum")
						if err != nil {
							o.Rollback()
							return err
						}
					} else { //添加
						item.FieldId = fieldItem.Id
						insertFieldItems = append(insertFieldItems, item)
					}
				}
			} else { //添加
				fieldItem.CateId = cateID
				_, err = o.Insert(&fieldItem)
				if err != nil {
					o.Rollback()
					return err
				}
				for _, item := range items {
					item.FieldId = fieldItem.Id
					insertFieldItems = append(insertFieldItems, item)
				}
			}
		}
		if len(insertFieldItems) > 0 {
			_, err = o.InsertMulti(len(insertFieldItems), insertFieldItems)
			if err != nil {
				o.Rollback()
				return err
			}
		}
	}
	//处理请求参数说明信息 end

	o.Commit()

	return nil
}

//更新分类信息
func UpdateCategroy(id int64, orderNum int64, name string) error {
	o := orm.NewOrm()

	var cate ApiCategory
	cate.Id = id
	cate.Name = name
	cate.OrderNum = orderNum

	//Update 默认更新所有的字段，可以更新指定的字段：
	_, err := o.Update(&cate, "Name", "OrderNum")

	return err
}

//删除分类信息
func DeleteCategroy(id, pid, orderNum, nodeLevel int64) (bool, error) {
	var buf bytes.Buffer
	buf.WriteString("DELETE FROM t_api_category")
	buf.WriteString(" WHERE ID =?")
	o := orm.NewOrm()
	o.Begin() //开事物
	_, err := o.Raw(buf.String(), id).Exec()
	if err != nil {
		o.Rollback()
		return false, err
	}
	var buf2 bytes.Buffer
	buf2.WriteString("UPDATE t_api_category SET order_num=order_num-1")
	buf2.WriteString(" WHERE pid = ? AND  order_num>?")
	_, err2 := o.Raw(buf2.String(), pid, orderNum).Exec()
	if err2 != nil {
		o.Rollback()
		return false, err2
	}

	if nodeLevel == 4 {
		delSqls := []string{"DELETE FROM t_request_params WHERE cate_id=?", "DELETE FROM t_form_data WHERE cate_id=?", "DELETE FROM t_request_codes WHERE cate_id=?", "DELETE FROM t_request_samples WHERE cate_id=?", "DELETE FROM t_response_params WHERE cate_id=?", "DELETE FROM t_response_samples WHERE cate_id=?", "DELETE FROM t_field_items WHERE field_id IN (SELECT id FROM  t_request_fields WHERE cate_id=?)", "DELETE FROM t_request_fields WHERE cate_id=?"}
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

//复制api信息
func CopyCategroy(id, pid int64) (int64, error) {
	o := orm.NewOrm()

	var currCate ApiCategory
	err := o.Raw("SELECT * FROM t_api_category WHERE status=1 AND id=?", id).QueryRow(&currCate)
	if err != nil || currCate.Name == "" {
		return 0, fmt.Errorf("未找到对就要复制的信息!")
	}
	var newCate ApiCategory
	newCate.Pid = currCate.Pid
	newCate.Name = currCate.Name + " 副本"
	newCate.Depth = currCate.Depth
	newCate.RequestUrl = currCate.RequestUrl
	newCate.Method = currCate.Method
	newCate.FullPath = currCate.FullPath
	newCate.Status = 1
	newCate.Remark = currCate.Remark
	newCate.OrderNum = currCate.OrderNum + 1

	o.Begin() //开事物
	_, err = o.Insert(&newCate)
	if err != nil {
		o.Rollback()
		return 0, err
	}
	newCateID := newCate.Id
	_, err = o.Raw("UPDATE t_api_category SET order_num=order_num+1 WHERE id<>? AND pid = ? AND  order_num>=?", newCateID, currCate.Pid, newCate.OrderNum).Exec()
	if err != nil {
		o.Rollback()
		return 0, err
	}

	//复制请求参数
	var requestParams []RequestParam
	o.Raw("SELECT * FROM t_request_params WHERE status=1 AND cate_id=? ORDER BY id ASC", id).QueryRows(&requestParams)
	if len(requestParams) > 0 {
		for i := range requestParams {
			requestParams[i].Id = 0
			requestParams[i].CateId = newCateID
			requestParams[i].CreateTime = time.Now()
		}
		_, err = o.InsertMulti(len(requestParams), &requestParams)
		if err != nil {
			o.Rollback()
			return 0, err
		}
	}
	//复制特定参数说明
	var fields []RequestField
	o.Raw("SELECT * FROM t_request_fields WHERE cate_id=? ORDER BY id ASC", id).QueryRows(&fields)
	if len(fields) > 0 {
		var items []FieldItem
		o.Raw("SELECT * FROM t_field_items WHERE field_id IN (SELECT id FROM  t_request_fields WHERE cate_id=?) ORDER BY id ASC", id).QueryRows(&items)
		if len(items) > 0 {
			for _, field := range fields {
				oldFieldID := field.Id
				field.Id = 0
				field.CateId = newCateID
				field.CreateTime = time.Now()
				_, err = o.Insert(&field)
				if err != nil {
					o.Rollback()
					return 0, err
				}
				for i := range items {
					if items[i].FieldId == oldFieldID {
						items[i].Id = 0
						items[i].FieldId = field.Id
						items[i].CreateTime = time.Now()
					}
				}
			}
			_, err = o.InsertMulti(len(items), &items)
			if err != nil {
				o.Rollback()
				return 0, err
			}
		}
	}

	//复制表单构造参数
	var formDatas []FormData
	o.Raw("SELECT * FROM t_form_data WHERE status=1 AND cate_id=? ORDER BY id ASC", id).QueryRows(&formDatas)
	if len(formDatas) > 0 {
		for i := range formDatas {
			formDatas[i].Id = 0
			formDatas[i].CateId = newCateID
			formDatas[i].CreateTime = time.Now()
		}
		_, err = o.InsertMulti(len(formDatas), &formDatas)
		if err != nil {
			o.Rollback()
			return 0, err
		}
	}
	//复制请求实例
	var reqSamples []RequestSample
	o.Raw("SELECT * FROM t_request_samples WHERE cate_id=? ORDER BY id ASC", id).QueryRows(&reqSamples)
	if len(reqSamples) > 0 {
		for i := range reqSamples {
			reqSamples[i].Id = 0
			reqSamples[i].CateId = newCateID
			reqSamples[i].CreateTime = time.Now()
		}
		_, err = o.InsertMulti(len(reqSamples), &reqSamples)
		if err != nil {
			o.Rollback()
			return 0, err
		}
	}

	//复制请求代码
	var reqCodes []RequestCode
	o.Raw("SELECT * FROM t_request_codes WHERE cate_id=? ORDER BY id ASC", id).QueryRows(&reqCodes)
	if len(reqCodes) > 0 {
		for i := range reqCodes {
			reqCodes[i].Id = 0
			reqCodes[i].CateId = newCateID
			reqCodes[i].CreateTime = time.Now()
		}
		_, err = o.InsertMulti(len(reqCodes), &reqCodes)
		if err != nil {
			o.Rollback()
			return 0, err
		}
	}

	//复制响应实例
	var respSamples []ResponseSample
	o.Raw("SELECT * FROM t_response_samples WHERE cate_id=? ORDER BY id ASC", id).QueryRows(&respSamples)
	if len(respSamples) > 0 {
		for i := range respSamples {
			respSamples[i].Id = 0
			respSamples[i].CateId = newCateID
			respSamples[i].CreateTime = time.Now()
		}
		_, err = o.InsertMulti(len(respSamples), &respSamples)
		if err != nil {
			o.Rollback()
			return 0, err
		}
	}

	//复制响应参数说明
	var respParams []ResponseParam
	o.Raw("SELECT * FROM t_response_params WHERE status=1 AND cate_id=? ORDER BY parent_id,idx,order_num ASC", id).QueryRows(&respParams)
	if len(respParams) > 0 {
		for i := range respParams {
			oldID := respParams[i].Id

			respParams[i].Id = 0
			respParams[i].CateId = newCateID
			respParams[i].CreateTime = time.Now()
			_, err = o.Insert(&respParams[i])
			if err != nil {
				o.Rollback()
				return 0, err
			}
			replaceID(respParams[i].Id, oldID, respParams)
		}

	}

	o.Commit() //所有事物统一提交

	return newCateID, nil
}

func replaceID(newID, oldID int64, respParams []ResponseParam) {
	for i := range respParams {
		if respParams[i].ParentId == oldID {
			respParams[i].ParentId = newID
		}
	}
}

//获取接口主题信息
/*func QueryApiTheme(id string) []ApiThemeItem {
	o := orm.NewOrm()
	var maps []orm.Params
	themeItems := make([]ApiThemeItem, 0)
	sql := "select id,pid,name,depth from t_api_category where status=1 and depth in(2,3) order by pid,order_num asc"
	if id != "" {
		sql = "select id,pid,name,depth from t_api_category where status=1 and depth in(2,3) and pid=" + id + " order by pid,order_num asc"
	}
	num, err := o.Raw(sql).Values(&maps)
	if err == nil && num > 0 {
		for _, m := range maps {
			depth, _ := strconv.Atoi(fmt.Sprintf("%v", m["depth"]))
			if depth == 2 || id != "" {
				id, _ := strconv.ParseInt(fmt.Sprintf("%v", m["id"]), 10, 64)
				pid, _ := strconv.ParseInt(fmt.Sprintf("%v", m["pid"]), 10, 64)
				var themItem ApiThemeItem
				themItem.Id = id
				themItem.Pid = pid
				themItem.Title = fmt.Sprintf("%v", m["name"])
				themItem.Items = findApiThemeItem(id, maps)
				themeItems = append(themeItems, themItem)
			}
		}
	}
	return themeItems
}*/
//获取接口主题信息
func QueryApiTheme(id string, lang string) []ApiThemeItem {
	o := orm.NewOrm()
	var maps []orm.Params
	themeItems := make([]ApiThemeItem, 0)
	sqlEnglish := `select id,pid,name,depth from t_api_category where status=1 and lang_type = 2 and depth in(2,3) order by pid,order_num asc`
	sqlEnglishWithId := "select id,pid,name,depth from t_api_category where status=1 and lang_type = 2 and depth in(2,3) and pid=" + id + " order by pid,order_num asc"

	sqlChinese := `select id,pid,name,depth from t_api_category where status=1 and lang_type = 1 and depth in(2,3) order by pid,order_num asc`
	sqlChineseWithId := "select id,pid,name,depth from t_api_category where status=1 and lang_type = 1 and depth in(2,3) and pid=" + id + " order by pid,order_num asc"
	var sql string
	if lang == "zh" {
		sql = sqlChinese
		if id != "" {
			sql = sqlChineseWithId
		}
	} else {
		sql = sqlEnglish
		if id != "" {
			sql = sqlEnglishWithId
		}
	}

	num, err := o.Raw(sql).Values(&maps)
	if err == nil && num > 0 {
		for _, m := range maps {
			depth, _ := strconv.Atoi(fmt.Sprintf("%v", m["depth"]))
			if depth == 2 || id != "" {
				id, _ := strconv.ParseInt(fmt.Sprintf("%v", m["id"]), 10, 64)
				pid, _ := strconv.ParseInt(fmt.Sprintf("%v", m["pid"]), 10, 64)
				var themItem ApiThemeItem
				themItem.Id = id
				themItem.Pid = pid
				themItem.Title = fmt.Sprintf("%v", m["name"])
				themItem.Items = findApiThemeItem(id, maps)
				themeItems = append(themeItems, themItem)
			}
		}
	}
	return themeItems
}
func findApiThemeItem(id int64, maps []orm.Params) []ApiThemeItem {
	items := make([]ApiThemeItem, 0)
	for _, m := range maps {
		pid, _ := strconv.ParseInt(fmt.Sprintf("%v", m["pid"]), 10, 64)
		currID, _ := strconv.ParseInt(fmt.Sprintf("%v", m["id"]), 10, 64)
		if pid == id {
			var item ApiThemeItem
			item.Id = currID
			item.Pid = pid
			item.Title = fmt.Sprintf("%v", m["name"])
			item.Items = make([]ApiThemeItem, 0)
			items = append(items, item)
		}
	}
	return items
}

//初始化模型
func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(ApiCategory), new(RequestParam), new(RequestSample), new(RequestCode), new(ResponseSample), new(FormData), new(ResponseParam), new(RequestField), new(FieldItem))
}
