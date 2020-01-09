package deliveries

import (
	r "GraduationProject/deliveries/router"
	"GraduationProject/utils"
	"github.com/gin-gonic/gin"
)

func Run() {
	var conf = utils.GetConfig()
	var router = r.Router()
	//测试路由
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	//关闭debug
	if !conf.Debug {
		gin.SetMode(gin.ReleaseMode)
	}
	//服务运行
	_ = router.Run(conf.Server)
}
