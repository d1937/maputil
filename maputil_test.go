package maputil

import (
	"fmt"
	"testing"
)

func TestInArray(t *testing.T) {
	m := make(map[string]string)
	m["name"] = "samy"

	if InArray("samy", m) {
		fmt.Println("YES")
	}
}

func TestArrayKeyExists(t *testing.T) {
	m := make(map[string]string)
	m["name"] = "samy"
	if ok := ArrayKeyExists(11, m); ok {
		fmt.Println("YES")
	}
}
