package main

import (
	"github.com/eric-jxl/go/blog/search"
	"log"
	"os"
)

func init() {
	log.SetOutput(os.Stdout)
}
func main() {
	search.Run("president")
}
