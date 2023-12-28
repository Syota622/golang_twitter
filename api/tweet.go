// api/tweet.go
package api

import (
	"golang_twitter/db"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func PostTweetHandler(dbQueries *db.Queries) gin.HandlerFunc {
	return func(c *gin.Context) {
		// ログインしているユーザーのセッション情報を取得
		session := sessions.Default(c)
		userIDValue := session.Get("user_id")
		if userIDValue == nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "ログインが必要です"})
			return
		}

		// userID を int32 にキャスト
		userID, ok := userIDValue.(int32)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "ユーザーIDの型が不正です"})
			return
		}

		// リクエストからツイートの内容を取得
		text := c.PostForm("text")

		// ツイートの文字数制限をチェック
		if len(text) > 140 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ツイートは140文字以内である必要があります"})
			return
		}

		// ツイートをデータベースに保存
		tweet, err := dbQueries.CreateTweet(c, db.CreateTweetParams{
			UserID: userID, // キャストした userID を使用
			Text:   text,
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "ツイートの投稿に失敗しました"})
			return
		}

		// レスポンスとしてツイートを返す
		c.JSON(http.StatusOK, tweet)
	}
}
