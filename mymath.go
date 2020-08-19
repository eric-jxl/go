package main

import "fmt"

func math(a,b int) int {
	return a*b
}

func main()  {
	a := 100
	b := 900
	fmt.Println(math(a,b))
}