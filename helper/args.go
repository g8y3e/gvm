package helper

import "os"

func GetMethodFromArgs() (string, string) {
	args := os.Args
	method_args := ""

	if len(args) < 2 {
		return "help", method_args
	}

	if len(args) > 2 {
		method_args = args[2]
	}

	return args[1], method_args
}