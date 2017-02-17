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

	// 采集端 handler
	sv.AddHandler("SourceConnecting", func(conn *net.Conn, chanConn chan string) {
		_conn := *conn
		remoteAddr := _conn.RemoteAddr().String()
		fmt.Println("采集端连接成功", remoteAddr)

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
			receiveCont = "采集端::" + remoteAddr + "::" + receiveCont
			chanConn <- receiveCont
			// fmt.Println("chan ----> ", <-chanConn)
			go sv.ParseChanStr(receiveCont)
		}

	})

	sv.AddHandler("ClientConnecting", func(conn *net.Conn, chanConn chan string) {
		_conn := *conn
		remoteAddr := _conn.RemoteAddr().String()
		fmt.Println("客户端连接成功")

		// 轮询 发送接收到的数据
		go sv.ClientCrossChanleContent(conn, chanConn)

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
			receiveCont = remoteAddr + "::" + receiveCont
			// fmt.Println("chan ----> ", <-chanConn)
			go sv.ParseChanStr(receiveCont)
		}
	})
}

// 客户端发送接收到的内容
func (sv *TCPServer) ClientCrossChanleContent(conn *net.Conn, chanConn chan string) {
	_conn := *conn
	// 轮询 发送接收到的数据
	for {
		sourceContent, isClosed := <-chanConn
		fmt.Println(isClosed)
		_conn.Write([]byte(sourceContent))
	}
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
func (sv *TCPServer) Run(handlerName string, servsChan chan string, crossChan chan string) {
	servsChan <- "started"

	fmt.Println()

	sv.autoinit()

	fmt.Println("Server is running on::", sv.Port)
	ln, err := net.Listen("tcp", ":"+sv.Port)
	CheckError(err)

	//chanConn := make(chan string, 10)

	handler, ok := sv.handlers[handlerName]
	if ok == false {
		fmt.Println("handler ", handlerName, " is not exist")
		return
	}

	for {
		conn, err := ln.Accept()
		//defer conn.Close()
		CheckError(err)

		go handler(&conn, crossChan)
	}
}
