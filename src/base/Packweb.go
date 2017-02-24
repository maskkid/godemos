package base

import (
	"fmt"
)

func PackfileRun() {
	filepathmap := map[string]string{
    "index" : "website/index.html",
	}
	wf := WebViewFile{
		filepaths: filepathmap,
	}
	fmt.Println("index::", wf.S("index"))
}

type WebViewFile struct {
	filepaths map[string]string
}

func (wf WebViewFile)S(filepath string) string {
	data, err := Asset(wf.filepaths[filepath])
	if err != nil {
		return ""
	}
	return string(data)
}
