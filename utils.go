package main

import (
	"math/rand"
	"os"
	"time"
)

var (
	randChars = []rune{}
	randBase  = 0
)

func init() {
	for c := '0'; c <= '9'; c++ {
		randChars = append(randChars, c)
	}
	for c := 'a'; c <= 'z'; c++ {
		randChars = append(randChars, c)
	}
	for c := 'A'; c <= 'Z'; c++ {
		randChars = append(randChars, c)
	}

	randBase = len(randChars)
}

func getEnv(key string, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue
	}
	return value
}

func randomString(length int) (string, error) {
	buf := make([]byte, length)
	randomInit()

	for idx, _ := range buf {
		randIdx := randomNumber(randBase)
		buf[idx] = byte(randChars[randIdx])
	}

	return string(buf), nil
}

func randomInit() {
	rand.Seed(time.Now().Unix())
}

func randomNumber(max int) int {
	return rand.Intn(max)
}
