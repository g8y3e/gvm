package command

import (
	"fmt"
	"github.com/g8y3e/gvm/registry"
)

func Version(reg *registry.Registry) {
	fmt.Println(reg.AppVersion)
}
