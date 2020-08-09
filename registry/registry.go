package registry

import (
	"fmt"
	"github.com/g8y3e/gvm/helper"
	"github.com/g8y3e/gvm/registry/config"
	"os"
	"runtime"
	"sort"
)

type Registry struct {
	AppVersion   string
	SystemArch string
	GoRoot     string
	GvmRoot    string
	SystemName string
	Versions map[string]interface{}
	VersionKeys []string
}

func New() *Registry {
	goRoot := os.Getenv("GOROOT")
	gvmRoot := os.Getenv("GVM_ROOT")

	// get OS architecture
	systemArch := helper.GetSystemArch()

	versions, ok := config.GoVersions[systemArch].(map[string]interface{})
	if !ok {
		fmt.Println("Don't have any Go versions for your architecture:", systemArch)
		os.Exit(1)
	}

	versionKeys := make([]string, 0, len(versions))
	for version := range versions {
		versionKeys = append(versionKeys, version)
	}
	sort.Strings(versionKeys)

	// get system name
	systemName := runtime.GOOS

	return &Registry{
		AppVersion:    config.AppVersion,
		SystemArch:  systemArch,
		GoRoot:      goRoot,
		GvmRoot:     gvmRoot,
		SystemName:  systemName,
		Versions:    versions,
		VersionKeys: versionKeys,
	}
}
