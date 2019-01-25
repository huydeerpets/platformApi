package model

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type Token struct {
	TokenID      string `orm:"pk;size(32);column(token_id)" json:"tokenID" description:"token标识"`
	Name         string `orm:"size(100);column(name)" json:"name" description:"token名称"`
	TokenSymbol  string `orm:"size(40);column(token_symbol)" json:"tokenSymbol" description:"token名称简写"`
	IconUrl      string `orm:"size(200);column(icon_url)" json:"iconUrl" description:"token图标"`
	IsBaseCoin   bool   `orm:"column(is_base_coin)" json:"isBaseCoin" description:"是否主币"`
	DecimalUnits int    `orm:"size(11);column(decimal_units)" json:"decimalUnits" description:"最大小数点位数"`
	TotalNumber  string `orm:"size(100);column(total_number)" json:"totalNumber" description:"发行总量"`
	RestNumber   string `orm:"-" json:"restNumber" description:"余额"`
	IssuePrice   string `orm:"size(80);column(issue_price)" json:"issuePrice" description:"发行价"`
	IssueTime    string `orm:"type(datetime);column(issue_time)" json:"issueTime" description:"发行时间"`
	Status       int    `orm:"size(4);column(status)" json:"status" description:"状态　1、启用　0、禁用"`
	OwnerAddress string `orm:"size(80);column(owner_address)" json:"ownerAddress" description:"token的发行者"`
}

type TokenFlex struct {
	TokenID          string `json:"tokenID" description:"token标识"`
	Name             string `json:"name" description:"token名称"`
	TokenSymbol      string `json:"tokenSymbol" description:"token名称简写"`
	Status           int    `json:"status" description:"token状态　1、启用　0、禁用"`
	HasMutliSign     bool   `json:"hasMutliSign" description:"是否已设置多重签名"`
	ManagerCount     int    `json:"managerCount" description:"manager个数"`
	MasterThreshold  int    `json:"masterThreshold" description:"master相关操作所需的签名次数"`
	ManagerThreshold int    `json:"managerThreshold" description:"manager相关操作所需的签名次数"`
}

type TokenItem struct {
	TokenId     string `json:"tokenId" description:"token标识"`
	Name        string `json:"name" description:"token名称"`
	TokenSymbol string `json:"tokenSymbol" description:"token名称简写"`
}

func (t *Token) TableName() string {
	return "t_token_info"
}

func UpdateTokenIcon(tokenID, iconUrl string) error {
	o := orm.NewOrm()
	_, err := o.Raw("UPDATE t_token_info SET icon_url=? WHERE token_id=?", iconUrl, tokenID).Exec()
	return err
}

func QueryTokenInfo(tokenID string) (*Token, error) {
	o := orm.NewOrm()
	token := new(Token)
	err := o.Raw("SELECT * FROM t_token_info WHERE token_id = ?", tokenID).QueryRow(token)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func QueryTokenDecimalUnits(tokenIds []string) (map[string]interface{}, error) {
	o := orm.NewOrm()
	res := make(orm.Params)
	condition := ""
	for _, v := range tokenIds {
		if len(condition) > 0 {
			condition += " OR "
		}
		condition += "token_id='" + v + "'"
	}
	_, err := o.Raw("SELECT token_id, decimal_units FROM t_token_info WHERE "+condition).RowsToMap(&res, "token_id", "decimal_units")

	return res, err
}

func QueryAllTokens() ([]TokenItem, error) {
	o := orm.NewOrm()
	var items []TokenItem
	_, err := o.Raw("SELECT token_id,name,token_symbol FROM t_token_info WHERE status=1 ORDER BY name ASC").QueryRows(&items)

	return items, err
}

func SearchTokens(pageSize, pageNo int, search string) map[string]interface{} {
	o := orm.NewOrm()
	qs := o.QueryTable("t_token_info")

	if search != "" {
		cond := orm.NewCondition()
		cond = cond.And("TokenID__icontains", search).Or("Name__icontains", search).Or("TokenSymbol__icontains", search).Or("OwnerAddress__icontains", search)
		qs = qs.SetCond(cond)
	}
	tokens := make([]Token, 0)
	resultMap := make(map[string]interface{})
	cnt, err := qs.Count()

	if err == nil && cnt > 0 {

		_, err := qs.OrderBy("-IssueTime").Limit(pageSize, (pageNo-1)*pageSize).All(&tokens)

		if err == nil {
			resultMap["total"] = cnt
			resultMap["data"] = tokens
			return resultMap
		}
	}

	resultMap["total"] = 0
	resultMap["data"] = tokens

	return resultMap
}

//初始化模型
func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(Token))
}
