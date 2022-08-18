package main

import (
	gogithubpackageclean "github.com/interlock/go-github-package-clean"
)

func main() {
	config := gogithubpackageclean.NewConfig()

	config.RunFatal()

}
