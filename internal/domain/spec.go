package domain

// Spec описывает основную структуру YAML файла.
type Spec struct {
	Name              string            `yaml:"name"`
	Description       string            `yaml:"description"`
	On                []Repository      `yaml:"on"`
	Steps             []Step            `yaml:"steps"`
	ChangesetTemplate ChangesetTemplate `yaml:"changesetTemplate"`
}

// Repository представляет запись репозитория.
type Repository struct {
	Repository string `yaml:"repository"`
}

// Step описывает один шаг выполнения.
type Step struct {
	Run       string `yaml:"run"`
	Container string `yaml:"container"`
}

// ChangesetTemplate описывает шаблон для создания changeset.
type ChangesetTemplate struct {
	Title  string `yaml:"title"`
	Body   string `yaml:"body"`
	Branch string `yaml:"branch"`
	Commit Commit `yaml:"commit"`
}

// Commit содержит информацию о коммите.
type Commit struct {
	Author  Author `yaml:"author"`
	Message string `yaml:"message"`
}

// Author описывает автора коммита.
type Author struct {
	Name  string `yaml:"name"`
	Email string `yaml:"email"`
}
