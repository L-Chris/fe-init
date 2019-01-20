package yarn

import (
	"fmt"
	"os"
	"os/exec"
	"github.com/L-Chris/fe-init/utils"
)

func Install(version string)  {
	var url = fmt.Sprintf("https://github.com/yarnpkg/yarn/releases/download/v%s/yarn-%s.msi", version, version)

	// download
	var isDownloaded = utils.Download(url, "yarn.msi")

	// install
	if (isDownloaded) {
		exec.Command("msiexec", "/i", "yarn.msi").Run()
		exec.Command("yarn", "config", "set", "registry", "https://registry.npm.taobao.org").Run()
		exec.Command("yarn", "config", "set", "sass_binary_site", "http://cdn.npm.taobao.org/dist/node-sass").Run()
		exec.Command("yarn", "config", "set", "ELECTRON_MIRROR", "http://npm.taobao.org/mirrors/electron").Run()
		exec.Command("yarn", "config", "set", "PUPPETEER_DOWNLOAD_HOST", "https://storage.googleapis.com.cnpmjs.org").Run()

		exec.Command("yarn", "global", "add", "@vue/cli", "whistle").Run()

		// config whistle boot
		const bootPath = "%USERPROFILE%/AppData/Roaming/Microsoft/Windows/Start Menu/Programs/Startup"
		const whistleBootFileName = "whistle.bat"
		const whistleBootFilePath = bootPath + "/" + whistleBootFileName

		f, err := os.Create(whistleBootFileName)
		if (err != nil) {
			fmt.Println(err.Error())
		} else {
			_, err = f.Write([]byte("w2 start"))
		}

		f.Close()
		os.Rename(whistleBootFileName, whistleBootFilePath)
	}
}