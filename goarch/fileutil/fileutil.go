package fileutil

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func CreateDirectory(absPath, name string) error {
	return exec.Command("mkdir", absPath+"/"+name).Run()
}

func CreateFileWithContent(absPath, name, content string) error {
	workingDir, _ := os.Getwd()
	file, err := os.OpenFile(workingDir+"/"+absPath+"/"+name, os.O_CREATE|os.O_WRONLY, 0644)
	defer file.Close()
	if err != nil {
		return err
	}
	_, err = fmt.Fprint(file, content)
	if err != nil {
		return err
	}
	return nil
}

func IsGoFilename(filename string) bool {
	return strings.HasSuffix(filename, ".go")
}
func IsGoTestFilename(filename string) bool {
	return strings.HasSuffix(filename, "_test.go")
}
func IsYamlFilename(filename string) bool {
	return strings.HasSuffix(filename, ".yaml")
}
func IsSqlFilename(filename string) bool {
	return strings.HasSuffix(filename, ".sql")
}
func IsMarkdownFilename(filename string) bool {
	return strings.HasSuffix(filename, ".md")
}
func IsGohtmlFilename(filename string) bool {
	return strings.HasSuffix(filename, ".gohtml")
}
