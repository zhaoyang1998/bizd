package utils

import (
	"bizd/metion/global"
	"strconv"
	"time"
)

func GetNowTime() string {
	timeString := time.Now().Format(global.TimeFormat)
	return timeString
}

func GetCurDayTime() string {
	t := time.Now()
	tm := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	timeString := tm.Format(global.TimeFormat)
	return timeString
}

func GetNextDayTime() string {
	t := time.Now().Add(time.Hour * 24)
	tm := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	timeString := tm.Format(global.TimeFormat)
	return timeString
}

func DateConversionCron(tmpTime time.Time) string {
	var cronTime = strconv.Itoa(tmpTime.Second()) + " " + strconv.Itoa(tmpTime.Minute()) + " " + strconv.Itoa(tmpTime.Hour()) + " " + strconv.Itoa(tmpTime.Day()) + " " + strconv.Itoa(int(tmpTime.Month())) + " *"
	return cronTime
}

// 获取本周一时间

func GetCurWeekStart() time.Time {
	t := time.Now()
	offset := int(time.Monday - t.Weekday())
	if offset > 0 {
		offset = -6
	}
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local).
		AddDate(0, 0, offset)
}

func GetCurWeekStartAndEnd() string {
	t := time.Now()
	offset := int(time.Monday - t.Weekday())
	if offset > 0 {
		offset = -6
	}

	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local).
		AddDate(0, 0, offset).Format(global.TimeDayFormat) + "~" + time.Now().Format(global.TimeDayFormat)
}

func GetPrevWeeksStartAndEnd(num int) string {
	tmp := GetCurWeekStart()
	t := tmp.AddDate(0, 0, -7*num)
	return t.Format(global.TimeDayFormat) + "~" + t.AddDate(0, 0, 6).Format(global.TimeDayFormat)
}
