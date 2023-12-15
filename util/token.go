package util

import (
	"crypto/rand"
	"encoding/hex"
	"log" // ログ出力のためのパッケージをインポート
)

// GenerateActivationToken はランダムなアクティベーショントークンを生成します。
func GenerateActivationToken() (string, error) {
	byteSize := 16

	bytes := make([]byte, byteSize)
	_, err := rand.Read(bytes)
	if err != nil {
		log.Printf("トークン生成時にエラー: %v", err)
		return "", err
	}

	token := hex.EncodeToString(bytes)
	log.Printf("生成されたトークン: %s", token) // トークンのログ出力
	return token, nil
}
