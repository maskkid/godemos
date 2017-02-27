package pholcus_lib

// 基础包
import (
	// "github.com/henrylee2cn/pholcus/common/goquery" //DOM解析
	"github.com/henrylee2cn/pholcus/app/downloader/request" //必需
	"github.com/henrylee2cn/pholcus/logs"           //信息输出
	. "github.com/henrylee2cn/pholcus/app/spider"        //必需
	. "github.com/henrylee2cn/pholcus/app/spider/common" //选用

	// net包
	"net/http" //设置http.Header
	// "net/url"
	// 编码包
	// "encoding/xml"
	// "encoding/json"
	// 字符串处理包
	// "regexp"
	// "strconv"
	// "strings"
	// 其他包
	// "fmt"
	// "math"
	// "time"
)

func init() {
	Lewa.Register()
}

var Lewa = &Spider{
	Name:        "订货网登录",
	Description: "订货网登录测试 [Auto Page] [http://accounts.dinghuo123.com]",
	// Pausetime: 300,
	// Keyin:   KEYIN,
	// Limit:        LIMIT,
	EnableCookie: true,
	RuleTree: &RuleTree{
		Root: func(ctx *Context) {
			ctx.AddQueue(&request.Request{Url: "http://sso.dinghuo123.com/login", Rule: "登录页"})
		},

		Trunk: map[string]*Rule{

			"登录页": {
				ParseFunc: func(ctx *Context) {
					// ctx.AddQueue(&request.Request{
					// 	Url:    "http://accounts.dinghuo123.com",
					// 	Rule:   "登录后",
					// 	Method: "POST",
					// 	PostData: "username=123456@qq.com&password=123456&login_btn=login_btn&submit=login_btn",
					// })
					NewForm(
						ctx,
						"选择入口",
						"http://sso.dinghuo123.com/accountList",
						ctx.GetDom().Find("#myForm"),
					).Inputs(map[string]string{
						"username": "13920402226",
						"password": "Snake13920402226",
					}).Submit()
				},
			},
      "选择入口": {
				ParseFunc: func(ctx *Context) {
					// 结果存入Response中转
					ctx.Output(map[string]interface{}{
						"Body":   ctx.GetText(),
						"Cookie": ctx.GetCookie(),
					})
          logs.Log.Critical(ctx.GetText())
					ctx.AddQueue(&request.Request{
						Url:    "http://accounts.dinghuo123.com/member",
						Rule:   "登录后",
						Header: http.Header{"Referer": []string{ctx.GetUrl()}},
					})
				},
			},
			"登录后": {
				ParseFunc: func(ctx *Context) {
					// 结果存入Response中转
					ctx.Output(map[string]interface{}{
						"Body":   ctx.GetText(),
						"Cookie": ctx.GetCookie(),
					})
          logs.Log.Critical(ctx.GetText())
					ctx.AddQueue(&request.Request{
						Url:    "http://agent.dinghuo123.com/index",
						Rule:   "个人中心",
						Header: http.Header{"Referer": []string{ctx.GetUrl()}},
					})
				},
			},
			"个人中心": {
				ParseFunc: func(ctx *Context) {
					// 结果存入Response中转
					ctx.Output(map[string]interface{}{
						"Body":   ctx.GetText(),
						"Cookie": ctx.GetCookie(),
					})
          logs.Log.Critical(ctx.GetText())
					ctx.AddQueue(&request.Request{
						Url:    "http://agent.dinghuo123.com/index",
						Rule:   "商品列表",
						Header: http.Header{"Referer": []string{ctx.GetUrl()}},
					})
				},
			},

      "商品列表": {
				ParseFunc: func(ctx *Context) {
					// 结果存入Response中转
					ctx.Output(map[string]interface{}{
						"Body":   ctx.GetText(),
						"Cookie": ctx.GetCookie(),
					})
          logs.Log.Debug("-------------------------商品列表")
					/*
					ctx.AddQueue(&request.Request{
						Url:    "http://agent.dinghuo123.com/index",
						Rule:   "个人中心",
						Header: http.Header{"Referer": []string{ctx.GetUrl()}},
					})
					*/
				},
			},

		},
	},
}
