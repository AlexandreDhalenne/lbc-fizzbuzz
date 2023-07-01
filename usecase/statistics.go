package usecase

import (
	"errors"
	"lbc/fizzbuzz/model"
	"lbc/fizzbuzz/stats"
	"sort"
)

// Strcut to hold pair of {request, count}
type RequestAndCount struct {
	Request model.FizzbuzzRequest
	Count   int
}

// Return the most frequent request and the number of times it has been called
// Sort the map by value
func GetMostFrequentRequest(s *stats.Stats) (model.FizzbuzzRequest, int, error) {
	if len(s.StatsMap()) == 0 {
		return model.FizzbuzzRequest{}, 0, errors.New("No requests cached yet !")
	}
	listOfRequestAndCount := make([]RequestAndCount, len(s.StatsMap()))
	for request, count := range s.StatsMap() {
		listOfRequestAndCount = append(listOfRequestAndCount, RequestAndCount{Request: request, Count: count})
	}
	sort.Slice(listOfRequestAndCount, func(i, j int) bool {
		return listOfRequestAndCount[i].Count > listOfRequestAndCount[j].Count
	})
	keys := make([]model.FizzbuzzRequest, len(listOfRequestAndCount))
	for i, pair := range listOfRequestAndCount {
		keys[i] = pair.Request
	}
	return keys[0], s.StatsMap()[keys[0]], nil
}
