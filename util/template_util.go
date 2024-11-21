package util

import (
	"embed"
	"fmt"
	"html/template"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/nilotpaul/gospur/config"
	tmpls "github.com/nilotpaul/gospur/template"
)

// processTemplate represents a template that needs to be processed(parsed)
// and writen to the specified target path in the project directory.
type processTemplate struct {
	targetFilePath string
	template       *template.Template
}

// CreateProject takes a `targetDir` and any optional data.
// It creates the necessary folders and files for the entire project.
func CreateProject(targetDir string, cfg StackConfig, data interface{}) error {
	// Ranging over files in base dir which doesn't depend on `StackConfig`
	for targetPath, templatePath := range config.ProjectBaseFiles {
		// Getting the embeded folder containing all base template files.
		baseTmplFS := tmpls.GetBaseFiles()

		// `targetFilePath` is the final path where the file will be stored.
		// It's joined with the (project or target) dir.
		targetFilePath := filepath.Join(targetDir, targetPath)

		// Parsing the raw tempate to get the processed template which will contain
		// the `targetFilePath`(location where the target file will be written) and
		// actual `template` itself.
		processedTmpl, err := parseTemplate(targetFilePath, templatePath, baseTmplFS)
		if err != nil {
			return fmt.Errorf("Template Parsing Error (pls report): %v", err)
		}

		// Creating the file with the parsed template.
		err = createFileFromTemplate(
			processedTmpl.targetFilePath,
			processedTmpl.template,
			data,
		)
		if err != nil {
			return fmt.Errorf("Failed to create file -> '%s' due to %v", processedTmpl.targetFilePath, err)
		}
	}

	// Ranging over files in API dir which depend on `StackConfig`
	for targetPath, templatePath := range config.ProjectAPIFiles {
		// Getting the embeded folder containing all API template files.
		apiTmplFS := tmpls.GetAPIFiles()

		// `targetFilePath` is the final path where the file will be stored.
		// It's joined with the project/target dir.
		targetFilePath := filepath.Join(targetDir, targetPath)

		// Parsing the raw tempate to get the processed template which will contain
		// the `targetFilePath`(location where the target file will be written) and
		// actual `template` itself.
		processedTmpl, err := parseTemplate(targetFilePath, templatePath, apiTmplFS)
		if err != nil {
			return fmt.Errorf("Template Parsing Error (pls report): %v", err)
		}

		// Creating the file with the parsed template.
		err = createFileFromTemplate(
			processedTmpl.targetFilePath,
			processedTmpl.template,
			data,
		)
		if err != nil {
			return fmt.Errorf("Failed to create file -> '%s' due to %v", processedTmpl.targetFilePath, err)
		}
	}

	// Ranging over files in page dir which depend on `StackConfig`.
	//
	// These needs to be processed seperately as it needs to be written
	// as template files itself, thus parsing isn't required.
	for targetPath := range config.ProjectPageFiles {
		var (
			paths = strings.Split(targetPath, "/")
			name  = paths[len(paths)-1]
		)

		// `targetFilePath` is the final path where the file will be stored.
		// It's joined with the project/target dir.
		targetFilePath := filepath.Join(targetDir, targetPath)

		// Generation the page content with `StackConfig`.
		fileBytes := generatePageContent(name, cfg)

		// Creating the file with the raw template.
		if err := writeRawTemplateFile(targetFilePath, fileBytes); err != nil {
			return fmt.Errorf("Failed to create file -> '%s' due to %v", targetFilePath, err)
		}
	}

	// Create an example public asset
	if err := createExamplePublicAsset(targetDir); err != nil {
		return fmt.Errorf("Failed to create the public directory %v", err)
	}

	return nil
}

// parseTemplate takes `fullWritePath`, template path and template embed.
//
// `fullWritePath` -> has to be joined with the project or targetPath. (eg. gospur/config/env.go)
// `tmplPath` -> path where the template is stored
// `tmplFS` -> template embed FS which contains all template files.
func parseTemplate(fullWritePath, tmplPath string, tmplFS embed.FS) (*processTemplate, error) {
	baseTmplBytes, err := tmplFS.ReadFile(tmplPath)
	if err != nil {
		return nil, err
	}

	// Parsing the tmpl bytes(file contents) to get the actual template.
	tmpl, err := template.New(filepath.Base(tmplPath)).Parse(string(baseTmplBytes))
	if err != nil {
		return nil, err
	}

	return &processTemplate{targetFilePath: fullWritePath, template: tmpl}, nil
}

// createFileFromTemplate writes the output of a parsed template to a specified file path,
// creating directories as needed.
//
// `fullWritePath`: The full path where the file will be created (e.g., "project/config/env.go").
// `tmpl`: The parsed template to execute and write to the file.
// `data`: Dynamic data for the template; use `nil` if not required.
func createFileFromTemplate(fullWritePath string, tmpl *template.Template, data interface{}) error {
	// Create parent directories for the target file.
	// Here second arg of `CreateTargetDir` is false which depicts write even
	// if the directory is not empty.
	if err := CreateTargetDir(filepath.Dir(fullWritePath), false); err != nil {
		return err
	}

	// Create the file in the target file path.
	destFile, err := os.Create(fullWritePath)
	if err != nil {
		return err
	}
	defer destFile.Close()

	// Execute the template and write the output to the file.
	if err := tmpl.Execute(destFile, data); err != nil {
		return err
	}

	return nil
}

// writeRawTemplateFile writes the raw contents of a template file directly to a specified path.
//
// `fullWritePath`: The full path where the file will be created (e.g., "project/templates/index.html").
// `templatePath`: The path of the static template file within the embedded filesystem.
// `tmplFS`: The embedded filesystem containing the template files.
func writeRawTemplateFile(fullWritePath string, bytes []byte) error {
	if err := CreateTargetDir(filepath.Dir(fullWritePath), false); err != nil {
		return err
	}

	// Write the file directly
	return os.WriteFile(fullWritePath, bytes, fs.ModePerm)
}

// createExamplePublicAsset takes a project dir path and creates a example public
// asset in the created project template.
func createExamplePublicAsset(projectDir string) error {
	fullFilePath := filepath.Join(projectDir, "public", "golang.jpg")
	assetBytes := tmpls.GetGolangImage()

	if err := CreateTargetDir(filepath.Dir(fullFilePath), false); err != nil {
		return err
	}

	// Create the file in the public folder in the project dir.
	destFile, err := os.Create(fullFilePath)
	if err != nil {
		return err
	}
	defer destFile.Close()

	// Write the (file contents) -> []byte to the created file.
	if _, err := destFile.Write(assetBytes); err != nil {
		return err
	}

	return nil
}
