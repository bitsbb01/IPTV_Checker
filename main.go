package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	// load and parse m3u file
	channels, err := ParseAddresses("channels.m3u")
	if err != nil {
		fmt.Println(err)
		return
	}

	// new pool
	pool := New(runtime.NumCPU())
	
	var wg sync.WaitGroup
	num := len(channels)
	wg.Add(num)

	for i := 0; i < num; i++ {
		ck := NewChecker(channels[i])
		go func()  {
			pool.Run(ck)
			wg.Done()
		}()
	}
	wg.Wait()

	pool.Shutdown()
}