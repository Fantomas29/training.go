package main

import (
	"math/rand"
)

func GetRandomChar() byte {

	return byte(ALLOWED_CHARMAP[rand.Intn(len(ALLOWED_CHARMAP))])
}
