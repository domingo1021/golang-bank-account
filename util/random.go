package util

import (
	"math/rand"
	"strings"
	"time"
)

var currencies []string
const alphabet = "abcdefghijklmnopqrstuvwxys"

// to solve no const slice literal problem in golang.
func immutableCurrency() []string {
	return []string{"USD", "EUR", "TWD"}
}

func init() {
	currencies = immutableCurrency()
	rand.Seed(time.Now().UnixNano())
}


// RandomInt generate a random integer between min and max.
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max - min + 1)
}

func RandomUpperCase() byte {
	return byte(RandomInt(65, 90))
}

func RandomLowerCase() byte {
	return byte(RandomInt(97, 122))
}

var capitalWiseFactory = map[int64]func()(byte){
	0: RandomLowerCase,
	1: RandomUpperCase,
}

// RandomString generate a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	
	for i := 0; i < n; i++ {
		capitalWise := RandomInt(0, 1)
		c := capitalWiseFactory[capitalWise]()
		sb.WriteByte(c)
	}
	return sb.String();
}

func RandomOwner() string {
	return RandomString(6)
}

func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

func RandomCurrency() string {
	n := len(currencies)
	return currencies[rand.Intn(n)]
}