package main

import (
	"fmt"
	"os"
)

func main() {
	if (len(os.Args)) != 2 {
		fmt.Println("Wrong arguments.")
		os.Exit(1)
	}
	fileName := os.Args[1]
	file, err := os.Open(fileName)

	if err != nil {
		fmt.Println("Error openning file.")
		os.Exit(1)
	}

	b1 := make([]byte, 100000)
	n1, err := file.Read(b1)

	if err != nil {
		fmt.Println("Error reading file.")
		os.Exit(1)
	}

	if n1 >= 0 {
		fileData := string(b1)
		fmt.Println(fileData)
	} else {
		fmt.Println("The file is empty.")
	}
}

// File implements
