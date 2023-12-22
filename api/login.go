package api

import (
	"golang_twitter/db"
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid" // UUIDを生成するためのパッケージ
	"golang.org/x/crypto/bcrypt"
)

// LoginHandler は、ログインフォームのハンドラー
func LoginHandler(dbQueries *db.Queries) gin.HandlerFunc {
	return func(c *gin.Context) {
		email := c.PostForm("email")
		password := c.PostForm("password")

		// メールアドレスが存在するかどうかを確認
		user, err := dbQueries.GetUserByEmail(c, email)

		// メールアドレスが存在しない場合は、ログインを許可しない
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

		// セッションの取得、UUIDを生成
		session := sessions.Default(c)   // &{mysession 0xc0004e2000 0xc00049c018 <nil> false 0xc0004e2100}
		sessionId := uuid.New().String() // 847b656b-4238-44e0-8985-0ee943b4af97

		session.Set("session_id", sessionId) // セッションIDをセット
		session.Set("user_id", user.ID)      // ユーザーIDをセット

		// 変更されたセッションをRedisに保存します。
		if err := session.Save(); err != nil {
			// セッションの保存に失敗した場合の処理
			c.HTML(http.StatusInternalServerError, "error.html", gin.H{"Error": "セッションの保存に失敗しました"})
			return
		}

		// ログインに成功した場合は、ルートページにリダイレクト
		c.Redirect(http.StatusFound, "/")
	}
}
