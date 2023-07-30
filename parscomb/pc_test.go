package parscomb

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// write compressor: 'AATTAAAARYYY' -> '2A2T4AR3Y'
func Test_CompressStr(t *testing.T) {
	var ts = []struct {
		input    string
		expected string
	}{{"AATTAAAARYYY", "2A2T4AR3Y"}}

	for i := range ts {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			result := parseAndCompress(ts[i].input)
			assert.Equal(t, ts[i].expected, result)
		})
	}
}
