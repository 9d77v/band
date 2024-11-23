package id_util

import (
	"math"
	"strings"
)

const base64Chars = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz_-"

func ToInt64IDs(ids []string) []int64 {
	result := []int64{}
	for _, id := range ids {
		result = append(result, ToInt64ID(id))
	}
	return result
}

func ToStringIDs(ids []int64) []string {
	result := []string{}
	for _, id := range ids {
		result = append(result, ToStringID(id))
	}
	return result
}

func ToInt64ID(id string) int64 {
	var result int64
	for i := len(id) - 1; i >= 0; i-- {
		charIndex := strings.Index(base64Chars, string(id[i]))
		result += int64(charIndex) * int64(math.Pow(64, float64(len(id)-1-i)))
	}
	return result
}

func ToStringID(id int64) string {
	if id == 0 {
		return "0"
	}
	var result strings.Builder
	for id > 0 {
		remainder := id % 64
		result.WriteByte(base64Chars[remainder])
		id = id / 64
	}
	return reverseString(result.String())
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
