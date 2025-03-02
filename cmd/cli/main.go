package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/charmbracelet/lipgloss"

	"github.com/devem-tech/batch-changes/internal/adapter/parser"
)

func main() {
	// Флаг для указания пути к YAML файлу (-f)
	filePath := flag.String("f", "", "Путь к YAML спецификации")
	flag.Parse()

	// Определяем стили для ошибок и успеха
	errorStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FF5555")).
		Padding(0, 1)

	successStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#50FA7B")).
		Padding(0, 1)

	// Если флаг не указан, выводим ошибку
	if *filePath == "" {
		fmt.Println(errorStyle.Render("Пожалуйста, укажите путь к YAML файлу с помощью флага -f"))
		os.Exit(1)
	}

	// Парсинг YAML файла
	_, err := parser.ParseSpec(*filePath)
	if err != nil {
		fmt.Println(errorStyle.Render("Ошибка при разборе спецификации: " + err.Error()))
		os.Exit(1)
	}

	// Сообщение об успешном парсинге
	fmt.Println(successStyle.Render("Спецификация успешно обработана"))
}
