package geyser

import (
	"fmt"
	"net/url"
	"strconv"
)

func SteamRemoteStorageGetCollectionDetailsFormData(collectionIDs []uint64) url.Values {
	values := url.Values{}

	values.Set("collectioncount", strconv.Itoa(len(collectionIDs)))

	for i, id := range collectionIDs {
		key := fmt.Sprintf("publishedfileids[%d]", i)
		values.Set(key, strconv.FormatUint(id, 10))
	}

	return values
}

func SteamRemoteStorageGetPublishedFileDetailsFormData(fileIDs []uint64) url.Values {
	values := url.Values{}

	values.Set("itemcount", strconv.Itoa(len(fileIDs)))

	for i, id := range fileIDs {
		key := fmt.Sprintf("publishedfileids[%d]", i)
		values.Set(key, strconv.FormatUint(id, 10))
	}

	return values
}
