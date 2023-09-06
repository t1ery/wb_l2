package main

import (
	"reflect"
	"testing"
)

func TestFindAnagrams(t *testing.T) {
	// Тестовые случаи для функции findAnagrams
	testCases := []struct {
		input    []string            // Входные данные: массив слов
		expected map[string][]string // Ожидаемый результат: мапа с анаграммами
	}{
		{
			input: []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "кот", "ток"},
			expected: map[string][]string{
				"акптя":  {"пятак", "пятка", "тяпка"},    // Анаграммы "пятак", "пятка", "тяпка"
				"иклост": {"листок", "слиток", "столик"}, // Анаграммы "листок", "слиток", "столик"
				"кот":    {"кот", "ток"},                 // Анаграммы "кот", "ток"
			},
		},
		{
			input: []string{"абв", "вба", "где", "едг"},
			expected: map[string][]string{
				"абв": {"абв", "вба"}, // Анаграммы "абв", "вба"
				"где": {"где", "едг"}, // Анаграммы "где", "едг"
			},
		},
		{
			input:    []string{"апельсин", "банан", "вишня"},
			expected: map[string][]string{}, // Нет анаграмм
		},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			result := findAnagrams(tc.input)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Ожидается %v, но получено %v", tc.expected, result)
			}
		})
	}
}
