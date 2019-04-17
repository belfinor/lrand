package lrand

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.000
// @date    2019-04-17

import (
	"time"
)

var defGen *Mersenne
var output chan int64

func init() {
	defGen = New()
	defGen.Seed(time.Now().UnixNano())

	output = make(chan int64, 100)

	go generator()
}

func Next() int64 {
	return <-output
}

func generator() {
	for {
		output <- defGen.Int63()
	}
}
