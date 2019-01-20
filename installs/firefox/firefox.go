package firefox

import (
	"os/exec"
	"github.com/L-Chris/fe-init/utils"
)

func Install()  {
	const url = "https://download-ssl.firefox.com.cn/releases-sha2/stub/official/zh-CN/Firefox-latest.exe"

	// download
	var isDownloaded = utils.Download(url, "Firefox-latest.exe")

	// install
	if (isDownloaded) {
		exec.Command("Firefox-latest.exe").Run()
	}
}