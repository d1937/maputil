package maputil

import (
	"reflect"
)

//获取map中所有key 返回slice
func MapKeys(elements interface{}) []interface{} {
	val := reflect.ValueOf(elements)
	i, keys := 0, make([]interface{}, val.Len())
	switch val.Kind() {
	case reflect.Map:
		for _, v := range val.MapKeys() {
			keys[i] = v
			i++
		}
	}
	return keys
}

// 判断key指定key是否在map中
func MapKeyExists(needle interface{}, haystack interface{}) bool {
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
func InMap(needle interface{}, haystack interface{}) bool {
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
