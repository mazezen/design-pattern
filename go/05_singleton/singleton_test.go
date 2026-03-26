package singleton

import (
	"fmt"
	"sync"
	"testing"
)

// go test -v singleton_test.go singleton.go

const parCount = 100

func TestSingleton(t *testing.T) {
	ins1 := GetInstance()
	ins2 := GetInstance()
	fmt.Printf("ins1: %p\n", ins1)
	fmt.Printf("ins2: %p\n", ins2)
	if ins1 != ins2 {
		t.Fatal("instance is not queal")
	}
}

func TestParallelSingleton(t *testing.T) {
	start := make(chan struct{})
	wg := sync.WaitGroup{}
	wg.Add(parCount)

	instances := [parCount]Singleton{}
	for i := 0; i < parCount; i++ {
		go func (index int)  {
			<-start
			instances[index] = GetInstance()
			wg.Done()
		}(i)
	}

	close(start)
	wg.Wait()
	for i := 1; i < parCount; i++ {
		fmt.Printf("instance [%d]: %p\n", i, instances[i])
		if instances[i] != instances[i-1] {
			t.Fatal("instance is not equal")
		}
	}
}