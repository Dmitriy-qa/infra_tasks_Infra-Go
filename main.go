package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	td := testData{}
	var wg sync.WaitGroup
	generate(100, &td, &wg)
	wg.Wait()
	fmt.Println(len(td.phones))

}

type testData struct {
	phones []int
	mu     sync.Mutex
}

func (td *testData) add() {
	td.mu.Lock()
	defer td.mu.Unlock()
	td.phones = append(td.phones, randPhone())
}

func generate(n int, td *testData, wg *sync.WaitGroup) *testData {
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			td.add()
		}()
	}
	return td
}

func randPhone() int {
	return 89000000000 + rand.Intn(1000000000)

}
