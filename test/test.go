package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("whistle.bat")
	if (err != nil) {
		fmt.Println(err.Error())
	} else {
		_, err = f.Write([]byte("w2 start"))
	}

	f.Close()

	os.Rename("whistle.bat", "test/whistle.bat")
}