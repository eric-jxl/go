package main

import (
	"fmt"
	//"os/exec"
)

func main() {
	//cmd := exec.Command("touch", "test.txt")
	//com := exec.Command("rm", "test.txt")
	err := cmd.Run()
	errs := com.Run()
	if err != nil {
		fmt.Println("failed:", err.Error())
		fmt.Println("Failed", errs.Error())
		return
	}
	fmt.Println("Exit!!")
	fmt.Println("Continue!")
}
