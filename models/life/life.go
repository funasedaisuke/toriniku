package life

import (
	"toriniku/models/common"
)

const (
	// SeleniumURL seleniumAPI
	SeleniumURL = "http://selenium-python:5001/life/search"

	// ShopListURL ライフ店舗名、URL一括取得API
	ShopListURL = "http://selenium-python:5001/life/shoplist"

	// ShopURL ライフ商品取得API
	// ShopURL = "https://www.iy-net.jp/nspc/shoptop.do?shopcd="
	ShopURL = "https://www.iy-net.jp/nspc/shoptop.do?shopcd=00239"
)

// Group 店舗テーブル
type Group struct {
	common.Group
}

// TableName 店舗テーブル名
func (g Group) TableName() string {
	return "group_life"
}

// Stock 在庫テーブル
type Stock struct {
	common.Stock
}

// TableName 在庫テーブル名
func (s Stock) TableName() string {
	return "stock_life"
}

// Product 商品テーブル
type Product struct {
	common.Product
}

// TableName 商品テーブル名
func (p Product) TableName() string {
	return "product_life"
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
