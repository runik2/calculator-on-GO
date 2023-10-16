package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func isValidRoman(roman string) bool {
	romanPattern := regexp.MustCompile(`^(M{0,3})(CM|CD|D?C{0,3})(XC|XL|L?X{0,3})(IX|IV|V?I{0,3})$`)
	return romanPattern.MatchString(roman)
}

func romanToArabic(roman string) int {
	romanNumerals := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	result := 0
	prevValue := 0

	for i := len(roman) - 1; i >= 0; i-- {
		value := romanNumerals[rune(roman[i])]

		if value < prevValue {
			result -= value
		} else {
			result += value
		}

		prevValue = value
	}

	return result
}

func arabicToRoman(arabic int) string {
	if arabic < 1 || arabic > 3999 {
		return "Недопустимое значение"
	}

	romanNumerals := []struct {
		Value  int
		Symbol string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
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

func evaluateExpression(expression string) (interface{}, error) {
	parts := strings.Split(expression, " ")
	if len(parts) != 3 {
		return nil, fmt.Errorf("Неправильный формат выражения. Используйте формат 'число оператор число'.")
	}

	var firstOperand, secondOperand int
	var err error

	if firstOperand, err = strconv.Atoi(parts[0]); err != nil {
		if !isValidRoman(parts[0]) {
			return nil, fmt.Errorf("Неправильный формат числа")
		}
		firstOperand = romanToArabic(parts[0])
	}

	operator := parts[1]

	if secondOperand, err = strconv.Atoi(parts[2]); err != nil {
		if !isValidRoman(parts[2]) {
			return nil, fmt.Errorf("Неправильный формат числа")
		}
		secondOperand = romanToArabic(parts[2])
	}

	if (firstOperand < 1 || firstOperand > 10) || (secondOperand < 1 || secondOperand > 10) {
		return nil, fmt.Errorf("Число меньше 1 и больше 10, введите другое число")
	}
	switch operator {
	case "+":
		result := firstOperand + secondOperand
		if isValidRoman(parts[0]) && isValidRoman(parts[2]) {
			return arabicToRoman(result), nil
		}
		return result, nil
	case "-":
		result := firstOperand - secondOperand
		if result <= 0 {
			return nil, fmt.Errorf("Арабские числа не могут быть отрицательными или равными нулю")
		}
		if isValidRoman(parts[0]) && isValidRoman(parts[2]) {
			return arabicToRoman(result), nil
		}
		return result, nil
	case "*":
		result := firstOperand * secondOperand
		if isValidRoman(parts[0]) && isValidRoman(parts[2]) {
			return arabicToRoman(result), nil
		}
		return result, nil
	case "/":
		if secondOperand == 0 {
			return nil, fmt.Errorf("Деление на ноль невозможно")
		}
		result := firstOperand / secondOperand
		if isValidRoman(parts[0]) && isValidRoman(parts[2]) {
			return arabicToRoman(result), nil
		}
		return result, nil
	default:
		return nil, fmt.Errorf("Недопустимый оператор. Используйте: +, -, *, /")
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Введите выражение (например, 2 + 3 или II + III): ")
	expression, _ := reader.ReadString('\n')
	expression = strings.TrimSpace(expression)

	if !regexp.MustCompile(`^\d+\s[+\-*/]\s\d+$|^[IVXLCDM]+\s[+\-*/]\s[IVXLCDM]+$`).MatchString(expression) {
		fmt.Println("Неправильный формат ввода.")
		return
	}

	result, err := evaluateExpression(expression)
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
		return
	}

	if num, ok := result.(int); ok {
		fmt.Printf("Результат: %d\n", num)
	} else {
		romanResult := result.(string)
		fmt.Printf("Результат: %s\n", romanResult)
	}
}
