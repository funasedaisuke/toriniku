package router

import (
	"toriniku/controllers"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	// "github.com/me/toriniku/db"
)

func Router(dbConn *gorm.DB) {

	//初期化したデータベースをcontrollersに渡す
	nikuHandler := controllers.NikuHandler{
		Db: dbConn,
	}
	//Default() はLoggerとRecoveryというミドルウェア設定
	r := gin.Default()
	r.Static("/assets", "./assets")
	r.LoadHTMLGlob("templates/*")
	r.GET("/top", nikuHandler.GetAll)   // 一覧画面
	r.POST("/api", nikuHandler.Getjson) // json結果取得
	// r.GET("/json", nikuHandler.Getjson) // json画面
	// r.POST("/todo", nikuHandler.CreateTask)            // 新規作成
	// r.GET("/todo/:id", nikuHandler.EditTask)           // 編集画面
	// r.POST("/todo/edit/:id", nikuHandler.UpdateTask)   // 更新
	// r.POST("/todo/delete/:id", nikuHandler.DeleteTask) // 削除
	r.Run(":8000")
}
