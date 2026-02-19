package utils

import (
	"crypto/rand"
	"math/big"
)

const characters = "abcdefghijklmnopqrstuvwxyz0123456789"

func RandomString(n int) string {
	p := make([]byte, n)
	for i := range n {
		m, _ := rand.Int(rand.Reader, big.NewInt(36))
		p[i] = characters[m.Int64()]
	}
	return string(p)
}

func CountCharacters(s string) int {
	count := 0
	for _, _ = range s {
		count++
	}
	return count
}
