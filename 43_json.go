// json
package main

import (
	"encoding/json"
	"fmt"
)

type Astudent struct {
	Name string `json:"name"` //序列化指定key	-- 反射机制
	Age  int `json:"age"`
}

func main() {

	serialize()
}

func serialize() {

	// 使用json.Marshal
	var stu = Astudent{
		Name: "justin",
		Age:  24,
	}

	bytes, e := json.Marshal(&stu) // 注意序列化的跨包问题，结构体的私有类型需要注意
	if e != nil {
		fmt.Println("serialize error:", e)
	}
	fmt.Println("serialize result:", string(bytes))
}

func deSerialize() {

}
