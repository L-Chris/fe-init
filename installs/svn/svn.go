package svn

import (
	"os/exec"
	"github.com/L-Chris/fe-init/utils"
)

func Install()  {
	const url = "https://github-production-release-asset-2e65be.s3.amazonaws.com/23216272/d919d300-0076-11e9-8f8b-92a349e763c0?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=AKIAIWNJYAX4CSVEH53A%2F20190120%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20190120T025913Z&X-Amz-Expires=300&X-Amz-Signature=72de887509ae281b2a3a08f311492cb4d569ee94c491b83fea0c73795d44f0c3&X-Amz-SignedHeaders=host&actor_id=22991560&response-content-disposition=attachment%3B%20filename%3DGit-2.20.1-64-bit.exe&response-content-type=application%2Foctet-stream"

	// download
	var isDownloaded = utils.Download(url, "svn.exe")

	// install
	if (isDownloaded) {
		exec.Command("svn.exe").Run()
	}
}