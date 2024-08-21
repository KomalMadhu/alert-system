package config

type AlertConfig struct {
	Type           string
	Count          int
	WindowSizeSecs int64
}

type DispatchStrategy struct {
	Type    string
	Message string
	Subject string
}

type AlertConfigEntry struct {
	Client             string
	EventType          string
	AlertConfig        AlertConfig
	DispatchStrategies []DispatchStrategy
}

var AlertConfigList = []AlertConfigEntry{
	{
		Client:    "X",
		EventType: "PAYMENT_EXCEPTION",
		AlertConfig: AlertConfig{
			Type:           "TUMBLING_WINDOW",
			Count:          2,
			WindowSizeSecs: 5,
		},
		DispatchStrategies: []DispatchStrategy{
			{Type: "CONSOLE", Message: "Issue in payment"},
			{Type: "EMAIL", Subject: "Payment exception threshold breached"},
		},
	},
	{
		Client:    "X",
		EventType: "USERSERVICE_EXCEPTION",
		AlertConfig: AlertConfig{
			Type:           "SLIDING_WINDOW",
			Count:          2,
			WindowSizeSecs: 5,
		},
		DispatchStrategies: []DispatchStrategy{
			{Type: "CONSOLE", Message: "Issue in user service"},
		},
	},
}
