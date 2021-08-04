/**
 * @Author: long
 * @Description:公共函数类
 * @File:  Common
 * @Version: 1.0.0
 * @Date: 2019-08-15 11:36
 */
package sandan

import (
	"bytes"
	"crypto/md5"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"github.com/asaskevich/govalidator"
	"github.com/duolatech/xhttp"
	"github.com/go-redis/redis"
	"github.com/syyongx/go-wordsfilter"
	"github.com/xeipuuv/gojsonschema"
	"gopkg.in/ffmt.v1"
	"io"
	"io/ioutil"
	"math/rand"
	"mime/multipart"
	"net"
	"net/http"
	"os"
	"path"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
)

//中文字符串截取
func SubString(str string, start int, length int, flag bool) (newStr string) {

	stringRune := []rune(str)
	if len(stringRune) > length {
		newStr = string(stringRune[start:length])
		if flag {
			newStr += "..."
		}
	} else {
		newStr = string(stringRune)
	}
	return
}

//截取字符串（按照字节数）
func SubStringForByte(str string, length int) string {
	lth := len(str)
	if length >= lth {
		length = lth
	}
	var tmp int
	for k, v := range str {
		tmp += len(string(v))
		if tmp > length {
			return str[0:k]
		} else if tmp == length {
			return str[0:tmp]
		}
	}
	return str
}

//判断变量是否为空
func IsEmpty(a interface{}) bool {
	v := reflect.ValueOf(a)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	return v.Interface() == reflect.Zero(v.Type()).Interface()
}

//判断是否是手机号
func IsPhone(phone string) (b bool) {
	if m, _ := regexp.MatchString("^1[0-9]{10,}$", phone); !m {
		return false
	}
	return true
}

//判断是否是邮箱
func IsEmail(email string) (b bool) {
	if m, _ := regexp.MatchString(`\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*`, email); !m {
		return false
	}
	return true
}

//判断元素是否在切片中
func InArray(element int, arr []string) bool {
	str := strconv.Itoa(element)
	for _, v := range arr {
		if str == v {
			return true
		}
	}
	return false
}

//格式化时间,时间戳转日期
func FormatTime(timestamp int, status int) (date string) {
	switch status {
	//1 不含小时
	case 1:
		date = time.Unix(int64(timestamp), 0).Format("2006-01-02")
	//2 含小时
	case 2:
		date = time.Unix(int64(timestamp), 0).Format("2006-01-02 15:04")
	}
	return
}

// 通过map主键唯一的特性过滤重复元素
func RemoveRepByMap(slc []int) []int {
	result := []int{}
	tempMap := map[int]byte{} // 存放不重复主键
	for _, e := range slc {
		l := len(tempMap)
		tempMap[e] = 0
		if len(tempMap) != l { // 加入map后，map长度变化，则元素不重复
			result = append(result, e)
		}
	}
	return result
}

//检查用户权限是否存在
func CheckAuth(element string, auth []string) bool {

	flag := false
	for _, v := range auth {
		if element == v {
			flag = true
		}
	}

	return flag
}

//数字字符串转整型
func StringToInt(str string) (number int) {
	number, _ = strconv.Atoi(str)
	return
}

//字符串转float
func StringToFloat64(str string) float64 {
	a, _ := strconv.ParseFloat(str, 64)
	return a
}

//切片去重
func RemoveRepeatedElement(arr []string) (newArr []string) {
	newArr = make([]string, 0)
	for i := 0; i < len(arr); i++ {
		repeat := false
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				repeat = true
				break
			}
		}
		if !repeat {
			newArr = append(newArr, arr[i])
		}
	}
	return
}

//json解码
func JsonDecodetoMap(jsonStr string) (data map[string]interface{}) {

	data = map[string]interface{}{}
	json.Unmarshal([]byte(jsonStr), &data)

	return
}

//从字符串中去除html标签
func StripTags(str string) string {

	//将HTML标签全转换成小写
	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	str = re.ReplaceAllStringFunc(str, strings.ToLower)
	//去除STYLE
	re, _ = regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	str = re.ReplaceAllString(str, "")
	//去除SCRIPT
	re, _ = regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	str = re.ReplaceAllString(str, "")
	//去除所有尖括号内的HTML代码，并换成换行符
	re, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
	str = re.ReplaceAllString(str, "\n")
	//去除连续的换行符
	re, _ = regexp.Compile("\\s{2,}")
	str = re.ReplaceAllString(str, "\n")

	return strings.TrimSpace(str)
}

// 生成32位MD5
func MD5(text string) string {
	ctx := md5.New()
	ctx.Write([]byte(text))
	return hex.EncodeToString(ctx.Sum(nil))
}

//文件MD5
func FileMd5(file multipart.File) string {
	md5 := md5.New()
	io.Copy(md5, file)
	MD5Str := hex.EncodeToString(md5.Sum(nil))
	return MD5Str[8:24]
}

//随机生成字符串
func GetRandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

//随机生成字符串
func GetRandomUpperString(l int) string {
	//str := "0123456789abcdefghijklmnopqrstuvwxyz"
	str := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

//获取文件扩展名
func GetFileSuffix(file string) (fileSuffix string) {
	filenameWithSuffix := path.Base(file)     //获取文件名带后缀
	fileSuffix = path.Ext(filenameWithSuffix) //获取文件后缀
	return
}

//获取文件名
func GetFilename(file string) (filename, suffix string) {
	filenameWithSuffix := path.Base(file) //获取文件名带后缀
	temp := strings.Split(filenameWithSuffix, ".")

	if len(temp) > 0 {
		return temp[0], temp[1]

	}
	return "", ""
}

//创建目录
func CreateDir(path string) (bool, error) {
	exist, err := PathExists(path)
	if !exist {
		// 创建文件夹
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return false, err
		} else {
			return true, nil
		}
	}
	return false, err
}

// 判断文件夹是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

//读取文件全部内容
func ReadAll(filePth string) ([]byte, error) {
	f, err := os.Open(filePth)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(f)
}

//图片替换
func ReplaceImage(src string) string {

	if !strings.Contains(src, "tmp/") {
		return src
	}
	newstr := strings.Replace(src, "/tmp/", "/normal/", -1)
	dir := path.Dir(newstr)

	CreateDir("storage" + dir)

	_, err := CopyFile("storage"+newstr, "storage"+src)
	if err != nil {
		return ""
	} else {
		return newstr
	}
}

//文件复制
func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()
	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	defer dst.Close()
	return io.Copy(dst, src)
}

//设置连接超时
func TimeoutDialer(cTimeout time.Duration, rwTimeout time.Duration) func(net, addr string) (c net.Conn, err error) {
	return func(netw, addr string) (net.Conn, error) {
		conn, err := net.DialTimeout(netw, addr, cTimeout) //设置建立连接超时
		if err != nil {
			return nil, err
		}
		conn.SetDeadline(time.Now().Add(rwTimeout)) //设置发送接受数据超时
		return conn, nil
	}
}

//http get请求
func HttpRequest(httpurl string, token string) (data map[string]interface{}, err error) {

	h := xhttp.NewHttp()
	h.SetHeader(map[string]string{
		"token": token,
	})
	h.SetTimeout(3, 3)
	respBody, err := h.Get(httpurl).GetContent()
	data = map[string]interface{}{}
	json.Unmarshal([]byte(string(respBody)), &data)

	return data, err

}

//http post请求
func HttpPostRequest(httpurl string, headers map[string]string, data map[string]string) (result map[string]interface{}, err error) {

	h := xhttp.NewHttp()
	h.SetHeader(headers)
	h.SetTimeout(3, 3)
	respBody, err := h.Post(httpurl, data).GetContent()
	result = map[string]interface{}{}
	json.Unmarshal([]byte(string(respBody)), &result)

	return result, err

}

//post 请求
func HttpPost(httpurl string, data map[string]interface{}, headers map[string]interface{}) (result map[string]interface{}, err error) {

	client := &http.Client{}
	info, err := json.Marshal(data)
	req, err := http.NewRequest("POST", httpurl, strings.NewReader(string(info)))
	if err != nil {
		return nil, err
	}
	for k, va := range headers {
		req.Header.Set(k, va.(string))
	}

	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		return nil, err2
	}
	result = map[string]interface{}{}
	err3 := json.Unmarshal([]byte(string(body)), &result)

	return result, err3

}

//生成随机数
func RandInt(min, max int) int {
	if min >= max {
		return 0
	}
	rand.Seed(time.Now().UnixNano())
	randNum := rand.Intn(max-min) + min
	return randNum
}

//获取最大和最小值
func MaxAndMin(a, b int) (max, min int, c bool) {
	if a == b {
		return a, b, false
	} else if a > b {
		return a, b, true
	} else {
		return b, a, true
	}
}

//获取n天前日期
func GetDate(n int) string {

	nTime := time.Now()
	date := nTime.AddDate(0, 0, -n).Format("2006-01-02")

	return date
}


//去除字符中的特殊字符，\n 空格除外
func RegexpStrings(str string) string {
	re, _ := regexp.Compile("\v")
	str = re.ReplaceAllString(str, "")

	re, _ = regexp.Compile("\r")
	str = re.ReplaceAllString(str, "")

	re, _ = regexp.Compile("\f")
	str = re.ReplaceAllString(str, "")

	re, _ = regexp.Compile("\t")
	str = re.ReplaceAllString(str, "")
	return str
}

//记录\n空格所在的位置
func SaveIndex(str string) (index []int) {
	if len(str) == 0 {
		return nil
	}
	a := strings.Split(str, "")
	for _, v := range a {
		switch v {
		case "\n":
			index = append(index, 1) //\n标记1
		case " ":
			index = append(index, 2) //空格标记2
		default:
			index = append(index, 0) //其余标记0
		}
	}
	return
}

//重新将\n空格插入字符串
func SaveStrings(str string, index []int) string {
	str_temp_arr := strings.Split(str, "")
	var str_final string
	var i int = 0
	for k, v := range index {
		switch v {
		case 0:
			str_final = str_final + str_temp_arr[k-i]

		case 1:
			str_final = str_final + "\n"
			i++
		case 2:
			str_final = str_final + " "
			i++
		}
	}
	return str_final
}




//手机号加密
func PhoneSecret(phone string) string {
	if len([]rune(phone)) != 11 {
		return "188****8888"
	}
	return phone[0:3] + "****" + phone[7:11]
}

//群组UID
func GetGroupUid(uid int) string {
	return strconv.Itoa(uid) + strconv.Itoa(RandInt(1000, 9999))
}

//0000补全
func GetFourCount(a int) (b string) {
	alength := len(strconv.Itoa(a))
	switch {
	case alength == 4:
		return strconv.Itoa(a)
	case alength == 3:
		return "0" + strconv.Itoa(a)
	case alength == 2:
		return "00" + strconv.Itoa(a)
	case alength == 1:
		return "000" + strconv.Itoa(a)
	default:
		return "0000"
	}
}

//oss原图得到缩略图url
func ThumbnailUrlForHdUrl(hd_url string) string {
	fileName := strings.Split(hd_url, "/")
	img_len := len(fileName)
	img_name := fileName[img_len-1]
	ok, _ := regexp.MatchString(`\w(?i)(.gif|.jpeg|.png|.jpg|.bmp)`, img_name)
	if !ok {
		return ""
	}
	newName := hd_url + "?x-oss-process=style/thumbnail"
	return newName
}
func ToInterfaceSlice(slice interface{}) []interface{} {
	s := reflect.ValueOf(slice)
	if s.Kind() != reflect.Slice {
		ffmt.Puts("InterfaceSlice() given a non-slice type")
	}

	ret := make([]interface{}, s.Len())

	for i := 0; i < s.Len(); i++ {
		ret[i] = s.Index(i).Interface()
	}

	return ret
}

//点赞等数的处理
func DealCount(count int) string {
	if count >= 0 && count < 1000 {
		return strconv.Itoa(count)
	} else if count >= 1000 && count < 10000 {
		return strconv.FormatFloat(StringToFloat64(strconv.Itoa(count))/StringToFloat64("1000"), 'f', 1, 64) + "k"
	} else if count >= 10000 && count < 100000 {
		return strconv.FormatFloat(StringToFloat64(strconv.Itoa(count))/StringToFloat64("10000"), 'f', 1, 64) + "w"
	} else if count >= 100000 {
		return "10w+"
	}
	return ""
}

//ossurl去掉域名
func GetGroupAvatarOssUrl(regexp_string, oss_url string) string {
	regx := regexp.MustCompile(regexp_string + `/[\w\/\.]+`)
	res := regx.FindAllString(oss_url, -1)
	if len(res) > 0 {
		return res[0]
	}
	return ""
}

//redis存储替换禁词
func KeywordsReplaceRedis(words string, keywords []string) string {
	wf := wordsfilter.New()
	root := wf.Generate(keywords)
	r1 := wf.Replace(words, root)
	return r1
}

//群聊发图片处理url
func GroupMsgImage(url string) (filename string, oss_url string) {
	if len(url) == 0 {
		return "", ""
	}
	fileName := strings.Split(url, "/")
	img_len := len(fileName)
	filename = fileName[img_len-1]
	oss_url = GetGroupAvatarOssUrl("group/images", url)
	return
}

//对字符串切片去重
func RemoveSliceRepeat(des []string) (res []string) {
	temp := map[string]struct{}{}
	for _, v := range des {
		if _, ok := temp[v]; !ok {
			temp[v] = struct{}{}
			res = append(res, v)
		}

	}
	return
}


//获取视频时长、大小
func GetVideoMeta(localpath string) (duration, size int64) { //duration ms
	file, _ := os.Open(localpath)
	duration_temp, _ := GetMP4Duration(file)
	duration = int64(duration_temp) * 1000
	files, _ := os.Stat(localpath)
	size = files.Size()
	return
}

// BoxHeader 信息头
type BoxHeader struct {
	Size       uint32
	FourccType [4]byte
	Size64     uint64
}

// GetMP4Duration 获取视频时长，以秒计
func GetMP4Duration(reader io.ReaderAt) (lengthOfTime uint32, err error) {
	var info = make([]byte, 0x10)
	var boxHeader BoxHeader
	var offset int64 = 0
	// 获取moov结构偏移
	for {
		_, err = reader.ReadAt(info, offset)
		if err != nil {
			return
		}
		boxHeader = getHeaderBoxInfo(info)
		fourccType := getFourccType(boxHeader)
		if fourccType == "moov" {
			break
		}
		// 有一部分mp4 mdat尺寸过大需要特殊处理
		if fourccType == "mdat" {
			if boxHeader.Size == 1 {
				offset += int64(boxHeader.Size64)
				continue
			}
		}
		offset += int64(boxHeader.Size)
	}
	// 获取moov结构开头一部分
	moovStartBytes := make([]byte, 0x100)
	_, err = reader.ReadAt(moovStartBytes, offset)
	if err != nil {
		return
	}
	// 定义timeScale与Duration偏移
	timeScaleOffset := 0x1C
	durationOffest := 0x20
	timeScale := binary.BigEndian.Uint32(moovStartBytes[timeScaleOffset : timeScaleOffset+4])
	Duration := binary.BigEndian.Uint32(moovStartBytes[durationOffest : durationOffest+4])
	lengthOfTime = Duration / timeScale
	return
}

// getHeaderBoxInfo 获取头信息
func getHeaderBoxInfo(data []byte) (boxHeader BoxHeader) {
	buf := bytes.NewBuffer(data)
	binary.Read(buf, binary.BigEndian, &boxHeader)
	return
}

// getFourccType 获取信息头类型
func getFourccType(boxHeader BoxHeader) (fourccType string) {
	fourccType = string(boxHeader.FourccType[:])
	return
}

//jsonschema解析
type Metadata struct {
	Schema      string          `json:"$schema"`
	Type        string          `json:"type"`
	Description string          `json:"description"`
	Id          string          `json:"id"`
	Properties  json.RawMessage `json:"properties"`
}

//json_schema校验
func ValidateJsonSchema(json_schema, value string) (bool, error) {
	var info map[string]interface{}
	json.Unmarshal([]byte(json_schema), &info)

	var value_tem interface{}

	switch info["type"].(string) {
	case "integer":
		ok := CheckStringType(value, "integer")
		if !ok {
			return false, nil
		}
		value_tem = StringToInt(value)
	case "number":
		ok := CheckStringType(value, "number")
		if !ok {
			return false, nil
		}
		value_tem = StringToFloat64(value)
	case "boolean":
		ok := CheckStringType(value, "boolean")
		return ok, nil
	case "array":
		var tmp []interface{}
		json.Unmarshal([]byte(value), &tmp)
		value_tem = tmp
	case "object":
		var tmp map[string]interface{}
		json.Unmarshal([]byte(value), &tmp)
		value_tem = tmp

	default:
		value_tem = value
	}

	v := map[string]interface{}{
		"value": value_tem,
	}
	ffmt.Puts(v)
	value_json, _ := json.Marshal(v)

	m := map[string]interface{}{
		"type": "object",
		"properties": map[string]interface{}{
			"value": info,
		},
	}

	loader := gojsonschema.NewGoLoader(m)
	documentLoader := gojsonschema.NewStringLoader(string(value_json)) // 待校验的json数据

	result, err := gojsonschema.Validate(loader, documentLoader)

	ffmt.Puts(result.Errors())
	if err != nil {
		ffmt.Puts(err)
		return false, err
	}
	return result.Valid(), nil

}

func CheckStringType(str, mode_type string) bool {
	switch mode_type {
	case "integer":
		if str == "0" {
			return true
		}
		a := StringToInt(str)
		if a == 0 {
			return false
		}
		return true

	case "number":
		if str == "0" {
			return true
		}
		a := StringToFloat64(str)
		if a == 0 {
			return false
		}
		return true

	case "boolean":
		if str == "true" || str == "false" {
			return true
		}
		return false
	default:
		return true
	}
}

//获取根域名
func GetUrlDomain(url string) string {
	ok := govalidator.IsURL(url)
	if !ok {
		return ""
	}
	re, _ := regexp.Compile(`([a-zA-Z0-9--]{1,200})\.(ac.cn|bj.cn|sh.cn|tj.cn|cq.cn|he.cn|sn.cn|sx.cn|nm.cn|ln.cn|jl.cn|hl.cn|js.cn|zj.cn|ah.cn|fj.cn|jx.cn|sd.cn|ha.cn|hb.cn|hn.cn|gd.cn|gx.cn|hi.cn|sc.cn|gz.cn|yn.cn|gs.cn|qh.cn|nx.cn|xj.cn|tw.cn|hk.cn|mo.cn|xz.cn|com.cn|net.cn|org.cn|gov.cn|.com.hk|我爱你|在线|中国|网址|网店|中文网|公司|网络|集团|com|cn|cc|org|net|xin|xyz|vip|shop|top|club|wang|fun|info|online|tech|store|site|ltd|ink|biz|group|link|work|pro|mobi|ren|kim|name|tv|red|cool|team|live|pub|company|zone|today|video|art|chat|gold|guru|show|life|love|email|fund|city|plus|design|social|center|world|auto|.rip|.ceo|.sale|.hk|.io|.gg|.tm|.gs|.us)`)
	str := re.FindString(url)
	return str
}

//0000补全
func GetEighteenCount(str string) (b string) {
	alength := len(str)
	switch {
	case alength == 18:
		return str
	case alength == 17:
		return str + "0"
	case alength == 16:
		return str + "00"
	case alength == 15:
		return str + "000"
	default:
		return ""
	}
}

func GetPkcodeByUid(uid string) string {
	if len(uid) != 36 {
		return ""
	}
	tmp := GetPkcodeUid(uid[0:2]) + GetPkcodeUid(uid[2:4]) + GetPkcodeUid(uid[4:6]) + GetPkcodeUid(uid[6:8]) + GetPkcodeUid(uid[8:10]) +
		GetPkcodeUid(uid[10:12]) + GetPkcodeUid(uid[12:14]) + GetPkcodeUid(uid[14:16]) + GetPkcodeUid(uid[16:18]) + GetPkcodeUid(uid[18:20]) +
		GetPkcodeUid(uid[20:22]) + GetPkcodeUid(uid[22:24]) + GetPkcodeUid(uid[24:26]) + GetPkcodeUid(uid[26:28]) + GetPkcodeUid(uid[28:30]) +
		GetPkcodeUid(uid[30:32]) + GetPkcodeUid(uid[32:34]) + GetPkcodeUid(uid[34:36])
	return tmp

}

func GetPkcodeUid(code string) string {
	tmp := map[string]string{
		"00": "🆚", "01": "🌈", "02": "🌏", "03": "🌚", "04": "🌞", "05": "🌟", "06": "🌱", "07": "🌵", "08": "🌸", "09": "🌼",
		"10": "🌿", "11": "🍄", "12": "🍆", "13": "🍋", "14": "🍖", "15": "🍵", "16": "🍺", "17": "🎁", "18": "🎂", "19": "🎃",
		"20": "🎉", "21": "🎓", "22": "🎤", "23": "🎮", "24": "🎹", "25": "🏀", "26": "🏆", "27": "🏥", "28": "🐂", "29": "🐇",
		"30": "🐉", "31": "🐍", "32": "🐎", "33": "🐒", "34": "🐔", "35": "🐛", "36": "🐟", "37": "🐢", "38": "🐮", "39": "🐯",
		"40": "🐰", "41": "🐱", "42": "🐲", "43": "🐴", "44": "🐵", "45": "🐶", "46": "🐷", "47": "🐸", "48": "🐹", "49": "👀",
		"50": "👂", "51": "👃", "52": "👅", "53": "👊", "54": "👋", "55": "👌", "56": "👍", "57": "👏", "58": "👨", "59": "👩",
		"60": "👮", "61": "👴", "62": "👶", "63": "👻", "64": "👽", "65": "👿", "66": "💄", "67": "💉", "68": "💊", "69": "💋",
		"70": "💏", "71": "💔", "72": "💡", "73": "💣", "74": "💤", "75": "💦", "76": "💩", "77": "💪", "78": "💭", "79": "💯",
		"80": "🔔", "81": "🔞", "82": "🔥", "83": "😂", "84": "😈", "85": "😉", "86": "😋", "87": "😎", "88": "😏", "89": "😐",
		"90": "😘", "91": "😝", "92": "😣", "93": "😤", "94": "😭", "95": "😱", "96": "😳", "97": "🙈", "98": "🙉", "99": "🙊",
	}
	return tmp[code]
}

func GetPkcodeAndN(pkcode string) string {
	if len(pkcode) != 72 {
		return ""
	}
	return pkcode[0:24] + "\n" + pkcode[24:48] + "\n" + pkcode[48:72]
}

type PkTeamSortStruct struct {
	Uid   string  `json:"uid"`
	Score float64 `json:"score"`
	Sort  int     `json:"sort"`
}

//获取排序
func PkTeamSort(teamSort []redis.Z, uid string) (list []PkTeamSortStruct, user PkTeamSortStruct, scoreMap map[float64]int) {
	if len(teamSort) == 0 {
		return
	}

	var tmp PkTeamSortStruct
	var tmp_score float64
	var tmp_sort int
	scoreMap = map[float64]int{}
	for k, v := range teamSort {
		tmp.Uid = v.Member.(string)
		tmp.Score = v.Score
		switch k {
		case 0:
			tmp.Sort = 1
			tmp_sort = 1
			scoreMap[v.Score] = tmp.Sort
		default:
			if v.Score == tmp_score { //和上一排名的分数比较
				tmp.Sort = tmp_sort
			} else {
				tmp.Sort = k + 1
				tmp_sort = k + 1
				scoreMap[v.Score] = tmp.Sort
			}
		}
		tmp_score = v.Score
		if v.Member.(string) == uid {
			user = tmp
		}
		list = append(list, tmp)
	}
	return
}

//总星数和段位号、每段位星数映射关系
func GetDivisionByStar(all_star int) (division, star int) {
	switch all_star {
	case 0:
		return 1, 1
	case 1, 2, 3:
		return 1, all_star
	case 4, 5, 6, 7:
		return 2, all_star - 4
	case 8, 9, 10, 11:
		return 3, all_star - 8
	case 12, 13, 14, 15:
		return 4, all_star - 12
	case 16, 17, 18, 19:
		return 5, all_star - 16
	case 20, 21, 22, 23:
		return 6, all_star - 20
	case 24, 25, 26, 27:
		return 7, all_star - 24
	case 28, 29, 30, 31:
		return 8, all_star - 28
	case 32, 33, 34, 35:
		return 9, all_star - 32
	case 36, 37, 38, 39, 40:
		return 10, all_star - 36
	case 41, 42, 43, 44, 45:
		return 11, all_star - 41
	case 46, 47, 48, 49, 50:
		return 12, all_star - 46
	case 51, 52, 53, 54, 55, 56:
		return 13, all_star - 51
	case 57, 58, 59, 60, 61, 62:
		return 14, all_star - 57
	case 63, 64, 65, 66, 67, 68:
		return 15, all_star - 63
	case 69, 70, 71, 72, 73, 74:
		return 16, all_star - 69
	case 75, 76, 77, 78, 79, 80, 81:
		return 17, all_star - 75
	case 82, 83, 84, 85, 86, 87, 88:
		return 18, all_star - 82
	case 89, 90, 91, 92, 93, 94, 95:
		return 19, all_star - 89
	case 96, 97, 98, 99, 100, 101, 102:
		return 20, all_star - 96
	case 103, 104, 105, 106, 107, 108, 109:
		return 21, all_star - 103
	default:
		return 22, all_star - 109

	}
}

//总星数增加（）
func DivisionIncr(star, incr int) int {
	switch star {
	case 3, 7, 11, 15, 19, 23, 27, 31, 35, 40, 45, 50, 56, 62, 68, 74, 81, 88, 95, 102, 109: //段位满级升级多增1星
		return 1
	case 2, 6, 10, 14, 18, 22, 26, 30, 34, 39, 44, 49, 55, 61, 67, 73, 80, 87, 94, 101, 108: //段位满级-1升级多增1星
		if incr == 2 {
			return 2
		}
	default:
		return 0
	}

	return 0
}

//总星数减少(段位0星降级多减1星)
func DivisionDecr(star int) int {
	switch star {
	case 4, 8, 12, 16, 20, 24, 28, 32, 36, 41, 46, 51, 57, 63, 69, 75, 82, 89, 96, 103, 110:
		return -1
	default:
		return 0
	}
}

//段位继承
func DivisionInherit(star int) (int, int) {
	switch star {
	case 0:
		return 1, 1
	case 1, 2, 3:
		return 1, star
	case 4, 5, 6, 7:
		return 2, star
	case 8, 9, 10, 11:
		return 3, star
	case 12, 13, 14, 15:
		return 4, star
	case 16, 17, 18, 19:
		return 5, star
	case 20, 21, 22, 23:
		return 6, star
	case 24:
		return 7, star
	case 25, 26, 27:
		return 7, 25
	case 28, 29, 30, 31:
		return 8, 25
	case 32, 33, 34, 35:
		return 9, 29
	case 36, 37, 38, 39, 40:
		return 10, 33
	case 41, 42, 43, 44, 45:
		return 11, 33
	case 46, 47, 48, 49, 50:
		return 12, 37
	case 51, 52, 53, 54, 55, 56:
		return 13, 42
	case 57, 58, 59, 60, 61, 62:
		return 14, 42
	case 63, 64, 65, 66, 67, 68:
		return 15, 47
	case 69, 70, 71, 72, 73, 74:
		return 16, 47
	case 75, 76, 77, 78, 79, 80, 81:
		return 17, 47
	case 82, 83, 84, 85, 86, 87, 88:
		return 18, 52
	case 89, 90, 91, 92, 93, 94, 95:
		return 19, 52
	case 96, 97, 98, 99, 100, 101, 102:
		return 20, 52
	case 103, 104, 105, 106, 107, 108, 109:
		return 21, 58
	case 110, 111, 112, 113, 114, 115, 116, 117, 118, 119:
		return 22, 64
	case 120, 121, 122, 123, 124, 125, 126, 127, 128, 129:
		return 22, 70
	case 130, 131, 132, 133, 134, 135, 136, 137, 138, 139:
		return 22, 76
	case 140, 141, 142, 143, 144, 145, 146, 147, 148, 149:
		return 22, 83
	default:
		return 22, 97
	}
}


// 将获取的纳秒转为毫秒
func GetUnixNanoToSecondStringTime(timeStng string) string{
	SecondStringTime := strconv.Itoa(StringToInt(timeStng)/1000)
	return SecondStringTime
}

// 获取当前时间戳 秒数
func GetCurentSecondStringTime() string{
	current := strconv.FormatInt(time.Now().UnixNano()/1e6, 10)
	return current

}
