package model

type Result struct {
	Status int64
	Msg    string
	Data   interface{}
}

type CCQueryResult struct {
	Status bool
	Msg    string
	Data   map[string]interface{}
}

type QuestionTypeResult struct {
	Status bool              `json:"status" description:"true 表示成功；false 表示失败"`
	Msg    string            `json:"msg" description:"status为false时的错误信息"`
	Data   map[string]string `json:"data" description:"数据信息"`
}

type ApiThemeResponse struct {
	Status bool           `json:"status" description:"true 表示成功；false 表示失败"`
	Msg    string         `json:"msg" description:"status为false时的错误信息"`
	Data   []ApiThemeItem `json:"data" description:"API主题信息"`
}

type RespParamResult struct {
	Title      string          `json:"title" description:"标题"`
	RespParams []ResponseParam `json:"respParams" description:"响应结果"`
}

type ApiMainInfo struct {
	Id                 int64              `json:"id" description:"接口ID"`
	Name               string             `json:"name" description:"接口名称"`
	RequestUrl         string             `json:"requestUrl" description:"请求url"`
	Method             string             `json:"method" description:"请求方法"`
	Remark             string             `json:"remark" description:"接口说明"`
	RequestParams      []RequestParam     `json:"requestParams" description:"接口请求所需参数"`
	RequestParamsMemos []RequestFieldData `json:"requestParamsMemos" description:"对某一请求参数的说明"`
	FormParams         []FormData         `json:"formParams" description:"测试接口所需的表单参数"`
	RequestSamples     []RequestSample    `json:"requestSamples" description:"请求例子"`
	RequestCodes       []RequestCode      `json:"RequestCodes" description:"请求实例代码"`
	RespParamResults   []RespParamResult  `json:"respParamResults" description:"响应结果"`
	RespSamples        []ResponseSample   `json:"respSamples" description:"响应实例"`
}

type ApiInfoResponse struct {
	Status bool        `json:"status" description:"true 表示成功；false 表示失败"`
	Msg    string      `json:"msg" description:"status为false时的错误信息"`
	Data   ApiMainInfo `json:"data" description:"API详情"`
}

type CommonResult struct {
	Status bool     `json:"status" description:"true 表示成功；false 表示失败"`
	Msg    string   `json:"msg" description:"status为false时的错误信息"`
	Data   []string `json:"data" description:"数据信息"`
}

type JsonResult struct {
	Status bool   `json:"status" description:"true 表示成功；false 表示失败"`
	Msg    string `json:"msg" description:"status为false时的错误信息"`
	Data   string `json:"data" description:"数据信息"`
}

type ReturnGasConfigResponse struct {
	Status bool            `json:"status" description:"true 表示成功；false 表示失败"`
	Msg    string          `json:"msg" description:"status为false时的错误信息"`
	Data   ReturnGasConfig `json:"data" description:"数据信息"`
}

type ReturnGasConfig struct {
	InitReleaseRatio string `json:"initReleaseRatio" description:"立即返还比例（按百分比）"`
	Interval         string `json:"interval" description:"每次释放间隔时间以秒为单位"`
	ReleaseRatio     string `json:"releaseRatio" description:"释放比例（按百分比）"`
}

type SignInfoResult struct {
	Status bool      `json:"status" description:"true 表示成功；false 表示失败"`
	Msg    string    `json:"msg" description:"status为false时的错误信息"`
	Data   *SignData `json:"data" description:"签名数据"`
}

type ContractResult struct {
	Status bool      `json:"status" description:"true 表示成功；false 表示失败"`
	Msg    string    `json:"msg" description:"status为false时的错误信息"`
	Data   *Contract `json:"data" description:"合约信息"`
}

type SignThemeInfo struct {
	MasterAddress    string   `json:"masterAddress" description:"master地址"`
	ManagerAddresses []string `json:"managerAddresses" description:"manager地址"`
	MasterThreshold  int      `json:"masterThreshold" description:"master操作所需签名次数"`
	ManagerThreshold int      `json:"managerThreshold" description:"manager操作所需签名次数"`
	ManagerCount     int      `json:"managerCount" description:"manager个数"`
}

type TokenResult struct {
	Status    bool           `json:"status" description:"true 表示成功；false 表示失败"`
	Msg       string         `json:"msg" description:"status为false时的错误信息"`
	Data      *Token         `json:"data" description:"token信息"`
	SignTheme *SignThemeInfo `json:"signInfo" description:"签名信息"`
}

type SignThemeResult struct {
	Status bool           `json:"status" description:"true 表示成功；false 表示失败"`
	Msg    string         `json:"msg" description:"status为false时的错误信息"`
	Data   *SignThemeInfo `json:"data" description:"签名信息"`
}

type ContractListResult struct {
	Status bool       `json:"status" description:"true 表示成功；false 表示失败"`
	Msg    string     `json:"msg" description:"status为false时的错误信息"`
	Data   []Contract `json:"data" description:"合约列表"`
}

type ManagerCountResult struct {
	Status bool   `json:"status" description:"true 表示成功；false 表示失败"`
	Msg    string `json:"msg" description:"status为false时的错误信息"`
	Count  string `json:"count" description:"manager个数"`
}

type TokenItemListResult struct {
	Status bool        `json:"status" description:"true 表示成功；false 表示失败"`
	Msg    string      `json:"msg" description:"status为false时的错误信息"`
	Data   []TokenItem `json:"data" description:"token列表"`
}

type FaqListResult struct {
	Status bool       `json:"status" description:"true 表示成功；false 表示失败"`
	Msg    string     `json:"msg" description:"status为false时的错误信息"`
	Data   []FaqModel `json:"data" description:"faq列表"`
}

type TechnologyListResult struct {
	Status bool         `json:"status" description:"true 表示成功；false 表示失败"`
	Msg    string       `json:"msg" description:"status为false时的错误信息"`
	Data   []Technology `json:"data" description:"搜索结果"`
}

type FaqResult struct {
	Status bool      `json:"status" description:"true 表示成功；false 表示失败"`
	Msg    string    `json:"msg" description:"status为false时的错误信息"`
	Data   *FaqModel `json:"data" description:"faq详情"`
}

//合约调用返回结果
type CCInvokeResult struct {
	Status          bool   `json:"status" description:"true 表示成功；false 表示失败"`
	Msg             string `json:"msg" description:"status为false时的错误信息"`
	TxId            string `json:"txId" description:"交易ID"`
	TokenId         string `json:"tokenId" description:"token标识"`
	ContractAddress string `json:"contractAddress" description:"合约地址"`
}

type TokenListResult struct {
	Status bool        `json:"status" description:"true 表示成功；false 表示失败"`
	Total  int         `json:"total" description:"总记录数"`
	Msg    string      `json:"msg" description:"status为false时的错误信息"`
	Data   []TokenFlex `json:"data" description:"token列表"`
}

type ChainApplyResult struct {
	Status bool            `json:"status" description:"true 表示成功；false 表示失败"`
	Msg    string          `json:"msg" description:"status为false时的错误信息"`
	Data   *ChainApplyInfo `json:"data" description:"token列表"`
}

type DocInfoResult struct {
	Status bool        `json:"status" description:"true 表示成功；false 表示失败"`
	Msg    string      `json:"msg" description:"status为false时的错误信息"`
	Data   *DocContent `json:"data" description:"文档详情"`
}

type DocThemeListResult struct {
	Status bool           `json:"status" description:"true 表示成功；false 表示失败"`
	Msg    string         `json:"msg" description:"status为false时的错误信息"`
	Data   []DocThemeItem `json:"data" description:"文档主题列表"`
}

type ChainApplyListResult struct {
	Status bool             `json:"status" description:"true 表示成功；false 表示失败"`
	Total  int              `json:"total" description:"总记录数"`
	Msg    string           `json:"msg" description:"status为false时的错误信息"`
	Data   []ChainApplyInfo `json:"data" description:"token列表"`
}

type VersionResult struct {
	Status bool    `json:"status" description:"true 表示成功；false 表示失败"`
	Msg    string  `json:"msg" description:"status为false时的错误信息"`
	Data   Version `json:"data" description:"版本信息"`
}

type RespParamResults []RespParamResult

// Len()方法和Swap()方法不用变化
// 获取此 slice 的长度
func (respRet RespParamResults) Len() int { return len(respRet) }

// 交换数据
func (respRet RespParamResults) Swap(i, j int) { respRet[i], respRet[j] = respRet[j], respRet[i] }

// 根据父ID升序排序
func (respRet RespParamResults) Less(i, j int) bool {
	var start, end int64
	if len(respRet[i].RespParams) > 0 {
		start = respRet[i].RespParams[0].ParentId
	}

	if len(respRet[j].RespParams) > 0 {
		end = respRet[j].RespParams[0].ParentId
	}

	return start < end
}
