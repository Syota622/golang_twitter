package main

import (
	"database/sql"
	"golang_twitter/api"
	"golang_twitter/db"
	"golang_twitter/util"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq" // PostgreSQLドライバー

	"github.com/gin-contrib/sessions"       // Ginフレームワーク用のセッションミドルウェアを提供するパッケージ
	"github.com/gin-contrib/sessions/redis" // セッションミドルウェアがRedisをバックエンドとして使用するための拡張パッケージ
)

func main() {
	route := gin.Default()

	// 第一パラメータ（"/static"）はURLパス
	// 第二パラメータ（"./view"）はローカルのファイルパス
	route.Static("/static", "./view")

	// 'view' ディレクトリ内のHTMLテンプレートをロード
	route.LoadHTMLGlob("view/*.html")

	// redis.NewStore() は、Redisをバックエンドとして使用するセッションストアを作成します。
	// "10" は セッションストアに保存できるセッションの最大数、"redis" は docker-compose.yml 内で定義したサービス名に対応します。
	store, err := redis.NewStore(10, "tcp", "redis:6379", "", []byte("secret"))
	if err != nil {
		log.Fatalf("Redisストアの設定に失敗しました: %v", err)
	}

	// セッションのオプションを設定
	store.Options(sessions.Options{
		MaxAge: 1000, // 有効期限
	})

	// セッションのミドルウェアを使用
	route.Use(sessions.Sessions("mysession", store))

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

	// ログインフォームのページを提供するルート
	route.GET("/login", func(c *gin.Context) {
		// c.HTML を使用してテンプレートをレンダリングする
		c.HTML(http.StatusOK, "login.html", gin.H{})
	})

	// ログイン処理のルート
	route.POST("/login", api.LoginHandler(dbQueries))

	// ツイートを投稿するルート
	// 認証が必要なAPIのためのグループ
	authGroup := route.Group("/auth")
	authGroup.Use(AuthMiddleware())
	authGroup.GET("/tweet", func(c *gin.Context) {
		c.HTML(http.StatusOK, "tweet.html", gin.H{})
	})
	authGroup.POST("/tweet", api.PostTweetHandler(dbQueries))

	// ツイートの取得API
	authGroup.GET("/tweets", api.GetTweetsHandler(dbQueries))

	// ツイート一覧ページを提供するルート
	authGroup.GET("/tweets-page", func(c *gin.Context) {
		c.HTML(http.StatusOK, "tweets.html", gin.H{})
	})

	if err := route.Run(":8080"); err != nil {
		log.Fatalf("起動に失敗しました: %v", err)
	}
}

// AuthMiddleware は認証を担当するミドルウェア
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		userID := session.Get("user_id")
		if userID == nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "認証が必要です"})
			c.Abort()
			return
		}
		c.Next()
	}
}
