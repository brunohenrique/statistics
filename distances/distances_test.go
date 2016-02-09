package distances_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/brunohenrique/statistics/distances"
)

var equalVectorsTests = []struct {
	vectorX  []float64
	expected float64
}{
	{[]float64{1}, 0.},
	{[]float64{2}, 0.},
	{[]float64{3}, 0.},
}

func TestEuclideanWithOneDimensionAndEqualVectors(t *testing.T) {
	for _, v := range equalVectorsTests {
		distance, _ := distances.Euclidean(v.vectorX, v.vectorX)

		if distance != v.expected {
			t.Errorf("Euclidean(%f, %f): expect %.1f, got %.1f", v.vectorX, v.vectorX, v.expected, distance)
		}
	}
}

func TestSqEuclideanWithOneDimensionAndEqualVectors(t *testing.T) {
	for _, v := range equalVectorsTests {
		distance, _ := distances.SqEuclidean(v.vectorX, v.vectorX)

		if distance != v.expected {
			t.Errorf("SqEuclidean(%f, %f): expect %.1f, got %.1f", v.vectorX, v.vectorX, v.expected, distance)
		}
	}
}

func TestCityBlockWithOneDimensionAndEqualVectors(t *testing.T) {
	for _, v := range equalVectorsTests {
		distance, _ := distances.CityBlock(v.vectorX, v.vectorX)

		if distance != v.expected {
			t.Errorf("CityBlock(%f, %f): expect %.1f, got %.1f", v.vectorX, v.vectorX, v.expected, distance)
		}
	}
}

func TestChebyshevWithOneDimensionAndEqualVectors(t *testing.T) {
	for _, v := range equalVectorsTests {
		distance, _ := distances.Chebyshev(v.vectorX, v.vectorX)

		if distance != v.expected {
			t.Errorf("Chebyshev(%f, %f): expect %.1f, got %.1f", v.vectorX, v.vectorX, v.expected, distance)
		}
	}
}

var diferrentVectorsTests = []struct {
	vectorX, vectorY []float64
	expected         float64
}{
	{[]float64{1}, []float64{2}, 1.},
	{[]float64{2}, []float64{3}, 1.},
	{[]float64{3}, []float64{4}, 1.},
}

func TestEuclideanWithOneDimensionDiffentVectors(t *testing.T) {
	for _, v := range diferrentVectorsTests {
		distance, _ := distances.Euclidean(v.vectorX, v.vectorY)

		if distance != v.expected {
			t.Errorf("Euclidean(%f, %f): expect %.1f, got %.1f", v.vectorX, v.vectorY, v.expected, distance)
		}
	}
}

func TestSqEuclideanWithOneDimensionDiffentVectors(t *testing.T) {
	for _, v := range diferrentVectorsTests {
		distance, _ := distances.SqEuclidean(v.vectorX, v.vectorY)

		if distance != v.expected {
			t.Errorf("SqEuclidean(%f, %f): expect %.1f, got %.1f", v.vectorX, v.vectorY, v.expected, distance)
		}
	}
}

func TestCityBlockWithOneDimensionDiffentVectors(t *testing.T) {
	for _, v := range diferrentVectorsTests {
		distance, _ := distances.CityBlock(v.vectorX, v.vectorY)

		if distance != v.expected {
			t.Errorf("CityBlock(%f, %f): expect %.1f, got %.1f", v.vectorX, v.vectorY, v.expected, distance)
		}
	}
}

func TestChebyshevWithOneDimensionDiffentVectors(t *testing.T) {
	for _, v := range diferrentVectorsTests {
		distance, _ := distances.Chebyshev(v.vectorX, v.vectorY)

		if distance != v.expected {
			t.Errorf("Chebyshev(%f, %f): expect %.1f, got %.1f", v.vectorX, v.vectorY, v.expected, distance)
		}
	}
}

func TestEuclideanWithNDimensions(t *testing.T) {
	var tests = []struct {
		vectorX, vectorY []float64
		expected         float64
	}{
		{[]float64{4, 5, 6, 7}, []float64{3, 9, 8, 1}, math.Sqrt(57.)},
		{[]float64{1, 2, 3, 4}, []float64{1, 2, 3, 1}, 3.},
	}

	for _, v := range tests {
		distance, _ := distances.Euclidean(v.vectorX, v.vectorY)

		if distance != v.expected {
			t.Errorf("Euclidean(%f, %f): expect %.1f, got %.1f", v.vectorX, v.vectorY, v.expected, distance)
		}
	}
}

func TestSqEuclideanWithNDimensions(t *testing.T) {
	var tests = []struct {
		vectorX, vectorY []float64
		expected         float64
	}{
		{[]float64{4, 5, 6, 7}, []float64{3, 9, 8, 1}, 57.},
	}

	for _, v := range tests {
		distance, _ := distances.SqEuclidean(v.vectorX, v.vectorY)

		if distance != v.expected {
			t.Errorf("SqEuclidean(%f, %f): expect %.1f, got %.1f", v.vectorX, v.vectorY, v.expected, distance)
		}
	}
}

func TestCityBlockWithNDimensions(t *testing.T) {
	var tests = []struct {
		vectorX, vectorY []float64
		expected         float64
	}{
		{[]float64{4, 5, 6, 7}, []float64{3, 9, 8, 1}, 13.},
	}

	for _, v := range tests {
		distance, _ := distances.CityBlock(v.vectorX, v.vectorY)

		if distance != v.expected {
			t.Errorf("CityBlock(%f, %f): expect %.1f, got %.1f", v.vectorX, v.vectorY, v.expected, distance)
		}
	}
}

func TestChebyshevkWithNDimensions(t *testing.T) {
	var tests = []struct {
		vectorX, vectorY []float64
		expected         float64
	}{
		{[]float64{4, 5, 6, 7}, []float64{3, 9, 8, 1}, 6.},
	}

	for _, v := range tests {
		distance, _ := distances.Chebyshev(v.vectorX, v.vectorY)

		if distance != v.expected {
			t.Errorf("Chebyshev(%f, %f): expect %.1f, got %.1f", v.vectorX, v.vectorY, v.expected, distance)
		}
	}
}

var wrongDimensionsTests = []struct {
	vectorX, vectorY []float64
	expected         error
}{
	{[]float64{1, 1, 1, 1}, []float64{1, 1, 1}, fmt.Errorf(distances.ErrDimensionMismatch, 4, 3)},
	{[]float64{1, 1, 1}, []float64{1, 1, 1, 1}, fmt.Errorf(distances.ErrDimensionMismatch, 3, 4)},
	{[]float64{1}, []float64{1, 1, 1, 1}, fmt.Errorf(distances.ErrDimensionMismatch, 1, 4)},
	{[]float64{1, 1, 1, 1}, []float64{1}, fmt.Errorf(distances.ErrDimensionMismatch, 4, 1)},
}

func TestEuclideanWithErrorDimensionMismatch(t *testing.T) {
	for _, v := range wrongDimensionsTests {
		_, err := distances.Euclidean(v.vectorX, v.vectorY)

		if err.Error() != v.expected.Error() {
			t.Errorf("Euclidean(%f, %f): expect %v, got %v", v.vectorX, v.vectorY, v.expected, err)
		}
	}
}

func TestSqEuclideanWithErrorDimensionMismatch(t *testing.T) {
	for _, v := range wrongDimensionsTests {
		_, err := distances.SqEuclidean(v.vectorX, v.vectorY)

		if err.Error() != v.expected.Error() {
			t.Errorf("SqEuclidean(%f, %f): expect %v, got %v", v.vectorX, v.vectorY, v.expected, err)
		}
	}
}

func TestCityBlockWithErrorDimensionMismatch(t *testing.T) {
	for _, v := range wrongDimensionsTests {
		_, err := distances.CityBlock(v.vectorX, v.vectorY)

		if err.Error() != v.expected.Error() {
			t.Errorf("CityBlock(%f, %f): expect %v, got %v", v.vectorX, v.vectorY, v.expected, err)
		}
	}
}

func TestChebyshevWithErrorDimensionMismatch(t *testing.T) {
	for _, v := range wrongDimensionsTests {
		_, err := distances.Chebyshev(v.vectorX, v.vectorY)

		if err.Error() != v.expected.Error() {
			t.Errorf("Chebyshev(%f, %f): expect %v, got %v", v.vectorX, v.vectorY, v.expected, err)
		}
	}
}
