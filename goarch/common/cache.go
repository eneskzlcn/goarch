package common

import (
	_ "embed"
	"github.com/eneskzlcn/goarch/goarch/directory"
	"github.com/eneskzlcn/goarch/goarch/file"
)

// MARK: cache directory initialization

//go:embed templates/cache/cache_go.arch
var cacheGoFileContent string

//go:embed templates/cache/gcache-adapter_go.arch
var cacheGcacheAdapterGoFileContent string

var CacheDirectory = directory.Directory{
	Files: file.Files{
		"cache":          file.NewGoFile(cacheGoFileContent),
		"gcache_adapter": file.NewGoFile(cacheGcacheAdapterGoFileContent),
	},
}
