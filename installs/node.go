package installs

import (
	"fmt"
	"../utils"
)

func installNode(version string)  {
	var nodeUrl = fmt.Sprintf("https://nodejs.org/download/release/v%s/node-v%s-x64.msi", version, version)
	var yarnUrl = fmt.Sprintf("https://github.com/yarnpkg/yarn/releases/download/v%s/yarn-%s.msi", version, version)

	// download
	utils.download(nodeUrl, "node.msi")
	utils.download(yarnUrl, "yarn.msi")

	// install
	exec.Command("msiexec", "/i", "node.msi").Run()
	exec.Command("msiexec", "/i", "yarn.msi").Run()
}