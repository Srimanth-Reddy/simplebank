package util

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var names = []string{
	"Alice", "Bob", "Charlie", "David", "Eve", "Frank", "Grace",
	"Hannah", "Ivy", "Jack", "Kevin", "Laura", "Mike", "Nina",
	"Oscar", "Paul", "Quinn", "Rachel", "Steve", "Tina", "Uma",
	"Victor", "Wendy", "Xander", "Yara", "Zane",
}

// RandomInt generates a random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomOwner generates a random owner name
func RandomOwner() string {
	n := len(names)
	return names[rand.Intn(n)]
}

// RandomMoney generates a random amount of money
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

// RandomCurrency generates a random currency code
func RandomCurrency() string {
	currencies := []string{"USD", "EUR", "INR"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}
