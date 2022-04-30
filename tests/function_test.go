package tests

import (
	"reflect"
	"testing"

	"github.com/DylanMeeus/hasgo2/pkg"
)


type VectorTest[T any] struct {
    input []T
}


type ScalarTest[T any] struct {
    head pkg.Maybe[T]
}

// listManipulationTest contains tests for ops that can be performed on any type: []T
type anyListManipulationTest[T any] struct {
    input []T
    tail []T
}

type mathTest[T pkg.SignedNumber | pkg.UnsignedNumber] struct {
    input []T
    abs []T
    min T
    max T
}


var (
	intVectorTests  =  []mathTest[int]{
		{
			input: []int{-1, 2, -3, 4},
			abs: []int{1, 2, 3, 4},
		},
	}

	floatVectorTests = []mathTest[float64] {
	    {
		input: []float64{-1.5, -2.},
		abs: []float64{1.5, 2.},
	    },
	}


	intListManipulationTests = []anyListManipulationTest[int] {
	    {
			input: []int{-1, 2, -3, 4},
			tail: []int{2,-3,4},
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


func Test_AnyListManipulation(t *testing.T) {
    for _,test := range intListManipulationTests {
	t.Run("", func(t *testing.T) {
	    res := pkg.Tail(test.input)
	    if !reflect.DeepEqual(res, test.tail) {
		t.Errorf("expected %v but got %v\n", test.tail, res)
	    }
	})
    }
}


func Test_Head(t *testing.T) {
    
}
