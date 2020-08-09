package helper

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

func CheckValidVersion(version string, versions map[string]interface{}) bool {
	_, ok := versions[version]
	return ok
}

func GetGoDirVersion(dirPath string) string {
	version := ""
	callback := func(f os.FileInfo) bool {
		if f.Name() == "VERSION" {
			versionData, err := ioutil.ReadFile(filepath.Join(dirPath, f.Name()))
			if err != nil {
				return false
			}

			version = string(versionData)
			return true
		}
		return false
	}

	LoopDir(dirPath, callback)

	return version
}
