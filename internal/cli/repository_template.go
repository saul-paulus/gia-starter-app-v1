package cli

import "fmt"

func repositoryTemplate(name string, pascalName string) string {
	return fmt.Sprintf(`package repositories

import "gorm.io/gorm"

type %sRepository interface {}

type %sRepositoryImpl struct {
    db *gorm.DB
}

func New%sRepository(db *gorm.DB) %sRepository {
    return &%sRepositoryImpl{db: db}
}
`, pascalName, pascalName, pascalName, pascalName, pascalName)
}
