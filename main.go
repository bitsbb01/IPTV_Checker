package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	// new pool
	pool := New(runtime.NumCPU())

	addrs, err := ParseAddresses("")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	
	var wg sync.WaitGroup
	num := len(addrs)
	wg.Add(num)

	for i := 0; i < num; i++ {
		ck := NewChecker(addrs[i])
		go func()  {
			pool.Run(ck)
			wg.Done()
		}()
	}
	wg.Wait()

	pool.Shutdown()
}