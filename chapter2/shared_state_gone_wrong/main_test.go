package main
/*
	Use 'go test -race' to detect the race condition
*/
import (
	"testing"
)

func TestPackItems(t *testing.T) {
	totalItems := PackItems(2000)
	expectedTotal := 2000
	if totalItems != expectedTotal {
		t.Errorf("Expected total: %d, Actual total: %d",
			expectedTotal, totalItems)
	}
}
