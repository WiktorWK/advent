package main

import (
	"adventgo/utils"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func main() {
	input, errI := utils.GetInput(2025, 2)
	if errI != nil {
		fmt.Println(errI.Error())
	}

	arr := strings.Split(strings.Trim(input, "\n"), ",")

	cache := make([]int, 0)

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

			for j := 1; j <= len(str)/2; j++ {

				check := strings.ReplaceAll(str, str[:j], "")

				if check == "" {
					cache = append(cache, int(i))
				}
			}
		}
	}

	slices.Sort(cache)

	cache = slices.Compact(cache)

	for i := range cache {
		sum += cache[i]
	}

	fmt.Printf("\nDay 2 part 2 Answer: %v\n", sum)
}
