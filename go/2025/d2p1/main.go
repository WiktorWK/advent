package main

import (
	"adventgo/utils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input, errI := utils.GetInput(2025, 2)
	if errI != nil {
		fmt.Println(errI.Error())
	}

	arr := strings.Split(strings.Trim(input, "\n"), ",")

	sum := 0

	for i := range arr {
		r := strings.Split(arr[i], "-")

		a, errA := strconv.ParseInt(r[0], 0, 0)
		if errA != nil {
			fmt.Println(errA)
			return
		}

		b, errB := strconv.ParseInt(r[1], 0, 0)
		if errB != nil {
			fmt.Println(errB)
			return
		}

		for i := a; i <= b; i++ {
			str := strconv.FormatInt(i, 10)

			if len(str)%2 == 0 && strings.Contains(str[len(str)/2:], str[:len(str)/2]) {
				sum += int(i)
			}
		}
	}

	fmt.Printf("\nDay 2 part 1 Answer: %v\n", sum)
}
