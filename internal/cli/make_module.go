package cli

import (
	"fmt"
	"os"
	"strings"
)

func MakeModule(args []string) {
	if len(args) < 3 {
		fmt.Println("Usage: make:module <name>")
		return
	}

	name := strings.ToLower(args[2])
	pascalName := toPascal(name)
	basePath := "internal/modules/" + name

	dirs := []string{
		basePath + "/http",
		basePath + "/services",
		basePath + "/repositories",
		basePath + "/models",
		basePath + "/domain",
		basePath + "/dto",
	}

	for _, dir := range dirs {
		os.MkdirAll(dir, os.ModePerm)
	}

	createFile(basePath+"/module.go", moduleTemplate(name, pascalName))
	createFile(basePath+"/http/"+name+"_handler.go", handlerTemplate(name, pascalName))
	createFile(basePath+"/services/"+name+"_service.go", serviceTemplate(name, pascalName))
	createFile(basePath+"/repositories/"+name+"_repository.go", repositoryTemplate(name, pascalName))

	fmt.Println("✅ Module created:", name)
}
