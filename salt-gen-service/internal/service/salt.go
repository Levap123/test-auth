package service

import (
	"math/rand"
	"time"
)

type Salt struct{}

func NewSalt() *Salt {
	return &Salt{}
}

func (s *Salt) GenerateSalt() string {
	rand.Seed(time.Now().UnixNano())
	chars := []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	result := make([]byte, 12)
	for i := range result {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return string(result)
}
