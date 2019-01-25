package model

import (
	"fmt"
	"platformApi/src/common"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

//基础数据
type BaseData struct {
	Id            int64     `orm:"auto" from:"id" json:"id"`
	DataName      string    `orm:"size(100)" valid:"Required" form:"dataName" json:"dataName"`
	DataCode      string    `orm:"size(40)" valid:"Required" form:"dataCode" json:"dataCode"`
	DataType      string    `orm:"size(40)" form:"dataType" json:"dataType"`
	DataDesc      string    `orm:"size(100)" form:"dataDesc" json:"dataDesc"`
	Status        int       `orm:"size(11)" json:"status"`
	LangType      int       `orm:"size(2)" json:"langType" description:"语言　1:中文 2:英文"`
	CreateTime    time.Time `orm:"auto_now_add;type(datetime)" json:"createTime"` //创建时间
	UpdateTime    time.Time `orm:"auto_now;type(datetime)" json:"updateTime"`     //修改时间
	CreateTimeFmt string    `orm:"-" json:"createTimeFmt"`
	UpdateTimeFmt string    `orm:"-" json:"updateTimeFmt"`
}

func (u *BaseData) TableName() string {
	return "t_platform_basedata"
}

//添加基础数据
func AddBaseData(baseData *BaseData) error {
	o := orm.NewOrm()

	tempObj := new(BaseData)
	o.Raw("SELECT * FROM t_platform_basedata WHERE data_type=? AND data_code=?", baseData.DataType, baseData.DataCode).QueryRow(tempObj)
	if tempObj != nil && tempObj.DataName != "" {
		return fmt.Errorf("已经存存，不能重复添加！")
	}

	baseData.Status = 1
	_, err := o.Insert(baseData)

	return err
}

//更新基础数据
func UpdateBaseData(baseData *BaseData) error {
	o := orm.NewOrm()
	tempObj := new(BaseData)
	o.Raw("SELECT * FROM t_platform_basedata WHERE data_type=? AND data_code=? AND id<>?", baseData.DataType, baseData.DataCode, baseData.Id).QueryRow(tempObj)
	if tempObj != nil && tempObj.DataName != "" {
		return fmt.Errorf("已经存存，不能将代码修改成［%s］！", baseData.DataCode)
	}
	bData := BaseData{Id: baseData.Id}
	err := o.Read(&bData)
	if err != nil {
		return err
	}
	if bData.DataType == "" {
		_, err = o.Raw("UPDATE t_platform_basedata SET  data_type=? WHERE data_type=?", baseData.DataCode, bData.DataCode).Exec()
		if err != nil {
			return err
		}
	}
	_, err = o.Update(baseData, "DataType", "DataCode", "DataName", "DataDesc", "UpdateTime")
	return err
}

//删除基础数据
func DeleteBaseData(id int64) error {
	o := orm.NewOrm()
	baseData := new(BaseData)
	err := o.Raw("SELECT * FROM t_platform_basedata WHERE id=?", id).QueryRow(baseData)
	if err != nil {
		return err
	}
	beginTx := false
	if baseData.DataType == "" { //如果删除的是大类,需要删除所有子类
		o.Begin()
		_, err = o.Raw("DELETE FROM t_platform_basedata WHERE data_type = ?", baseData.DataCode).Exec()
		if err != nil {
			o.Rollback()
			return err
		}
		beginTx = true
	}
	_, err = o.Delete(baseData)
	if beginTx {
		if err != nil {
			o.Rollback()
		} else {
			o.Commit()
		}
	}
	return err
}

//禁用基础数据
func DisabledBaseData(id int64) error {
	o := orm.NewOrm()
	baseData := new(BaseData)
	err := o.Raw("SELECT * FROM t_platform_basedata WHERE id=?", id).QueryRow(baseData)
	if err != nil {
		return err
	}
	beginTx := false
	if baseData.DataType == "" { //如果禁用的是大类,需要禁用所有子类
		o.Begin()
		_, err = o.Raw("UPDATE t_platform_basedata SET status=2,update_time=now() WHERE data_type = ?", baseData.DataCode).Exec()
		if err != nil {
			o.Rollback()
			return err
		}
		beginTx = true
	}
	baseData.Status = 2
	_, err = o.Update(baseData, "Status", "UpdateTime")
	if beginTx {
		if err != nil {
			o.Rollback()
		} else {
			o.Commit()
		}
	}
	return err
}

//启用基础数据
func EnabledBaseData(id int64) error {
	o := orm.NewOrm()
	baseData := new(BaseData)
	err := o.Raw("SELECT * FROM t_platform_basedata WHERE id=?", id).QueryRow(baseData)
	if err != nil {
		return err
	}
	baseData.Status = 1
	_, err = o.Update(baseData, "Status", "UpdateTime")
	return err
}

//根据id查询详细信息
func GetBaseDataInfo(id string) (*BaseData, error) {
	o := orm.NewOrm()
	baseData := new(BaseData)
	err := o.Raw("SELECT * FROM t_platform_basedata WHERE id = ?", id).QueryRow(baseData)
	if err != nil {
		return nil, err
	}
	return baseData, nil
}

//根据数据类型获取字典列表
func QueryBaseDataByType(dataType, langType string) (map[string]interface{}, error) {
	o := orm.NewOrm()
	res := make(orm.Params)
	var err error
	if dataType == "" || dataType == "0" {
		_, err = o.Raw("SELECT data_code, data_name FROM t_platform_basedata WHERE status=1 AND lang_type='"+langType+"' AND (ISNULL(data_type) OR LENGTH(trim(data_type))<1) ORDER BY data_code ASC").RowsToMap(&res, "data_code", "data_name")
	} else {
		typeStr := ""
		arr := strings.Split(dataType, ",")
		for _, v := range arr {
			if len(typeStr) > 0 {
				typeStr += ","
			}
			typeStr += "'" + v + "'"
		}
		_, err = o.Raw("SELECT data_code, data_name FROM t_platform_basedata WHERE status=1 AND lang_type='"+langType+"' AND data_type IN("+typeStr+") ORDER BY data_type,data_code ASC").RowsToMap(&res, "data_code", "data_name")
	}

	return res, err
}

//根据数据类型获取字典列表
func QueryBaseDataByTypeExt(dataType, langType string) (map[string]([]map[string]string), error) {
	o := orm.NewOrm()
	var lists []orm.ParamsList
	var err error
	if dataType == "" || dataType == "0" {
		_, err = o.Raw("SELECT data_type,data_code, data_name FROM t_platform_basedata WHERE status=1 AND lang_type='"+langType+"' AND (ISNULL(data_type) OR LENGTH(trim(data_type))<1) ORDER BY data_type,data_code ASC").ValuesList(&lists, "data_type", "data_code", "data_name")
	} else {
		typeStr := ""
		arr := strings.Split(dataType, ",")
		for _, v := range arr {
			if len(typeStr) > 0 {
				typeStr += ","
			}
			typeStr += "'" + v + "'"
		}
		_, err = o.Raw("SELECT data_type,data_code, data_name FROM t_platform_basedata WHERE status=1 AND lang_type='"+langType+"' AND data_type IN("+typeStr+") ORDER BY data_type,data_code ASC").ValuesList(&lists, "data_type", "data_code", "data_name")
	}
	if err == nil {
		retMap := make(map[string]([]map[string]string), 0)
		for _, row := range lists {
			dataType := fmt.Sprintf("%v", row[0])
			arr := retMap[dataType]

			item := make(map[string]string, 0)
			item["data_code"] = fmt.Sprintf("%v", row[1])
			item["data_name"] = fmt.Sprintf("%v", row[2])
			arr = append(arr, item)

			retMap[dataType] = arr
		}
		return retMap, nil
	}
	return nil, err
}

func SearchBaseData(pageSize, pageNo int, dataType, search, langType string) map[string]interface{} {
	o := orm.NewOrm()
	qs := o.QueryTable("t_platform_basedata")

	if dataType != "" && search != "" {
		cond := orm.NewCondition()
		cond = cond.And("DataName__icontains", search).Or("DataCode__icontains", search).Or("DataDesc__icontains", search)

		cond2 := orm.NewCondition()
		cond2 = cond2.And("DataType__exact", dataType).AndCond(cond)
		qs = qs.SetCond(cond2)
	} else {
		if dataType != "" {
			cond := orm.NewCondition()
			cond = cond.And("DataType__exact", dataType)
			qs = qs.SetCond(cond)
		} else if search != "" {
			cond := orm.NewCondition()
			cond = cond.And("DataName__icontains", search).Or("DataCode__icontains", search).Or("DataDesc__icontains", search)
			qs = qs.SetCond(cond)
		}
	}
	qs = qs.Filter("LangType", langType)

	resultMap := make(map[string]interface{})
	cnt, err := qs.Count()

	if err == nil && cnt > 0 {
		var baseDatas []BaseData
		_, err := qs.Limit(pageSize, (pageNo-1)*pageSize).OrderBy("DataType", "-UpdateTime").All(&baseDatas)

		if err == nil {
			resultMap["total"] = cnt
			for i, v := range baseDatas {
				v.CreateTimeFmt = common.FormatTime(v.CreateTime)
				v.UpdateTimeFmt = common.FormatTime(v.UpdateTime)
				baseDatas[i] = v
			}
			resultMap["data"] = baseDatas
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
	orm.RegisterModel(new(BaseData))
}
