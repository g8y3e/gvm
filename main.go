package main

import (
	"github.com/g8y3e/gvm/command"
	"github.com/g8y3e/gvm/helper"
	"github.com/g8y3e/gvm/registry"
)

func main() {
	reg := registry.New()

	method, args := helper.GetMethodFromArgs()
	command.Run(method, args, reg)
}