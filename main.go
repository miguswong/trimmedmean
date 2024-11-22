package trimmedMean

import "sort"

// Create Type that takes floats and ints
type Number interface {
	~int | ~float64
}

//function for prcessing int or float slices, provide slice of ints and the number of elements to trim within the slice
func TrimMean[T Number](vals []T, trim int) float64 {

	//ensure the 2*trim value is not greater than the length of the slice
	if 2*trim >= len(vals) {
		panic("Trim value is too large")
	}

	// Create a float64 slice to store converted values
	floatVals := make([]float64, len(vals))
	for i, v := range vals {
		floatVals[i] = float64(v)
	}

	// Sort the float64 slice
	sort.Float64s(floatVals)

	//trim the slice
	trimVals := vals[trim : len(vals)-trim]

	//calculate the sum of the trimmed slice
	trimmedSum := 0.0
	for _, val := range trimVals {
		trimmedSum += float64(val)
	}

	//calculate the mean of the trimmed slice
	trimmedMean := float64(trimmedSum) / float64(len(trimVals))
	return trimmedMean
}
