package twoSum

import(
	"github.com/stretchr/testify/assert"
)
func TestTwoSum(t *testing.T) {
	assert.Equal(t, TwoSum("3", "55"), "58")
	assert.Equal(t, TwoSum("444", "44"), "488")
}
