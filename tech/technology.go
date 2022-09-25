package tech

import "strings"

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
func OptionsFromStr(str string) Options {
	strs := strings.Split(str, "")
	options := make(Options, 0)
	for _, possibleOption := range strs {
		if allOptions.Has((Option)(possibleOption)) {
			options = append(options, (Option)(possibleOption))
		}
	}
	return options
}

const (
	Fiber    Option = "f"
	Gin      Option = "g"
	Kafka    Option = "k"
	Rabbitmq Option = "r"
	Postgres Option = "p"
	Mongo    Option = "m"
	Logger   Option = "l"
	Client   Option = "c"
)

var allOptions = Options{Fiber, Gin, Kafka, Rabbitmq, Postgres, Mongo}
var DefaultOptions = Options{Fiber, Logger, Rabbitmq, Postgres, Client, Kafka}
