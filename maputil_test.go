package maputil

import (
	"fmt"
	"testing"
)

func TestKeys(t *testing.T) {
	m := make(map[string]string)
	m["name"] = "samy"
	m["addr"] = "gz"
	fmt.Println(Keys(m))
}

func TestKeyExists(t *testing.T) {
	m := make(map[string]string)
	m["name"] = "samy"
	m["addr"] = "gz"
	if KeyExists("name1", m) {
		fmt.Println("YES")
	}
}

func TestInMap(t *testing.T) {
	m := make(map[string]string)
	m["name"] = "samy"
	m["addr"] = "gz"

	if InMap("gz", m) {
		fmt.Println("YES")
	}

}
