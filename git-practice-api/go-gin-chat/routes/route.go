package routes

import (
	"git-practice-api/go-gin-chat/controller"
	"git-practice-api/go-gin-chat/services/session"
	"git-practice-api/go-gin-chat/static"
	"git-practice-api/go-gin-chat/ws/primary"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
)

func InitRoute() *gin.Engine {
	//router := gin.Default()
	router := gin.Default()

	if viper.GetString(`app.debug_mod`) == "false" {
		// live 模式 打包用
		router.StaticFS("/static", http.FS(static.EmbedStatic))
	} else {
		// dev 开发用 避免修改静态资源需要重启服务
		router.StaticFS("/static", http.Dir("static"))
	}

	sr := router.Group("/", session.EnableCookieSession())
	{
		sr.GET("/", controller.Index)

		sr.POST("/login", controller.Login)
		sr.GET("/logout", controller.Logout)
		sr.GET("/ws", primary.Start)

		authorized := sr.Group("/", session.AuthSessionMiddle())
		{
			//authorized.GET("/ws", ws.Run)
			authorized.GET("/home", controller.Home)
			authorized.GET("/room/:room_id", controller.Room)
			authorized.GET("/private-chat", controller.PrivateChat)
			authorized.POST("/img-kr-upload", controller.ImgKrUpload)
			authorized.GET("/pagination", controller.Pagination)
		}

	}

	return router
}
