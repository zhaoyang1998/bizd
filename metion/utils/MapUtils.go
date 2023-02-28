package utils

import "bizd/metion/global"

func GetValueByKey(key string) string {
	if global.SystemParameters[key] == "" {
		return global.WxUrlDefault
	}
	return global.SystemParameters[key]
}
