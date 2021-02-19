package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"syscall"
)

func main() {

	cmd := exec.Command("/bin/bash", "-c", "ls -l") //不加第一个第二个参数会报错

	//cmd.Stdout = os.Stdout // cmd.Stdout -> stdout  重定向到标准输出，逐行实时打印
	//cmd.Stderr = os.Stderr // cmd.Stderr -> stderr
	//也可以重定向文件 cmd.Stderr= fd (文件打开的描述符即可)

	stdout, _ := cmd.StdoutPipe() //创建输出管道
	defer stdout.Close()
	if err := cmd.Start(); err != nil {
		log.Fatalf("cmd.Start: %v")
	}

	fmt.Println(cmd.Args) //查看当前执行命令

	cmdPid := cmd.Process.Pid //查看命令pid
	fmt.Println(cmdPid)

	result, _ := ioutil.ReadAll(stdout) // 读取输出结果
	resdata := string(result)
	fmt.Println(resdata)

	var res int
	if err := cmd.Wait(); err != nil {
		if ex, ok := err.(*exec.ExitError); ok {
			fmt.Println("cmd exit status")
			res = ex.Sys().(syscall.WaitStatus).ExitStatus() //获取命令执行返回状态，相当于shell: echo $?
		}
	}

	fmt.Println(res)
}
