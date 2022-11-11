package main

//int sum(int a, int b) { return a+b; };
//int abs(int a){
//	if (a <0) return -a ;
//else return a;
//};
import "C"
import (
	Ca "Workspace/gocache"
	"crypto/md5"
	"fmt"
	"io"
	"time"
)

const (
	name = "Eric"
	_  = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
	x = iota // x == 0
	y = iota // y == 1
	z = iota // z == 2
	w
	DefaultExpiration time.Duration = 0
)

func md5encrypt(){
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
}


func main() {
	/**
		@param:
	 */
	//c := [...]int{4, 5, 6}
	//d := `hello
	//world`
	//fmt.Printf("%v,%v,%s,%d", name, c, d, len(c))
	//
	//slice := []byte{'a', 'b', 'c', 'd'}
	//fmt.Println(len(slice),string(slice))

	//fmt.Println(KB);fmt.Println(MB);fmt.Println(GB);fmt.Println(TB);fmt.Println(PB)
	//md5encrypt()
	tc := Ca.New(DefaultExpiration,1)
	s := struct {
		name string
		age  int
	}{"nike",11}
	tc.Set("a",s,0)
	a ,ok := tc.Get("v")
	if ok{
		fmt.Println(a)
	}else{
		fmt.Println("Key is not Exists!")
	}
	if err:=  tc.Increment("tint8", 2); err != nil{
		e := fmt.Errorf("Error: %s",err)
		fmt.Println(e.Error())
	}
	fmt.Println(C.sum(1,2))
	fmt.Println(C.abs(-2))
}
