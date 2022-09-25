package architecture

type Type string

const (
	NLayeredBackend Type = "nlb"
	NLayeredWebApp  Type = "nlwa"
	Microservice    Type = "ms"
	Default              = Microservice
)
