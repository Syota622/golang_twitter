// mail/sender.go
package validation

import (
	"net/mail"
	"regexp"
	"strings"
)

// メールアドレスはフォーマットのチェックを行うことで、有効なメールアドレスかどうかを確認する
func ValidateEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

// パスワードは8文字以上以下を満たすこと
// 半角英数字が含まれる（英字だけはx,数字だけもx)
// 英字は小文字大文字混合(小文字だけはx)
// 以下の記号が1文字以上含まれる
// !?-_
func ValidatePassword(password string) string {
	var (
		minLenRegex      = regexp.MustCompile(`^.{8,}$`)
		numberRegex      = regexp.MustCompile(`[0-9]`)
		upperRegex       = regexp.MustCompile(`[A-Z]`)
		lowerRegex       = regexp.MustCompile(`[a-z]`)
		specialCharRegex = regexp.MustCompile(`[!?\-_]`)
	)

	var errorMsgs []string

	if !minLenRegex.MatchString(password) {
		errorMsgs = append(errorMsgs, "最低8文字以上である必要があります")
	}
	if !numberRegex.MatchString(password) {
		errorMsgs = append(errorMsgs, "少なくとも1つの数字を含む必要があります")
	}
	if !upperRegex.MatchString(password) {
		errorMsgs = append(errorMsgs, "少なくとも1つの大文字を含む必要があります")
	}
	if !lowerRegex.MatchString(password) {
		errorMsgs = append(errorMsgs, "少なくとも1つの小文字を含む必要があります")
	}
	if !specialCharRegex.MatchString(password) {
		errorMsgs = append(errorMsgs, "特殊文字 (!, ?, -, _) のうち少なくとも1つを含む必要があります")
	}

	if len(errorMsgs) == 0 {
		return "" // バリデーションチェックに成功
	}
	return strings.Join(errorMsgs, ", ")
}
