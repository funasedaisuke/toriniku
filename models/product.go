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
	Product string
	Price   string
}

type PostCode struct {
	Message interface{} `json:"message"`
	Results []struct {
		Address1 string `json:"address1"`
		Address2 string `json:"address2"`
		Address3 string `json:"address3"`
		Kana1    string `json:"kana1"`
		Kana2    string `json:"kana2"`
		Kana3    string `json:"kana3"`
		Prefcode string `json:"prefcode"`
		Zipcode  string `json:"zipcode"`
	} `json:"results"`
	Status int `json:"status"`
}

type Items struct {
	ContentType string `json:"Content-Type"`
	Total_item  []Item `json:"total_item"`
}

type Item struct {
	Product          string `json:"product"`
	Price            int    `json:"price"`
	TaxIncludedPrice string `json:"tax_included_price"`
	Per100G          int    `json:"per_100g"`
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
