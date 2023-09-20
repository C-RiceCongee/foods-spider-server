package pkg

import "golang.org/x/crypto/bcrypt"

// EncryptionPassword 给密码加密！
func EncryptionPassword(pwd string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
}

/*
ComparePasswords 对比明文和加密密码是否匹配！
*/
func ComparePasswords(hashedPassword string, plaintextPassword string) bool {
	// Returns true on success, pwd1 is for the database.
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plaintextPassword))
	if err != nil {
		return false
	} else {
		return true
	}
}
