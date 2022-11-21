package main

import (
	"Workspace/blog/search"
	"log"
	"os"
)

func init() {
	log.SetOutput(os.Stdout)
}
func main() {
	search.Run("president")
}
