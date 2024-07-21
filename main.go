package main

import (
	"fmt"
)

func main() {
	fmt.Println(MAIN_TITLE)

	for {
		var fisrCurrency, secondCurrency string

		fisrCurrency = getCurrencyType(CURRENCY_TYPE_1)
		value := getCurrencyValue()
		secondCurrency = getCurrencyType(CURRENCY_TYPE_2)

		res := convertation(fisrCurrency, secondCurrency, value)

		fmt.Println(res)

		if !isContinue() {
			break
		}
	}
}

func getCurrencyType(text string) string {
	fmt.Println(text)

	var currency string

	for {
		fmt.Scan(&currency)

		if !checkCurrency(currency) {
			fmt.Println(WRONG_CURRENCY)
			continue
		}

		return currency
	}

}

func getCurrencyValue() float64 {
	fmt.Println(CURRENCY_VALUE)

	var value float64

	for {
		fmt.Scan(&value)

		if value == 0 {
			continue
		}

		return value
	}
}

func convertation(fc, sc string, val float64) float64 {
	IS_USD_RUB := (fc == USD && sc == RUB) || (fc == RUB && sc == USD)
	IS_EUR_RUB := (fc == EUR && sc == RUB) || (fc == RUB && sc == EUR)
	IS_USD_EUR := (fc == USD && sc == EUR) || (fc == EUR && sc == USD)

	switch {
	case IS_USD_RUB:
		return val * USD_RUB
	case IS_EUR_RUB:
		return val * EUR_RUB
	case IS_USD_EUR:
		return val * USD_EUR
	default:
		return val
	}
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
