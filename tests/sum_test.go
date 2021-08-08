/*~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
~ Most likely it was written by Faralaks!
~ Use this code for your own purposes, if they do not contradict the license installed on the project.
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~*/

/*~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
~ Most likely it was written by Faralaks!
~ Use this code for your own purposes, if they do not contradict the license installed on the project.
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~*/

/*~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
~ Most likely it was written by Faralaks!
~ Use this code for your own purposes, if they do not contradict the license installed on the project.
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~*/

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
	require.Equal(t, 3, s.Z) // Почининая строка

	require.True(t, s.AfterCalc)
}
