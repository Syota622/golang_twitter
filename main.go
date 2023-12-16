package main

import (
	"database/sql"
	"golang_twitter/api"
	"golang_twitter/db"
	"golang_twitter/util"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq" // PostgreSQLドライバー
)

func main() {
	route := gin.Default()

	// 第一パラメータ（"/static"）はURLパス
	// 第二パラメータ（"./view"）はローカルのファイルパス
	route.Static("/static", "./view")

	// データベース設定の取得
	dbConfig := util.NewDBConfig()

	// sql.Open() は、データベースへの接続を確立する
	conn, err := sql.Open("postgres", dbConfig.ConnectionString)

	if err != nil {
		log.Fatalf("データベースの接続に失敗: %v", err)
	}
	defer conn.Close()

	log.Println("データベースの接続に成功")

	// ルートのページを提供するルート
	route.GET("/", func(c *gin.Context) {
		c.File("./view/index.html")
	})

	// サインアップフォームのページを提供するルート
	route.GET("/signup", func(c *gin.Context) {
		c.File("./view/signup.html")
	})

	// db.Queriesオブジェクトを作成
	// db.go には、データベース操作を行うための基本的な機能が定義される。このファイルには、DBTX インターフェースや New 関数などが含まれる
	dbQueries := db.New(conn)

	// signup.go には、サインアップフォームのハンドラーが定義される
	route.POST("/signup", api.SignupHandler(dbQueries))

	// アクティベーションリンクの処理を行うルート
	route.GET("/activate", api.ActivateUserHandler(dbQueries))

	if err := route.Run(":8080"); err != nil {
		log.Fatalf("起動に失敗しました: %v", err)
	}
}
