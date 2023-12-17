package util

import (
	"fmt"
	"time"
)

func WithTimings(f func() interface{}) interface{} {
	start := time.Now()
	ret := f()
	fmt.Println("done in", time.Now().Sub(start))
	return ret
}
