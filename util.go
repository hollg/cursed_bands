package main

import (
	"math/rand"
	"time"
)

func genBool() bool {
	source := rand.NewSource(time.Now().Unix())
	r := rand.New(source)
	i := r.Int()
	return i%2 == 0
}
