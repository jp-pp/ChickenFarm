package rand

import (
	"math/rand"
	"time"
)

func RandomUint(min uint, max uint) uint {
	rand.Seed(time.Now().Unix())
	if min > max {
		return min
	} else {
		return uint(rand.Int63n(int64(max-min))) + min
	}
}