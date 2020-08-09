package command

import (
	"fmt"
	"github.com/g8y3e/gvm/helper"
	"github.com/g8y3e/gvm/registry"
	"path/filepath"
)

func Use(args string, reg *registry.Registry) {
	if len(reg.GoRoot) == 0 {
		fmt.Println("Go not installed yet" )
		fmt.Println("For install Go please run `gvm install <version>` adn then `gvm use <version>`")
		fmt.Println("For more information run `gvm help`")
		return
	}

	if len(args) == 0 {
		fmt.Println("Version must be specified.")
		return
	}
	searchVersion := "go" + args
	goRootDir := filepath.Join(reg.GvmRoot, "go" + args, "go")
	version := helper.GetGoDirVersion(goRootDir)
	if version == searchVersion {
		helper.SetGoRootEnvPath(goRootDir)
		fmt.Println("Now using Go version:", args)
		fmt.Println("Note: You'll have to start another prompt to see the changes.")
		return
	}
	fmt.Println("Couldn't use Go version:", args, ". Check Go versions with `gvm list`")
}
