package main

import (
	"github.com/vumanskyi/data-object-convert/version"
	"log"
)

func main() {
	log.Printf(
		"Service is starting, version is %s, commit is %s, time is %s...",
		version.Release, version.Commit, version.BuildTime,
	)
}
