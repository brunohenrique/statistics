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

	var total float64

	zipped := zip(x, y)
	for _, v := range zipped {
		total += math.Pow(v[0]-v[1], 2)
	}

	return total, nil
}

func CityBlock(x, y []float64) (float64, error) {
	if len(x) != len(y) {
		return 0., fmt.Errorf(ErrDimensionMismatch, len(x), len(y))
	}

	var total float64

	zipped := zip(x, y)
	for _, v := range zipped {
		total += math.Abs(v[0] - v[1])
	}

	return total, nil
}

func Chebyshev(x, y []float64) (float64, error) {
	if len(x) != len(y) {
		return 0., fmt.Errorf(ErrDimensionMismatch, len(x), len(y))
	}

	var max, abs float64

	zipped := zip(x, y)
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

	var total, abs float64

	zipped := zip(x, y)
	for _, v := range zipped {
		abs = math.Abs(v[0] - v[1])
		total += math.Pow(abs, p)
	}

	return math.Pow(total, 1/p), nil
}

func Hamming(x, y []float64) (float64, error) {
	if len(x) != len(y) {
		return 0., fmt.Errorf(ErrDimensionMismatch, len(x), len(y))
	}

	var dist float64

	zipped := zip(x, y)
	for _, v := range zipped {
		if v[0] != v[1] {
			dist++
		}
	}

	return dist, nil
}

func Cosine(x, y []float64) (float64, error) {
	if len(x) != len(y) {
		return 0., fmt.Errorf(ErrDimensionMismatch, len(x), len(y))
	}

	var dotXY, dotXX, dotYY float64

	zipped := zip(x, y)
	for _, v := range zipped {
		dotXY += v[0] * v[1]
		dotXX += math.Pow(v[0], 2)
		dotYY += math.Pow(v[1], 2)
	}

	return 1 - dotXY/(math.Sqrt(dotXX*dotYY)), nil
}

func zip(x, y []float64) [][]float64 {
	var zipped [][]float64

	for i := 0; i < len(x); i++ {
		zipped = append(zipped, []float64{x[i], y[i]})
	}

	return zipped
}
