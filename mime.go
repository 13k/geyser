package geyser

import "mime"

var (
	mimeJSON string
)

func init() {
	mimeJSON = mime.TypeByExtension(".json")

	if mimeJSON == "" {
		mimeJSON = "application/json"
	}
}
