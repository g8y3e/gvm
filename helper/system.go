package helper

import (
	"io/ioutil"
	"os"
	"strings"
)

type LoopCallback func(file os.FileInfo) bool

func GetSystemArch() string {
	systemArch := strings.ToLower(os.Getenv("PROCESSOR_ARCHITECTURE"))
	res := systemArch
	switch systemArch {
	case "x86":
		res = "386"
	}

	return res
}

func LoopDir(dirPath string, fn LoopCallback) bool {
	files, _ := ioutil.ReadDir(dirPath)
	for _, f := range files {
		if fn(f) {
			return true
		}
	}
	return false
}