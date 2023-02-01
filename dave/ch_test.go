package dave

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Ch1(t *testing.T) {
	do_ch1()
	assert.Equal(t, 1, 0)
}

func Test_Ch2(t *testing.T) {
	do_ch2()
	assert.Equal(t, 1, 0)
}

func Test_Ch3(t *testing.T) {
	do_ch3()
	assert.Equal(t, 1, 0)
}

func Test_Ch4(t *testing.T) {
	do_ch4()
	assert.Equal(t, 1, 0)
}

func Test_Ch6(t *testing.T) {
	do_ch6()
	assert.Equal(t, 1, 0)
}

func Test_Ch6A(t *testing.T) {
	readLine()
	assert.Equal(t, 1, 0)
}

func Test_Ch10(t *testing.T) {
	do_ch10()
	assert.Equal(t, 1, 0)
}

func Test_Ch10A(t *testing.T) {
	do_ch10A()
	assert.Equal(t, 1, 0)
}

func Test_Ch10B(t *testing.T) {
	do_ch10B()
	assert.Equal(t, 1, 0)
}

func Test_Ch10C(t *testing.T) {
	do_ch10C()
	assert.Equal(t, 1, 0)
}

func Test_ChClose(t *testing.T) {
	do_close()
	assert.Equal(t, 1, 0)
}
