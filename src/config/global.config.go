package config

import (
	"fmt"
	"reflect"
	"strings"
)

type Conf map[string]interface{}

/**
 * 支持第归获取多级配置
 * @param keystr string 多级key字符串，eq. net.x.a.b.c
 * @param confs Conf    父级配置
 *
 */
func (this Conf) Get(keystr string, confs ...Conf) interface{} {
	/*
		fmt.Println("----------------------")
		fmt.Println("keystr", keystr, "    ", confs)
		fmt.Println("----------------------")
	*/
	ks := strings.Split(keystr, ".")
	var conf Conf
	if len(confs) == 0 {
		conf = this
	} else {
		conf = confs[0]
	}
	for _, key := range ks {
		v := conf[key]
		//fmt.Println("Get:: ", key, v, reflect.TypeOf(v), reflect.TypeOf(this))
		vtp := reflect.TypeOf(v) // 根据值类型递归or返回
		if vtp == reflect.TypeOf(this) {
			//nkeystr := strings.Join(ks[1:], ".")
			return conf.Get(strings.Join(ks[1:], "."), v.(Conf)) // v.(Conf) 用于转换interface{}类型到对应的数据类型
		} else {
			return v
			break
		}
	}
	return nil
}

func Interface() Conf {
	fmt.Println("config loaded")
	return Conf{
		"TCPServer": Conf{
			"port": "20100",
		},
		"ClientServer": Conf{
			"port": "20101",
		},
		"version": "v0.0.1",
		"author":  "simo",
	}
}
