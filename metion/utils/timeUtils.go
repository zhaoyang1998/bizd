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
