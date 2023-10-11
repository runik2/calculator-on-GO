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

func evaluateExpression(expression string) (interface{}, error) {
	parts := strings.Split(expression, " ")
	if len(parts) != 3 {
		return nil, fmt.Errorf("Неправильный формат выражения. Используйте формат 'число оператор число'.")
	}

	firstOperand, err := strconv.Atoi(parts[0])
	if err != nil {
		// Если не удалось преобразовать в арабское число, попробуем римское
		firstOperand = romanToArabic(parts[0])
	}

	operator := parts[1]

	secondOperand, err := strconv.Atoi(parts[2])
	if err != nil {
		// Если не удалось преобразовать в арабское число, попробуем римское
		secondOperand = romanToArabic(parts[2])
	}

	switch operator {
	case "+":
		result := firstOperand + secondOperand
		return result, nil
	case "-":
		result := firstOperand - secondOperand
		if result <= 0 {
			return nil, fmt.Errorf("Арабские числа не могут быть отрицательными или равными нулю")
		}
		return result, nil
	case "*":
		result := firstOperand * secondOperand
		return result, nil
	case "/":
		if secondOperand == 0 {
			return nil, fmt.Errorf("Деление на ноль невозможно")
		}
		result := firstOperand / secondOperand
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
	} else {
		// Если результат - арабское число, выводим как арабское
		if num, ok := result.(int); ok {
			fmt.Printf("Результат: %d\n", num)
		}
	}
}
