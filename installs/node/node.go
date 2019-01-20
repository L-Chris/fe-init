package node

import (
	"fmt"
	"os/exec"
	"github.com/L-Chris/fe-init/utils"
)

func Install(nodeVersion string, yarnVersion string)  {
	var nodeUrl = fmt.Sprintf("https://nodejs.org/download/release/v%s/node-v%s-x64.msi", nodeVersion, nodeVersion)
	var yarnUrl = fmt.Sprintf("https://github.com/yarnpkg/yarn/releases/download/v%s/yarn-%s.msi", yarnVersion, yarnVersion)

	// download
	var isNodeDownloaded = utils.Download(nodeUrl, "node.msi")
	var isYarnDownloaded = utils.Download(yarnUrl, "yarn.msi")

	// install
	if (isNodeDownloaded) {
		exec.Command("msiexec", "/i", "node.msi").Run()
	}
	if (isYarnDownloaded) {
		exec.Command("msiexec", "/i", "yarn.msi").Run()
	}
}