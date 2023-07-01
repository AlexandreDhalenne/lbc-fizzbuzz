package usecase_test

import (
	"lbc/fizzbuzz/usecase"
	"reflect"
	"testing"
)

type fizzbuzzParameters struct {
	firstInt     int
	secondInt    int
	limit        int
	firstString  string
	secondString string
}

func TestFizzbuzz(t *testing.T) {
	testcases := []struct {
		in   fizzbuzzParameters
		want []string
	}{
		{fizzbuzzParameters{firstInt: 3, secondInt: 5, limit: 15, firstString: "fizz", secondString: "buzz"}, []string{"1", "2", "fizz", "4", "buzz", "fizz", "7", "8", "fizz", "buzz", "11", "fizz", "13", "14", "fizzbuzz"}},
		{fizzbuzzParameters{firstInt: 2, secondInt: 7, limit: 21, firstString: "lebon", secondString: "coin"}, []string{"1", "lebon", "3", "lebon", "5", "lebon", "coin", "lebon", "9", "lebon", "11", "lebon", "13", "leboncoin", "15", "lebon", "17", "lebon", "19", "lebon", "coin"}},
		{fizzbuzzParameters{firstInt: 2, secondInt: 5, limit: 20, firstString: "whatis", secondString: "goingon"}, []string{"1", "whatis", "3", "whatis", "goingon", "whatis", "7", "whatis", "9", "whatisgoingon", "11", "whatis", "13", "whatis", "goingon", "whatis", "17", "whatis", "19", "whatisgoingon"}},
	}
	for _, tc := range testcases {
		fizzbuzzed := usecase.Fizzbuzz(tc.in.firstInt, tc.in.secondInt, tc.in.limit, tc.in.firstString, tc.in.secondString)
		if !reflect.DeepEqual(fizzbuzzed, tc.want) {
			t.Errorf("Fizzbuzz: %q, want %q", fizzbuzzed, tc.want)
		}
	}
}

func TestFizzbuzzNegativeLimit(t *testing.T) {
	params := fizzbuzzParameters{firstInt: 3, secondInt: 5, limit: -1, firstString: "fizz", secondString: "buzz"}
	fizzbuzzed := usecase.Fizzbuzz(params.firstInt, params.secondInt, params.limit, params.firstString, params.secondString)
	if len(fizzbuzzed) != 0 {
		t.Errorf("Length should be 0")
	}
}
