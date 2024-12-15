package main

import (
	"fmt"
	filehandle "lem-in/fileHandle"
	"os"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("Usage: go run main.go <filename>")
		return
	}
	filehandle.Filehandler(args)
}
