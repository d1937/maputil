package maputil

import (
	"reflect"
)

//获取map中所有key 返回slice
func Keys(mp interface{}) (keys []string) {
	rftVal := reflect.ValueOf(mp)
	if rftVal.Type().Kind() == reflect.Ptr {
		rftVal = rftVal.Elem()
	}

	if rftVal.Kind() != reflect.Map {
		return
	}

	for _, key := range rftVal.MapKeys() {
		keys = append(keys, key.String())
	}
	return
}

// 判断指定key是否在map中
func KeyExists(needle interface{}, haystack interface{}) bool {
	//_, ok := m[key]
	rftVal := reflect.ValueOf(haystack)
	if rftVal.Type().Kind() == reflect.Ptr {
		rftVal = rftVal.Elem()
	}
	if rftVal.Kind() != reflect.Map {
		return false
	}
	for _, key := range rftVal.MapKeys() {
		if key.Interface() == needle {
			return true
		}
	}
	return false
}

// 判断value值是否在map中
func InMap(needle interface{}, haystack interface{}) bool {
	val := reflect.ValueOf(haystack)
	if val.Kind() == reflect.Map {
		for _, k := range val.MapKeys() {
			if reflect.DeepEqual(needle, val.MapIndex(k).Interface()) {
				return true
			}
		}
	}
	return false
}

// 获取map全部values
func Values(mp interface{}) (values []interface{}) {
	rftTyp := reflect.TypeOf(mp)
	if rftTyp.Kind() == reflect.Ptr {
		rftTyp = rftTyp.Elem()
	}

	if rftTyp.Kind() != reflect.Map {
		return
	}

	rftVal := reflect.ValueOf(mp)
	for _, key := range rftVal.MapKeys() {
		values = append(values, rftVal.MapIndex(key).Interface())
	}
	return
}
