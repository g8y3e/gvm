package command

import (
	"fmt"
	"github.com/g8y3e/gvm/helper"
	"github.com/g8y3e/gvm/registry"
)

func Root(args string, reg *registry.Registry) {
	if len(args) == 0 {
		folderName := reg.GvmRoot
		if len(folderName) == 0 {
			fmt.Println("GVM root wasn't set." )
			fmt.Println("Set a GVM root with `gvm root <path>`")
		} else {
			fmt.Println("gvm root folder:", reg.GvmRoot)
		}
		return
	}

	helper.SetEnvVar("GVM_ROOT", args)
	fmt.Println("gvm root initiated at:", args)
}