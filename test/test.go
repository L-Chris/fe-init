package main

import (
	"os/exec"
)

func main() {
	exec.Command("yarn", "global", "add", "@vue/cli", "whistle").Run()
}