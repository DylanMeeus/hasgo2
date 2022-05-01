package tests

import (
	"reflect"
	"testing"

	"github.com/DylanMeeus/hasgo2/pkg"
)

// listManipulationTest contains tests for ops that can be performed on any type: []T
type anyListManipulationTest[T any] struct {
    input []T
    tail []T
    reverse []T
    head pkg.Maybe[T]
}

type mathTest[T pkg.SignedNumber | pkg.UnsignedNumber] struct {
    input []T
    abs []T
    min T
    max T
}


var (
	intMathTests =  []mathTest[int]{
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
			reverse: []int{4,-3,2,-1},
			tail: []int{2,-3,4},
			head: pkg.Just(-1),
	    },
	}
)

func Test_Abs(t *testing.T) {
	for _, test := range intMathTests {
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
	// test tail
	t.Run("", func(t *testing.T) {
	    res := pkg.Tail(test.input)
	    if !reflect.DeepEqual(res, test.tail) {
		t.Errorf("expected %v but got %v\n", test.tail, res)
	    }
	})

	// test head
	t.Run("", func(t *testing.T) {
	    res := pkg.Head(test.input)
	    if res != test.head {
		t.Errorf("expected %v but got %v\n", test.tail, res)
	    }
	})

	// test reverse
	t.Run("", func(t *testing.T) {
	    res := pkg.Reverse(test.input)
	    if !reflect.DeepEqual(res, test.reverse) {
		t.Errorf("expected %v but got %v\n", test.tail, res)
	    }
	})
    }
}

