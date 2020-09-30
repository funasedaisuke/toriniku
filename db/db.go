package db

import (
	"toriniku/models/aeon"
	"toriniku/models/itoyokado"
	"toriniku/models/life"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

// Init 接続、マイグレーション
func Init() *gorm.DB {

	db = gormConnect()

	// ロガーを有効にすると、詳細なログを表示します
	db.LogMode(true)

	// イトーヨーカドーテーブル
	db.AutoMigrate(&itoyokado.Group{})
	db.AutoMigrate(&itoyokado.Product{})
	db.AutoMigrate(&itoyokado.Stock{})

	// ライフテーブル
	db.AutoMigrate(&life.Group{})
	db.AutoMigrate(&life.Product{})
	db.AutoMigrate(&life.Stock{})

	// イオンテーブル
	db.AutoMigrate(&aeon.Group{})
	db.AutoMigrate(&aeon.Product{})
	db.AutoMigrate(&aeon.Stock{})

	return db
}

// gormConnect　DBのコネクションを接続
func gormConnect() *gorm.DB {
	DBMS := "mysql"
	USER := "test"
	PASS := "12345678"
	DBNAME := "test"
	CONTAINER := "toriniku_mysql"

	// MySQLだと文字コードの問題で"?parseTime=true"を末尾につける必要がある
	CONNECT := USER + ":" + PASS + "@tcp(" + CONTAINER + ")/" + DBNAME + "?parseTime=true"

	db, err := gorm.Open(DBMS, CONNECT)
	if err != nil {
		panic(err.Error())
	}
	return db
}

// Get DBのコネクションを接続
func Get() *gorm.DB {
	return db
}

// Close DBのコネクションを切断する
func Close() {
	db.Close()
}
