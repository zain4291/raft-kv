package main

import "time"

func main() {
	store := NewStore()
	for i := 0; i < 100; i++ {
		go func() {
			store.Put("key", "value")
		}()
	}

	time.Sleep(time.Second)

}
