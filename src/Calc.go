package main

import (
	. "fmt"
	"strconv"
	"strings"
)

const operators = "+-/*"

func testOp(operator string) string {
	if strings.ContainsAny(operator, operators) {
		return operator
	}
	panic("Ошибка! Введен неверный арифметический знак. Допускается ввод только этих арифметических знаков: * / + -")
}

func calculatingArabNumbers(x int, operator string, y int) int {
	resultCalcArab := 0
	if x <= 10 && y <= 10 {
		switch operator {
		case "+":
			resultCalcArab = x + y
		case "-":
			resultCalcArab = x - y
		case "*":
			resultCalcArab = x * y
		case "/":
			resultCalcArab = x / y
		}
		return resultCalcArab
	} else {
		panic("Ошибка! Используются одновременно разные системы счисления (арабские и римские цифры одновременно) ИЛИ А или Б больше 10")
	}

}

func calculatingRomeNumbers(a, operator, b string) string {
	romeNumbers := map[string]int{
		"I":    1,
		"II":   2,
		"III":  3,
		"IV":   4,
		"V":    5,
		"VI":   6,
		"VII":  7,
		"VIII": 8,
		"IX":   9,
		"X":    10,
	}
	x, okX := romeNumbers[a]
	y, okY := romeNumbers[b]
	if okX && okY {
		testCalcRome := calculatingArabNumbers(x, operator, y)
		if testCalcRome < 0 {
			panic("Ошибка! В римской счетной системе нет отрицательных чисел")
		}
		symbols := []string{"C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
		values := []int{100, 90, 50, 40, 10, 9, 5, 4, 1}
		var resultCalcRome strings.Builder
		for i := range symbols {
			for testCalcRome >= values[i] {
				resultCalcRome.WriteString(symbols[i])
				testCalcRome = testCalcRome - values[i]
			}
		}
		return resultCalcRome.String()
	} else {
		panic("Ошибка! Используются одновременно разные системы счисления (арабские и римские цифры одновременно) ИЛИ А или Б больше 10")
	}
}

func calculatingResult(a, operator, b string) string {
	testOp(operator)
	x, errX := strconv.Atoi(a)
	y, errY := strconv.Atoi(b)
	if errX == nil && errY == nil {
		return strconv.Itoa(calculatingArabNumbers(x, operator, y))
	} else {
		return calculatingRomeNumbers(a, operator, b)
	}
}

func main() {
	var a, operator, b string
	Scan(&a, &operator, &b)
	Printf(`Результат вычисления выражения: %s %s %s = %s`, a, operator, b, calculatingResult(a, operator, b))
}
