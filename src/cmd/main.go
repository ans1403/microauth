package main

import (
	"golang-devcontainer/src/shell"
)

func main() {
	shell := shell.NewCalculateShell()
	shell.Execute()
}
