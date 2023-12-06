package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// Ginのデフォルトエンジンを初期化
	r := gin.Default()

	// 静的ファイルのディレクトリを設定
	r.Static("/", "./view")

	// サーバーを起動
	err := r.Run(":8080") // デフォルトで8080ポートでリッスン
	if err != nil {
		panic(err)
	}
}
