package command

import "github.com/g8y3e/gvm/registry"

func Run(method, args string, reg *registry.Registry) {
	switch method {
	case "arch":
		Arch(reg)
	case "install":
		Install(args, reg)
	case "uninstall":
		Uninstall(args, reg)
	case "list":
		List(args, reg)
	case "root":
		Root(args, reg)
	case "use":
		Use(args, reg)
	case "version":
		Version(reg)
	case "help":
		fallthrough
	default:
		Help(reg)
	}
}
