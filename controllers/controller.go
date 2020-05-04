package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/me/toriniku/models"
	"net/http"
	// "strconv"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type NikuHandler struct {
	Db *gorm.DB
}

// 一覧表示
func (h *NikuHandler) GetAll(c *gin.Context) {
	go Getjson(h)
	var products []models.Product
	//データベース内の最新情報を格納
	h.Db.Last(&products)
	//index.htmlに最新情報を渡す
	c.HTML(http.StatusOK, "index.html", gin.H{
		"products": products,
	})
}

func Getjson(h *NikuHandler) {
	url := "https://zip-cloud.appspot.com/api/search?zipcode=7830060"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Get(url) error")
	}
	defer resp.Body.Close()
	byteArray, _ := ioutil.ReadAll(resp.Body)

	jsonBytes := ([]byte)(byteArray)
	data := new(models.PostCode)
	fmt.Println(byteArray)

	fmt.Println("before error")
	if err := json.Unmarshal(jsonBytes, data); err != nil {
		fmt.Println("JSON Unmarshal error:", err)
		return
	}
	//データベースに保存する
	h.Db.Create(&models.Product{Product: data.Results[0].Address1})
}

// func (h *TodoHandler) EditTask(c *gin.Context) {
// 	todo := models.Todo{}
// 	id := c.Param("id")
// 	h.Db.First(&todo, id)
// 	c.HTML(http.StatusOK, "edit.html", gin.H{
// 		"todo": todo,
// 	})
// }
// func (h *TodoHandler) UpdateTask(c *gin.Context) {
// 	todo := models.Todo{}
// 	id := c.Param("id")
// 	text, _ := c.GetPostForm("text")
// 	status, _ := c.GetPostForm("status")
// 	istatus, _ := strconv.ParseUint(status, 10, 32)
// 	h.Db.First(&todo, id)
// 	todo.Text = text
// 	todo.Status = istatus
// 	h.Db.Save(&todo)
// 	c.Redirect(http.StatusMovedPermanently, "/todo")
// }
// func (h *TodoHandler) DeleteTask(c *gin.Context) {
// 	todo := models.Todo{}
// 	id := c.Param("id")
// 	h.Db.First(&todo, id)
// 	h.Db.Delete(&todo)
// 	c.Redirect(http.StatusMovedPermanently, "/todo")
// }

//一覧
// router.GET("/", func(c *gin.Context) {
// 	tweets := dbGetAll()
// 	c.HTML(200, "index.html", gin.H{"tweets": tweets})
// })

// // 新規作成
// func (handler *NikuHandler) Create(c *gin.Context) {
// 	text, _ := c.GetPostForm("text")            // index.htmlからtextを取得
// 	handler.Db.Create(&models.Task{Text: text}) // レコードを挿入する
// 	c.Redirect(http.StatusMovedPermanently, "/")
// }

// // 編集画面
// func (handler *NikuHandler) Edit(c *gin.Context) {
// 	task := models.Task{}                                   // Task構造体の変数宣言
// 	id := c.Param("id")                                     // index.htmlからidを取得
// 	handler.Db.First(&task, id)                             // idに一致するレコードを取得する
// 	c.HTML(http.StatusOK, "edit.html", gin.H{"task": task}) // edit.htmlに編集対象のレコードを渡す
// }

// 更新
// func (handler *NikuHandler) Update(c *gin.Context) {
// 	task := models.Task{}            // Task構造体の変数宣言
// 	id := c.Param("id")              // edit.htmlからidを取得
// 	text, _ := c.GetPostForm("text") // edit.htmlからtextを取得
// 	handler.Db.First(&task, id)      // idに一致するレコードを取得する
// 	task.Text = text                 // textを上書きする
// 	handler.Db.Save(&task)           // 指定のレコードを更新する
// 	c.Redirect(http.StatusMovedPermanently, "/")
// }

// 削除
// func (handler *NikuHandler) Delete(c *gin.Context) {
// 	task := models.Task{}       // Task構造体の変数宣言
// 	id := c.Param("id")         // index.htmlからidを取得
// 	handler.Db.First(&task, id) // idに一致するレコードを取得する
// 	handler.Db.Delete(&task)    // 指定のレコードを削除する
// 	c.Redirect(http.StatusMovedPermanently, "/")
// }

// router := gin.Default()
// 	router.LoadHTMLGlob("views/*.html")

// 	dbInit()

// 	//登録
// 	router.POST("/new", func(c *gin.Context) {
// 		var form Tweet
// 		// ここがバリデーション部分
// 		if err := c.Bind(&form); err != nil {
// 			tweets := dbGetAll()
// 			c.HTML(http.StatusBadRequest, "index.html", gin.H{"tweets": tweets, "err": err})
// 			c.Abort()
// 		} else {
// 			content := c.PostForm("content")
// 			dbInsert(content)
// 			c.Redirect(302, "/")
// 		}
// 	})

// 	//投稿詳細
// 	router.GET("/detail/:id", func(c *gin.Context) {
// 		n := c.Param("id")
// 		id, err := strconv.Atoi(n)
// 		if err != nil {
// 			panic(err)
// 		}
// 		tweet := dbGetOne(id)
// 		c.HTML(200, "detail.html", gin.H{"tweet": tweet})
// 	})

// 	//更新
// 	router.POST("/update/:id", func(c *gin.Context) {
// 		n := c.Param("id")
// 		id, err := strconv.Atoi(n)
// 		if err != nil {
// 			panic("ERROR")
// 		}
// 		tweet := c.PostForm("tweet")
// 		dbUpdate(id, tweet)
// 		c.Redirect(302, "/")
// 	})

// 	//削除確認
// 	router.GET("/delete_check/:id", func(c *gin.Context) {
// 		n := c.Param("id")
// 		id, err := strconv.Atoi(n)
// 		if err != nil {
// 			panic("ERROR")
// 		}
// 		tweet := dbGetOne(id)
// 		c.HTML(200, "delete.html", gin.H{"tweet": tweet})
// 	})

// 	//削除
// 	router.POST("/delete/:id", func(c *gin.Context) {
// 		n := c.Param("id")
// 		id, err := strconv.Atoi(n)
// 		if err != nil {
// 			panic("ERROR")
// 		}
// 		dbDelete(id)
// 		c.Redirect(302, "/")

// 	})

// 	router.Run()
// }
