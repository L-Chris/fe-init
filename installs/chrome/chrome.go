package chrome

import (
	"os/exec"
	"github.com/L-Chris/fe-init/utils"
)

func Install()  {
	const url = "https://dl.google.com/tag/s/appguid%3D%7B8A69D345-D564-463C-AFF1-A69D9E530F96%7D%26iid%3D%7B83760A2C-4DB4-2DFA-CCA7-DCC3230233FD%7D%26lang%3Dzh-CN%26browser%3D4%26usagestats%3D1%26appname%3DGoogle%2520Chrome%26needsadmin%3Dprefers%26ap%3Dx64-stable-statsdef_1%26installdataindex%3Dempty/update2/installers/ChromeSetup.exe"

	// download
	var isDownloaded = utils.Download(url, "ChromeSetup.exe")

	// install
	if (isDownloaded) {
		exec.Command("ChromeSetup.exe").Run()
	}
}