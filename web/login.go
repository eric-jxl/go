package main

import (
	"fmt"
	//"reflect"
)

type Monster struct {
	Name  string  `json:"name"`
	Age   int     `json:"monster_age"`
	Score float32 `json:"成绩"`
	Sex   string
}

func (s Monster) GetSum(n1, n2 int) int {
	return n1 + n2
}

func (s Monster) Set(name string, age int, score float32, sex string) {
	s.Name = name
	s.Age = age
	s.Score = score
	s.Sex = sex

}
func (s Monster) Print() {
	fmt.Println("---start---")
	fmt.Println(s)
	fmt.Println("---end-----")
}

//func TestStruct() {
//
//	typ := reflect.TypeOf()
//	val := reflect.ValueOf(a) //获取reflect.Type类型
//	kd := val.Kind()          //获取到a对应的类别
//	if kd != reflect.Struct {
//		fmt.Println("expect struct")
//		return
//	}
//	//获取到该结构体有几个字段
//	num := val.NumField()
//	fmt.Printf("该结构体有%d个字段\n", num) //4个
//
//	//遍历结构体的所有字段
//	for i := 0; i < num; i++ {
//		fmt.Printf("Field %d:值=%v\n", i, val.Field(i))
//		//获取到struct标签，需要通过reflect.Type来获取tag标签的值
//		tagVal := typ.Field(i).Tag.Get("json")
//		//如果该字段有tag标签就显示，否则就不显示
//		if tagVal != "" {
//			fmt.Printf("Field %d:tag=%v\n", i, tagVal)
//		}
//	}
//	//获取到该结构体有多少个方法
//	numOfMethod := val.NumMethod()
//	fmt.Printf("struct has %d methods\n", numOfMethod)
//
//	//方法的排序默认是按照函数名的顺序（ASCII码）
//	val.Method(1).Call(nil) //获取到第二个方法，调用它
//
//	//调用结构体的第1个方法，Method(0)
//	var params []reflect.Value
//	params = append(params, reflect.ValueOf(10))
//	params = append(params, reflect.ValueOf(40))
//	//传入的参数是 []reflect.Value，返回 []reflect.Value
//	res := val.Method(0).Call(params)
//	//返回结果，返回的结果是 []reflect.Value
//	fmt.Println("返回的结果 res=", res[0].Int())
//
//}

func main() {
	_ = Monster{
		Name:  "Eric",
		Age:   18,
		Score: 93.9,
		Sex:   "男",
	}
}
