package main

import (
	"log"
	"runtime"

	"github.com/mouuff/go-rocket-update/pkg/provider"
	"github.com/mouuff/go-rocket-update/pkg/updater"
)

func main() {
	u := &updater.Updater{
		Provider: &provider.Github{
			RepositoryURL: "github.com/mouuff/go-rocket-update-example",
			ZipName:       "binaries" + runtime.GOOS + ".zip",
		},
		BinaryName: "go-rocket-update-example",
		Version:    "v0.2",
	}
	log.Println(u.Version)
	err := u.Run()
	if err != nil {
		log.Fatal(err)
	}
}
