package main

import (
	"base"
	//"config"
	"fmt"
	//"reflect"
	//"demos"
	//"os"
)

func main() {
	fmt.Println("Main process running...")
	// demos.GinSimple() // gin web框架测试
	// demos.PholcusDemoRun() // 蜘蛛测试

	// demos.BeegoDemoRun()
	base.DDosRun()
	//fmt.Println("type::", os.Args[1], reflect.TypeOf(os.Args[1]))
}
