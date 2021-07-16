package maputil

import (
	"reflect"
)

//获取map中所有key 返回slice
func ArrayKeys(elements map[interface{}]interface{}) []interface{} {
	i, keys := 0, make([]interface{}, len(elements))
	for key, _ := range elements {
		keys[i] = key
		i++
	}
	return keys
}

// 判断key指定key是否在map中
func ArrayKeyExists(needle interface{}, haystack interface{}) bool {
	//_, ok := m[key]
	val := reflect.ValueOf(haystack)
	switch val.Kind() {
	case reflect.Map:
		for _, v := range val.MapKeys() {
			if v.Interface() == needle {
				return true
			}
		}
	}
	return false
}

// 判断value值是否在map中
func InArray(needle interface{}, haystack interface{}) bool {
	val := reflect.ValueOf(haystack)
	switch val.Kind() {
	case reflect.Map:
		for _, k := range val.MapKeys() {
			if reflect.DeepEqual(needle, val.MapIndex(k).Interface()) {
				return true
			}
		}

	}
	return false
}
