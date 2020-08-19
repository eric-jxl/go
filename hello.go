package main

import "os/exec"

func main() {
	cmd := exec.Command("echo", "hello")
	buf, _ := cmd.Output()
	print(string(buf))

}
