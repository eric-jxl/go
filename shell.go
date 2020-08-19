package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("/bin/bash", "-c", `nohup /Applications/WeChat.app/Contents/MacOS/WeChat > /dev/null 2>&1 &Killall Terminal
`)
	err := cmd.Run()
	if err != nil {
		fmt.Println("failed:", err.Error())
		return
	}
	fmt.Println("Continue!")
}
