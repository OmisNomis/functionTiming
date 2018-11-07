 package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

type times struct {
	mu      *sync.RWMutex
	timeMap map[int]time.Time
}

var a = times{
	mu:      new(sync.RWMutex),
	timeMap: make(map[int]time.Time),
}

func add(key int, value time.Time) {
	a.mu.Lock()
	defer a.mu.Unlock()

	a.timeMap[key] = value
}

func get(key int) time.Time {
	a.mu.RLock()
	defer a.mu.RUnlock()

	return a.timeMap[key]
}

func main() {
	ticker := time.NewTicker(500 * time.Millisecond)
	for range ticker.C {
		go func() {
			defer elapsed("main")()
			time.Sleep(8 * time.Second)
		}()
	}
}

func getRand() int {
	rand.Seed(time.Now().UnixNano())
	min := 100000000
	max := 999999999
	return rand.Intn(max-min) + min
}

func elapsed(what string) func() {
	start := time.Now()
	r := getRand()
	add(r, start)

	log.Printf("%d started at %s", r, start)

	return func() {
		log.Printf("%s (%d) took %v\n", what, r, time.Since(get(r)))
	}
}
