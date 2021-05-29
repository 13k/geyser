package web

import (
	"net/url"
)

func mustParseURL(s string) *url.URL {
	u, err := url.Parse(s)

	if err != nil {
		panic(err)
	}

	return u
}

func mustParseURLs(ss ...string) []*url.URL {
	urls := make([]*url.URL, len(ss))

	for i, s := range ss {
		urls[i] = mustParseURL(s)
	}

	return urls
}
