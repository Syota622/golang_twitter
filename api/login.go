package api

import (
	"golang_twitter/db"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// LoginHandler は、ログインフォームのハンドラー
func LoginHandler(dbQueries *db.Queries) gin.HandlerFunc {
	return func(c *gin.Context) {
		email := c.PostForm("email")
		password := c.PostForm("password")

		// メールアドレスが存在しない場合は、ログインを許可しない
		user, err := dbQueries.GetUserByEmail(c, email)
		if err != nil {
			log.Printf("ログイン中にエラーが発生しました(GetUserByEmail): %v", err)
			c.HTML(http.StatusUnauthorized, "login.html", gin.H{"Error": "無効なメールアドレスまたはパスワード"})
			return
		}

		// is_activeがfalseまたは無効の場合はログインを許可しない
		if !user.IsActive.Valid || !user.IsActive.Bool {
			log.Printf("非アクティブなユーザーのログイン試行: %v", email)
			c.HTML(http.StatusUnauthorized, "login.html", gin.H{"Error": "このアカウントはアクティブではありません。"})
			return
		}

		// パスワードが一致しない場合は、ログインを許可しない
		err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
		if err != nil {
			log.Printf("ログイン中にエラーが発生しました(CompareHashAndPassword): %v", err)
			c.HTML(http.StatusUnauthorized, "login.html", gin.H{"Error": "無効なメールアドレスまたはパスワード"})
			return
		}

		// ログインに成功した場合は、ルートページにリダイレクト
		c.Redirect(http.StatusFound, "/")
	}
}
