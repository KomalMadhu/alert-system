package main

import (
	"alert-system/event"
	"alert-system/processor"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	go processor.ProcessEvents(&wg)

	// Simulate events
	go func() {
		for i := 0; i < 5; i++ {
			e1 := event.Event{
				Client:    "X",
				EventType: "PAYMENT_EXCEPTION",
				Timestamp: time.Now().Unix(),
				Details:   "An error occurred in payment",
			}
			processor.EventQueue <- e1

			e2 := event.Event{
				Client:    "X",
				EventType: "USERSERVICE_EXCEPTION",
				Timestamp: time.Now().Unix(),
				Details:   "An error occurred in user service",
			}
			processor.EventQueue <- e2

			time.Sleep(1 * time.Second)
		}
		close(processor.EventQueue)
	}()

	wg.Wait()
}
