package api

import (
	"context"
	"log"
	"net/http"

	"golang_twitter/db"

	"github.com/gin-gonic/gin"
)

func ActivateUserHandler(dbQueries *db.Queries) gin.HandlerFunc {
	return func(c *gin.Context) {
		activationToken := c.Query("token")

		// トークンを使ってユーザーを検索
		user, err := dbQueries.GetUserByActivationToken(context.Background(), activationToken)
		if err != nil {
			log.Printf("アクティベーション中にエラーが発生しました(GetUserByActivationToken): %v", err)
			c.String(http.StatusInternalServerError, "アクティベーション中にエラーが発生しました。")
			return
		}

		if user == nil {
			c.String(http.StatusBadRequest, "無効なアクティベーショントークンです。")
			return
		}

		// ユーザーのアクティブ状態を更新
		err = dbQueries.ActivateUser(context.Background(), int(user.ID))
		if err != nil {
			log.Printf("アクティベーション中にエラーが発生しました(ActivateUser): %v", err)
			c.String(http.StatusInternalServerError, "ユーザーのアクティベーション中にエラーが発生しました。")
			return
		}

		c.String(http.StatusOK, "アカウントが正常にアクティベートされました。")
	}
}
