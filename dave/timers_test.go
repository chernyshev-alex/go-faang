package dave

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Timer(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)
	do_timer(&wg)
	wg.Wait()
	assert.Equal(t, 1, 0)
}

func Test_Ticker(t *testing.T) {
	do_ticker()
	assert.Equal(t, 1, 0)
}
