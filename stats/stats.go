package stats

import (
	"lbc/fizzbuzz/model"
)

// Singleton reference
var statsInstance *Stats

// Structure to hold the cache to count number of time a request was received
type Stats struct {
	statsMap map[model.FizzbuzzRequest]int
}

// Singleton
func GetInstance() *Stats {
	if statsInstance == nil {
		statsInstance = &Stats{statsMap: make(map[model.FizzbuzzRequest]int)}
	}
	return statsInstance
}

// Getter for statsMap
func (s *Stats) StatsMap() map[model.FizzbuzzRequest]int {
	return s.statsMap
}

func (s *Stats) SetMap(m map[model.FizzbuzzRequest]int) {
	s.statsMap = m
}

func (s *Stats) Clear() {
	s.statsMap = make(map[model.FizzbuzzRequest]int)
}

// Increase by 1 the number of times a request was received
func (s *Stats) AddFor(request model.FizzbuzzRequest) {
	if _, exists := s.statsMap[request]; exists {
		s.statsMap[request] += 1
	} else {
		s.statsMap[request] = 1
	}
}
