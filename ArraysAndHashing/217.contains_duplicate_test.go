package ArraysAndHashing

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_containsDuplicate(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(true, containsDuplicate([]int{1, 2, 3, 1}))
	assert.Equal(false, containsDuplicate([]int{1, 2, 3, 4}))
	assert.Equal(true, containsDuplicate([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}))
}
