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

func DateConversionCron(tmpTime time.Time) string {
	var cronTime = strconv.Itoa(tmpTime.Second()) + " " + strconv.Itoa(tmpTime.Minute()) + " " + strconv.Itoa(tmpTime.Hour()) + " " + strconv.Itoa(tmpTime.Day()) + " " + strconv.Itoa(int(tmpTime.Month())) + " *"
	return cronTime
}
