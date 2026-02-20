package main

import (
	"adventgo/utils"
	"fmt"
)

func main() {
	fmt.Println("hi")

	input, errI := utils.GetInput(2025, 1)
	if errI != nil {
		fmt.Println(errI.Error())
	}

	fmt.Println(input)

}
