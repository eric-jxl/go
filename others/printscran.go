package main

import "fmt"


func main() {
	//要求： 可以从控制台接受用户的信息 【姓名，年龄，薪水，是否通过考试】
	var name string
	var age byte
	var sal float32
	var ispass bool
	fmt.Println("请输入姓名：")
	fmt.Scanln(&name)
	fmt.Println("请输入年龄：")
	fmt.Scanln(&age)
	fmt.Println("请输入薪水：")
	fmt.Scanln(&sal)
	fmt.Println("请输入是否通过：")
	fmt.Scanln(&ispass)
	fmt.Printf(" The name is:%s,age:%d,sal:%f, ispass :%t", name, age, sal, ispass)

}
