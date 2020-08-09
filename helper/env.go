package helper

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

import "golang.org/x/sys/windows/registry"

func SetGoRootEnvPath(dirName string) {
	// check if windows
	if true {
		envPath := filepath.FromSlash(dirName)
		goRootPath := filepath.FromSlash(os.Getenv("GOROOT"))
		machineEnvPath := "SYSTEM\\CurrentControlSet\\Control\\Session Manager\\Environment"
		userEnvPath := "Environment"

		SetEnvVar("GOROOT", envPath)

		//Also update path for user and local machine
		updateWinPathVar("PATH", goRootPath, dirName, machineEnvPath, true)
		updateWinPathVar("PATH", goRootPath, dirName, userEnvPath, false)
	}
}

func SetEnvVar(varName, varValue string) {
	machineEnvPath := "SYSTEM\\CurrentControlSet\\Control\\Session Manager\\Environment"
	userEnvPath := "Environment"

	setWinEnvVar(varName, varValue, machineEnvPath, true)
	setWinEnvVar(varName, varValue, userEnvPath, false)
}

func setWinEnvVar(envVar, envVal, envPath string, machine bool) {
	//this sets the environment variable (GOROOT in this case) for either LOCAL_MACHINE or CURRENT_USER.
	//They are set in the registry. both must be set since the GOROOT could be used from either location.
	regPlace := registry.CURRENT_USER
	if machine {
		regPlace = registry.LOCAL_MACHINE
	}

	key, err := registry.OpenKey(regPlace, envPath, registry.ALL_ACCESS)
	if err != nil {
		fmt.Println("error open reg key:", err)
		return
	}
	defer key.Close()

	err = key.SetStringValue(envVar, envVal)
	if err != nil {
		fmt.Println("error set reg key value: ", err)
	}
}

func updateWinPathVar(envVar string, oldVal string, newVal string, envPath string, machine bool) {
	//this sets the environment variable for either LOCAL_MACHINE or CURRENT_USER.
	//They are set in the registry. both must be set since the GOROOT could be used from either location.
	regPlace := registry.CURRENT_USER
	if machine {
		regPlace = registry.LOCAL_MACHINE
	}
	key, err := registry.OpenKey(regPlace, envPath, registry.ALL_ACCESS)
	if err != nil {
		fmt.Println("error", err)
		return
	}
	defer key.Close()

	val, _, kerr := key.GetStringValue(envVar)
		if kerr != nil {
		fmt.Println("error", err)
		return
	}
	pathVars := strings.Split(val, ";")
	newPath := make([]string, 0)
	for _, pathVar := range pathVars {
		if pathVar == newVal + "\\bin" || pathVar == oldVal + "\\bin" {
			//the requested new value already exists in PATH, do nothing
			continue
		}

		newPath = append(newPath, pathVar)
	}

	newPath = append(newPath, newVal + "\\bin")
	val = strings.Join(newPath, ";")

	err = key.SetStringValue(envVar, val)
	if err != nil {
		fmt.Println("error set env value:", err)
	}
}
