package tests

import (
	"experemental/code/super_functions"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSum(t *testing.T) {
	x := 1
	y := 2
	s := super_functions.SumObj{}
	s.Sum(x, y)

	require.Equal(t, 1, s.X)
	require.Equal(t, 2, s.Y)
	require.Equal(t, 3, s.Z)

	require.True(t, s.AfterCalc)
}
