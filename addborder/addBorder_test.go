package addborder

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func addBorder(picture []string) []string {
	matrixPicture := make([]string, len(picture)+2)
	var firstAndLastMatrixArray string
	for i := 0; i < len(picture[0])+2; i++ {
		firstAndLastMatrixArray += "*"
	}
	matrixPicture[0] = firstAndLastMatrixArray
	for i := 0; i < len(picture); i++ {
		toUpdate := picture[i]
		var updated string
		updated += "*"
		updated += toUpdate
		updated += "*"
		matrixPicture[i+1] = updated
	}
	matrixPicture[len(matrixPicture)-1] = firstAndLastMatrixArray
	return matrixPicture
}

func TestAddBorder(t *testing.T) {
	var cases = []struct {
		input    []string
		expected []string
	}{
		{
			input: []string{"abc",
				"ded"},
			expected: []string{"*****",
				"*abc*",
				"*ded*",
				"*****"},
		},
	}
	for _, c := range cases {
		result := addBorder(c.input)
		require.Equal(t, c.expected, result)
	}
}
