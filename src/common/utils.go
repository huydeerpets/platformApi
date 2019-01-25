package common

import (
	"bytes"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"
	uuid "github.com/satori/go.uuid"
)

const USER_INFO = "userInfo"

const DATETIMEFORMAT = "2006-01-02 15:04"
const DATETIMEFULLFORMAT = "2006-01-02 15:04:05"
const DATEFORMAT = "2006-01-02"
const PLATSIGN = "platsign"

var tenToAny map[int64]string = map[int64]string{0: "0", 1: "1", 2: "2", 3: "3", 4: "4", 5: "5", 6: "6", 7: "7", 8: "8", 9: "9", 10: "a", 11: "b", 12: "c", 13: "d", 14: "e", 15: "f", 16: "g", 17: "h", 18: "i", 19: "j", 20: "k", 21: "l", 22: "m", 23: "n", 24: "o", 25: "p", 26: "q", 27: "r", 28: "s", 29: "t", 30: "u", 31: "v", 32: "w", 33: "x", 34: "y", 35: "z", 36: ":", 37: ";", 38: "<", 39: "=", 40: ">", 41: "?", 42: "@", 43: "[", 44: "]", 45: "^", 46: "_", 47: "{", 48: "|", 49: "}", 50: "A", 51: "B", 52: "C", 53: "D", 54: "E", 55: "F", 56: "G", 57: "H", 58: "I", 59: "J", 60: "K", 61: "L", 62: "M", 63: "N", 64: "O", 65: "P", 66: "Q", 67: "R", 68: "S", 69: "T", 70: "U", 71: "V", 72: "W", 73: "X", 74: "Y", 75: "Z"}

var MultiSignMethod map[int]string = map[int]string{
	1:  "deployeCC",
	2:  "initCC",
	3:  "upgradeCC",
	4:  "publishToken",
	5:  "provideAuthority",
	6:  "confirm",
	7:  "addManager",
	8:  "replaceManager",
	9:  "removeManager",
	10: "setMajorityThreshold",
	11: "publishTokenRequireNum",
	12: "publishCCRequireNum",
	13: "returnGasConfig",
	14: "deleteCC",
	15: "setMasterThreshold",
	16: "updateTokenIcon",
}

type Callback func(result interface{}, err error)

func GetMD5Str(str string) string {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(str))
	cipherStr := md5Ctx.Sum(nil)
	return hex.EncodeToString(cipherStr)
	//return string(cipherStr[:])
}

func GetUUID() string {
	uuid, err := uuid.NewV4()
	if err != nil {
		return ""
	}
	str := uuid.String()
	str = strings.Replace(str, "-", "", -1)
	return str
}

//http get request
func HttpGet(url string, callback Callback) {
	resp, err := http.Get(url)
	if err != nil {
		// handle error
		callback(nil, err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		callback(nil, err)
		return
	}
	if resp.StatusCode == 200 {
		callback(body, err)
	} else {
		callback(body, errors.New("请求失败!"))
	}
}
//udo-api http get async request
func UDOHttpGet(url string, lang string, callback Callback) {
	client := &http.Client{}
	req,err:=http.NewRequest("GET",url,nil)
	if err!=nil{
		// handle error
		callback(nil, err)
		return
	}
	req.Header.Set("language",lang)
	resp, err := client.Do(req)
	if err != nil {
		// handle error
		callback(nil, err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		callback(nil, err)
		return
	}
	if resp.StatusCode == 200 {
		callback(body, err)
	} else {
		callback(body, errors.New("请求失败!"))
	}
}
//http get request
func SyncHttpGet(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		// handle error
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		return nil, err
	}
	if resp.StatusCode == 200 {
		return body, err
	} else {
		return body, errors.New("请求失败!")
	}
}
//udo-api http get sync request
func UDOSyncHttpGet(url string,lang string) ([]byte, error) {
	client := &http.Client{}
	req,err:=http.NewRequest("GET",url,nil)
	if err!=nil{
		// handle error
		return nil, err
	}
	req.Header.Set("language",lang)
	resp, err := client.Do(req)
	if err != nil {
		// handle error
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		return nil, err
	}
	if resp.StatusCode == 200 {
		return body, err
	} else {
		return body, errors.New("请求失败!")
	}
}
//http post request
func HttpPost(url string, data url.Values, callback Callback) {
	//data := url.Values{"apikey": {blockNum}, "mobile": {mobile}, "tpl_value": {tpl_value}}
	resp, err := http.PostForm(url, data)
	if err != nil {
		// handle error
		callback(nil, err)
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		callback(nil, err)
		return
	}

	if resp.StatusCode == 200 {
		callback(body, err)
	} else {
		callback(body, errors.New("请求失败!"))
	}
}
//udo-api http post async request
func UDOHttpPost(url string, lang string, data url.Values, callback Callback) {
	//data := url.Values{"apikey": {blockNum}, "mobile": {mobile}, "tpl_value": {tpl_value}}

	client := &http.Client{}
	req,err:=http.NewRequest("POST",url,strings.NewReader(data.Encode()))
	if err!=nil{
		// handle error
		callback(nil, err)
		return
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("language",lang)
	resp, err := client.Do(req)
	if err != nil {
		// handle error
		callback(nil, err)
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		callback(nil, err)
		return
	}

	if resp.StatusCode == 200 {
		callback(body, err)
	} else {
		callback(body, errors.New("请求失败!"))
	}
}
//获取文件大小
func GetFileSize(url string) int64 {
	var fsize int64

	//创建一个http client
	client := new(http.Client)
	//get方法获取资源
	resp, err := client.Get(url)
	if err != nil {
		return 0
	}
	//读取服务器返回的文件大小
	fsize, err = strconv.ParseInt(resp.Header.Get("Content-Length"), 10, 32)
	if err != nil {
		return 0
	}

	defer resp.Body.Close()

	return fsize
}

//格式化文件大小
func FormatFileSize(fileSize int64) string {
	tempFloat, _ := strconv.ParseFloat(fmt.Sprintf("%d", fileSize), 64)
	if fileSize < 1024 {
		return fmt.Sprintf("%dB", fileSize)
	} else if fileSize < (1024 * 1024) {
		var temp = tempFloat / 1024
		return fmt.Sprintf("%.2fKB", temp)
	} else if fileSize < (1024 * 1024 * 1024) {
		var temp = tempFloat / (1024 * 1024)
		return fmt.Sprintf("%.2fMB", temp)
	} else {
		var temp = tempFloat / (1024 * 1024 * 1024)
		return fmt.Sprintf("%.2fGB", temp)
	}
}

//发送远程推送消息
func SendMessage(url string, reqData []byte) ([]byte, error) {
	reader := bytes.NewReader(reqData)
	appkey := beego.AppConfig.String("jpushappkey")
	secret := beego.AppConfig.String("jpushsecret")
	request, err := http.NewRequest("POST", url, reader)
	if err != nil {
		return nil, err
	}
	input := []byte(appkey + ":" + secret)
	// 演示base64编码
	encodeString := base64.StdEncoding.EncodeToString(input)
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	request.Header.Set("Authorization", "Basic "+encodeString)
	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	//byte数组直接转成string，优化内存
	//str := (*string)(unsafe.Pointer(&respBytes))
	//fmt.Println(*str)
	return respBytes, nil
}

//假定字符串的每节数都在5位以下
func ToNum(a string) string {
	c := strings.Split(a, ".")
	ret := make([]string, len(c))
	r := []string{"", "0", "00", "000", "0000"}
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	for i := 0; i < len(c); i++ {
		len := len(c[i])
		ret = append(ret, r[len]+c[i])
	}
	return strings.Join(ret, "")
}

//验证版本格式是否正确
func IsValidVersion(version string) bool {
	//不能省掉括号()
	r, err := regexp.Compile(`(^\d+\.\d+\.\d+$)|(^\d+$)|(^\d+\.\d+$)`)
	if err != nil {
		return false
	}
	return r.MatchString(version)
}

func CmpVersion(a, b string) int {
	_a := ToNum(a)
	_b := ToNum(b)
	//比较字符串的大小即可, 不能转换成数字进行比较
	if _a == _b {
		return 0
	} else if _a > _b {
		return 1
	} else {
		return -1
	}
}

//精确获取两个大整型数据相除的结果
func BigIntDiv(aV, b *big.Int) string {
	bigA := big.NewInt(0)
	ltZero := false
	if aV.Cmp(big.NewInt(0)) == -1 {
		bigA = big.NewInt(0).Abs(aV)
		ltZero = true
	} else {
		bigA = aV
	}
	ip := big.NewInt(1)
	r := ip.Div(bigA, b)

	ip = big.NewInt(1)
	c := ip.Mul(r, b)

	ip = big.NewInt(1)
	d := ip.Sub(bigA, c)
	e := d.Cmp(big.NewInt(0))
	if e > 0 {
		n := len(b.String()) - len(d.String()) - 1
		var buffer bytes.Buffer
		for i := 0; i < n; i++ {
			buffer.WriteString("0")
		}
		buffer.WriteString(d.String())
		if ltZero {
			return fmt.Sprintf("-%v.%s", r, buffer.String())
		}
		return fmt.Sprintf("%v.%s", r, buffer.String())
	}
	if ltZero {
		return fmt.Sprintf("-%v", r)
	}
	return fmt.Sprintf("%v", r)
}

//根据指定的位数生成数字，如3位，会生成对应的1000
func GetNumberWithDigits(n int) *big.Int {
	if n > 0 {
		var buffer bytes.Buffer //Buffer是一个实现了读写方法的可变大小的字节缓冲
		buffer.WriteString("1")
		for i := 0; i < n; i++ {
			buffer.WriteString("0")
		}
		result := new(big.Int)
		result.SetString(buffer.String(), 10)
		return result
	}
	return big.NewInt(1)
}

func Decimal(value float64) float64 {
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.4f", value), 64)
	return value
}

//将20060102150405格式转换成2006-01-02 15:04:05
func FormatDate(dateStr string) string {
	timeLayout := "20060102150405"                               //转化所需模板
	loc, _ := time.LoadLocation("Local")                         //重要：获取时区
	theTime, _ := time.ParseInLocation(timeLayout, dateStr, loc) //使用模板在对应时区转化为time.time类型

	return theTime.Format(DATETIMEFULLFORMAT)
}

//获取某个月份起始日期和结止日期
func StartAndEndDayOfMonth(yearMonth string) (startDay, endDay string) {
	y, _ := strconv.Atoi(yearMonth[0:4])
	m, _ := strconv.Atoi(yearMonth[5:7])
	m = m + 1
	if m > 12 {
		y++
		m = 1
	}
	mStr := ""
	if m < 10 {
		mStr = fmt.Sprintf("0%d", m)
	} else {
		mStr = fmt.Sprintf("%d", m)
	}
	toBeCharge := fmt.Sprintf("%d-%s-01 00:00:00", y, mStr)                 //待转化为时间戳的字符串 注意 这里的小时和分钟还要秒必须写 因为是跟着模板走的 修改模板的话也可以不写
	loc, _ := time.LoadLocation("Local")                                    //重要：获取时区
	theTime, _ := time.ParseInLocation(DATETIMEFULLFORMAT, toBeCharge, loc) //使用模板在对应时区转化为time.time类型

	year, month, _ := theTime.Date()
	thisMonth := time.Date(year, month, 1, 0, 0, 0, 0, time.Local)
	start := thisMonth.AddDate(0, -1, 0).Format(DATEFORMAT)
	end := thisMonth.AddDate(0, 0, -1).Format(DATEFORMAT)

	return start, end
}

//获取两月份之间的年月
func MonthsBetweenYM(startYM, endYM string) []string {
	startY, _ := strconv.Atoi(startYM[0:4])
	startM, _ := strconv.Atoi(startYM[5:7])

	endY, _ := strconv.Atoi(endYM[0:4])
	endM, _ := strconv.Atoi(endYM[5:7])

	var arr []string
	if startY == endY {
		for i := startM; i <= endM; i++ {
			yearMonth := ""
			if i < 10 {
				yearMonth = fmt.Sprintf("%d-0%d", startY, i)
			} else {
				yearMonth = fmt.Sprintf("%d-%d", startY, i)
			}
			arr = append(arr, yearMonth)
		}
	} else if startY < endY {
		i := startY
		for i <= endY {
			s := 1
			e := 12
			if i == startY {
				s = startM
			}
			if i == endY {
				e = endM
			}
			for j := s; j <= e; j++ {
				yearMonth := ""
				if j < 10 {
					yearMonth = fmt.Sprintf("%d%s%d", i, "-0", j)
				} else {
					yearMonth = fmt.Sprintf("%d-%d", i, j)
				}
				arr = append(arr, yearMonth)
				if j == e {
					i++
				}
			}
		}
	}

	return arr
}

// 10进制转16进制
func DecimalToHex(num *big.Int) string {
	return "0x" + decimalToAny(num, 16)
}

// 10进制转任意进制
func decimalToAny(num *big.Int, n int64) string {
	newNumStr := ""
	var remainder *big.Int
	var remainderString string
	for num.Cmp(big.NewInt(0)) != 0 {
		remainder = big.NewInt(1).Mod(num, big.NewInt(n))
		r76 := remainder.Cmp(big.NewInt(76))
		r9 := remainder.Cmp(big.NewInt(9))
		if r76 == -1 && r9 > 0 {
			remainderString = tenToAny[remainder.Int64()]
		} else {
			remainderString = remainder.String()
		}
		newNumStr = remainderString + newNumStr
		num = big.NewInt(1).Div(num, big.NewInt(n))
	}
	if newNumStr == "" {
		newNumStr = "0"
	}
	return newNumStr
}

func FloatNumber(number, tokenID string, decimalUnits int) string {
	digits := GetNumberWithDigits(decimalUnits)
	bigAmount := new(big.Int)
	bigAmount.SetString(number, 10)

	end := BigIntDiv(bigAmount, digits)
	floatAmount, _ := strconv.ParseFloat(end, 64)
	return fmt.Sprintf("%.4f", floatAmount)
}

// 16进制转10进制
func HexToBigInt(num string) *big.Int {
	if num == "" || num == "0" {
		return big.NewInt(0)
	}
	r := strings.Index(num, "0x")
	if r == 0 {
		num = num[2:]
	}
	return anyToDecimal(num, 16)
}

//map根据value找key
func findkey(in string) int64 {
	var result int64 = -1
	for k, v := range tenToAny {
		if in == v {
			result = k
		}
	}
	return result
}

// 任意进制转10进制
func anyToDecimal(num string, n int64) *big.Int {
	newNum := big.NewInt(0)
	nNum := len(strings.Split(num, "")) - 1
	for _, value := range strings.Split(num, "") {
		tmp := big.NewInt(findkey(value))
		if tmp.Int64() != -1 {
			ip := big.NewInt(1)
			newNum = ip.Mul(tmp, BigIntPow(n, int64(nNum))).Add(ip, newNum)
			nNum = nNum - 1
		} else {
			break
		}
	}
	return newNum
}

/**
* n 如果为16
* m 如果为9
*结果为 16的9次方
 */
func BigIntPow(n, m int64) *big.Int {
	bigSum := big.NewInt(1)
	var i int64
	for i = 0; i < m; i++ {
		bigSum = big.NewInt(1).Mul(bigSum, big.NewInt(n))
	}
	return bigSum
}

//获取source的子串,如果start小于0或者end大于source长度则返回""
//start:开始index，从0开始，包括0
//end:结束index，以end结束，但不包括end
func Substring(source string, start int, end int) string {
	var r = []rune(source)
	length := len(r)

	if start < 0 || end > length || start > end {
		return ""
	}

	if start == 0 && end == length {
		return source
	}

	return string(r[start:end])
}

//查找某个字符串在整个字符串中的位置
func UnicodeIndex(str, substr string) int {
	// 子串在字符串的字节位置
	result := strings.Index(str, substr)
	if result >= 0 {
		// 获得子串之前的字符串并转换成[]byte
		prefix := []byte(str)[0:result]
		// 将子串之前的字符串转换成[]rune
		rs := []rune(string(prefix))
		// 获得子串之前的字符串的长度，便是子串在字符串的字符位置
		result = len(rs)
	}
	return result
}

//将email部分设为*号
func HideEmail(email string) string {
	idx := UnicodeIndex(email, "@")
	if idx > -1 {
		endPrefix := Substring(email, idx, len(email))
		prefix := Substring(email, 0, idx)
		if len(prefix) > 4 {
			return prefix[0:2] + "****" + prefix[len(prefix)-2:] + endPrefix
		} else if len(prefix) == 4 {
			return prefix[0:1] + "****" + prefix[3:] + endPrefix
		}
	}
	return email
}

//将联系电话部分设为*号
func HideTel(tel string) string {
	if len(tel) > 6 {
		return tel[0:3] + "****" + tel[len(tel)-4:]
	} else if len(tel) == 6 {
		return tel[0:2] + "****" + tel[4:]
	}
	return tel
}

//将日期转换成字符串
func FormatTime(t time.Time) string {
	local, _ := time.LoadLocation("Local")
	return t.In(local).Format(DATETIMEFORMAT)
}

//将日期转换成字符串
func FormatFullTime(t time.Time) string {
	local, _ := time.LoadLocation("Local")
	return t.In(local).Format(DATETIMEFULLFORMAT)
}
