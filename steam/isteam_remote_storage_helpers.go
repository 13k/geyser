package steam

import (
	"fmt"
	"net/url"
	"strconv"
)

// ISteamRemoteStorageGetCollectionDetailsFormData generates request form data for method
// ISteamRemoteStorage/GetCollectionDetails.
func ISteamRemoteStorageGetCollectionDetailsFormData(collectionIDs []uint64) url.Values {
	values := url.Values{}

	values.Set("collectioncount", strconv.Itoa(len(collectionIDs)))

	for i, id := range collectionIDs {
		key := fmt.Sprintf("publishedfileids[%d]", i)
		values.Set(key, strconv.FormatUint(id, 10))
	}

	return values
}

// ISteamRemoteStorageGetPublishedFileDetailsFormData generates request form data for method
// ISteamRemoteStorage/GetPublishedFileDetails.
func ISteamRemoteStorageGetPublishedFileDetailsFormData(fileIDs []uint64) url.Values {
	values := url.Values{}

	values.Set("itemcount", strconv.Itoa(len(fileIDs)))

	for i, id := range fileIDs {
		key := fmt.Sprintf("publishedfileids[%d]", i)
		values.Set(key, strconv.FormatUint(id, 10))
	}

	return values
}
