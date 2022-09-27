package common

import (
	_ "embed"
	"github.com/eneskzlcn/goarch/goarch/directory"
	"github.com/eneskzlcn/goarch/goarch/file"
)

// MARK: mail directory initialization

//go:embed templates/mail/mail_go.arch
var mailGoFileContent string

//go:embed templates/mail/gomail_adapter_go.arch
var mailGomailAdapterGoFileContent string

var MailDirectory = directory.Directory{
	Files: file.Files{
		"mail":           file.NewGoFile(mailGoFileContent),
		"gomail_adapter": file.NewGoFile(mailGomailAdapterGoFileContent),
	},
}
