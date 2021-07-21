package decorater

import (
	"fmt"
	"log"
	"reflect"
	"runtime"
	"sync"
	"time"
)

type fibFunc func(int) int

func Fib(n int) int {
	if n <= 1 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}

func FibInParal(n int) int {
	if n <= 1 {
		return n
	}

	i1 := make(chan int)
	i2 := make(chan int)

	go func() {
		i1 <- FibInParal(n - 1)
	}()

	go func() {
		i2 <- FibInParal(n - 2)
	}()

	var res int
	for i := 0; i < 2; i++ {
		select {
		case i := <-i1:
			res += i
		case i := <-i2:
			res += i
		}
	}

	return res
}

func WithLogger(f fibFunc, logger *log.Logger) fibFunc {
	return func(i int) int {
		fn := func(n int) (result int) {
			defer func(t time.Time) {
				funcName := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
				logger.Printf("%s took=%v, n=%v, result=%v", funcName, time.Since(t), i, result)
			}(time.Now())

			return f(n)
		}

		return fn(i)
	}
}

func WithCache(f fibFunc, cache *sync.Map) fibFunc {
	return func(i int) int {
		fn := func(n int) int {
			key := fmt.Sprintf("n=%d", n)
			val, ok := cache.Load(key)
			if ok {
				return val.(int)
			}
			result := f(n)
			cache.Store(key, result)
			return result
		}

		return fn(i)
	}
}
