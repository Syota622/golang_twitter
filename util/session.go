package util

import (
	"errors"

	"github.com/gin-contrib/sessions"
)

// GetUserIDFromSession はセッションからユーザーIDを取得し、int32型にキャストして返す。
func GetUserIDFromSession(session sessions.Session) (int32, error) {
	userIDValue := session.Get("user_id")
	if userIDValue == nil {
		return 0, errors.New("ログインが必要です")
	}

	userID, ok := userIDValue.(int32)
	if !ok {
		return 0, errors.New("ユーザーIDの型が不正です")
	}

	return userID, nil
}
