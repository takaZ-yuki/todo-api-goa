package database

import (
	"fmt"
	// "os"

	// "github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DBを使い回すことで、DBへのConnectとCloseを毎回しないようにする
var DB *gorm.DB

func Connect() {
	var err error

	//TODO: envファイルを読み込むように修正したい・・・
	// err := godotenv.Load()
	// if err != nil {
	// 	panic(err.Error())
	// }
	// user := os.Getenv("DB_USER")
	// password := os.Getenv("DB_PASSWORD")
	// host := os.Getenv("DB_HOST")
	// port := os.Getenv("DB_PORT")
	// database_name := os.Getenv("DB_DATABASE_NAME")

	user := "webuser"
	password := "webpass"
	host := "go-mysql"
	port := "3306"
	database_name := "go_mysql8_development"


	// [ユーザ名]:[パスワード]@tcp([ホスト名]:[ポート番号])/[データベース名]?charset=[文字コード]
	dsn := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + database_name + "?charset=utf8mb4" // "webuser:webpass@tcp(db:3306)/go_mysql8_development?charset=utf8mb4"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
	}
	// 実行ログ出力
	// DB.LogMode(true)
}
