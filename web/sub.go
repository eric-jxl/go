package main

import "fmt"

func main() {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%2d ", j, i, j*i)
		}
		fmt.Println()
	}

	for i := 1; i < 10; i++ {
		for m := 9; m >= i; m-- {
			fmt.Print("  ")
		}
		for j := 1; j <= i; j++ {
			fmt.Print("★")
		}
		fmt.Println()
	}

}
