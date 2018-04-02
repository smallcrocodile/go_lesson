package persist

import "log"

func ItemSaver() chan interface{} {
	out := make(chan interface{})
	go func() {
		itemCount := 0
		for {
			itemCount++
			item := <-out
			log.Printf("ItemSaver got Item: #%d:%v", itemCount, item)
		}
	}()
	return out
}
