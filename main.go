package main

import (
	"fmt"
	"skripsi/cmd"
)

func main() {
	selected, err := cmd.SelectSequence()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(selected)
}
