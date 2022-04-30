package tests

import (
	"testing"

	"github.com/DylanMeeus/hasgo2/pkg"
)

type testTypes interface {
    int | string
}


type maybeTest[T any] struct {
    m pkg.Maybe[T]
    value T
    present bool
}

var (
    intMaybeTests = []maybeTest[int] {
	{
	    m: pkg.Just(16),
	    value: 16,
	    present: true,
	},
	{
	    m: pkg.Just(0),
	    value: 0,
	    present: true,
	},
	{
	    m: pkg.Nothing[int](),
	    value: 0,
	    present: false,
	},
    }


    stringMaybeTests = []maybeTest[string] {
	{
	    m: pkg.Just("Hello World"),
	    value: "Hello World",
	    present: true,
	},
    }
)

func Test_Maybe(t *testing.T) {
    for _, test := range intMaybeTests {
	t.Run("", func(t *testing.T) {
	    runTest(test, t)
	})
    }


    for _,test := range stringMaybeTests {
	t.Run("", func(t *testing.T) {
	    runTest(test, t)
	})
    }
}


func runTest[T testTypes](test maybeTest[T], t *testing.T) {
	    if test.present != test.m.Present {
		t.Errorf("expected present to be %v but got %v\n", test.present, test.m.Present)
	    }
	    if test.value!= test.m.Value {
		t.Errorf("expected value to be %v but got %v\n", test.value, test.m.Value)
	    }
}
