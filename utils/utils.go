package utils

import (
	"fmt"
	"github.com/cavaliercoder/grab"
)

func Download(src string, dst string)(bool) {
	resp, err := grab.Get(dst, src)
	if err != nil {
		fmt.Println(src, " download failed!")
		return false
	}

	fmt.Println(resp.Filename, " download success!")
	return true
}