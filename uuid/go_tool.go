package main

/*
#include <stdint.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
*/
import "C"
import (
	"fmt"

	"github.com/bwmarrin/snowflake"
)

// AddNumbers Exported Go function to C
//
//export AddNumbers
func AddNumbers(num1, num2 int) int {
	return num1 + num2
}

//export Snowflake1
func Snowflake1() *C.char {
	node, err := snowflake.NewNode(1) //❄️雪花算法生成 分布式ID
	if err != nil {
		return C.CString(fmt.Sprintf("%s", err.Error()))
	}
	id := node.Generate().String()
	return C.CString(fmt.Sprintf("id:%s", id))
}

func main() {} // Go's main function can be empty
