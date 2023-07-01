package usecase

import (
	"strconv"

	log "github.com/sirupsen/logrus"
)

// Fizzbuzz will write all numbers from 1 to limit, and :
//   - replace all multiple of firstMultiple by firstString,
//   - all multiples of secondMultiple by secondString
//   - all multiples of firstMultiple and secondMultiple by firstString+secondString
func Fizzbuzz(firstMultiple int, secondMultiple int, limit int, firstString string, secondString string) []string {
	var result []string
	log.Debugf("Fizzbuzz invoked with parameters : %v %v %v %v %v", firstMultiple, secondMultiple, limit, firstString, secondString)
	for i := 1; i <= limit; i++ {
		var currentValue string
		if i%firstMultiple == 0 {
			currentValue += firstString
		}
		if i%secondMultiple == 0 {
			currentValue += secondString
		}
		if i%firstMultiple != 0 && i%secondMultiple != 0 {
			currentValue += strconv.Itoa(i)
		}
		result = append(result, currentValue)
	}
	return result
}
