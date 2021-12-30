package services

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func Hash(p string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)
	return string(hash)
}

func ComparePasswords(hashedPwd string, plainPwd string) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, []byte(plainPwd))
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
