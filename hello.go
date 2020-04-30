package main

import (
	"os/exec"
)

//func main()  {
//	fmt.Printf("hello%d" ,55)
//}

func main() {
	cmd := exec.Command("echo", "hello")
	buf, _ := cmd.Output()
	print(string(buf))
}
