package base

import (
	"fmt"
)

func PackfileRun() {
	filepathmap := map[string]string{

	}
	wf := WebViewFile{
		filepaths:  filepathmap
	}
	fmt.Println("index::", wf.S("index"))
}

type WebViewFile struct {
	filepaths map[string]string
}

func S(filepath string) string {
	data, err := Asset(filepaths[filepath])
	if err != nil {
		return ""
	}
	return string(data)
}
