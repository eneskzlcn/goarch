package common

import (
	_ "embed"
	"github.com/eneskzlcn/goarch/goarch/directory"
	"github.com/eneskzlcn/goarch/goarch/file"
)

// MARK: session directory initialization.

//go:embed templates/session/session_go.arch
var sessionGoFileContent string

//go:embed templates/session/college_sessions_adapter_go.arch
var sessionCollegeSessionAdapterGoFileContent string

var SessionDirectory = directory.Directory{
	Files: file.Files{
		"session":                 file.NewGoFile(sessionGoFileContent),
		"college_session_adapter": file.NewGoFile(sessionCollegeSessionAdapterGoFileContent),
	},
}
