package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strconv"
	"toriniku/config"
	"toriniku/models/itoyokado"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// NikuHandler db操作構造体
type NikuHandler struct {
	Db *gorm.DB
}

// GetAll 一覧表示
func (h *NikuHandler) GetAll(c *gin.Context) {

	var products []itoyokado.Product
	//データベース内の最新情報を格納
	h.Db.Last(&products)
	//index.htmlに最新情報を渡す
	c.HTML(http.StatusOK, "index.html", gin.H{
		"products": products,
	})
}

// Search １店舗の鶏肉情報を取得
func (h *NikuHandler) Search(c *gin.Context) {

	var (
		shopInfo     itoyokado.Group
		shopCode     int
		strShopCode  string = c.PostForm("shopcode")
		URL          string = itoyokado.SeleniumURL
		ResponseData itoyokado.Items
	)
	shopCode, _ = strconv.Atoi(strShopCode)

	h.Db.First(&shopInfo, shopCode)

	Body := config.ShopReq{
		URL: shopInfo.URL,
	}
	byteBody, _ := json.Marshal(Body)

	req, err := http.NewRequest(
		"POST",
		URL,
		bytes.NewBuffer(byteBody),
	)
	if err != nil {
		fmt.Println("NewRequest error ->", err)
	}
	req.Header.Add("Content-Type", "application/json")

	res, reserr := http.DefaultClient.Do(req)
	if reserr != nil {
		fmt.Println("Post error->", reserr)
	}

	defer res.Body.Close()

	byteArray, _ := ioutil.ReadAll(res.Body)

	if err := json.Unmarshal(byteArray, &ResponseData); err != nil {
		fmt.Println("JSON Unmarshal error:", err)
		return
	}

	for _, item := range ResponseData.TotalItem {
		h.Db.Create(&itoyokado.Product{
			ShopName: ResponseData.ShopName,
			Product:  item.Product,
			Price:    item.Price,
			Per100G:  item.Per100G},
		)
	}
	c.Redirect(http.StatusMovedPermanently, "/top")
}

// GetShopURL 各店舗のURLを取得
func (h *NikuHandler) GetShopURL(c *gin.Context) {

	var (
		URL          string = itoyokado.ShopListURL
		ResponseData config.Shops
	)

	resp, error := http.Get(URL)
	if error != nil {
		fmt.Println(error)
	}

	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	jsonBytes := ([]byte)(byteArray)

	fmt.Println(jsonBytes)

	if err := json.Unmarshal(jsonBytes, &ResponseData); err != nil {
		fmt.Println("JSON Unmarshal error:", err)
		return
	}

	fmt.Println("shop_list", ResponseData.ShopList)
	for _, shop := range ResponseData.ShopList {

		// データベースに保存する
		h.Db.Create(&itoyokado.Group{
			ShopName:   shop.ShopName,
			URL:        shop.URL,
			Prefecture: shop.Prefecture,
		})
	}
	c.Redirect(http.StatusMovedPermanently, "/top")
}

// Compare 別店舗との価格比較
func (h *NikuHandler) Compare(c *gin.Context) {

	var (
		products  []itoyokado.Product
		mapresult map[string]itoyokado.Product
		result    []itoyokado.Product
	)
	//データベース内の最新情報を格納
	h.Db.Where("product LIKE ?", "%若鶏もも肉%").Find(&products)

	mapresult = map[string]itoyokado.Product{}

	for _, product := range products {
		if _, ok := mapresult[product.ShopName]; ok {
			if mapresult[product.ShopName].Per100G > product.Per100G {
				mapresult[product.ShopName] = product
			}
		} else {
			mapresult[product.ShopName] = product
		}
	}
	for _, val := range mapresult {
		result = append(result, val)
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].Per100G > result[j].Per100G
	})

	//index.htmlに最新情報を渡す
	c.HTML(http.StatusOK, "index.html", gin.H{
		"products": result,
	})
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
