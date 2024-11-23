package trimmedmean

import (
	"testing"
)

func TestTrimmedMean(t *testing.T) {
	data := []float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0}
	lowerTrim := 0.1
	upperTrim := 0.1

	result, err := TrimmedMean(data, lowerTrim, upperTrim)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	expected := 5.5
	if result != expected {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}
