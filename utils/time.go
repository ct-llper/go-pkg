package utils

import "time"

const (
	DefaultDate       = "2006-01-02"          // 时间初始值
	DefaultDateHi     = "2006-01-02 15:04"    // 时间初始值到分
	DefaultDateHis    = "2006-01-02 15:04:05" // 时间初始值到秒
	DefaultDateLayout = "2006-01-02 15:04:05" // 时间到秒模板
)

var TimeUtil = new(timeUtil)

type timeUtil struct{}

// TimeStampFormat int64类型时间戳 转换为 字符串类型
func (u *timeUtil) TimeStampFormat(timeStamp int64, format string) string {
	t := time.Unix(timeStamp, 0)

	return t.Format(format)
}

// StringToTime 字符串格式时间 转换为 时间格式
func (u *timeUtil) StringToTime(in string, format string) (err error, out time.Time) {
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		return err, out
	}

	out, err = time.ParseInLocation(format, in, loc)
	if err != nil {
		return err, out
	}

	return err, out
}

// StringToInt64 字符串时间 转换为 时间戳
func (u *timeUtil) StringToInt64(in string, format string) (err error, out int64) {
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		return err, out
	}

	timeObj, err := time.ParseInLocation(format, in, loc)
	if err != nil {
		return err, out
	}
	out = timeObj.Unix()

	return err, out
}

// DiffDaysBySecond 计算日期相差天数-参数为时间戳时
func (u *timeUtil) DiffDaysBySecond(t1, t2 int64) int {
	time1 := time.Unix(t1, 0)
	time2 := time.Unix(t2, 0)

	// 调用上面的函数
	return u.DiffDays(time1, time2)
}

// DiffDays 计算日期相差天数
func (u *timeUtil) DiffDays(t1, t2 time.Time) int {
	t1 = time.Date(t1.Year(), t1.Month(), t1.Day(), 0, 0, 0, 0, time.Local)
	t2 = time.Date(t2.Year(), t2.Month(), t2.Day(), 0, 0, 0, 0, time.Local)

	return int(t1.Sub(t2).Hours() / 24)
}
