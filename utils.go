package main

import (
	"fmt"
	"strconv"
	"strings"
)

func isEven(num int) bool {
	return num%2 == 0
}

func isContinue() bool {
	var char string

	fmt.Println(IS_CONTINUE)

	for {
		fmt.Scanln(&char)

		if char == EMPTY_STR {
			continue
		}

		if char == YES || char == NO {
			return char == YES
		}

		fmt.Println(CONTINUE_ERROR)
	}
}

func getArrayFromStr() []float64 {
	fmt.Println(GET_NUMBERS)

	var str string
	fmt.Scan(&str)

	trimStr := strings.Trim(str, " ")
	arrStr := strings.Split(trimStr, ",")

	arr := make([]float64, len(arrStr), cap(arrStr))

	for i, value := range arrStr {
		tmp, err := strconv.ParseFloat(value, 64)

		if err != nil {
			continue
		}

		arr[i] = tmp
	}

	return arr
}
