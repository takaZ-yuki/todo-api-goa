package database

import (
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DBを使い回すことで、DBへのConnectとCloseを毎回しないようにする
var DB *gorm.DB

func Connect() {
	var err error

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	database_name := os.Getenv("DB_DATABASE_NAME")

	// [ユーザ名]:[パスワード]@tcp([ホスト名]:[ポート番号])/[データベース名]?charset=[文字コード]
	dsn := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + database_name + "?charset=utf8mb4&parseTime=true"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("FAILED TO CONNECT TO DB")
	} else {
		log.Println("CONNECTED TO DB")
	}
}
