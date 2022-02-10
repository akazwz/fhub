package crypt

import "golang.org/x/crypto/bcrypt"

// HashAndSortPwd 加密密码
func HashAndSortPwd(pwdStr string) (err error, hashedPwd string) {
	pwd := []byte(pwdStr)
	password, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		return err, ""
	}
	hashedPwd = string(password)
	return
}

// ComparePassword 验证密码
func ComparePassword(hashPwd, plainPwd string) bool {
	byteHash := []byte(hashPwd)
	bytePwd := []byte(plainPwd)

	err := bcrypt.CompareHashAndPassword(byteHash, bytePwd)
	if err != nil {
		return false
	}
	return true
}
