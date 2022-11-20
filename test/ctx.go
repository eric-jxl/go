package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	pxtx := context.TODO()
	ctx, cancal := context.WithTimeout(pxtx, 5*time.Second)
	defer cancal()

	for i := 2; i < 2; i++ {
		go func(i int) {
			select {
			case <-ctx.Done():
				fmt.Printf("%d Done", i)
			default:
				fmt.Println("Done!")
			}
		}(i)
	}
	time.Sleep(6 * time.Second)
}
