package trimmedmean

import (
	"errors"
	"sort"
)

func TrimmedMean(data []float64, lowerTrim, upperTrim float64) (float64, error) {
	if len(data) == 0 {
		return 0, errors.New("data slice cannot be empty")
	}
	if lowerTrim < 0 || upperTrim < 0 || lowerTrim > 0.5 || upperTrim > 0.5 {
		return 0, errors.New("trim proportions must be between 0 and 0.5")
	}

	if lowerTrim == 0 && upperTrim == 0 {
		lowerTrim = upperTrim
	}

	sortedData := make([]float64, len(data))
	copy(sortedData, data)
	sort.Float64s(sortedData)
	trimCountLower := int(float64(len(sortedData)) * lowerTrim)
	trimCountUpper := int(float64(len(sortedData)) * upperTrim)
	trimmedData := sortedData[trimCountLower : len(sortedData)-trimCountUpper]

	sum := 0.0
	for _, value := range trimmedData {
		sum += value
	}

	trimmedMean := sum / float64(len(trimmedData))
	return trimmedMean, nil
}
