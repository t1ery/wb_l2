package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func unpackString(s string) string {
	var result strings.Builder
	escape := false

	for i := 0; i < len(s); i++ {
		if escape {
			// Если включен режим escape, добавляем символ в результат и выключаем режим escape.
			result.WriteByte(s[i])
			escape = false
		} else if s[i] == '\\' {
			// Если обнаружен символ '\' (обратный слеш), включаем режим escape.
			escape = true
		} else if unicode.IsDigit(rune(s[i])) {
			// Если обнаружена цифра, собираем все последующие цифры, чтобы определить количество повторений предыдущего символа.
			count := ""
			for j := i; j < len(s); j++ {
				if unicode.IsDigit(rune(s[j])) {
					count += string(s[j])
				} else {
					break
				}
			}
			num, _ := strconv.Atoi(count)
			if num > 0 {
				// Если число больше 0, добавляем предыдущий символ в результат нужное количество раз (за исключением самого символа).
				if i > 0 {
					result.WriteString(strings.Repeat(string(s[i-1]), num-1))
				}
			}
			i += len(count) - 1 // Перемещаем указатель i на последнюю цифру числа.
		} else {
			// В остальных случаях добавляем текущий символ в результат.
			result.WriteByte(s[i])
		}
	}

	return result.String()
}

func main() {
	// Тестовые примеры
	testCases := []string{"a4bc2d5e", "abcd", "45", "", "qwe\\4\\5", "qwe\\45", "qwe\\\\5"}
	for _, testCase := range testCases {
		result := unpackString(testCase)
		fmt.Printf("%s => %s\n", testCase, result)
	}
}
