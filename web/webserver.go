package main

import "fmt"

func main() {
	s := "dd"
	var a int = 10
	b := a
	var ip *int16
	fmt.Printf("%s,%v,%T,%p,%x,%p,%d", s, s, s, &a, &b,&ip,a)
}
