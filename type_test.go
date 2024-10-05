package goutil

import (
	"fmt"
	"testing"
)

func TestIntPtr(t *testing.T) {
	type Test struct {
		Num *int `json:"num"`
	}
	data := Test{Num: IntPtr(10)}
	fmt.Println(*data.Num)
}
