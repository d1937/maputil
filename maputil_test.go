package maputil

import (
	"fmt"
	"testing"
)

func TestInArray(t *testing.T) {
	m := make(map[string]string)
	m["name"] = "samy"

	if InMap("samy", m) {
		fmt.Println("YES")
	}
}

func TestArrayKeyExists(t *testing.T) {
	m := make(map[string]string)
	m["name"] = "samy"
	if ok := MapKeyExists("name", m); ok {
		fmt.Println("YES")
	}
}

func TestArrayKeys(t *testing.T) {
	m := make(map[string]string)
	m["a"] = "15"
	m["b"] = "samy"
	fmt.Println(MapKeys(m))
}
