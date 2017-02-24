package pholcus_lib

// 基础包
import (
  "fmt"
	"github.com/henrylee2cn/pholcus/app/downloader/request" //必需
	"github.com/henrylee2cn/pholcus/common/goquery"         //DOM解析
	"github.com/henrylee2cn/pholcus/logs"              //信息输出
	. "github.com/henrylee2cn/pholcus/app/spider" //必需
	// . "github.com/henrylee2cn/pholcus/app/spider/common" //选用
	// net包
	// "net/http" //设置http.Header
	// "net/url"
	// 编码包
	// "encoding/xml"
	// "encoding/json"
	// 字符串处理包
	// "regexp"
	"strconv"
	// "strings"
	// 其他包
	// "fmt"
	// "math"
	// "time"
)

func init() {
	Youzan.Register()
}

// 考拉海淘,海外直采,7天无理由退货,售后无忧!考拉网放心的海淘网站!
var Youzan = &Spider{
	Name:        "有赞-一清香店铺",
	Description: "有赞-一清香店铺 所有商品详情数据",
	// Pausetime: 300,
	// Keyin:   KEYIN,
	// Limit:        LIMIT,
	EnableCookie: false,
	RuleTree: &RuleTree{
		Root: func(ctx *Context) {
      ctx.AddQueue(&request.Request{Url: "https://h5.youzan.com/v2/showcase/homepage?kdt_id=18593404&reft=1487908761842_1487908763635&spm=f46395154_ag18593404", Rule: "获取版块URL"})
		},

		Trunk: map[string]*Rule{

			"获取版块URL": {
			  /*
        AidFunc: func(ctx *Context, aid map[string]interface{}) interface{} {
          for i:=0;i<50;i++ {
            logs.Log.Critical("aidFunc running")
            ctx.AddQueue(&request.Request{
              Url: "https://h5.youzan.com/v2/showcase/goods/allgoods?kdt_id=18593404&t=0&p=" + strconv.Itoa(i+1),
              Rule: "商品列表",
              Temp: map[string]interface{}{"goodsType": 123},
            })
          }
          return nil
        },
        */
				ParseFunc: func(ctx *Context) {
				  logs.Log.Critical("商品入口解析ing....")
					query := ctx.GetDom()
          total := query.Find(".js-all-goods span").Eq(0).Text()
          fmt.Println(total)

          logs.Log.Critical("[消息提示：| 任务：%v | KEYIN：%v | 规则：%v] 由于跳转AJAX问题，目前只能每个子类抓取 1 页……\n", ctx.GetName(), ctx.GetKeyin(), ctx.GetRuleName())

          for i:=0;i<50;i++ {
            logs.Log.Critical("https://h5.youzan.com/v2/showcase/goods/allgoods?kdt_id=18593404&t=0&p=" + strconv.Itoa(i+1))
            ctx.AddQueue(&request.Request{
              Url: "https://h5.youzan.com/v2/showcase/goods/allgoods?kdt_id=18593404&t=0&p=" + strconv.Itoa(i+1),
              Rule: "商品列表",
              Temp: map[string]interface{}{"goodsType": total},
            })
          }
					/*
					lis := query.Find(".js-all-goods a")

					lis.Each(func(i int, s *goquery.Selection) {
						if i == 0 {
							return
						}
						if url, ok := s.Attr("href"); ok {
							ctx.AddQueue(&request.Request{Url: url, Rule: "商品列表", Temp: map[string]interface{}{"goodsType": s.Text()}})
						}
					})
					*/
				},
			},

			"商品列表": {
				ParseFunc: func(ctx *Context) {
				  logs.Log.Critical("商品列表解析ing....")
					query := ctx.GetDom()
					query.Find(".js-goods-card").Each(func(i int, s *goquery.Selection) {
						if url, ok := s.Find("a").Attr("href"); ok {
							ctx.AddQueue(&request.Request{
								Url:  url,
								Rule: "商品详情",
								Temp: map[string]interface{}{"goodsType": ctx.GetTemp("goodsType", "").(string)},
							})
						}
					})
				},
			},

			"商品详情": {
				//注意：有无字段语义和是否输出数据必须保持一致
				ItemFields: []string{
					"标题",
					"图片",
					"价格",
					"运费",
					"库存",
					"描述",
				},
				ParseFunc: func(ctx *Context) {
					query := ctx.GetDom()

					// 获取标题
					title := query.Find(".goods-title").Text()

					// 获取图片
					img, _ := query.Find(".goods-img img").Eq(0).Attr("src")

					// 获取价格
					price := query.Find(".goods-current-price").Text()

					// 获取运费
					yunfei := query.Find(".goods-meta tr").Eq(0).Find("td").Eq(1).Text()

					// 库存
					stock := query.Find(".goods-meta tr").Eq(1).Find("td").Eq(1).Text()

					// 获取描述
					descripe := query.Find(".js-goods-detail .rich-text").Text()

					// 结果存入Response中转
					ctx.Output(map[int]interface{}{
						0: title,
						1: img,
						2: price,
						3: yunfei,
						4: stock,
            5: descripe,
						//5: ctx.GetTemp("goodsType", ""),
					})
				},
			},
		},
	},
}
