// api/tweet.go
package api

import (
	"golang_twitter/db"
	"golang_twitter/util"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func PostTweetHandler(dbQueries *db.Queries) gin.HandlerFunc {
	return func(c *gin.Context) {
		// ログインしているユーザーのセッション情報を取得
		// セッションからユーザーIDを取得
		session := sessions.Default(c)
		userID, err := util.GetUserIDFromSession(session)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
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
