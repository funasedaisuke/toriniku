package router

import (
	"toriniku/config"
	"toriniku/controllers"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// Router APIルーティング
func Router(dbConn *gorm.DB) {

	YokadoHandler := controllers.YokadoHandler{
		Db: dbConn,
	}

	LifeHandler := controllers.LifeHandler{
		Db: dbConn,
	}

	AeonHandler := controllers.AeonHandler{
		Db: dbConn,
	}

	CommonHandler := controllers.CommonHandler{
		Db: dbConn,
	}

	router := gin.Default()
	router.Static("/assets", "./assets")
	router.LoadHTMLGlob("templates/*")

	router.GET("/top", YokadoHandler.GetAll)

	itoyokado := router.Group("/itoyokado")
	{
		itoyokado.POST("/search", YokadoHandler.Search)
		itoyokado.POST("/shoplist", YokadoHandler.GetShopURL)
		itoyokado.POST("/compare", YokadoHandler.Compare)
	}

	life := router.Group("/life")
	{
		life.POST("/search", LifeHandler.Search)
		life.POST("/shoplist", LifeHandler.GetShopURL)
		life.POST("/compare", LifeHandler.Compare)
	}

	aeon := router.Group("/aeon")
	{
		aeon.POST("/search", AeonHandler.Search)
		aeon.POST("/shoplist", AeonHandler.GetShopURL)
		aeon.POST("/compare", AeonHandler.Compare)
	}

	common := router.Group("/common")
	{
		common.POST("/compare", CommonHandler.Compare)
	}

	router.Run(config.ServerPort)
}
