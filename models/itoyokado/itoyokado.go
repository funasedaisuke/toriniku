package itoyokado

import (
	"github.com/jinzhu/gorm"
)

const (
	// SeleniumURL seleniumAPI
	SeleniumURL = "http://selenium-python:5001/yokado_search"

	// ShopListURL イトーヨーカドー店舗名、URL一括取得API
	ShopListURL = "http://selenium-python:5001/shoplist"

	// ShopURL イトーヨーカドー商品取得API
	// ShopURL = "https://www.iy-net.jp/nspc/shoptop.do?shopcd="
	ShopURL = "https://www.iy-net.jp/nspc/shoptop.do?shopcd=00239"
)

// Group 店舗テーブル
type Group struct {
	ID         uint `gorm:"primary_key"`
	ShopName   string
	URL        string
	Prefecture string
}

// TableName 店舗テーブル名
func (g Group) TableName() string {
	return "group_itoyokado"
}

// Stock 在庫テーブル
type Stock struct {
	gorm.Model
	ShopID      uint
	ProductID   uint
	ProductName string
}

// TableName 在庫テーブル名
func (s Stock) TableName() string {
	return "stock_itoyokado"
}

// Product 商品テーブル
type Product struct {
	gorm.Model
	ShopName string
	Product  string
	Price    int
	Per100G  int
}

// TableName 在庫テーブル名
func (p Product) TableName() string {
	return "product_itoyokado"
}

// Items seliniumレスポンス
type Items struct {
	ContentType string `json:"Content-Type"`
	ShopName    string `json:"shop_name"`
	TotalItem   []Item `json:"total_item"`
}

// Item Items要素
type Item struct {
	Per100G          int    `json:"per_100g"`
	Price            int    `json:"price"`
	Product          string `json:"product"`
	TaxIncludedPrice int    `json:"tax_included_price"`
}