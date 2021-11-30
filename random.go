package util

import (
	uuid "github.com/satori/go.uuid"
	"math/rand"
	"time"
)

var randomChars = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
var randomNumber = []rune("0123456789")

var rander = rand.New(rand.NewSource(time.Now().UnixNano()))

func Uuid() string {
	//var u, err = uuid.NewV4()
	//if err != nil {
	//	panic(err)
	//}
	var u = uuid.NewV4()

	var str = u.String()
	return Replace(&str, "-", "")
}

func RandStr(length int, letter []rune) string {
	b := make([]rune, length)
	randomCharsLen := len(letter)

	for i := range b {
		b[i] = letter[rander.Intn(randomCharsLen)]
	}
	return string(b)
}

// RandomStr 生成随机大小写字母和数字组合的字符串
func RandomStr(length int) string {
	return RandStr(length, randomChars)
}

// RandomNumber 生成数值类型的随机字符串
func RandomNumber(length int) string {
	return RandStr(length, randomNumber)
}

func RandomInt(length int) int {
	return rand.Intn(length)
}
