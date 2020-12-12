package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Rotate(t *testing.T) {

	s := NewShipT()
	s.process(&cmdT{cmd: "R", value: 180})
	assert.Equal(t, int64(180), s.w)
	s.process(&cmdT{cmd: "R", value: 180})
	assert.Equal(t, int64(0), s.w)
	s.process(&cmdT{cmd: "R", value: 90})
	s.process(&cmdT{cmd: "R", value: 90})
	s.process(&cmdT{cmd: "R", value: 90})
	assert.Equal(t, int64(90), s.w)
	s.process(&cmdT{cmd: "L", value: 90})
	assert.Equal(t, int64(180), s.w)
	s.process(&cmdT{cmd: "L", value: 90})
	s.process(&cmdT{cmd: "L", value: 90})
	s.process(&cmdT{cmd: "L", value: 90})
	assert.Equal(t, int64(90), s.w)

}
