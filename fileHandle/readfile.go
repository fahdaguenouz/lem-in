package filehandle

import (
	"bufio"
	"fmt"
	"os"
)

func ReadFile(filename string)[]string {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return nil
	}
	defer file.Close()

	var data []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("error reading the file: %v\n ", err)
		return nil
	}
	return data

}
