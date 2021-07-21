package main

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/tao-yi/golang-design-patterns/decorater"
)

func main() {
	logger := log.New(os.Stdout, "test: ", 1)
	fibWithLogger := decorater.WithLogger(decorater.Fib, logger)
	fibWithLogger(40)

	var cache sync.Map
	fibWithCache := decorater.WithLogger(decorater.WithCache(decorater.Fib, &cache), logger)
	fibWithCache(40)

	fmt.Println(cache)

	// decoratedFibInParal := decorater.WithLogger(decorater.FibInParal, logger)
	// decoratedFibInParal(40)

}
