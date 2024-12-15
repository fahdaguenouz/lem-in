package filehandle

import (
	"fmt"
	"strings"
)

func Filehandler(args []string) {
	filename := args[1]
	if strings.Contains(filename, "/") {
		fmt.Println("error: file name should not contain '/'")
		return
	}
	if !strings.HasSuffix(filename, ".txt") || len(filename) <= 4 {
		fmt.Println("Error: Input file must have a .txt extension")
		return
	}
	data:=ReadFile(filename)
	// var rooms []string
	// var tunnels []string
	// var start, end string
	// var numAnts int
	for _, line:=range data{
		parts := strings.Split(line," ")
		
        if len(parts) == 0 {
			fmt.Println("Error there is a line empty")
            break
        }
		for _, part:=range parts{
			fmt.Println("inside",part)
		} 
		if len(parts)==1 {
			fmt.Println("parts :\n",parts)
		}
		
	}
	


}
