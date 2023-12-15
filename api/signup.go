// api/signup.go
package api

import (
	"context"
	"database/sql"
	"log"
	"net/http"

	"golang_twitter/db"
	"golang_twitter/util"
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

		// アクティベーショントークンの生成
		activationToken, err := util.GenerateActivationToken()
		if err != nil {
			log.Fatalf("アクティベーショントークンの生成に失敗しました: %v", err)
		}
		if activationToken == "" {
			log.Fatalf("生成されたアクティベーショントークンが空です")
		}

		// ユーザーを作成
		err = dbQueries.CreateUser(context.Background(), db.CreateUserParams{
			Email:           email,
			PasswordHash:    string(hashedPassword),
			IsActive:        sql.NullBool{Bool: false, Valid: true}, // ユーザーはまだ非アクティブ
			ActivationToken: sql.NullString{String: activationToken, Valid: true},
		})
		if err != nil {
			log.Fatalf("ユーザーの作成に失敗しました: %v", err)
		}

		// http://localhost:8080/ にリダイレクト
		c.Redirect(http.StatusFound, "/")
	}
}
