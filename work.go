package main

import (
	"fmt"
	"sync"
)

// Worker must be implemented by types that want to use
// the work pool.
 type Worker interface {
     Do() error
 }

 // Pool provides a pool of goroutines that can execute any Worker
 // tasks that are submitted.
 type Pool struct {
     works chan Worker
     wg   sync.WaitGroup
 }

 // New creates a new work pool.
 func New(maxGoroutines int) *Pool {
     p := Pool{
         works: make(chan Worker),
     }

     p.wg.Add(maxGoroutines)
     for i := 0; i < maxGoroutines; i++ {
         go func() {
             for w := range p.works {
                 err := w.Do()
				 if err != nil {
					 fmt.Println(err)
				 }
             }
             p.wg.Done()
         }()
     }

     return &p
 }

 // Run submits work to the pool.
 func (p *Pool) Run(w Worker) {
     p.works <- w
 }

 // Shutdown waits for all the goroutines to shutdown.
 func (p *Pool) Shutdown() {
     close(p.works)
     p.wg.Wait()
 }