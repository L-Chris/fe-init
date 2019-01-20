package main

import (
	"fmt"
	"gopkg.in/AlecAivazis/survey.v1"
	"github.com/L-Chris/fe-init/installs/node"
	"github.com/L-Chris/fe-init/installs/yarn"
	"github.com/L-Chris/fe-init/installs/git"
	"github.com/L-Chris/fe-init/installs/svn"
	"github.com/L-Chris/fe-init/installs/pxCook"
	"github.com/L-Chris/fe-init/installs/chrome"
	"github.com/L-Chris/fe-init/installs/firefox"
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
		Name: "browsers",
		Prompt: &survey.MultiSelect{
			Message: "Choose browser",
			Options: []string{"Chrome", "Firefox"},
		},
	},
	{
		Name: "tools",
		Prompt: &survey.MultiSelect{
			Message: "Choose tools",
			Options: []string{"Photoshop", "PxCook"},
		},
	},
}

func main() {
	answers := struct {
		NodeVersion string
		YarnVersion string
		RevisionControl []string
		Browsers []string
		Tools []string
	}{}

	err := survey.Ask(qs, &answers)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	node.Install(answers.NodeVersion)
	yarn.Install(answers.YarnVersion)
	git.Install()
	svn.Install()
	pxCook.Install()
	chrome.Install()
	firefox.Install()
}