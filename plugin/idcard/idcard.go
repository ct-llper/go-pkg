package idcard

import (
	"errors"
	"strconv"
	"time"
)

// 身份证解析

type (
	CardStruct struct {
		Province string `json:"province,omitempty"` // 省
		City     string `json:"city,omitempty"`     // 市
		County   string `json:"county,omitempty"`   // 县
		Sex      string `json:"sex,omitempty"`      // 性别
		Age      string `json:"age,omitempty"`      // 年龄
		Birth    string `json:"birth,omitempty"`    // 出生日期
		Zodiac   string `json:"zodiac,omitempty"`   // 生肖
	}

	BirthdayStruct struct {
		Year  string `json:"year"`  // 年份
		Month string `json:"month"` // 月份
		Day   string `json:"day"`   // 日期
	}
)

func ParseIdCard(idCard string) (err error, out *CardStruct) {
	out = &CardStruct{}
	// 检查：长度
	if len(idCard) != 18 {
		return errors.New("身份证长度异常"), out
	}
	// 检查：有效字符
	for i, ch := range idCard {
		if ch < 48 || ch > 57 {
			if i == len(idCard)-1 {
				if ch == 88 {
					continue
				}
			}
			return errors.New(idCard + " 身份证格式有误"), out
		}
	}
	// 校验码验证
	checkCode, err := getCheckCode(idCard)
	if err != nil {
		return err, out
	}
	if string(idCard[17]) != checkCode {
		return errors.New(idCard + " 身份证校验码错误，" + checkCode), out
	}

	// 性别
	// 生日
	// 年龄
	// 出生地
	// 生肖

	return nil, out
}

// 校验码
func getCheckCode(idCard string) (string, error) {
	// 十七位数字本体码加权求和
	// s = Sum(Ai * Wi)；
	// i = 0, ... , 16 ，分别对应身份证的前17位数字；
	// Ai，表示第i位置上的身份证号码数字值；
	// Wi，表示第i位置上的加权因子，分别是 7，9，10，5，8，4，2，1，6，3，7，9，10，5，8，4，2；
	var s int
	yzArr := []int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
	for i := 0; i < 17; i++ {
		ich, err := strconv.ParseInt(string(idCard[i]), 10, 0)
		if err != nil {
			return idCard, errors.New("身份证异常")
		}
		s += int(ich) * yzArr[i]
	}
	// 计算模,两数相除的余数 Y = mod(S, 11)
	y := s % 11
	// 通过模得到对应的校验码
	// Y: 0 1 2 3 4 5 6 7 8 9 10
	// 校验码: 1 0 X 9 8 7 6 5 4 3 2
	checkCodes := []string{"1", "0", "X", "9", "8", "7", "6", "5", "4", "3", "2"}

	return checkCodes[y], nil
}

// GetSex 性别：身份证第17位、能整除-女、否则-男
func GetSex(idCard string) (out string) {

	s := idCard[16:17]
	sInt, _ := strconv.ParseInt(s, 10, 0)

	out = "男"
	if sInt%2 == 0 {
		out = "女"
	}

	return out
}

// GetBirthday 生日：年、月、日
func GetBirthday(idCard string) *BirthdayStruct {
	birthday := idCard[6:14]
	return &BirthdayStruct{
		Year:  birthday[:4],
		Month: birthday[4:6],
		Day:   birthday[6:8],
	}
}

// GetAge 年龄
func GetAge(year int) (age int) {
	if year <= 0 {
		age = -1
	}

	nowYear := time.Now().Year()
	age = nowYear - year
	return
}

// GetAgeReal 年龄：真实年龄
func GetAgeReal(year, month, day string) (age int64) {
	yearInt, _ := strconv.ParseInt(year, 10, 0)
	monthInt, _ := strconv.ParseInt(month, 10, 0)
	dayInt, _ := strconv.ParseInt(day, 10, 0)

	// 当前-年、月、日
	now := time.Now()
	var (
		nowYear  = int64(now.Year())
		nowMonth = int64(now.Month())
		nowDay   = int64(now.Day())
	)

	age = nowYear - yearInt - 1
	if nowMonth > monthInt {
		age = age + 1
		return
	} else if nowMonth < monthInt {
		return age
	}

	if nowDay >= dayInt {
		age = age + 1
	}

	return age
}

// GetZodiac 生肖
func GetZodiac(year string) string {
	zodiac := []string{"子鼠", "丑牛", "寅虎", "卯兔", "辰龙", "巳蛇", "午马", "未羊", "申猴", "酉鸡", "戌狗", "亥猪"}
	// 2020:鼠
	yearInt, _ := strconv.ParseInt(year, 10, 0)
	var diff int64
	if yearInt < 2020 {
		diff = 2020 - yearInt - 1
	} else {
		diff = yearInt - 2020 - 1
	}
	return zodiac[11-(diff%12)]
}

// GetStar 星座
func GetStar(month, day int) (star string) {
	if month <= 0 || month >= 13 {
		star = "-1"
	}
	if day <= 0 || day >= 32 {
		star = "-1"
	}
	if (month == 1 && day >= 20) || (month == 2 && day <= 18) {
		star = "水瓶座"
	}
	if (month == 2 && day >= 19) || (month == 3 && day <= 20) {
		star = "双鱼座"
	}
	if (month == 3 && day >= 21) || (month == 4 && day <= 19) {
		star = "白羊座"
	}
	if (month == 4 && day >= 20) || (month == 5 && day <= 20) {
		star = "金牛座"
	}
	if (month == 5 && day >= 21) || (month == 6 && day <= 21) {
		star = "双子座"
	}
	if (month == 6 && day >= 22) || (month == 7 && day <= 22) {
		star = "巨蟹座"
	}
	if (month == 7 && day >= 23) || (month == 8 && day <= 22) {
		star = "狮子座"
	}
	if (month == 8 && day >= 23) || (month == 9 && day <= 22) {
		star = "处女座"
	}
	if (month == 9 && day >= 23) || (month == 10 && day <= 22) {
		star = "天秤座"
	}
	if (month == 10 && day >= 23) || (month == 11 && day <= 21) {
		star = "天蝎座"
	}
	if (month == 11 && day >= 22) || (month == 12 && day <= 21) {
		star = "射手座"
	}
	if (month == 12 && day >= 22) || (month == 1 && day <= 19) {
		star = "魔蝎座"
	}

	return star
}
