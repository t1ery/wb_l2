package main

import (
	"sort"
	"strings"
)

// Функция для поиска анаграмм
func findAnagrams(words []string) map[string][]string {
	anagramSets := make(map[string][]string)

	for _, word := range words {
		// Приводим слово к нижнему регистру и сортируем его буквы
		word = strings.ToLower(word)
		letters := strings.Split(word, "")
		sort.Strings(letters)
		sortedWord := strings.Join(letters, "")

		// Добавляем слово в соответствующее множество анаграмм
		if _, found := anagramSets[sortedWord]; !found {
			anagramSets[sortedWord] = []string{word}
		} else {
			anagramSets[sortedWord] = append(anagramSets[sortedWord], word)
		}
	}

	// Удаляем множества из одного элемента
	for key, value := range anagramSets {
		if len(value) < 2 {
			delete(anagramSets, key)
		}
	}

	return anagramSets
}

func main() {
	words := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "кот", "ток"}
	anagramSets := findAnagrams(words)

	// Выводим результат
	for key, value := range anagramSets {
		sort.Strings(value)
		println(key, ":", strings.Join(value, ", "))
	}
}
