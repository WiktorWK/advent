package main

import (
	"adventgo/utils"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	input, errI := utils.GetInput(2025, 1)
	if errI != nil {
		log.Fatal(errI)
	}

	input = strings.Trim(input, "\n")

	inputArr := strings.Split(input, "\n")

	pointer, max := 50, 100

	minus, plus := "L", "R"

	counter := 0

	for i := range inputArr {
		num, errPI := strconv.ParseInt(inputArr[i][1:], 0, 0)
		if errPI != nil {
			fmt.Println(errPI)
			return
		}

		if strings.Contains(inputArr[i], minus) {
			diff := pointer - int(num)
			rest := diff % max
			pointer = (rest + max) % max
		}

		if strings.Contains(inputArr[i], plus) {
			sum := pointer + int(num)
			pointer = sum % max
		}

		if pointer == 0 {
			fmt.Println(pointer)
			counter++
		}
	}

	fmt.Printf("\nDay 1 Answer: %v\n", counter)
}
