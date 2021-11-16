package jtime

import (
	"encoding/json"
	"fmt"
	"testing"
)

type JT struct {
	Tm JsonTime `json:"tm"` //注意大写才能导出
}

func TestJsonTime_UnmarshalJSON(b *testing.T) {

	var t = JT{NewNowJsonTime()}
	fmt.Printf("系统默认的时间格式: %s \n", t.Tm)

	bt, err := json.Marshal(t)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("转成常用时间格式: %s \n", string(bt))
	err = json.Unmarshal(bt, &t)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("转换成系统默认格式: %s \n", t.Tm)
}

func TestZeroTime(t *testing.T) {
	var j JsonTime
	marshal, err := json.Marshal(j)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(marshal))
	var text string
	fmt.Println(text)
}
