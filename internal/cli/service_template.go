package cli

import "fmt"

func serviceTemplate(name string, pascalName string) string {
	return fmt.Sprintf(`package services

import "gia-starter-app-V1/internal/modules/%s/repositories"

type %sService struct {
    repo repositories.%sRepository
}

func New%sService(repo repositories.%sRepository) *%sService {
    return &%sService{repo: repo}
}
`, name, pascalName, pascalName, pascalName, pascalName, pascalName, pascalName)
}
