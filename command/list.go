package command

import (
	"fmt"
	"github.com/g8y3e/gvm/helper"
	"github.com/g8y3e/gvm/registry"
	"os"
	"path/filepath"
	"sort"
)

func List(args string, reg *registry.Registry) {
	if len(args) > 0 && (args == "--global" || args == "-g") {
		for _, version := range reg.VersionKeys {
			fmt.Println("  ",  version)
		}
		return
	}

	if len(reg.GvmRoot) == 0 {
		fmt.Println("GVM root not set yet.")
		fmt.Println("Run `gvm root <path>` to gvm root folder.")
		return
	}

	versions := make([]string, 0)
	callback := func(f os.FileInfo) bool {
		if f.IsDir() {
			version := helper.GetGoDirVersion(filepath.Join(reg.GvmRoot, f.Name(), "go"))
			if len(version) > 3 {
				versions = append(versions, version[2:])
			}
		}
		return false
	}
	helper.LoopDir(reg.GvmRoot, callback)
	sort.Strings(versions)
	for _, version := range versions {
		fmt.Println("  ",  version)
	}
}
