package main

import (
	"context"
	"database/sql"
	"golang_twitter/db"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt" // パスワードのハッシュ化

	_ "github.com/lib/pq" // PostgreSQLドライバー
)

func main() {
	route := gin.Default()

	// 第一パラメータ（"/static"）はURLパス
	// 第二パラメータ（"./view"）はローカルのファイルパス
	route.Static("/static", "./view")

	// PostgreSQLデータベースに接続。dbは、compose.ymlで定義したコンテナ名
	connStr := "postgres://postgres:Passw0rd@db:5432/db?sslmode=disable"

	// sql.Open() は、データベースへの接続を確立する
	conn, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatalf("データベースの接続に失敗: %v", err)
	}
	defer conn.Close()

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

	log.Println("データベースの接続に成功")

	route.POST("/signup", func(c *gin.Context) {
		email := c.PostForm("email")
		password := c.PostForm("password") // パスワードはハッシュ化する

		// パスワードをハッシュ化
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			log.Fatalf("パスワードのハッシュ化に失敗しました: %v", err)
		}

		// ユーザーを作成
		err = dbQueries.CreateUser(context.Background(), db.CreateUserParams{
			Email:        email,
			PasswordHash: string(hashedPassword),
		})
		if err != nil {
			log.Fatalf("ユーザーの作成に失敗しました: %v", err)
		}

		// http://localhost:8080/ にリダイレクトされるため、index.html が表示される
		c.Redirect(http.StatusFound, "/")
	})

	if err := route.Run(":8080"); err != nil {
		log.Fatalf("起動に失敗しました: %v", err)
	}
}
