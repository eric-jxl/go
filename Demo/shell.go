package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("echo", "Hello World")
	err := cmd.Run()
	if err != nil {
		fmt.Println("failed:", err.Error())
		return
	}
	fmt.Println("Exit!!")
	fmt.Println("Continue!")
}
