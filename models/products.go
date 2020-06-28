package models

import (
	"github.com/jinzhu/gorm"
)

// Product 商品詳細
type Product struct {
	gorm.Model
	ShopName string
	Product  string
	Price    int
	Per100G  int
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

// Shops seleniumレスポンス
type Shops struct {
	ContentType string `json:"Content-Type"`
	ShopList    []Shop `json:"shop_list"`
}

// Shop Shops要素
type Shop struct {
	Prefecture string `json:"prefecture"`
	ShopName   string `json:"shop_name"`
	URL        string `json:"url"`
}
