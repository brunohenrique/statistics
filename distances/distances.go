package distances

import (
	"fmt"
	"math"
)

var ErrDimensionMismatch = "first array has length %v which does not match the length of the second, %v."

func Euclidean(x, y []float64) (float64, error) {
	sq, err := SqEuclidean(x, y)
	if err != nil {
		return 0., err
	}

	return math.Sqrt(sq), nil
}

func SqEuclidean(x, y []float64) (float64, error) {
	if len(x) != len(y) {
		return 0., fmt.Errorf(ErrDimensionMismatch, len(x), len(y))
	}

	zipped := zip(x, y)
	var total float64

	for _, v := range zipped {
		total += math.Pow(v[0]-v[1], 2)
	}

	return total, nil
}

func CityBlock(x, y []float64) (float64, error) {
	if len(x) != len(y) {
		return 0., fmt.Errorf(ErrDimensionMismatch, len(x), len(y))
	}

	zipped := zip(x, y)
	var total float64

	for _, v := range zipped {
		total += math.Abs(v[0] - v[1])
	}

	return total, nil
}

func Chebyshev(x, y []float64) (float64, error) {
	if len(x) != len(y) {
		return 0., fmt.Errorf(ErrDimensionMismatch, len(x), len(y))
	}

	zipped := zip(x, y)
	var max float64 = 0
	var abs float64

	for _, v := range zipped {
		abs = math.Abs(v[0] - v[1])
		max = math.Max(max, abs)
	}

	return max, nil
}

func Minkowski(x, y []float64, p float64) (float64, error) {
	if len(x) != len(y) {
		return 0., fmt.Errorf(ErrDimensionMismatch, len(x), len(y))
	}

	if p == math.Inf(0) {
		return Chebyshev(x, y)
	}

	zipped := zip(x, y)
	var total float64
	var abs float64

	for _, v := range zipped {
		abs = math.Abs(v[0] - v[1])
		total += math.Pow(abs, p)
	}
	total = math.Pow(total, 1/p)

	return total, nil
}

func zip(x, y []float64) [][]float64 {
	var zipped [][]float64

	for i := 0; i < len(x); i++ {
		zipped = append(zipped, []float64{x[i], y[i]})
	}

	return zipped
}
