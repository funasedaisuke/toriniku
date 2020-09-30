package common

import (
	"github.com/jinzhu/gorm"
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
	return "group_common"
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
	return "stock_common"
}

// Product 商品テーブル
type Product struct {
	gorm.Model
	ShopName  string
	Name      string
	Price     int
	Per100G   int
	GroupName string
}

// TableName 在庫テーブル名
func (p Product) TableName() string {
	return "product_common"
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
