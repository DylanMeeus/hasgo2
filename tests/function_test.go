package testing

import (
	"reflect"
	"testing"

	"github.com/DylanMeeus/hasgo2/pkg"
)


type VectorTest[T any] struct {
    input []T
    abs []T
}

var (
	intVectorTests  =  []VectorTest[int]{
		{
			input: []int{-1, 2, -3, 4},
			abs: []int{1, 2, 3, 4},
		},
	}

	floatVectorTests = []VectorTest[float64] {
	    {
		input: []float64{-1.5, -2.},
		abs: []float64{1.5, 2.},
	    },
	}
)

func Test_Abs(t *testing.T) {
	for _, test := range intVectorTests {
		t.Run("", func(t *testing.T) {
			res := pkg.Abs(test.input)
			if !reflect.DeepEqual(res, test.abs) {
				t.Errorf("expected %v but got %v", test.abs, res)
			}
		})
	}

	for _, test := range floatVectorTests {
		t.Run("", func(t *testing.T) {
			res := pkg.Abs(test.input)
			if !reflect.DeepEqual(res, test.abs) {
				t.Errorf("expected %v but got %v", test.abs, res)
			}
		})
	}
}
