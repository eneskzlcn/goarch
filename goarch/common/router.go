package common

import (
	_ "embed"
	"github.com/eneskzlcn/goarch/goarch/directory"
	"github.com/eneskzlcn/goarch/goarch/file"
)

//MARK: router directory initialization.

//go:embed templates/router/router_go.arch
var routerGoFileContent string

//go:embed templates/router/mux_router_adapter_go.arch
var routerMuxRouterAdapterGoFileContent string

var RouterDirectory = directory.Directory{
	Files: file.Files{
		"router":             file.NewGoFile(routerGoFileContent),
		"mux_router_adapter": file.NewGoFile(routerMuxRouterAdapterGoFileContent),
	},
}
