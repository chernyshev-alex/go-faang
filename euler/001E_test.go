package euler

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Eu001(t *testing.T) {
	var ts = []struct {
		input int
		exp   int
	}{{10, 23}}

	for i := range ts {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			assert.Equal(t, ts[i].exp, eu001(ts[i].input))
		})
	}
}

func Test_Eu012(t *testing.T) {
	var ts = []struct {
		divisors int
		expected int
	}{{5, 28}}

	for i := range ts {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			assert.Equal(t, ts[i].expected, eu012(ts[i].divisors))
		})
	}
}
