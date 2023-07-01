package model_test

import (
	"lbc/fizzbuzz/model"
	"testing"
)

func TestValidFizzbuzzRequests(t *testing.T) {
	validRequest := &model.FizzbuzzRequest{
		FirstInteger:  3,
		SecondInteger: 5,
		Limit:         100,
		FirstString:   "fizz",
		SecondString:  "buzz",
	}
	if err := model.ValidateFizzbuzzRequest(*validRequest); err != nil {
		t.Fatalf("Invalid request %+v", validRequest)
	}
}

func TestInvalidFizzBuzzRequests(t *testing.T) {
	testcases := []struct {
		in model.FizzbuzzRequest
	}{
		{model.FizzbuzzRequest{FirstInteger: -1, SecondInteger: 5, Limit: 15, FirstString: "fizz", SecondString: "buzz"}},
		{model.FizzbuzzRequest{FirstInteger: 3, SecondInteger: -99, Limit: 15, FirstString: "fizz", SecondString: "buzz"}},
		{model.FizzbuzzRequest{FirstInteger: 3, SecondInteger: 5, Limit: -14, FirstString: "fizz", SecondString: "buzz"}},
		{model.FizzbuzzRequest{FirstInteger: 3, SecondInteger: 5, Limit: 100, FirstString: "", SecondString: "buzz"}},
		{model.FizzbuzzRequest{FirstInteger: 3, SecondInteger: 5, Limit: 100, FirstString: "fizz", SecondString: ""}},
	}
	for _, tc := range testcases {
		if err := model.ValidateFizzbuzzRequest(tc.in); err == nil {
			t.Fatalf("Expected invalid request, but was considered as valid : %+v", tc.in)
		}
	}
}
