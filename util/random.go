package util

import (
	"math/rand"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())

}

//RandomInt generates a random Integer bw min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}
