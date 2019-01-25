package model

//contract info
type Contract struct {
	Name            string `json:"name" description:"合约名称"`
	ContractAddress string `json:"contractAddress" description:"合约地址"`
	ContractSymbol  string `json:"contractSymbol" description:"合约名称简写"`
	MAddress        string `json:"mAddress" description:"钱包地址"`
	Version         string `json:"version" description:"版本号"`
	CcPath          string `json:"ccPath" description:"合约路径"`
	Remark          string `json:"remark" description:"合约简介"`
	Status          string `json:"status" description:"合约状态 -1、已删除 1、待初始化 2、正在运行 3、余额不足 4、合约已禁用 5、已弃用"`
	CreateTime      string `json:"createTime" description:"合约发布时间"`
	UpdateTime      string `json:"updateTime" description:"合约更新时间"`
}
