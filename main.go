package main

import (
	"fmt"
)

func main() {
	fmt.Println(MAIN_TITLE)

	for {
		action := getAction()
		arr := getArrayFromStr()
		res := calculate(action, arr)

		fmt.Println(res)

		if !isContinue() {
			break
		}
	}
}

func getAction() string {
	fmt.Println(ACTION_TYPE)
	var str string

	for {
		fmt.Scan(&str)

		switch str {
		case AVG:
			return AVG
		case MED:
			return MED
		case SUM:
			return SUM
		default:
			fmt.Println(WRONG_ACTION)
			continue
		}
	}
}

func getAVG(arr []float64) float64 {
	var res float64

	for _, value := range arr {
		res += value
	}

	return res / float64(len(arr))
}

func getMED(arr []float64) float64 {
	if !isEven(len(arr)) {
		mid := len(arr) / 2

		return arr[mid]
	}

	rIndex := len(arr) / 2
	lIndex := rIndex - 1

	return (arr[rIndex] + arr[lIndex]) / 2
}

func getSUM(arr []float64) float64 {
	var res float64

	for _, value := range arr {
		res += value
	}

	return res
}

func calculate(action string, arr []float64) float64 {
	switch action {
	case AVG:
		return getAVG(arr)
	case MED:
		return getMED(arr)
	default:
		return getSUM(arr)
	}
}
