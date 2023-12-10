// api/signup.go
package api

import (
	"context"
	"log"
	"net/http"

	"golang_twitter/db"
	"golang_twitter/validation"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// サインアップフォームのハンドラー
func SignupHandler(dbQueries *db.Queries) gin.HandlerFunc {
	return func(c *gin.Context) {
		email := c.PostForm("email")
		password := c.PostForm("password")

		// メールアドレスのバリデーション
		if !validation.ValidateEmail(email) {
			c.String(http.StatusBadRequest, "無効なメールアドレスです。")
			return
		}

		// パスワードのバリデーション
		validationMsg := validation.ValidatePassword(password)
		if validationMsg != "" {
			c.String(http.StatusBadRequest, "パスワードの要件を満たしていません: %s", validationMsg)
			return
		}

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

		// http://localhost:8080/ にリダイレクト
		c.Redirect(http.StatusFound, "/")
	}
}
