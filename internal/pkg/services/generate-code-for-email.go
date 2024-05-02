package services

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateCodeForEmail() string {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	return fmt.Sprintf("%04d", rand.Intn(10000))
}
