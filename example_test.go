package geyser_test

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"time"

	"github.com/13k/geyser"
)

func Example_basic() {
	client, err := geyser.New(
		geyser.WithDebug(),
		geyser.WithKey("<steam_api_key>"),
		geyser.WithLanguage("en_US"),
	)

	if err != nil {
		log.Fatal(err)
	}

	dota2match, err := client.DOTA2Match(570)

	if err != nil {
		log.Fatal(err)
	}

	req, err := dota2match.GetTeamInfoByTeamID()

	if err != nil {
		log.Fatal(err)
	}

	result := make(map[string]interface{})

	req.SetResult(result)
	req.SetOptions(geyser.RequestOptions{
		Params: url.Values{"teams_requested": []string{"3"}},
	})

	resp, err := req.Execute()

	if err != nil {
		log.Fatal(err)
	}

	if !resp.IsSuccess() {
		log.Fatalf("HTTP error: %s", resp.Status())
	}

	fmt.Printf("%#v\n", result)
}

func ExampleNew_inline() {
	// Inline options
	geyser.New(
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

	geyser.New(options...)
}
