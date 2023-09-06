package main

import (
	"io/ioutil"
	"os"
	"os/exec"
	"testing"
)

func TestCutUtility(t *testing.T) {
	// Подготовка временного файла с тестовыми данными
	inputData := "1\tJohn\tDoe\n2\tJane\tSmith"
	expectedOutput := "1\tDoe\n2\tSmith\n"
	inputFile := "input.txt"

	err := ioutil.WriteFile(inputFile, []byte(inputData), 0644)
	if err != nil {
		t.Fatalf("Ошибка создания файла с тестовыми данными: %v", err)
	}

	defer func() {
		err := os.Remove(inputFile)
		if err != nil {
			t.Fatalf("Ошибка удаления временного файла: %v", err)
		}
	}()

	// Запуск вашей утилиты с заданными флагами и входными данными
	cmd := exec.Command("go", "run", "mycut.go", "-f", "1,3", inputFile)
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Ошибка выполнения утилиты: %v", err)
	}

	// Сравнение фактического вывода с ожидаемым
	if string(output) != expectedOutput {
		t.Errorf("Ожидаемый вывод: %s, Фактический вывод: %s", expectedOutput, string(output))
	}
}
