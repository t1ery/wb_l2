package main

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestChangeDirectory(t *testing.T) {
	// Создаем временную директорию для тестов
	tempDir := t.TempDir()
	err := os.Mkdir(tempDir+"/testdir", 0755)
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		args   []string
		errMsg string
	}{
		{[]string{"cd", tempDir}, ""},
		{[]string{"cd", "nonexistent"}, "chdir nonexistent: no such file or directory"},
		{[]string{"cd"}, "Использование: cd <директория>"},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("Args: %v", test.args), func(t *testing.T) {
			err := changeDirectory(test.args)
			if err != nil {
				if test.errMsg == "" {
					t.Errorf("Ожидалось отсутствие ошибки, но получено: %v", err)
				} else if !strings.Contains(err.Error(), test.errMsg) {
					t.Errorf("Ожидалось: %v, получено: %v", test.errMsg, err.Error())
				}
			}
		})
	}
}

func TestPrintWorkingDirectory(t *testing.T) {
	dir, err := printWorkingDirectory()
	if err != nil {
		t.Fatalf("Ошибка при получении текущей директории: %v", err)
	}
	t.Logf("Текущая директория: %s", dir)
}

func TestEcho(t *testing.T) {
	tests := []struct {
		args     []string
		expected string
		errMsg   string
	}{
		{[]string{"echo", "Привет, мир!"}, "Привет, мир!", ""},
		{[]string{"echo"}, "", "Использование: echo <текст>"},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("Args: %v", test.args), func(t *testing.T) {
			output, err := echo(test.args)
			if err != nil {
				if test.errMsg == "" {
					t.Errorf("Ожидалось отсутствие ошибки, но получено: %v", err)
				} else if !strings.Contains(err.Error(), test.errMsg) {
					t.Errorf("Ожидалось: %v, получено: %v", test.errMsg, err.Error())
				}
			}

			if output != test.expected {
				t.Errorf("Ожидалось: %v, получено: %v", test.expected, output)
			}
		})
	}
}

func TestListProcesses(t *testing.T) {
	output, err := listProcesses()
	if err != nil {
		t.Fatalf("Ошибка при выполнении команды 'ps': %v", err)
	}
	t.Logf("Вывод 'ps':\n%s", output)
}

func TestKillProcess(t *testing.T) {
	// Здесь в pid нужно указать имя процесса, который хотим убить.
	pid := "12345"

	tests := []struct {
		args   []string
		errMsg string
	}{
		{[]string{"kill", pid}, ""},
		{[]string{"kill"}, "Использование: kill <PID>"},
		{[]string{"kill", "invalid"}, "Ошибка: неверный формат PID"},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("Args: %v", test.args), func(t *testing.T) {
			err := killProcess(test.args)
			if err != nil {
				if test.errMsg == "" {
					t.Errorf("Ожидалось отсутствие ошибки, но получено: %v", err)
				} else if !strings.Contains(err.Error(), test.errMsg) {
					t.Errorf("Ожидалось: %v, получено: %v", test.errMsg, err.Error())
				}
			}
		})
	}
}

func TestExecuteExternalCommand(t *testing.T) {
	// Сохраняем текущую директорию
    originalDir, err := os.Getwd()
    if err != nil {
        t.Fatalf("Ошибка при получении текущей директории: %v", err)
    }

    // Устанавливаем временную директорию для теста
    tempDir := t.TempDir()
    err = os.Chdir(tempDir)
    if err != nil {
        t.Fatalf("Ошибка при смене директории: %v", err)
    }

    defer func() {
        // Возвращаемся в исходную директорию после теста
        err = os.Chdir(originalDir)
        if err != nil {
            t.Fatalf("Ошибка при возвращении в исходную директорию: %v", err)
        }
    }()

	tests := []struct {
		args     []string
		expected string
		errMsg   string
	}{
		{[]string{"echo", "Hello, World!"}, "Hello, World!\n", ""},
		{[]string{"ls", "-l"}, "total 0\n", ""},
		{[]string{"nonexistent-command"}, "", "exec: \"nonexistent-command\": executable file not found"},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("Args: %v", test.args), func(t *testing.T) {
			output, err := executeExternalCommand(test.args)
			if err != nil {
				if test.errMsg == "" {
					t.Errorf("Ожидалось отсутствие ошибки, но получено: %v", err)
				} else if !strings.Contains(err.Error(), test.errMsg) {
					t.Errorf("Ожидалось: %v, получено: %v", test.errMsg, err.Error())
				}
			}

			if string(output) != test.expected {
				t.Errorf("Ожидалось: %v, получено: %v", test.expected, string(output))
			}
		})
	}
}
