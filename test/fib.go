package main

import "C"
import (
	"fmt"
	"time"
)

func Fib(n int) int {
	if n == 1 || n == 2{
		return 1
	} else {
		return Fib(n - 1) + Fib(n - 2)
	}
}

func main() {
	start := time.Now()
	result := Fib(40)
	end := time.Now()
	time_diff :=  end.Sub(start)
	fmt.Println("斐切拉波数列第40项的值为：",result,"用时: ",time_diff)
}
