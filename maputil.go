package maputil

func ArrayKeys(elements map[interface{}]interface{}) []interface{} {
	i, keys := 0, make([]interface{}, len(elements))
	for key, _ := range elements {
		keys[i] = key
		i++
	}
	return keys
}
