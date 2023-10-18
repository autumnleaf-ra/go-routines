package goroutines

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestGomaxProcs(t *testing.T) {
	group := sync.WaitGroup{}

	for i := 0; i < 1003; i++ {
		go func() {
			defer group.Done()
			group.Add(1)
			time.Sleep(3 * time.Second)
		}()
	}

	cpu := runtime.NumCPU()
	fmt.Println("Total cpu", cpu)

	thread := runtime.GOMAXPROCS(-1)
	fmt.Println("Total thread", thread)

	totalGo := runtime.NumGoroutine()
	fmt.Println("Total Goroutine", totalGo)

}

func TestChange(t *testing.T) {
	group := sync.WaitGroup{}

	for i := 0; i < 1003; i++ {
		go func() {
			defer group.Done()
			group.Add(1)
			time.Sleep(3 * time.Second)
		}()
	}

	cpu := runtime.NumCPU()
	fmt.Println("Total cpu", cpu)

	runtime.GOMAXPROCS(20)
	thread := runtime.GOMAXPROCS(-1)
	fmt.Println("Total thread", thread)

	totalGo := runtime.NumGoroutine()
	fmt.Println("Total Goroutine", totalGo)

}
