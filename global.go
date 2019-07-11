package lrand

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.001
// @date    2019-07-11

import (
	"time"
)

var (
	defGen       *Mersenne
	output       chan int64
	GlobalReseed int
)

func init() {
	defGen = New()
	defGen.Seed(time.Now().UnixNano())

	output = make(chan int64, 1024)

	GlobalReseed = 10000

	go generator()
}

func Next() int64 {
	return <-output
}

func generator() {

	i := 0

	for {
		output <- defGen.Int63()
		i = (i + 1) % GlobalReseed
		if i == 0 {
			defGen.Seed(time.Now().UnixNano())
		}
	}
}
