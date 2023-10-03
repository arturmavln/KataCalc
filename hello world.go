package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func romanToArabic(roman string) int {
	switch roman {
	case "I":
		return 1
	case "II":
		return 2
	case "III":
		return 3
	case "IV":
		return 4
	case "V":
		return 5
	case "VI":
		return 6
	case "VII":
		return 7
	case "VIII":
		return 8
	case "IX":
		return 9
	case "X":
		return 10
	default:
		return -1
	}
}

func arabicToRoman(arabic int) string {
	if arabic <= 0 || arabic > 100 {
		fmt.Println("Римские числа не могут быть меньше 1")
	}

	var romanNumerals = []struct {
		Value  int
		Symbol string
	}{
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	roman := ""

	for _, numeral := range romanNumerals {
		for arabic >= numeral.Value {
			roman += numeral.Symbol
			arabic -= numeral.Value
		}
	}

	return roman
}

var result int
var usedRomanToArabic bool

func main() {
	fmt.Println("Калькулятор")
	fmt.Println("Введите арифметическое выражение (например, 2 + 3):")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	elements := strings.Fields(input)

	if len(elements) != 3 {
		fmt.Println("Неверный формат выражения")
		return
	}

	a := elements[0]
	operator := elements[1]
	b := elements[2]

	numA, errA := strconv.Atoi(a)
	numB, errB := strconv.Atoi(b)

	if errA != nil && errB != nil {
		numA = romanToArabic(a)
		numB = romanToArabic(b)
		usedRomanToArabic = true
	} else if errA != nil || errB != nil {
		fmt.Println("Оба числа должны быть в арабской или римской форме записи")
		return
	}

	if numA < 1 || numA > 10 || numB < 1 || numB > 10 {
		fmt.Println("Можно ввести только числа от 1 до 10 включительно")
		return
	}

	switch operator {
	case "+":
		result = numA + numB
	case "-":
		result = numA - numB
	case "*":
		result = numA * numB
	case "/":
		if numB == 0 {
			fmt.Println("делить на 0 нельзя, получится бесконечность")
			return
		}
		result = numA / numB
	default:
		fmt.Println("можно использовать только операторы + - * /")
		return
	}
	if usedRomanToArabic {
		fmt.Println(arabicToRoman(result))
	} else {
		fmt.Println(result)
	}
}
