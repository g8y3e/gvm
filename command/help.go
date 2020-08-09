package command

import (
	"fmt"
	"github.com/g8y3e/gvm/registry"
)

func Help(reg *registry.Registry) {
	fmt.Println("\nRunning version " + reg.AppCofig.GetString("APP_VERSION") + ".")
	fmt.Println("\nUsage:")
	fmt.Println(" ")
	fmt.Println("  gvm arch                     : Show architecture of OS.")
	fmt.Println("  gvm install <version>        : The version must be a version of Go.")
	fmt.Println("  gvm root [path]            : Sets/appends GOROOT/PATH. Without the extra arg just shows current GOROOT.")
	fmt.Println("  gvm list                     : List the Go installations at or adjacent to GOROOT. Aliased as ls.")
	fmt.Println("  gvm uninstall <version>      : Uninstall specified version of Go. If it was your GOROOT/PATH, make sure to set a new one after.")
	fmt.Println("  gvm use <version>            : Switch to use the specified version. This will set your GOROOT and PATH.")
	fmt.Println("  gvm version                  : Displays the current running version of gvm for Windows. Aliased as v.")
}