package utils

import "time"

//带毫秒的日期格式化字符串
const FormatDateStringToMs = "2006-01-02 15:04:05.999999999"

//正常到秒的日期格式化字符串
const FormatDateString = "2006-01-02 15:04:05"

func FormatDate(t time.Time) string {
	return t.Format("2006-01-02")
}

func FormatDataString(t time.Time) string {
	return t.Format(FormatDateString)
}

func FormatDataStringMs(t time.Time) string {
	return t.Format(FormatDateStringToMs)
}

// 获取星座
func GetConstellation(t time.Time) string {
	arr1 := []int{21, 20, 21, 21, 22, 22, 23, 24, 24, 24, 23, 22}
	arr2 := []string{"摩羯座", "水瓶座", "双鱼座", "白羊座", "金牛座", "双子座", "巨蟹座", "狮子座", "处女座", "天秤座", "天蝎座", "射手座", "摩羯座"}
	if t.Day() < arr1[t.Month()-1] {
		return arr2[t.Month()-1]
	} else {
		return arr2[t.Month()]
	}
}

// 获取今天剩余秒数
func GetTodayOverSec() int {
	return 86400 - time.Now().Hour()*60*60 + time.Now().Minute()*60 + time.Now().Second()
}

func NewDateNow() Date {
	return Date{time.Now()}
}
func NewDate(t time.Time) Date {
	return Date{t}
}

func NewDateByString(date string) Date {
	t, _ := time.ParseInLocation("2006-01-02 15:04:05 -0700 MST", date, time.Local)
	return Date{t}
}

func NewDateByStringDate(date string) Date {
	t, _ := time.ParseInLocation("2006-01-02", date, time.Local)
	return Date{t}
}
func NewDateByStringDateTime(date string) Date {
	t, _ := time.ParseInLocation("2006-01-02 15:04:05", date, time.Local)
	return Date{t}
}

type Date struct {
	time time.Time
}

func (this Date) Self() time.Time {
	return this.time
}

func (this Date) AddDay(day int) Date {
	this.time.AddDate(0, 0, day)
	return this
}

func (this Date) Short() string {
	return this.time.Format("2006-01-02")
}
func (this Date) Long() string {
	return this.time.Format("2006-01-02 15:04:05")
}
func (this Date) LongMs() string {
	return this.time.Format("2006-01-02 15:04:05.999999999")
}
