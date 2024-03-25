package util

import (
	"math/rand"
	"time"
)

func Random4Num() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(10000)
}
