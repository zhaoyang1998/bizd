package utils

import "time"

func GetNowTime() string {
	timeString := time.Now().Format("2006-01-02 15:04:05")
	return timeString
}
