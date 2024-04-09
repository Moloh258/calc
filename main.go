package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Введите выражение в формате X + Y")
		inputText, _ := reader.ReadString('\n')
		inputText = strings.TrimSpace(inputText)
		splittedInputText := strings.Split(inputText, " ")

		if len(splittedInputText) == 1 {
			panic("Выдача паники, так как строка не является математической операцией.")
		}

		if len(splittedInputText) != 3 {
			panic("Выдача паники, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
		}

		firstNumber := parseNumber(splittedInputText[0])
		secondNumber := parseNumber(splittedInputText[2])
		operator := splittedInputText[1]

		if firstNumber == 0 || secondNumber == 0 {
			panic("Ошибка: неверный формат числа")
		}

		if isRomanNumeral(splittedInputText[0]) != isRomanNumeral(splittedInputText[2]) {
			panic("Выдача паники, так как используются одновременно разные системы счисления.")
		}

		result := 0
		switch operator {
		case "+":
			result = firstNumber + secondNumber
		case "-":
			result = firstNumber - secondNumber
		case "*":
			result = firstNumber * secondNumber
		case "/":
			result = firstNumber / secondNumber
		default:
			panic("Выдача паники, так как какой-то не такой у вас оператор")
			continue
		}

		if isRomanNumeral(splittedInputText[0]) {
			romanResult := arabicToRoman(result)
			if result < 0 {
				panic("Выдача паники, так как в римской системе нет отрицательных чисел.")
			} else {
				fmt.Println("Результат: ", romanResult)
			}
		} else {
			fmt.Println("Результат: ", result)
		}
	}
}

func parseNumber(number string) int {
	arabicNumber, err := strconv.Atoi(number)
	if err == nil && arabicNumber >= 0 && arabicNumber <= 10 {
		return arabicNumber
	} else {
		romanMap := map[string]int{"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10}
		arabicNumber, isRoman := romanMap[number]
		if isRoman {
			return arabicNumber
		}
	}
	return 0
}

func isRomanNumeral(input string) bool {
	roman := []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}
	for _, numeral := range roman {
		if input == numeral {
			return true
		}
	}
	return false
}

func arabicToRoman(num int) string {
	if num <= 0 || num > 10 {
		return ""
	}

	romanNumerals := []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}
	return romanNumerals[num-1]
}
