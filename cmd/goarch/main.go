package main

import (
	_ "embed"
	"fmt"
	"github.com/eneskzlcn/goarch/arch"
	"github.com/eneskzlcn/goarch/goarch"
	"github.com/eneskzlcn/goarch/tech"
	"os"
)

//func CreateCmdDirectory() error {
//	if err := CreateDirectory("cmd"); err != nil {
//		return err
//	}
//
//	currentDirectory, _ := os.Getwd()
//	projectDirName := path.Base(currentDirectory)
//	if err := CreateDirectory(fmt.Sprintf("cmd/%s", projectDirName)); err != nil {
//		return err
//	}
//
//	if err := CreateFile(fmt.Sprintf("cmd/%s/main.go", projectDirName)); err != nil {
//
//	}
//	return nil
//}

func main() {
	var err error
	if len(os.Args) <= 1 {
		err = goarch.Create(arch.Default, tech.DefaultOptions...)
	} else {
		//err = CreateArchitecture(os.Args[1], DefaultDirectory)
		fmt.Println("not valid args.")
	}
	if err != nil {
		fmt.Println("error occurred when creating architecture", err.Error())
	}
}
