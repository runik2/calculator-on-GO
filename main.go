package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Функция для преобразования римских чисел в арабские
func romanToArabic(roman string) int {
	romanNumerals := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
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

// Функция для преобразования арабских чисел в римские
func arabicToRoman(arabic int) string {
	if arabic < 1 || arabic > 3999 {
		return "Недопустимое значение"
	}

	romanNumerals := []string{"X", "IX", "V", "IV", "I"}
	arabicValues := []int{10, 9, 5, 4, 1}

	roman := ""

	for i := 0; i < len(arabicValues); i++ {
		for arabic >= arabicValues[i] {
			roman += romanNumerals[i]
			arabic -= arabicValues[i]
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
		firstOperand = romanToArabic(parts[0])
	}

	operator := parts[1]

	if secondOperand, err = strconv.Atoi(parts[2]); err != nil {
		secondOperand = romanToArabic(parts[2])
	}

	if (firstOperand < 1 || firstOperand > 10) || (secondOperand < 1 || secondOperand > 10) {
		return nil, fmt.Errorf("Число меньше 1 и больше 10, введите другое число")
	}

	switch operator {
	case "+":
		result := firstOperand + secondOperand
		if regexp.MustCompile(`^[IVXLCDM]+$`).MatchString(parts[0]) && regexp.MustCompile(`^[IVXLCDM]+$`).MatchString(parts[2]) {
			return arabicToRoman(result), nil
		}
		return result, nil
	case "-":
		result := firstOperand - secondOperand
		if result <= 0 {
			return nil, fmt.Errorf("Арабские числа не могут быть отрицательными или равными нулю")
		}
		if regexp.MustCompile(`^[IVXLCDM]+$`).MatchString(parts[0]) && regexp.MustCompile(`^[IVXLCDM]+$`).MatchString(parts[2]) {
			return arabicToRoman(result), nil
		}
		return result, nil
	case "*":
		result := firstOperand * secondOperand
		if regexp.MustCompile(`^[IVXLCDM]+$`).MatchString(parts[0]) && regexp.MustCompile(`^[IVXLCDM]+$`).MatchString(parts[2]) {
			return arabicToRoman(result), nil
		}
		return result, nil
	case "/":
		if secondOperand == 0 {
			return nil, fmt.Errorf("Деление на ноль невозможно")
		}
		result := firstOperand / secondOperand
		if regexp.MustCompile(`^[IVXLCDM]+$`).MatchString(parts[0]) && regexp.MustCompile(`^[IVXLCDM]+$`).MatchString(parts[2]) {
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

	// Проверяем, что введено корректное выражение
	if !regexp.MustCompile(`^\d+\s[+\-*/]\s\d+$|^[IVXLCDM]+\s[+\-*/]\s[IVXLCDM]+$`).MatchString(expression) {
		fmt.Println("Неправильный формат ввода.")
		return
	}

	result, err := evaluateExpression(expression)
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
		return
	}

	// Определяем, какой формат использовался в выражении и выводим
	if num, ok := result.(int); ok {
		fmt.Printf("Результат: %d\n", num)
	} else {
		romanResult := result.(string)
		fmt.Printf("Результат: %s\n", romanResult)
	}
}
