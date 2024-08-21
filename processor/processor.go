package processor

import (
	"alert-system/config"
	"alert-system/dispatcher"
	"alert-system/event"
	"fmt"
	"sync"
)

var EventQueue = make(chan event.Event, 100)
var eventWindows = make(map[string][]event.Event)
var mu sync.Mutex

func ProcessEvents(wg *sync.WaitGroup) {
	defer wg.Done()
	for e := range EventQueue {
		for _, conf := range config.AlertConfigList {
			if e.EventType == conf.EventType {
				processEvent(e, conf)
			}
		}
	}
}

func processEvent(e event.Event, conf config.AlertConfigEntry) {
	alertType := conf.AlertConfig.Type
	threshold := conf.AlertConfig.Count
	windowSize := conf.AlertConfig.WindowSizeSecs

	if alertType == "TUMBLING_WINDOW" {
		processTumblingWindow(e, threshold, windowSize, conf.DispatchStrategies)
	} else if alertType == "SLIDING_WINDOW" {
		processSlidingWindow(e, threshold, windowSize, conf.DispatchStrategies)
	}
}

func processTumblingWindow(e event.Event, threshold int, windowSize int64, strategies []config.DispatchStrategy) {
	mu.Lock()
	defer mu.Unlock()
	// fmt.Println(" started the process of TumblingWindow")
	currentTime := e.Timestamp
	windowStart := currentTime - (currentTime % windowSize)
	windowKey := fmt.Sprintf("%s_%d", e.EventType, windowStart)

	eventWindows[windowKey] = append(eventWindows[windowKey], e)

	// Clear old windows
	for key, events := range eventWindows {
		if len(events) > 0 && currentTime-int64(events[0].Timestamp) >= windowSize {
			delete(eventWindows, key)
		}
	}

	if len(eventWindows[windowKey]) >= threshold {
		dispatcher.TriggerAlert(e, strategies)
		eventWindows[windowKey] = []event.Event{}
	}
	// fmt.Println(" completed the process of TumblingWindow")
}

func processSlidingWindow(e event.Event, threshold int, windowSize int64, strategies []config.DispatchStrategy) {
	mu.Lock()
	defer mu.Unlock()
	// fmt.Println(" started the process of SlidingWindow")
	windowKey := e.EventType

	eventWindows[windowKey] = append(eventWindows[windowKey], e)

	// Remove events outside the sliding window
	for len(eventWindows[windowKey]) > 0 && e.Timestamp-int64(eventWindows[windowKey][0].Timestamp) > windowSize {
		eventWindows[windowKey] = eventWindows[windowKey][1:]
	}
	// fmt.Println("sliding window :eventWindows[windowKey]:", eventWindows[windowKey])
	// fmt.Println("sliding window threshold : ", threshold)
	if len(eventWindows[windowKey]) >= threshold {
		dispatcher.TriggerAlert(e, strategies)
	}
	// fmt.Println(" completed the process of SlidingWindow")
}
