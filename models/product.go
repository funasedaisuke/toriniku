package models

import (
	"github.com/jinzhu/gorm"
)

type Product struct {
	gorm.Model
	//gorm.Modelを追加すると下記カラムが追加される
	//ID        uint `gorm:"primary_key"`
	//CreatedAt time.Time
	//UpdatedAt time.Time
	//DeletedAt *time.Time
	ShopName string
	Product  string
	Price    int
	Per100G  int
}

type Items struct {
	ContentType string `json:"Content-Type"`
	ShopName    string `json:"shop_name"`
	Total_item  []Item `json:"total_item"`
}

type Item struct {
	Per100G          int    `json:"per_100g"`
	Price            int    `json:"price"`
	Product          string `json:"product"`
	TaxIncludedPrice int    `json:"tax_included_price"`
}

type Shops struct {
	ContentType string `json:"Content-Type"`
	Shop_list   []Shop `json:"shop_list"`
}

type Shop struct {
	Prefecture string `json:"prefecture"`
	Shop_name  string `json:"shop_name"`
	Url        string `json:"url"`
}
