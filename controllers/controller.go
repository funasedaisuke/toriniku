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
	"toriniku/models/aeon"
	"toriniku/models/itoyokado"
	"toriniku/models/life"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// YokadoHandler イトーヨーカドー構造体
type YokadoHandler struct {
	Db *gorm.DB
}

// LifeHandler ライフ構造体
type LifeHandler struct {
	Db *gorm.DB
}

// AeonHandler イオン構造体
type AeonHandler struct {
	Db *gorm.DB
}

// GetAll 一覧表示
func (h *YokadoHandler) GetAll(c *gin.Context) {

	var products []itoyokado.Product
	//データベース内の最新情報を格納
	h.Db.Last(&products)
	//index.htmlに最新情報を渡す
	c.HTML(http.StatusOK, "index.html", gin.H{
		"products": products,
	})
}

// Search １店舗の鶏肉情報を取得
func (h *YokadoHandler) Search(c *gin.Context) {

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
func (h *YokadoHandler) GetShopURL(c *gin.Context) {

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

	fmt.Println(string(jsonBytes))

	if err := json.Unmarshal(jsonBytes, &ResponseData); err != nil {
		fmt.Println("JSON Unmarshal error:", err)
		return
	}

	fmt.Println("shop_list", ResponseData.ShopList)
	for _, shop := range ResponseData.ShopList {

		var SearchShop itoyokado.Group
		h.Db.Where("shop_name = ?", shop.ShopName).First(&SearchShop)

		if len(SearchShop.ShopName) != 0 {
			fmt.Println("shop exist!!", SearchShop.ShopName)
			continue
		}

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
func (h *YokadoHandler) Compare(c *gin.Context) {

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

// GetAll 一覧表示
func (h *AeonHandler) GetAll(c *gin.Context) {

	var products []aeon.Product
	//データベース内の最新情報を格納
	h.Db.Last(&products)
	//index.htmlに最新情報を渡す
	c.HTML(http.StatusOK, "index.html", gin.H{
		"products": products,
	})
}

// Search １店舗の鶏肉情報を取得
func (h *AeonHandler) Search(c *gin.Context) {

	var (
		shopInfo     aeon.Group
		shopCode     int
		strShopCode  string = c.PostForm("shopcode")
		URL          string = aeon.SeleniumURL
		ResponseData aeon.Items
		shopURL      string
	)
	shopCode, _ = strconv.Atoi(strShopCode)

	h.Db.First(&shopInfo, shopCode)

	if len(shopInfo.URL) == 0 {
		shopURL = ""
	} else {
		shopURL = shopInfo.URL
	}

	Body := config.ShopReq{
		URL:  shopURL,
		Shop: shopInfo.ShopName,
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
		h.Db.Create(&aeon.Product{
			ShopName: ResponseData.ShopName,
			Product:  item.Product,
			Price:    item.Price,
			Per100G:  item.Per100G},
		)
	}

	if len(shopInfo.URL) == 0 {
		shopInfo.URL = ResponseData.ShopURL
		h.Db.Save(&shopInfo)
	}

	c.Redirect(http.StatusMovedPermanently, "/top")
}

// GetShopURL 各店舗のURLを取得
func (h *AeonHandler) GetShopURL(c *gin.Context) {

	var (
		URL          string = aeon.ShopListURL
		ResponseData config.Shops
	)

	resp, error := http.Get(URL)
	if error != nil {
		fmt.Println(error)
	}

	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	jsonBytes := ([]byte)(byteArray)

	fmt.Println(string(jsonBytes))

	if err := json.Unmarshal(jsonBytes, &ResponseData); err != nil {
		fmt.Println("JSON Unmarshal error:", err)
		return
	}

	fmt.Println("shop_list", ResponseData.ShopList)
	for _, shop := range ResponseData.ShopList {

		var SearchShop aeon.Group
		h.Db.Where("shop_name = ?", shop.ShopName).First(&SearchShop)

		if len(SearchShop.ShopName) != 0 {
			fmt.Println("shop exist!!", SearchShop.ShopName)
			continue
		}

		// データベースに保存する
		h.Db.Create(&aeon.Group{
			ShopName:   shop.ShopName,
			URL:        shop.URL,
			Prefecture: shop.Prefecture,
		})
	}
	c.Redirect(http.StatusMovedPermanently, "/top")
}

// Compare 別店舗との価格比較
func (h *AeonHandler) Compare(c *gin.Context) {

	var (
		products  []aeon.Product
		mapresult map[string]aeon.Product
		result    []aeon.Product
	)
	//データベース内の最新情報を格納
	h.Db.Where("product LIKE ?", "%若鶏もも肉%").Find(&products)

	mapresult = map[string]aeon.Product{}

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

// GetAll 一覧表示
func (h *LifeHandler) GetAll(c *gin.Context) {

	var products []life.Product
	//データベース内の最新情報を格納
	h.Db.Last(&products)
	//index.htmlに最新情報を渡す
	c.HTML(http.StatusOK, "index.html", gin.H{
		"products": products,
	})
}

// Search １店舗の鶏肉情報を取得
func (h *LifeHandler) Search(c *gin.Context) {

	var (
		shopInfo     life.Group
		shopCode     int
		strShopCode  string = c.PostForm("shopcode")
		URL          string = life.SeleniumURL
		ResponseData life.Items
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
		h.Db.Create(&life.Product{
			ShopName: ResponseData.ShopName,
			Product:  item.Product,
			Price:    item.Price,
			Per100G:  item.Per100G},
		)
	}
	c.Redirect(http.StatusMovedPermanently, "/top")
}

// GetShopURL 各店舗のURLを取得
func (h *LifeHandler) GetShopURL(c *gin.Context) {

	var (
		URL          string = life.ShopListURL
		ResponseData config.Shops
	)

	resp, error := http.Get(URL)
	if error != nil {
		fmt.Println(error)
	}

	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	jsonBytes := ([]byte)(byteArray)

	fmt.Println(string(jsonBytes))

	if err := json.Unmarshal(jsonBytes, &ResponseData); err != nil {
		fmt.Println("JSON Unmarshal error:", err)
		return
	}

	fmt.Println("shop_list", ResponseData.ShopList)
	for _, shop := range ResponseData.ShopList {

		var SearchShop life.Group
		h.Db.Where("shop_name = ?", shop.ShopName).First(&SearchShop)

		if len(SearchShop.ShopName) != 0 {
			fmt.Println("shop exist!!", SearchShop.ShopName)
			continue
		}

		// データベースに保存する
		h.Db.Create(&life.Group{
			ShopName:   shop.ShopName,
			URL:        shop.URL,
			Prefecture: shop.Prefecture,
		})
	}
	c.Redirect(http.StatusMovedPermanently, "/top")
}

// Compare 別店舗との価格比較
func (h *LifeHandler) Compare(c *gin.Context) {

	var (
		products  []life.Product
		mapresult map[string]life.Product
		result    []life.Product
	)
	//データベース内の最新情報を格納
	h.Db.Where("product LIKE ?", "%若鶏もも肉%").Find(&products)

	mapresult = map[string]life.Product{}

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
