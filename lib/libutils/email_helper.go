package libutils

import (
	"regexp"
)

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

func IsEmail(email string) bool {
	if email == "" {
		return false
	}

	// 邮箱长度限制检查，一般不超过254个字符
	if len(email) > 254 {
		return false
	}

	return emailRegex.MatchString(email)
}
