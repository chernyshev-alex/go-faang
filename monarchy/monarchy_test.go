package monarchy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Monarchy(t *testing.T) {
	mon := NewMonarchy("Jake")
	mon.Birth("Catherine", "Jake")
	mon.Birth("Tom", "Jake")
	mon.Birth("Celine", "Jake")
	mon.Birth("Peter", "Celine")
	mon.Birth("Jane", "Catherine")
	mon.Birth("Farah", "Jane")
	mon.Birth("Mark", "Catherine")

	assert.Equal(t, []string{"Jake", "Catherine", "Jane", "Farah", "Mark", "Tom", "Celine", "Peter"}, mon.orderOfSuccession())

	mon.Death("Jake")
	mon.Death("Jane")

	assert.Equal(t, []string{"Catherine", "Farah", "Mark", "Tom", "Celine", "Peter"}, mon.orderOfSuccession())
}
