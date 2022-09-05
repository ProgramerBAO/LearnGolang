package main

import (
	"fmt"
	"os"
)

func findFile() {
	file, err := os.Open("a.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(file.Name(), "opened successfully")
}

func main() {
	/*
		type error interface{
			Error() string
		}
	*/
	findFile()
}
