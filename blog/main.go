package main

import (
	"log"
	"os"

	"github.com/eric-jxl/go/blog/search"
)

func init() {
	log.SetOutput(os.Stdout)
}
func main() {
	search.Run("president")
}
