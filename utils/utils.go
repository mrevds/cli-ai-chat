package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// PrintError выводит ошибку в stderr с красным цветом (если терминал поддерживает цвета)
func PrintError(err error) {
	fmt.Fprintf(os.Stderr, "\033[31mError: %v\033[0m\n", err)
}

// PrintSuccess выводит успешное сообщение с зеленым цветом
func PrintSuccess(message string) {
	fmt.Printf("\033[32m%s\033[0m\n", message)
}

// FormatResponse форматирует ответ от API для красивого вывода
func FormatResponse(response string) string {
	// Убираем лишние пробелы и переносы строк
	response = strings.TrimSpace(response)

	// Добавляем отступы для каждого абзаца
	lines := strings.Split(response, "\n")
	for i, line := range lines {
		if line != "" {
			lines[i] = "  " + line
		}
	}

	return strings.Join(lines, "\n")
}

// IsEmptyInput проверяет, является ли ввод пользователя пустым
func IsEmptyInput(input string) bool {
	return strings.TrimSpace(input) == ""
}

// GetUserInput получает ввод от пользователя с обработкой многострочного ввода
func GetUserInput(prompt string) (string, error) {
	fmt.Print(prompt)

	var lines []string
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break // Пустая строка завершает ввод
		}
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("ошибка чтения ввода: %v", err)
	}

	return strings.Join(lines, "\n"), nil
}
