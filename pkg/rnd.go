package pkg

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func Rnd(size uint) []byte {
	b := make([]byte, size)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return b
}

func RndString(size uint) (string, error) {
	log.Print("message is visible only with verbose flag")
	if rand.Intn(2) == 1 {
		return "", fmt.Errorf("just a sample error")
	}
	return string(Rnd(size)), nil
}
