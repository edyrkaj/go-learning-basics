package checker

import "time"

// Result represents the status of a URL check.
type Result struct {
	URL     string
	Status  string
	Latency time.Duration
}
