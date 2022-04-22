package regexutil

import "regexp"

var RegMap = map[string]string{
	"spaceCheck": "\\s+",
	"en":         "[a-zA-Z0-9_]+",
	"zh":         "[\u4e00-\u9fa5_a-zA-Z0-9]+",
	"email":      "[a-zA-Z0-9_.-]+@[a-zA-Z0-9-]+(\\.[a-zA-Z0-9-]+)*\\.[a-zA-Z0-9]{2,6}",
	"mobile":     "1[3|4|5|7|8][0-9]{9}",
	"tel":        "(0\\d{2,3}-\\d{7,8})(-\\d{1,4})?",
	"url":        "http[s]?://[\\w.]+[\\w/]",
	"ip":         "(\\d{1,2}|1\\d\\d|2[0-4]\\d|25[0-5])\\.(\\d{1,2}|1\\d\\d|2[0-4]\\d|25[0-5])\\.(\\d{1,2}|1\\d\\d|2[0-4]\\d|25[0-5])\\.(\\d{1,2}|1\\d\\d|2[0-4]\\d|25[0-5])",
	"date":       "\\d{4}-\\d{1,2}-\\d{1,2}",
	"time":       "\\d{4}-\\d{1,2}-\\d{1,2}\\s\\d{1,2}:\\d{1,2}:\\d{1,2}",
	"datetime":   "\\d{4}-\\d{1,2}-\\d{1,2}\\s\\d{1,2}:\\d{1,2}:\\d{1,2}",
	"idcard":     "[1-9]\\d{5}[1-9]\\d{3}((0\\d)|(1[0-2]))(([0|1|2]\\d)|3[0-1])\\d{3}([0-9]|X)",
	"zipcode":    "[1-9]\\d{5}",
	"qq":         "[1-9]\\d{4,10}",
	"integer":    "[-\\+]?\\d+",
	"double":     "[-\\+]?\\d+(\\.\\d+)?",
	"english":    "[A-Za-z]+",
	"chinese":    "[\u4e00-\u9fa5]+",
	"username":   "[A-Za-z0-9_\u4e00-\u9fa5]+",
	"md5":        "[a-fA-F0-9]{32}",
	"sha1":       "[a-fA-F0-9]{40}",
	"sha256":     "[a-fA-F0-9]{64}",
	"sha512":     "[a-fA-F0-9]{128}",
	"mac":        "([0-9a-fA-F]{2})(([/\\s:-][0-9a-fA-F]{2}){5})",
}

var (
	// SpaceReg 空格匹配
	SpaceReg = regexp.MustCompile(RegMap["spaceCheck"])
	// ZhReg 字母汉字数字下划线
	ZhReg = regexp.MustCompile(RegMap["zh"])
	// EnReg 英文数字下划线
	EnReg = regexp.MustCompile(RegMap["En"])
	// EmailReg 邮箱匹配
	EmailReg = regexp.MustCompile(RegMap["email"])
	// MobileReg 手机号码匹配
	MobileReg = regexp.MustCompile(RegMap["mobile"])
	// TelReg 固定电话匹配
	TelReg = regexp.MustCompile(RegMap["tel"])
	// UrlReg 网址匹配
	UrlReg = regexp.MustCompile(RegMap["url"])
	// IpReg IP地址匹配
	IpReg = regexp.MustCompile(RegMap["ip"])
	// DateReg 日期匹配
	DateReg = regexp.MustCompile(RegMap["date"])
	// TimeReg 时间匹配
	TimeReg = regexp.MustCompile(RegMap["time"])
	// DatetimeReg 日期时间匹配
	DatetimeReg = regexp.MustCompile(RegMap["datetime"])
	// IdcardReg 身份证匹配
	IdcardReg = regexp.MustCompile(RegMap["idcard"])
	// ZipcodeReg 邮编匹配
	ZipcodeReg = regexp.MustCompile(RegMap["zipcode"])
	// QqReg QQ号码匹配
	QqReg = regexp.MustCompile(RegMap["qq"])
	// IntegerReg 整数匹配
	IntegerReg = regexp.MustCompile(RegMap["integer"])
	// DoubleReg 浮点数匹配
	DoubleReg = regexp.MustCompile(RegMap["double"])
	// EnglishReg 英文字符匹配
	EnglishReg = regexp.MustCompile(RegMap["english"])
	// ChineseReg 中文字符匹配
	ChineseReg = regexp.MustCompile(RegMap["chinese"])
	// UsernameReg 用户名匹配
	UsernameReg = regexp.MustCompile(RegMap["username"])
	// Md5Reg MD5匹配
	Md5Reg = regexp.MustCompile(RegMap["md5"])
	// Sha1Reg sha1匹配
	Sha1Reg = regexp.MustCompile(RegMap["sha1"])
	// Sha256Reg sha256匹配
	Sha256Reg = regexp.MustCompile(RegMap["sha256"])
	// Sha512Reg sha512匹配
	Sha512Reg = regexp.MustCompile(RegMap["sha512"])
	// MacReg mac地址匹配
	MacReg = regexp.MustCompile(RegMap["mac"])
)
