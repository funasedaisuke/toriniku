package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/me/toriniku/models"
)

//db
var db *gorm.DB

//*gorm.DBはreturnの型
func Init() *gorm.DB {
	db = gormConnect()
	// ロガーを有効にすると、詳細なログを表示します
	db.LogMode(true)
	//マイグレーションを実行するとテーブルが無い時は自動生成。あるときはなにもしない
	db.AutoMigrate(&models.Shop{})
	return db
}

func gormConnect() *gorm.DB {
	DBMS := "mysql"
	USER := "test"
	PASS := "12345678"
	DBNAME := "test"
	// MySQLだと文字コードの問題で"?parseTime=true"を末尾につける必要がある
	CONNECT := USER + ":" + PASS + "@/" + DBNAME + "?parseTime=true"
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}
	return db
}

// DBのコネクションを接続
func Get() *gorm.DB {
	return db
}

// DBのコネクションを切断する
func Close() {
	db.Close()
}

// func dbGetAll() []Product {
// 	db := gormConnect()

// 	defer db.Close()
// 	var products []Product
// 	// FindでDB名を指定して取得した後、orderで登録順に並び替え
// 	db.Order("created_at desc").Find(&products)
// 	return products
// }

// //DB一つ取得
// func dbGetOne(id int) Product {
// 	db := gormConnect()
// 	var product Product
// 	db.First(&product, id)
// 	db.Close()
// 	return product
// }

// //DB削除
// func dbDelete(id int) {
// 	db := gormConnect()
// 	var product Product
// 	db.First(&product, id)
// 	db.Delete(&product)
// 	db.Close()
// }
