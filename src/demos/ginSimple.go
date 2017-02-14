package demos

import "github.com/gin-gonic/gin"

func GinSimple() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		//c.HTML(200, "hahaha")
	})
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})

	})
	r.Run() // listen and serve on 0.0.0.0:8080

}

type Rcallback func(*gin.Context)

type Ritem struct {
	type 	string,
	handler	Rcallback
}

routers := map[string] Ritem {
	"/" : &Ritem{
		type : 'GET',
		handler: func(c *gin.Context) {

		}
	}
} 
