package file

import "github.com/eneskzlcn/goarch/architecture"

type ContentInformation struct {
	FileName        string
	FileContent     string
	TestFileContent string
}

func CreateEmptyGoAndTestFilesByGivenFileNames(packageName string, fileNames ...string) architecture.Files {
	files := make(architecture.Files, 0)
	for _, fileName := range fileNames {
		file := architecture.File{
			Name:    ToGoFileFormat(fileName),
			Content: EmptyGoFileContent(packageName),
		}
		testFile := architecture.File{
			Name:    ToGoTestFileFormat(fileName),
			Content: EmptyGoTestFileContent(packageName),
		}
		files = append(files, file, testFile)
	}
	return files
}
func CreateGoAndTestFilesByGivenContentInformations(packageName string, informations ...ContentInformation) architecture.Files {
	files := make(architecture.Files, 0)
	for _, information := range informations {
		file := architecture.File{
			Name:    ToGoFileFormat(information.FileName),
			Content: information.FileContent,
		}
		testFile := architecture.File{
			Name:    ToGoTestFileFormat(information.FileName),
			Content: information.TestFileContent,
		}
		files = append(files, file, testFile)
	}
	return files
}

func ToGoFileFormat(fileName string) string {
	return fileName + ".go"
}
func ToGoTestFileFormat(fileName string) string {
	return fileName + "_test.go"
}
func EmptyGoFileContent(packageName string) string {
	return "package " + packageName
}

func EmptyGoTestFileContent(packageName string) string {
	return "package " + packageName + "_test"
}
