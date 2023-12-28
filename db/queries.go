package db

import (
	"context"
	"database/sql"
)

// GetUserByActivationToken は指定されたアクティベーショントークンを持つユーザーを検索
func (q *Queries) GetUserByActivationToken(ctx context.Context, token string) (*User, error) {

	// models.go で定義された User 構造体を使用して、データベースからユーザーを取得
	var user User
	err := q.db.QueryRowContext(ctx, "SELECT * FROM users WHERE activation_token = $1", token).Scan(
		&user.ID,
		&user.Email,
		&user.PasswordHash,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.IsActive,
		&user.ActivationToken,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			// トークンに一致するユーザーが見つからない場合
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

// ActivateUser は指定されたユーザーIDのユーザーをアクティブ状態に更新
// $1 はユーザーID
func (q *Queries) ActivateUser(ctx context.Context, userID int) error {
	_, err := q.db.ExecContext(ctx, "UPDATE users SET is_active = true WHERE id = $1", userID)
	return err
}

// GetUserByEmail は指定されたメールアドレスを持つユーザーを検索
func (q *Queries) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	var user User
	err := q.db.QueryRowContext(ctx, "SELECT * FROM users WHERE email = $1", email).Scan(
		&user.ID,
		&user.Email,
		&user.PasswordHash,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.IsActive,
		&user.ActivationToken,
	)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
