package functiontiming

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

// TimesContainer holds the times of each call
type TimesContainer struct {
	mu      *sync.RWMutex
	timeMap map[int]time.Time
}

func (t *TimesContainer) get(key int) time.Time {
	t.mu.RLock()
	defer t.mu.RUnlock()

	return t.timeMap[key]
}

func (t *TimesContainer) add(key int, value time.Time) {
	t.mu.Lock()
	defer t.mu.Unlock()

	t.timeMap[key] = value
}

// NewTimeMap returns a TimesContainer
func NewTimeMap() TimesContainer {
	t := TimesContainer{
		mu:      new(sync.RWMutex),
		timeMap: make(map[int]time.Time),
	}

	return t
}

// Elapsed Comment
func (t *TimesContainer) Elapsed(functionName string) func() {
	start := time.Now()
	r := getRand()
	t.add(r, start)

	log.Printf("%d started at %s", r, start)

	return func() {
		log.Printf("%s (%d) took %v\n", functionName, r, time.Since(t.get(r)))
	}
}

func getRand() int {
	rand.Seed(time.Now().UnixNano())
	min := 100000000
	max := 999999999
	return rand.Intn(max-min) + min
}
