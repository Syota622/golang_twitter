// util/config.go
package util

// データベースの設定を保持
type DBConfig struct {
	ConnectionString string
}

// 新しいDBConfigを生成
func NewDBConfig() *DBConfig {
	return &DBConfig{
		ConnectionString: "postgres://postgres:Passw0rd@db:5432/db?sslmode=disable",
	}
}
