package tech

type Option string

type Options []Option

func (o *Options) Parse() []Option {
	return *o
}
func (o *Options) Has(option Option) bool {
	for _, opt := range *o {
		if opt == option {
			return true
		}
	}
	return false
}

const (
	Fiber    Option = "f"
	Gin      Option = "g"
	Kafka    Option = "k"
	Rabbitmq Option = "r"
	Postgres Option = "p"
	Mongo    Option = "m"
	Logger   Option = "l"
)

var DefaultOptions = Options{Fiber, Logger, Rabbitmq, Postgres}
