package utils

import (
	"fmt"
	"github.com/cavaliercoder/grab"
)

func download(src string, dst string) {
	resp, err := grab.Get(dst, src)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(resp.Filename, " download success!")
}