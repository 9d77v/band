package utils

import (
	"crypto/rand"
	"math/big"
)

const characters = "abcdefghijklmnopqrstuvwxyz0123456789"

func RandomString(n int) string {
	p := make([]byte, n)
	for i := 0; i < n; i++ {
		m, _ := rand.Int(rand.Reader, big.NewInt(36))
		p[i] = characters[m.Int64()]
	}
	return string(p)
}
