package sections

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var sequence1 = []int{1, -2, 4, 5, -1, -5, 2, 7}
var sequence2 = []int{1, -2, 4, 5, -1, -5, 2, 7, -100, 1, 41}
var sequence3 = []int{41, 1, -100, -2, 4, 5, -1, -5, 2, 7}

func TestTask0(t *testing.T) {
	result1 := Task0MaxSum3(sequence1)
	if result1 != 12 {
		t.Error("Expected 12, got ", result1)
	}

	result2 := Task0MaxSum3(sequence2)
	assert.Equal(t, 42, result2)

	result3 := Task0MaxSum3(sequence3)
	assert.Equal(t, 42, result3)
}

func TestTask1(t *testing.T) {
	var result, l, r int
	result, l, r = Task1MaxSum3WithIndex(sequence1)
	assert.Equal(t, 12, result)
	assert.Equal(t, 2, l)
	assert.Equal(t, 7, r)

	result, l, r = Task1MaxSum3WithIndex(sequence2)
	assert.Equal(t, 42, result)
	assert.Equal(t, 9, l)
	assert.Equal(t, 10, r)

	result, l, r = Task1MaxSum3WithIndex(sequence3)
	assert.Equal(t, 42, result)
	assert.Equal(t, 0, l)
	assert.Equal(t, 1, r)
}

func TestTask2(t *testing.T) {
	var result string
	result = Task2CharacterRepeated("aaaaahčeeskyoj")
	assert.Equal(t, "eskyoj", result)

	result = Task2CharacterRepeated("aaaaabbbbb")
	assert.Equal(t, "b", result)

	result = Task2CharacterRepeated("abcdefgh")
	assert.Equal(t, "abcdefgh", result)
}
