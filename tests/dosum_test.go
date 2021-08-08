package tests

import (
	"experemental/code/super_functions"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDoSum(t *testing.T) {
	x := 5
	y := 10
	m := super_functions.NewCanSumMock(t)
	m.SumMock.Expect(x, y).Return(x + y)

	err := super_functions.DoSum(x, y, m)

	require.Nil(t, err)

}
