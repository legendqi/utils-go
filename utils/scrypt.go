/* coding: utf-8
@Time :   2021/3/4 下午1:59
@Author : legend
@File :   scrypt.go
*/
package utils

import (
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

func ScryptPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return ""
	}
	return string(hash)
}

func ComparePassword(postPwd string, userPwd string) bool {
	compareResult := bcrypt.CompareHashAndPassword([]byte(userPwd), []byte(postPwd))
	return compareResult == nil
}

func ProductToken() string {
	uid := uuid.NewV4()
	token := strings.Replace(uid.String(), "-", "", -1)
	return token
}
