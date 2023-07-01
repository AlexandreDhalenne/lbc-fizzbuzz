package usecase_test

import (
	"lbc/fizzbuzz/model"
	"lbc/fizzbuzz/stats"
	"lbc/fizzbuzz/usecase"
	"testing"
)

func TestGetMostFrequentRequestNominal(t *testing.T) {
	stats := stats.GetInstance()
	mockStatsMap := make(map[model.FizzbuzzRequest]int)

	expectedRequest := model.NewFizzbuzzRequest(4, 10, 28, "nice", "test")
	expectedCount := 100
	mockStatsMap[expectedRequest] = expectedCount

	mockStatsMap[model.NewFizzbuzzRequest(3, 5, 15, "fizz", "buzz")] = 53
	mockStatsMap[model.NewFizzbuzzRequest(7, 11, 34, "weird", "values")] = 42
	mockStatsMap[model.NewFizzbuzzRequest(6, 8, 94, "please", "work")] = 3

	stats.SetMap(mockStatsMap)

	req, count, err := usecase.GetMostFrequentRequest(stats)

	if err != nil {
		t.Fatalf(err.Error())
	}
	if req != expectedRequest {
		t.Fatalf("Expected %v , found %v", expectedRequest, req)
	}
	if count != expectedCount {
		t.Fatalf("Expected %v, found %v", expectedCount, count)
	}

	stats.Clear()
}

func TestGetMostFrequentRequestEmptyMap(t *testing.T) {
	stats := stats.GetInstance()
	req, _, err := usecase.GetMostFrequentRequest(stats)

	if err == nil {
		t.Logf("req %v", req)
		t.Errorf("Expected to have an error if no map is set !")
	}

	stats.Clear()
}
