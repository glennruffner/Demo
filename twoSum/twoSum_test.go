package twoSum

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
	s := TwoSum("33", "55")
	fmt.Println(s)
	s = TwoSum("3", "55")
	fmt.Println(s)
	assert.Equal(t, TwoSum("3", "55"), "58")
	assert.Equal(t, TwoSum("444", "44"), "488")
	assert.Equal(t, TwoSum("", "44"), "44")
	assert.Equal(t, TwoSum("444", ""), "444")
	assert.Equal(t, TwoSum("4446666", "44"), "4446710")
}
