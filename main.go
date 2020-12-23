package main

import (
	"fmt"
	"log"

	"github.com/mouuff/easy-update/provider"
	"github.com/mouuff/easy-update/updater"
)

func main() {

	fmt.Println(updater.GetPlatformName())

	u := &updater.Updater{
		Provider: &provider.Github{
			RepositoryURL: "github.com/mouuff/go-rocket-update-example",
			ZipName:       "binaries.zip",
		},
		BinaryName: "rocket",
		Version:    "1.0",
	}
	log.Println(u.Version)
	err := u.Run()
	if err != nil {
		log.Fatal(err)
	}
}
