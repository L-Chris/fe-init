package main

import (
	"os/exec"
	"fmt"
	"gopkg.in/AlecAivazis/survey.v1"
	"github.com/cavaliercoder/grab"
)

var qs = []*survey.Question{
	{
		Name: "nodeVersion",
		Prompt: &survey.Input{Message: "Node version", Default: "10.4.1"},
	},
	{
		Name: "yarnVersion",
		Prompt: &survey.Input{Message: "Yarn version", Default: "1.12.3"},
	},
	{
		Name: "revisionControl",
		Prompt: &survey.MultiSelect{
			Message: "Choose revision control system",
			Options: []string{"Git", "Svn"},
		},
	},
	{
		Name: "tools",
		Prompt: &survey.MultiSelect{
			Message: "Choose tools",
			Options: []string{"Chrome", "Firefox", "Photoshop", "Pxcook"},
		},
	},
}

func download(src string, dst string) {
	resp, err := grab.Get(dst, src)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(resp.Filename, " download success!")
}

func main() {
	// interact
	answers := struct {
		NodeVersion string
		YarnVersion string
		RevisionControl []string
		Tools []string
	}{}

	err := survey.Ask(qs, &answers)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var nodeUrl = fmt.Sprintf("https://nodejs.org/download/release/v%s/node-v%s-x64.msi", answers.NodeVersion, answers.NodeVersion)
	var yarnUrl = fmt.Sprintf("https://github.com/yarnpkg/yarn/releases/download/v%s/yarn-%s.msi", answers.YarnVersion, answers.YarnVersion)

	// download
	download(nodeUrl, "node.msi")
	download(yarnUrl, "yarn.msi")

	// install
	exec.Command("msiexec", "/i", "node.msi").Run()
	exec.Command("msiexec", "/i", "yarn.msi").Run()
}