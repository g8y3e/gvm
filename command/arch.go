package command

import (
	"fmt"
	"github.com/g8y3e/gvm/registry"
)

func Arch(reg *registry.Registry) {
	fmt.Println("System architecture:", reg.SystemArch)
}
