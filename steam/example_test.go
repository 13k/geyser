package steam_test

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/13k/geyser/v2"
	"github.com/13k/geyser/v2/steam"
)

func Example_result() {
	client, err := steam.New(
		geyser.WithDebug(),
		geyser.WithKey("<steam_api_key>"),
		geyser.WithLanguage("en_US"),
	)

	if err != nil {
		log.Fatal(err)
	}

	iface, err := client.IDOTA2Match(570)

	if err != nil {
		log.Fatal(err)
	}

	req, err := iface.GetTeamInfoByTeamID()

	if err != nil {
		log.Fatal(err)
	}

	result := make(map[string]interface{})

	req.
		SetResult(result).
		SetQueryParam("teams_requested", "3")

	// response body is automatically deserialized into `result`
	resp, err := client.Execute(req)

	if err != nil {
		log.Fatal(err)
	}

	// resp.Body is nil
	// resp.BodyData contains the original response body data

	if !resp.IsSuccess() {
		log.Fatalf("HTTP error: %s", resp.Status())
	}

	fmt.Printf("%#v\n", result)
}

func Example_raw() {
	client, err := steam.New(
		geyser.WithDebug(),
		geyser.WithKey("<steam_api_key>"),
		geyser.WithLanguage("en_US"),
	)

	if err != nil {
		log.Fatal(err)
	}

	iface, err := client.IDOTA2Match(570)

	if err != nil {
		log.Fatal(err)
	}

	req, err := iface.GetTeamInfoByTeamID()

	if err != nil {
		log.Fatal(err)
	}

	req.SetQueryParam("teams_requested", "3")

	// response body is not deserialized
	resp, err := client.Execute(req)

	if err != nil {
		log.Fatal(err)
	}

	// response body must be manually closed
	defer resp.Body.Close()

	if !resp.IsSuccess() {
		log.Fatalf("HTTP error: %s", resp.Status())
	}

	// response body should be manually deserialized
	result := make(map[string]interface{})
	dec := json.NewDecoder(resp.Body)

	if err := dec.Decode(&result); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v\n", result)
}
