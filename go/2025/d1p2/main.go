package main

import (
	"adventgo/utils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input, errI := utils.GetInput(2025, 1)
	if errI != nil {
		fmt.Println(errI.Error())
		return
	}

	iarr := strings.Split(strings.Trim(input, "\n"), "\n")

	pointer, max, counter := 50, 100, 0

	minus, plus := "L", "R"

	for i := range iarr {
		num, errPI := strconv.ParseInt(iarr[i][1:], 0, 0)
		if errPI != nil {
			fmt.Println(errPI)
			return
		}

		if int(num) >= max {
			counter += int(num) / max
		}

		num %= int64(max)

		if strings.Contains(iarr[i], minus) {
			if pointer > 0 && pointer < max && pointer < int(num) {
				counter++
			}
			pointer -= int(num)
		}

		if strings.Contains(iarr[i], plus) {
			if pointer < 0 && pointer > -max && -pointer < int(num) {
				counter++
			}
			pointer += int(num)
		}

		if pointer < -max || pointer > max {
			counter++
		}

		pointer %= max

		if pointer == 0 {
			counter++
		}

	}

	fmt.Printf("\nDay 1 part 2 Answer: %v\n", counter)
}
