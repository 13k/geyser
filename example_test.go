package geyser_test

import (
	"os"
	"time"

	"github.com/13k/geyser/v2"
)

func ExampleNew_inline() {
	// Inline options
	_, _ = geyser.New(
		geyser.WithUserAgent("mylib 1.2.3"),
		geyser.WithKey("<api_key>"),
		geyser.WithTimeout(3*time.Second),
		geyser.WithRetryCount(5),
	)
}

func ExampleNew_conditional() {
	// Conditional options
	options := []geyser.ClientOption{
		geyser.WithKey("<api_key>"),
	}

	if os.Getenv("DEBUG") != "" {
		options = append(options, geyser.WithDebug())
	}

	_, _ = geyser.New(options...)
}
