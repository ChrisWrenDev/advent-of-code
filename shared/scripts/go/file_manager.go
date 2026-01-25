package scripts

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

func CreateAocDirectory(dayDir string) string {
	err := os.MkdirAll(dayDir, os.ModePerm)
	if err != nil {
		fmt.Println("Error creating directory: ", dayDir)
		return ""
	}

	return dayDir
}

func CreateAocFile(dayDir, fileName string, body []byte) {
	filePath := filepath.Join(dayDir, fileName)
	err := os.WriteFile(filePath, body, 0644)
	if err != nil {
		fmt.Println("Error creating file: ", fileName)
	}
}

func GenerateAocTemplate(dayDir string) {
	err := os.MkdirAll(dayDir, os.ModePerm)
	if err != nil {
		fmt.Println("Error creating directory: ", err)
	}

	// Create solution.go from day_template.go
	const solutionTemplatePath = "templates/day_template.go"
	solutionFilePath := filepath.Join(dayDir, "solution.go")
	err = createFileFromTemplate(solutionTemplatePath, solutionFilePath, nil)
	if err != nil {
		fmt.Println("error creating solution.go: ", err)
	}

	// Create solution_test.go from day_template_test.go
	const solutionTestTemplatePath = "templates/day_template_test.go"
	solutionTestFilePath := filepath.Join(dayDir, "solution_test.go")
	err = createFileFromTemplate(solutionTestTemplatePath, solutionTestFilePath, nil)
	if err != nil {
		fmt.Println("error creating solution_test.go: ", err)
	}
}

func createFileFromTemplate(templatePath, destinationPath string, data interface{}) error {
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return err
	}

	out, err := os.Create(destinationPath)
	if err != nil {
		return err
	}
	defer out.Close()

	if err := tmpl.Execute(out, data); err != nil {
		return err
	}

	return nil
}
