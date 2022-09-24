package main

import (
	"fmt"
	"github.com/eneskzlcn/goarch/anotherway"
	"github.com/eneskzlcn/goarch/inner/core/tech"
)

func main() {
	if err := anotherway.CreateMicroserviceArchitecture(tech.DefaultOptions); err != nil {
		fmt.Println(err)
	}
}
