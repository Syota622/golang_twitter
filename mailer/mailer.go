package mailer

import (
	"net/smtp"
)

func SendActivationEmail(to, activationToken string) error {
	from := "test@gmail.com"
	smtpHost := "mailcatcher"
	smtpPort := "1025" // Mailcatcherのデフォルトポート

	// メール本文の設定
	body := "アクティベーションリンクにクリックをお願いします:\n"
	body += "http://localhost:8080/activate?token=" + activationToken

	// メール送信
	msg := []byte("To: " + to + "\r\n" +
		"Subject: あなたのアカウント\r\n\r\n" +
		body + "\r\n")

	// smtpHost:port にメールを送信、smtpHost:port はメールサーバーのアドレス、nil は認証情報
	// from は送信元のメールアドレス、[]string{to} は送信先のメールアドレス、msg はメールの内容
	return smtp.SendMail(smtpHost+":"+smtpPort, nil, from, []string{to}, msg)
}
