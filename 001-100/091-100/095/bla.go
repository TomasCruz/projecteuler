package main

import "time"

type vStruct struct {
	v        string
	lastUsed time.Time
}

func newItem(s string, limit int, cache map[string]vStruct) {
	if len(cache) < limit {
		// insert to cache
		return
	}

	x := findNonused(cache)
	delete(cache, x)

	// insert to cache
}

func findNonused(cache map[string]vStruct) string {
	minK := ""
	minTime := time.Now()
	for _, v := range cache {
		if v.lastUsed.Before(minTime) {
			minTime = v.lastUsed
			minK = v.v
		}
	}

	return minK
}
