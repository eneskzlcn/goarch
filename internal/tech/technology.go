package tech

type Options string

//type ComposeTechnologies string

const (
	Fiber    Options = "f"
	Gin      Options = "g"
	Kafka    Options = "k"
	Rabbitmq Options = "r"
	Postgres Options = "p"
	Mongo    Options = "m"
	Logger   Options = "l"
)

var DefaultOptions = []Options{Fiber, Logger, Rabbitmq, Postgres}
