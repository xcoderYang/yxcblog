package utils

import (
	"math/rand"
	"strconv"
	"time"
)

func SessionGenerate(username string)string{
	part1 := strconv.FormatInt(time.Now().UnixNano(), 10)
	part2 := strconv.FormatInt(rand.Int63(), 10)
	part3 := username
	salt := "Session Generate"
	return Sha256(part1+part2+part3+salt)
}