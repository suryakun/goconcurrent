package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var cache = map[int]Book{}

var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {
	wg := &sync.WaitGroup{}
	mx := &sync.RWMutex{}
	cacheCh := make(chan Book)
	cacheDb := make(chan Book)
	for i := 0; i < 5; i++ {
		id := rnd.Intn(10) + 1
		wg.Add(2)
		go func(id int, wg *sync.WaitGroup, ch chan<- Book) {
			if b, ok := queryCache(id, mx); ok {
				ch <- b
			}
			wg.Done()
		}(id, wg, cacheCh)

		go func(id int, wg *sync.WaitGroup, ch chan<- Book) {
			if b, ok := queryDatabase(id, mx); ok {
				ch <- b
			}
			wg.Done()
		}(id, wg, cacheDb)

		go func(cacheCh, cacheDb <-chan Book) {
			select {
			case b := <-cacheCh:
				fmt.Println("From cache")
				fmt.Println(b)
			case b := <-cacheDb:
				fmt.Println("From database")
				fmt.Println(b)
			}
		}(cacheCh, cacheDb)

		fmt.Sprintf("Book not found with id %q", id)
		time.Sleep(150 * time.Millisecond)
	}
}

func queryCache(id int, mx *sync.RWMutex) (Book, bool) {
	mx.RLock()
	b, ok := cache[id]
	mx.RUnlock()
	return b, ok
}

func queryDatabase(id int, mx *sync.RWMutex) (Book, bool) {
	time.Sleep(100 * time.Millisecond)
	for _, b := range books {
		if b.ID == id {
			mx.Lock()
			cache[id] = b
			mx.Unlock()
			return b, true
		}
	}
	return Book{}, false
}
