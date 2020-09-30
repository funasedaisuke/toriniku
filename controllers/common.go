package controllers

import (
	"net/http"
	"sort"
	"toriniku/models/aeon"
	"toriniku/models/itoyokado"
	"toriniku/models/life"

	"toriniku/models/common"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// CommonHandler グループ共通構造体
type CommonHandler struct {
	Db *gorm.DB
}

// Compare 別店舗との価格比較
func (h *CommonHandler) Compare(c *gin.Context) {

	var (
		prodItoYokado []itoyokado.Product
		resItoYokado  []common.Product
		prodLife      []life.Product
		resLife       []common.Product
		prodAeon      []aeon.Product
		resAeon       []common.Product
	)
	//データベース内の最新情報を格納
	h.Db.Where("name LIKE ?", "%もも%").Find(&prodItoYokado)
	h.Db.Where("name LIKE ?", "%もも%").Find(&prodLife)
	h.Db.Where("name LIKE ?", "%もも%").Find(&prodAeon)

	if len(prodItoYokado) != 0 {
		resItoYokado = GetCheapestItoYokado(prodItoYokado)
	}
	if len(prodLife) != 0 {
		resLife = GetCheapestLife(prodLife)
	}
	if len(prodAeon) != 0 {
		resAeon = GetCheapestAeon(prodAeon)
	}

	// 共通化の余地あり↓

	// イトーヨーカドー のリストに追加
	resItoYokado = append(resItoYokado, resLife...)
	resItoYokado = append(resItoYokado, resAeon...)

	sort.Slice(resItoYokado, func(i, j int) bool {
		return resItoYokado[i].Per100G < resItoYokado[j].Per100G
	})

	//index.htmlに最新情報を渡す
	c.HTML(http.StatusOK, "index.html", gin.H{
		"products": resItoYokado[:3],
	})
}

// 共通化の余地あり↓

// GetCheapestItoYokado イトーヨーカドー の最安値
func GetCheapestItoYokado(products []itoyokado.Product) []common.Product {

	var (
		mapProduct map[string]common.Product
		result     []common.Product
	)
	mapProduct = map[string]common.Product{}

	for _, product := range products {
		if _, ok := mapProduct[product.ShopName]; ok {
			if mapProduct[product.ShopName].Per100G > product.Per100G {
				mapProduct[product.ShopName] = common.Product{
					ShopName:  product.ShopName,
					Name:      product.Name,
					Price:     product.Price,
					Per100G:   product.Per100G,
					GroupName: "イトーヨーカドー",
				}
			}
		} else {
			mapProduct[product.ShopName] = common.Product{
				ShopName:  product.ShopName,
				Name:      product.Name,
				Price:     product.Price,
				Per100G:   product.Per100G,
				GroupName: "イトーヨーカドー",
			}
		}
	}

	for _, val := range mapProduct {
		result = append(result, val)
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].Per100G > result[j].Per100G
	})

	return result
}

// GetCheapestLife ライフの最安値
func GetCheapestLife(products []life.Product) []common.Product {

	var (
		mapProduct map[string]common.Product
		result     []common.Product
	)
	mapProduct = map[string]common.Product{}

	for _, product := range products {
		if _, ok := mapProduct[product.ShopName]; ok {
			if mapProduct[product.ShopName].Per100G > product.Per100G {
				mapProduct[product.ShopName] = common.Product{
					ShopName:  product.ShopName,
					Name:      product.Name,
					Price:     product.Price,
					Per100G:   product.Per100G,
					GroupName: "ライフ",
				}
			}
		} else {
			mapProduct[product.ShopName] = common.Product{
				ShopName:  product.ShopName,
				Name:      product.Name,
				Price:     product.Price,
				Per100G:   product.Per100G,
				GroupName: "ライフ",
			}
		}
	}

	for _, val := range mapProduct {
		result = append(result, val)
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].Per100G > result[j].Per100G
	})

	return result
}

// GetCheapestAeon イオンの最安値
func GetCheapestAeon(products []aeon.Product) []common.Product {

	var (
		mapProduct map[string]common.Product
		result     []common.Product
	)
	mapProduct = map[string]common.Product{}

	for _, product := range products {
		if _, ok := mapProduct[product.ShopName]; ok {
			if mapProduct[product.ShopName].Per100G > product.Per100G {
				mapProduct[product.ShopName] = common.Product{
					ShopName:  product.ShopName,
					Name:      product.Name,
					Price:     product.Price,
					Per100G:   product.Per100G,
					GroupName: "イオン",
				}
			}
		} else {
			mapProduct[product.ShopName] = common.Product{
				ShopName:  product.ShopName,
				Name:      product.Name,
				Price:     product.Price,
				Per100G:   product.Per100G,
				GroupName: "イオン",
			}
		}
	}

	for _, val := range mapProduct {
		result = append(result, val)
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].Per100G > result[j].Per100G
	})

	return result
}
