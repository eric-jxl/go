package main

import (
	"crypto/md5"
	"fmt"
	"io"
)

const (
	name = "Eric"
)

func main() {
	c := [...]int{4, 5, 6}
	d := `hello
	world`
	fmt.Printf("%v,%v,%s,%d",name,c,d,len(c))
	//err := errors.New("emit macho dwarf: elf header corrupted")
	//if err != nil  {
	//	fmt.Println(err)
	//}

	const (
		x = iota // x == 0
		y = iota // y == 1
		z = iota // z == 2
		w
	)
	fmt.Println(x, y, z, w)
	var arr [10]int // 声明了一个int类型的数组
	arr[0] = 42     // 数组下标是从0开始的
	arr[1] = 13     // 赋值操作
	//c := [...]int{4, 5, 6}
	fmt.Printf("The first element is %d\n", arr[0]) // 获取数据，返回42
	fmt.Printf("The last element is %d\n", arr[9])
	fmt.Println(c)
	slice := []byte{'a', 'b', 'c', 'd'}
	fmt.Println(len(slice))

	//假设用户名 abc，密码 123456
	h := md5.New()
	var pwd string
	fmt.Println("输入密码:")
	fmt.Scanln(&pwd)
	io.WriteString(h, pwd)

	// pwmd5 等于 e10adc3949ba59abbe56e057f20f883e
	pwmd5 := fmt.Sprintf("%x", h.Sum(nil))

	// 指定两个 salt： salt1 = @#$%   salt2 = ^&*()
	salt1 := "@#$%"
	salt2 := "^&*()"

	// salt1 + 用户名 + salt2 + MD5 拼接
	io.WriteString(h, salt1)
	io.WriteString(h, "abc")
	io.WriteString(h, salt2)
	io.WriteString(h, pwmd5)

	last := fmt.Sprintf("%x", h.Sum(nil))
	fmt.Println(last)
	const (
		_  = iota
		KB = 1 << (10 * iota)
		MB = 1 << (10 * iota)
		GB = 1 << (10 * iota)
		TB = 1 << (10 * iota)
		PB = 1 << (10 * iota)
	)
	fmt.Println(KB)
	fmt.Println(MB)
	fmt.Println(GB)
	fmt.Println(TB)
	fmt.Println(PB)

}
