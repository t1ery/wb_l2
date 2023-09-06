package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"os/exec"
	"path/filepath"
	"testing"
)

func TestAppendLine(t *testing.T) {
	// Создаем тестовые случаи с разными входными данными
	testCases := []struct {
		lines        []string
		line         string
		lineNumbers  bool
		lineNumber   int
		expectedLine string
	}{
		{
			lines:        []string{},
			line:         "Тестовая строка",
			lineNumbers:  false,
			lineNumber:   0,
			expectedLine: "Тестовая строка",
		},
		{
			lines:        []string{},
			line:         "Другая строка",
			lineNumbers:  true,
			lineNumber:   42,
			expectedLine: "42: Другая строка",
		},
	}

	for _, testCase := range testCases {
		t.Run("AppendLine", func(t *testing.T) {
			result := appendLine(testCase.lines, testCase.line, testCase.lineNumbers, testCase.lineNumber)

			if len(result) != 1 {
				t.Errorf("Ожидается одна строка в результате, получено %d", len(result))
			}

			if result[0] != testCase.expectedLine {
				t.Errorf("Неверный результат. Ожидалось: %s, получено: %s", testCase.expectedLine, result[0])
			}
		})
	}
}

func TestGrepUtility(t *testing.T) {
	// Подготовка временного каталога и файлов с тестовыми данными
	inputFile := filepath.Join("testdata", "input.txt")
	expectedOutputFile := filepath.Join("testdata", "expected_output.txt")

	// Создание файлов с тестовыми данными
	err := ioutil.WriteFile(inputFile, []byte("Это простой тестовый текст.\nМы будем искать паттерн в этом тексте.\nПаттерн - это строка, которую мы ищем.\nТекст может содержать повторяющиеся паттерны.\nПаттерн может быть в разных регистрах."), 0644)
	require.NoError(t, err)

	err = ioutil.WriteFile(expectedOutputFile, []byte("Мы будем искать паттерн в этом тексте.\nПаттерн - это строка, которую мы ищем.\nТекст может содержать повторяющиеся паттерны.\nПаттерн может быть в разных регистрах.\n"), 0644)
	require.NoError(t, err)

	// Запуск вашей утилиты с заданными флагами и входными данными
	cmd := exec.Command("go", "run", "mygrep.go", "-i", "Паттерн", inputFile)

	output, err := cmd.CombinedOutput()
	t.Logf("Результат выполнения утилиты:\n%s", string(output))
	require.NoError(t, err)

	// Сравнение фактического вывода с ожидаемым
	expectedOutput, err := ioutil.ReadFile(expectedOutputFile)
	require.NoError(t, err)

	assert.Equal(t, string(expectedOutput), string(output))
}
