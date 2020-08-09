package command

import (
	"fmt"
	"github.com/g8y3e/gvm/registry"
)

func Help(reg *registry.Registry) {
	fmt.Println("\nUsage:")
	fmt.Println("  gvm arch                     : Show architecture of system.")
	fmt.Println("  gvm install <version>        : The version must be a version of Go - like 1.2.2")
	fmt.Println("  gvm root [path]              : Set GVM root folder for installing go versions")
	fmt.Println("  gvm list [--global|-g]       : List of installed Go. With `--global` or `-g` will show all possible versions for your arch")
	fmt.Println("  gvm uninstall <version>      : Uninstall specified version of Go")
	fmt.Println("  gvm use <version>            : Switch to use the specified of installed version.")
	fmt.Println("  gvm version                  : Displays the current running version of gvm for Windows.")
}