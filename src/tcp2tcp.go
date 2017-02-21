package main

import (
	"base"
	"config"
	"fmt"
	//"reflect"
	//"demos"
)

func main() {
	// demos.GinSimple() // gin web框架测试
	// demos.PholcusDemoRun() // 蜘蛛测试

	// demos.BeegoDemoRun()
	conf := config.Interface()

	port := conf.Get("TCPServer.port")
	fmt.Println("port:: ", port.(string))

	//TCPServerConf := conf["TCPServer"].(config.Conf) // 转换类型

	//fmt.Println("typeof TCPServerConf", reflect.TypeOf(TCPServerConf))
	//fmt.Println("config.TCPServer.port = ", TCPServerConf["port"])

	servsChan := make(chan string, 10)
	crossChan := make(chan string)

	// 采集端服务器
	tcpsv := &base.TCPServer{
		Port: conf.Get("TCPServer.port").(string), // interface{} 2 string
	}
	go tcpsv.Run("SourceConnecting", servsChan, crossChan)

	//  客户端服务器
	clientsv := &base.TCPServer{
		Port: conf.Get("ClientServer.port").(string),
	}
	go clientsv.Run("ClientConnecting", servsChan, crossChan)

	for i := 0; i < 10; i++ {
		<-servsChan
	}
}
