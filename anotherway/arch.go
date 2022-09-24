package anotherway

type Architecture struct {
	Directories Directories
}

func (a *Architecture) Create() error {
	return a.Directories.Create()
}
