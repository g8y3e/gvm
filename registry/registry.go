package registry

import (
	"fmt"
	"github.com/g8y3e/gvm/helper"
	"github.com/spf13/viper"
	"log"
	"os"
	"sort"
)

type Registry struct {
	AppCofig   *viper.Viper
	SystemArch string
	GoRoot     string
	GvmRoot    string
	SystemName string
	Versions map[string]interface{}
	VersionKeys []string
}

func New() *Registry {
	appConfig := newViper()
	goRoot := os.Getenv("GOROOT")
	gvmRoot := os.Getenv("GVM_ROOT")

	// get OS architecture
	systemArch := helper.GetSystemArch()

	archList := appConfig.GetStringMap("GO_VERSIONS")
	versions, ok := archList[systemArch].(map[string]interface{})
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
	systemName := "windows"

	return &Registry{
		AppCofig:   appConfig,
		SystemArch: systemArch,
		GoRoot:     goRoot,
		GvmRoot:    gvmRoot,
		SystemName: systemName,
		Versions: versions,
		VersionKeys: versionKeys,
	}
}

func newViper() *viper.Viper {
	v := viper.New()
	v.AddConfigPath("./registry/config")
	v.SetConfigName("config")
	err := v.ReadInConfig()

	if err != nil {
		log.Fatal(err)
	}

	return v
}