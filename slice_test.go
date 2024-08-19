package goutil

import (
	"fmt"
	"testing"
)

func TestSliceToMap(t *testing.T) {
	var (
		list = []int{1, 2, 3}
	)
	toMap := SliceToMap(list, func(i int) int {
		return i
	})
	fmt.Println(toMap)
}

func TestSliceToField(t *testing.T) {
	// 定义一个示例结构体
	type Person struct {
		Name string
		Age  int
	}
	// 示例数据
	people := []Person{
		{Name: "Alice", Age: 30},
		{Name: "Bob", Age: 25},
		{Name: "Charlie", Age: 35},
	}

	// 提取 Name 字段
	names := SliceToField(people, func(p Person) string {
		return p.Name
	})
	fmt.Println("Names:", names)

	// 提取 Age 字段
	ages := SliceToField(people, func(p Person) int {
		return p.Age
	})
	fmt.Println("Ages:", ages)
}
