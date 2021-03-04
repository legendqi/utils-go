/* coding: utf-8
@Time :   2021/3/4 下午1:59
@Author : legend
@File :   scrypt.go
*/
package scrypt

import (
	"crypto/md5"
	"encoding/hex"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"io"
	"math/rand"
	"strings"
	"time"
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

func Md5Sum(reader io.Reader) (string, error) {
	var returnMD5String string
	hash := md5.New()
	if _, err := io.Copy(hash, reader); err != nil {
		return returnMD5String, err
	}
	hashInBytes := hash.Sum(nil)[:16]
	returnMD5String = hex.EncodeToString(hashInBytes)
	return returnMD5String, nil
}

func Md5SumString(input string) (string, error) {
	buffer := strings.NewReader(input)
	return Md5Sum(buffer)
}

func GetRandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
