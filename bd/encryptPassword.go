package bd

import "golang.org/x/crypto/bcrypt"

// EncryptPassword is a function to encrypt a password
func EncryptPassword(password string) (string, error) {
	cost := 8

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)

	return string(bytes), err
}
