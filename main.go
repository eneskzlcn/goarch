/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"fmt"
	"github.com/eneskzlcn/goarch/goarch/microservice"
)

func main() {
	//cobra.Execute()
	ms := microservice.New(".")
	if err := ms.Create(); err != nil {
		fmt.Println("error occurred when creating the ms architecture")
	}
}
