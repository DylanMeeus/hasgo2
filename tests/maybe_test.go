package tests

import (
	"testing"

	"github.com/DylanMeeus/hasgo2/pkg"
)


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
)

func Test_Maybe(t *testing.T) {
    for _, test := range intMaybeTests {
	t.Run("", func(t *testing.T) {
	    if test.present != test.m.Present {
		t.Errorf("expected present to be %v but got %v\n", test.present, test.m.Present)
	    }
	    if test.value!= test.m.Value {
		t.Errorf("expected value to be %v but got %v\n", test.value, test.m.Value)
	    }
	})
    }
}
