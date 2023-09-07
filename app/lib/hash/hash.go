package hash

import "golang.org/x/crypto/bcrypt"

// Make 对字符串进行hash/**
func Make(text string) string {
	hashed, err := bcrypt.GenerateFromPassword([]byte(text), bcrypt.DefaultCost)
	if err != nil {
		return ""
	}
	return string(hashed)
}

// Check 检查传入的明文是否与hash过的密文一致
func Check(plainText string, hashedText string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedText), []byte(plainText))
	if err != nil {
		return false
	} else {
		return true
	}
}

func NeedHash(hashedText string) bool {
	hasCost, err := bcrypt.Cost([]byte(hashedText))
	return err != nil || hasCost != bcrypt.DefaultCost
}
