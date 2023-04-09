package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	TryCatch(func() {
		result, err := divide(10, 0)
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("The result is %d\n", result)
	}, func(err error) {
		// Handle the error here
		fmt.Printf("An error occurred: %s\n", err.Error())
	})
	atom()

}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

//TryCatch 捕获异常
func TryCatch(try func(), catch func(error)) {
	defer func() {
		if err := recover(); err != nil {
			catch(fmt.Errorf("%s", err))
		}
	}()

	try()
}

func atom() {
	var count int32 = 0
	for i := 0; i < 100; i++ {
		go func() {
			atomic.AddInt32(&count, 1)
		}()
	}
	// Wait for all goroutines to finish
	time.Sleep(time.Second)
	fmt.Println(count) // Output: 100
}
