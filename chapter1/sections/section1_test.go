package sections

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var sequence1 = []int{1, -2, 4, 5, -1, -5, 2, 7}
var sequence2 = []int{1, -2, 4, 5, -1, -5, 2, 7, -100, 1, 41}
var sequence3 = []int{41, 1, -100, -2, 4, 5, -1, -5, 2, 7}

func TestTask0_sequence1(t *testing.T) {
	v := MaxSum3(sequence1)
	if v != 12 {
		t.Error("Expected 12, got ", v)
	}
}

func TestTask0_sequence2(t *testing.T) {
	v := MaxSum3(sequence2)
	assert.Equal(t, 42, v)
}
