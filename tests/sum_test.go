package tests

import (
	"experemental/code/super_functions"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSum(t *testing.T) {
	t.Parallel()
	x := 1
	y := 1
	z := super_functions.Alinka(x, y)
	require.Equal(t, 2, z)
}
