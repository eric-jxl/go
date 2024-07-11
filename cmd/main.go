package main

import "fmt"

type Parent struct {
	Age  int    `json:"age"`
	Name string `json:"name"`
}
type Child struct {
	Parent
	Age int `json:"age" xml:"age" `
}

func (Parent) say(things Child) {
	fmt.Println(things.Parent, things.Age)
}
func main() {
	c := Child{Age: 18, Parent: Parent{
		Age:  20,
		Name: "Eric",
	}}
	c.say(c)
}
