package main

import (
	"fmt"
	"gopkg.in/AlecAivazis/survey.v1"
	"github.com/L-Chris/fe-init/installs/node"
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

	node.Install(answers.NodeVersion, answers.YarnVersion)
	// const gitUrl = "https://github-production-release-asset-2e65be.s3.amazonaws.com/23216272/d919d300-0076-11e9-8f8b-92a349e763c0?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=AKIAIWNJYAX4CSVEH53A%2F20190120%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20190120T025913Z&X-Amz-Expires=300&X-Amz-Signature=72de887509ae281b2a3a08f311492cb4d569ee94c491b83fea0c73795d44f0c3&X-Amz-SignedHeaders=host&actor_id=22991560&response-content-disposition=attachment%3B%20filename%3DGit-2.20.1-64-bit.exe&response-content-type=application%2Foctet-stream"
	// const svnUrl = ""

	// const chromeUrl = "https://dl.google.com/tag/s/appguid%3D%7B8A69D345-D564-463C-AFF1-A69D9E530F96%7D%26iid%3D%7B83760A2C-4DB4-2DFA-CCA7-DCC3230233FD%7D%26lang%3Dzh-CN%26browser%3D4%26usagestats%3D1%26appname%3DGoogle%2520Chrome%26needsadmin%3Dprefers%26ap%3Dx64-stable-statsdef_1%26installdataindex%3Dempty/update2/installers/ChromeSetup.exe"
	// const firefoxUrl = "https://download-ssl.firefox.com.cn/releases-sha2/stub/official/zh-CN/Firefox-latest.exe"
	// const photoshopUrl = ""
	// const adobeAirUrl = "https://airdownload.adobe.com/air/win/download/32.0/AdobeAIRInstaller.exe"
	// const pxCookUrl = "http://assets.fancynode.com.cn/pxcook/client/PxCook_v3.8.4_build_201812140017.air"
}