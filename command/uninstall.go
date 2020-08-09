package command

import (
	"fmt"
	"github.com/g8y3e/gvm/helper"
	"github.com/g8y3e/gvm/registry"
	"os"
	"path/filepath"
)

func Uninstall(args string, reg *registry.Registry) {
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
	if reg.GoRoot == goRootDir {
		fmt.Println("You trying remove currently used version" )
		fmt.Println("Please change current version with `gvm use <version>` and only then remove this version")
		return
	}

	version := helper.GetGoDirVersion(goRootDir)
	if version == searchVersion {
		os.RemoveAll(filepath.Join(reg.GvmRoot, "go" + args))
		fmt.Println("Uninstalled Go version:", args)
		fmt.Println("Note: If this was your GOROOT, make sure to set a new GOROOT with `gvm use <path>`")
		return
	}

	fmt.Println("Couldn't uninstall Go version:", args)
	fmt.Println("Check installed Go versions with `gvm list`")
}