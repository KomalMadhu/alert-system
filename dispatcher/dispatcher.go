package dispatcher

import (
	"alert-system/config"
	"alert-system/event"
	"fmt"
)

func TriggerAlert(e event.Event, strategies []config.DispatchStrategy) {
	// fmt.Println(" TriggerAlert called")
	for _, strategy := range strategies {
		if strategy.Type == "CONSOLE" {
			// fmt.Println(" TriggerAlert called of type CONSOLE")
			fmt.Println("console ouput ", strategy.Message)
		} else if strategy.Type == "EMAIL" {
			// fmt.Println(" TriggerAlert called of type email")
			sendEmail(e, strategy.Subject)
		}
	}
}

func sendEmail(e event.Event, subject string) {
	fmt.Printf("Sending email with subject: %s\n", subject)
}
