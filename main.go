package main

import (
	_ "embed"
	"fmt"
	"github.com/eneskzlcn/goarch/microservice"
	"github.com/eneskzlcn/goarch/tech"
	"os"
)

func main() {
	//var err error
	//if runtime.GOARCH != "arm64" || runtime.GOOS != "darwin" {
	//	fmt.Println("This application written and serving for arm64/darwin vehicles.")
	//	return
	//}
	//if len(os.Args) <= 1 {
	//	err = goarch.Create(arch.Default, tech.DefaultOptions)
	//} else {
	//	//err = CreateArchitecture(os.Args[1], DefaultDirectory)
	//	fmt.Println("not valid args.")
	//}
	//if err != nil {
	//	fmt.Println("error occurred when creating architecture", err.Error())
	//}
	if err := microservice.CreateArchitecture(tech.DefaultOptions); err != nil {
		fmt.Println(err.Error())
		os.Exit(5)
	}
	os.Exit(0)
}
