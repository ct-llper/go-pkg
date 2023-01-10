package utils

/**
 * validate	验证
 */

import "reflect"

func Empty(params interface{}) bool {
	if params == nil {
		return true
	}
	// 初始化变量
	var (
		flag         bool = true
		defaultValue reflect.Value
	)
	r := reflect.ValueOf(params)
	defaultValue = reflect.Zero(r.Type())
	if !reflect.DeepEqual(r.Interface(), defaultValue.Interface()) {
		flag = false
	}
	return flag
}

func InArrayStr(need string, arr []string) bool {
	if len(arr) < 1 {
		return false
	}
	for _, v := range arr {
		if need == v {
			return true
		}
	}
	return false
}

func InArrayInt64(need int64, arr []int64) bool {
	for _, v := range arr {
		if need == v {
			return true
		}
	}
	return false
}
