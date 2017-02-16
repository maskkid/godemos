package base

import (
	"fmt"
	"net"
	//"os"
)

type ConnHandler func(*net.Conn, chan string)

func CheckError(err error) {
	if err != nil {
		fmt.Println("error")
		panic(err)
		// os.Exit(1)
	}
}

type TCPServer struct {
	Port     string
	handlers map[string]ConnHandler
}

// init
func (sv *TCPServer) autoinit() {
	/*
		if len(sv.Port) > 0 {
			sv.Port = ":" + sv.Port
		} else {
			sv.Port = ":6666"
		}
	*/
	sv.handlers = make(map[string]ConnHandler)

	// client handler
	sv.AddHandler("serverConnect", func(conn *net.Conn, chanConn chan string) {
		_conn := *conn
		remoteAddr := _conn.RemoteAddr().String()
		fmt.Println("servver connected", remoteAddr)

		buf := make([]byte, 1024)

		// 读取字符串
		for {
			len, err := _conn.Read(buf)
			if err != nil {
				_conn.Close()
				fmt.Println(remoteAddr, " is closed........")
				break
			}

			receiveCont := string(buf[0:len])
			fmt.Println("receive::", receiveCont)
			chanConn <- remoteAddr + "::" + receiveCont
			// fmt.Println("chan ----> ", <-chanConn)
			go sv.ParseChanStr(<-chanConn)
		}

	})
}

//
func (sv *TCPServer) ReadReceiveFromConn(conn *net.Conn, chanConn chan string) chan string {
	chanConn <- "haha"
	return chanConn
}

// 解析数据
func (sv *TCPServer) ParseChanStr(str string) {
	fmt.Println("ParseChanstr :: ", str)
}

// add connection handler
func (sv *TCPServer) AddHandler(hname string, haction ConnHandler) {
	sv.handlers[hname] = haction
}

// server run
func (sv *TCPServer) Run() {
	sv.autoinit()

	fmt.Println("Server is running on::", sv.Port)
	ln, err := net.Listen("tcp", ":"+sv.Port)
	CheckError(err)

	chanConn := make(chan string, 10)

	for {
		conn, err := ln.Accept()
		//defer conn.Close()
		CheckError(err)

		go sv.handlers["serverConnect"](&conn, chanConn)
	}
}
