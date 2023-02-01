package dave

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_WorkersPool(t *testing.T) {
	run_pool()
	assert.Equal(t, 1, 0)
}

func Test_WaitGrp(t *testing.T) {
	var wg sync.WaitGroup
	wait_grp(&wg)
	wg.Wait()
	assert.Equal(t, 1, 0)
}
