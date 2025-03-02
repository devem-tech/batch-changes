package parser

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"

	"github.com/devem-tech/batch-changes/internal/domain"
)

// ParseSpec считывает YAML файл по указанному пути и возвращает структуру domain.Spec.
func ParseSpec(filePath string) (*domain.Spec, error) {
	data, err := os.ReadFile(filePath) // os.ReadFile – современный способ чтения файла.
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	var spec domain.Spec
	if err := yaml.Unmarshal(data, &spec); err != nil {
		return nil, fmt.Errorf("failed to unmarshal YAML: %w", err)
	}
	return &spec, nil
}
