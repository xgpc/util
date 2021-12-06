package util

import (
	"github.com/satori/go.uuid"
	"math/rand"
)

var randomChars = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
var randomNumber = []rune("0123456789")

func Uuid() string {
	//var u, err = uuid.NewV4()
	//if err != nil {
	//	panic(err)
	//}
	var u = uuid.NewV4()

	var str = u.String()
	return Replace(&str, "-", "")
}

func RandomStr(length int) string {
	b := make([]rune, length)
	randomCharsLen := len(randomChars)
	for i := range b {
		b[i] = randomChars[rand.Intn(randomCharsLen)]
	}
	return string(b)
}

func RandomNumber(length int) string {
	b := make([]rune, length)
	randomCharsLen := len(randomNumber)
	for i := range b {
		b[i] = randomNumber[rand.Intn(randomCharsLen)]
	}
	return string(b)
}

func RandomInt(length int) int {
	return rand.Intn(length)
}
