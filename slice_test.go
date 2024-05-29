package goutil

import (
	"fmt"
	"testing"
)

func TestSliceToMap(t *testing.T) {
	type TmpStruct struct {
		Key   string
		Value any
	}
	var (
		list = []TmpStruct{
			{"abc", 123},
			{"abc2", 1234},
			{"abc3", 12345},
		}
	)
	toMap := SliceToMap(list, func(item TmpStruct) string {
		return item.Key
	})
	fmt.Println(toMap)
}
