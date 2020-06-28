package router

import (
	"toriniku/config"
	"toriniku/controllers"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	// "github.com/me/toriniku/db"
)

// Router APIルーティング
func Router(dbConn *gorm.DB) {

	//初期化したデータベースをcontrollersに渡す
	nikuHandler := controllers.NikuHandler{
		Db: dbConn,
	}

	router := gin.Default()
	router.Static("/assets", "./assets")
	router.LoadHTMLGlob("templates/*")

	router.GET("/top", nikuHandler.GetAll)
	router.POST("/search", nikuHandler.Search)
	router.POST("/shoplist", nikuHandler.GetShopURL)
	router.POST("/compare", nikuHandler.Compare)

	// r.POST("/todo", nikuHandler.CreateTask)            // 新規作成
	// r.GET("/todo/:id", nikuHandler.EditTask)           // 編集画面
	// r.POST("/todo/edit/:id", nikuHandler.UpdateTask)   // 更新
	// r.POST("/todo/delete/:id", nikuHandler.DeleteTask) // 削除

	router.Run(config.ServerPort)
}
