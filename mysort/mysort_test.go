package main

import (
	"testing"
)

// Тест функции extractKey, которая извлекает ключ из строки.
func TestExtractKey(t *testing.T) {
	tests := []struct {
		line          string
		keyColumn     int
		monthSort     bool
		numericSuffix bool
		expectedKey   string
	}{
		{"яблоко 123 Январь", 1, false, false, "123"},    // Тест с числовым ключом и месяцем.
		{"яблоко 123.456", 1, false, true, "000123.456"}, // Тест с числовым ключом и суффиксом.
		{"яблоко Январь", 1, true, false, "01"},          // Тест с месяцем.
		{"яблоко", 1, false, false, ""},                  // Тест без ключа.
	}

	for _, test := range tests {
		key := extractKey(test.line, test.keyColumn, test.monthSort, test.numericSuffix)
		if key != test.expectedKey {
			t.Errorf("Строка: %s, Ожидаемый ключ: %s, Получено: %s", test.line, test.expectedKey, key)
		}
	}
}

// Тест функции compareStrings, которая сравнивает строки с учетом различных флагов.
func TestCompareStrings(t *testing.T) {
	tests := []struct {
		a                   string
		b                   string
		numeric             bool
		reverse             bool
		ignoreTrailingSpace bool
		expectedResult      bool
	}{
		{"123", "456", true, false, false, true},          // Сравнение чисел.
		{"яблоко", "банан", false, false, false, true},    // Сравнение строк.
		{"яблоко", "банан", false, true, false, false},    // Сравнение строк в обратном порядке.
		{"яблоко ", "яблоко", false, false, true, true},   // Сравнение строк с игнорированием конечных пробелов.
		{"яблоко ", "яблоко  ", false, false, true, true}, // Сравнение строк с игнорированием конечных пробелов и разной длиной.
	}

	for _, test := range tests {
		result := compareStrings(test.a, test.b, test.numeric, test.reverse, test.ignoreTrailingSpace)
		if result != test.expectedResult {
			t.Errorf("a: %s, b: %s, Ожидаемый результат: %v, Получено: %v", test.a, test.b, test.expectedResult, result)
		}
	}
}

// Тест функции compareHumanNumbers, которая сравнивает строки с числами и суффиксами.
func TestCompareHumanNumbers(t *testing.T) {
	tests := []struct {
		a              string
		b              string
		reverse        bool
		expectedResult bool
	}{
		{"123.456", "123.789", false, true},  // Сравнение чисел с суффиксами.
		{"123.789", "123.456", false, false}, // Сравнение чисел с суффиксами.
		{"123.xyz", "123.abc", false, true},  // Сравнение строк с суффиксами.
		{"яблоко", "банан", true, true},      // Сравнение строк в обратном порядке.
		{"яблоко", "банан", false, false},    // Сравнение строк.
	}

	for _, test := range tests {
		result := compareHumanNumbers(test.a, test.b, test.reverse)
		if result != test.expectedResult {
			t.Errorf("a: %s, b: %s, Ожидаемый результат: %v, Получено: %v", test.a, test.b, test.expectedResult, result)
		}
	}
}

// Тест функции isSorted, которая проверяет, отсортированы ли строки.
func TestIsSorted(t *testing.T) {
	lines := []Line{
		{text: "яблоко", key: "яблоко"},
		{text: "банан", key: "банан"},
		{text: "вишня", key: "вишня"},
	}

	tests := []struct {
		lines          []Line
		reverse        bool
		expectedResult bool
	}{
		{lines, false, true}, // Ожидается, что строки отсортированы по возрастанию.
		{lines, true, false}, // Ожидается, что строки не отсортированы по убыванию.
	}

	for _, test := range tests {
		result := isSorted(test.lines, test.reverse)
		if result != test.expectedResult {
			t.Errorf("Строки: %v, Обратный порядок: %v, Ожидаемый результат: %v, Получено: %v", test.lines, test.reverse, test.expectedResult, result)
		}
	}
}

// Тест функции removeDuplicates, которая удаляет дубликаты из списка строк.
func TestRemoveDuplicates(t *testing.T) {
	// Создаем список строк с дубликатами.
	lines := []Line{
		{text: "яблоко", key: "яблоко"},
		{text: "банан", key: "банан"},
		{text: "яблоко", key: "яблоко"},
		{text: "вишня", key: "вишня"},
	}

	// Ожидаемый результат после удаления дубликатов.
	expectedResult := []Line{
		{text: "яблоко", key: "яблоко"},
		{text: "банан", key: "банан"},
		{text: "вишня", key: "вишня"},
	}

	// Вызываем функцию удаления дубликатов.
	result := removeDuplicates(lines)

	// Проверяем, что длина результата совпадает с ожидаемой длиной.
	if len(result) != len(expectedResult) {
		t.Errorf("Ожидаемая длина: %d, Получено: %d", len(expectedResult), len(result))
	}

	// Проверяем, что каждая строка в результате совпадает с ожидаемой строкой.
	for i, line := range result {
		if line != expectedResult[i] {
			t.Errorf("Ожидаемый результат: %v, Получено: %v", expectedResult[i], line)
		}
	}
}
