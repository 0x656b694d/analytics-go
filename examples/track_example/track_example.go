// A basic example of using Amplitude Go SDK to track an event.

package main

// Import amplitude package
import (
	"github.com/amplitude/Amplitude-Go/amplitude"
)

// Define your callback function (optional)
func callbackFunc(e string, code int, message string) {
	println(e)
	println(code, message)
}

func main() {

	config := amplitude.NewConfig("your_api_key")

	// Config callback function (optional)
	client := amplitude.NewClient(config)

	client.Add(amplitude.NewContextPlugin())

	// Create a BaseEvent instance
	event := amplitude.Event{
		EventOptions: amplitude.EventOptions{DeviceID: "go-device-id", UserID: "go-user-id"},
		EventType:    "go-event-type",
	}

	// Track an event
	client.Track(event)

	// Revenue Tracking
	revenueObj := amplitude.Revenue{
		Price:    9.9,
		Quantity: 2,
	}
	client.Revenue(revenueObj, amplitude.EventOptions{UserID: "revenue-test-user-id"})

	// Track user properties
	identifyObj := amplitude.Identify{}
	identifyObj.Set("location", "LAX")
	client.Identify(identifyObj, amplitude.EventOptions{UserID: "identify-test-user-id"})

	// Flush the event buffer
	client.Flush()

	// Shutdown the client
	client.Shutdown()
}
