package tests

import (
	"experemental/code/super_functions"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSum(t *testing.T) {
	x := 1
	//y := 1
	m := super_functions.NewCanSumMock(t)
	m.SumMock.Return(2)
	r := m.Sum(x, 22)
	require.Equal(t, 2, r)
}
