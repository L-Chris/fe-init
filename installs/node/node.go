package node

import (
	"fmt"
	"os/exec"
	"github.com/L-Chris/fe-init/utils"
)

func Install(version string)  {
	var url = fmt.Sprintf("https://nodejs.org/download/release/v%s/node-v%s-x64.msi", version, version)

	// download
	var isDownloaded = utils.Download(url, "node.msi")

	// install
	if (isDownloaded) {
		exec.Command("msiexec", "/i", "node.msi").Run()
	}
}