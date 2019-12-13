package main

import (
	"strconv"
	"strings"
)

func stringJoinInt(ii []int, sep string) string {
	ss := make([]string, len(ii))

	for i, n := range ii {
		ss[i] = strconv.FormatInt(int64(n), 10)
	}

	return strings.Join(ss, sep)
}

func stringJoinUint32(ii []uint32, sep string) string {
	ss := make([]string, len(ii))

	for i, n := range ii {
		ss[i] = strconv.FormatUint(uint64(n), 10)
	}

	return strings.Join(ss, sep)
}
