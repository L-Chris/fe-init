package pxCook

import (
	"os/exec"
	"github.com/L-Chris/fe-init/utils"
)

func Install()  {
	const adobeAirUrl = "https://airdownload.adobe.com/air/win/download/32.0/AdobeAIRInstaller.exe"
	const pxCookUrl = "http://assets.fancynode.com.cn/pxcook/client/PxCook_v3.8.4_build_201812140017.air"

	// download
	var isAdobeAirDownloaded = utils.Download(adobeAirUrl, "AdobeAIRInstaller.exe")
	var isPxCookDownloaded = utils.Download(pxCookUrl, "PxCook.air")

	// install
	if (isAdobeAirDownloaded && isPxCookDownloaded) {
		exec.Command("AdobeAIRInstaller.exe").Run()
		exec.Command("cmd", "/k", "start", "PxCook.air").Run()
	}
}