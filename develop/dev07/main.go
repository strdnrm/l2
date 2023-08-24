package main

import (
	"fmt"
	"time"
)

func or(channels ...<-chan interface{}) <-chan interface{} {
	if len(channels) == 0 {
		return nil
	}
	orDone := make(chan interface{})

	go func() {
		defer close(orDone)
		for _, ch := range channels {
			fmt.Println("reading")
			_, ok := <-ch
			if !ok {
				return
			}
		}
	}()
	return orDone
}

func main() {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()

	sig1 := sig(2 * time.Second)
	sig2 := sig(5 * time.Second)
	sig3 := sig(1 * time.Second)
	sig4 := sig(1 * time.Hour)
	sig5 := sig(1 * time.Minute)

	<-or(sig1, sig2, sig3, sig4, sig5)
	fmt.Printf("Done after %v\n", time.Since(start))
}
