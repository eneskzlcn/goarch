package arch

type Architecture struct {
	Directories Directories
}

func New() *Architecture {
	arch := Architecture{Directories: Directories{}}
	return &arch
}
func (a *Architecture) Add(directories ...Directory) {
	for _, directory := range directories {
		a.Directories = append(a.Directories, directory)
	}
}
func (a *Architecture) Create() error {
	return a.Directories.Create()
}
func (a *Architecture) Print() {
	for _, dir := range a.Directories {
		dir.Print()
	}
}

const (
	CoreDirectory     = "./internal/core"
	RootDirectory     = "."
	InternalDirectory = "./internal"
)

var PathByArchAndDirName = map[Type]map[string]string{
	Microservice: {
		"kafka":    RootDirectory,
		"logger":   RootDirectory,
		"postgres": RootDirectory,
		"rabbitmq": RootDirectory,
		"server":   RootDirectory,
		"config":   InternalDirectory,
		"client":   InternalDirectory,
		"mocks":    InternalDirectory,
		"util":     InternalDirectory,
	},
	NLayeredWebApp: {
		"kafka":    CoreDirectory,
		"logger":   CoreDirectory,
		"postgres": CoreDirectory,
		"rabbitmq": CoreDirectory,
		"server":   CoreDirectory,
		"config":   InternalDirectory,
		"client":   CoreDirectory,
		"mocks":    InternalDirectory,
		"util":     CoreDirectory,
	},
	NLayeredBackend: {
		"kafka":    CoreDirectory,
		"logger":   CoreDirectory,
		"postgres": CoreDirectory,
		"rabbitmq": CoreDirectory,
		"server":   CoreDirectory,
		"config":   InternalDirectory,
		"client":   CoreDirectory,
		"mocks":    InternalDirectory,
		"util":     CoreDirectory,
	},
}
