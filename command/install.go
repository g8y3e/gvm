package command

import (
	"fmt"
	"github.com/g8y3e/gvm/helper"
	"github.com/g8y3e/gvm/registry"
)

func Install(args string, reg *registry.Registry) {
	if len(reg.GvmRoot) == 0 {
		fmt.Println("GVM root not set yet.")
		fmt.Println("Run `gvm root <path>` to gvm root folder.")
		return
	}

	if !helper.CheckValidVersion(args, reg.Versions) {
		fmt.Println("Invalid version set:", args)
		fmt.Println("Valid version format `<number>.<number>.<number>`")
		return
	}

	helper.Download(args, reg.SystemName, reg.SystemArch, reg.GvmRoot)
}