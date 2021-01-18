package main

import (
	_"errors"
	"fmt"
)

const (
	name = "Eric"
)
func main()  {
	//c := [...]int{4, 5, 6}
	//d := `hello
	//world`
	//fmt.Printf("%v,%v,%s",name,c,d)
	//err := errors.New("emit macho dwarf: elf header corrupted")
	//if err != nil  {
	//	fmt.Println(err)
	//}
	//const (
	//	x = iota // x == 0
	//	y = iota // y == 1
	//	z = iota // z == 2
	//	w
	//)
	//fmt.Println( x, y, z,w)
	var arr [10]int  // 声明了一个int类型的数组
	arr[0] = 42      // 数组下标是从0开始的
	arr[1] = 13      // 赋值操作
	c := [...]int{4, 5, 6}
	fmt.Printf("The first element is %d\n", arr[0])  // 获取数据，返回42
	fmt.Printf("The last element is %d\n", arr[9])
	fmt.Println(c)
	slice := []byte {'a', 'b', 'c', 'd'}
	fmt.Println(len(slice))
}
