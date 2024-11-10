package list

import (
	"github.com/stretchr/testify/assert"
	"github.com/wgzhulin/algo-go/internel/utils"
	"testing"
)

func TestReverse(t *testing.T) {
	head := utils.NewList([]int{1, 2, 3, 4, 5})

	r := reverse(head)

	assert.Equal(t, utils.NewList([]int{5, 4, 3, 2, 1}), r)
}
