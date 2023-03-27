package alternatingsums

func alternatingSums(a []int) []int {
	teamone := 0
	teamtwo := 0
	for i := 0; i < len(a); i++ {
		if i%2 == 0 {
			teamone += a[i]
		} else {
			teamtwo += a[i]
		}
	}

	return []int{teamone, teamtwo}
}

// func TestAlternatingSums(t *testing.T) {
// 	var cases = []struct {
// 		input    []int
// 		expected []int
// 	}{
// 		{
// 			input:    []int{},
// 			expected: []int{},
// 		},
// 	}
// 	for _, c := range cases {

// 		result := alternatingSums(c.input)
// 		require.Equal(t, c.expected, result)
// 	}
// }
