package productoflastk

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// https://leetcode.com/problems/product-of-the-last-k-numbers
type ProductOfNumbers struct {
	stream []int
}

func Constructor() ProductOfNumbers {
	return ProductOfNumbers{stream: []int{}}
}

func (this *ProductOfNumbers) Add(num int) {
	this.stream = append(this.stream, num)
}

func (this *ProductOfNumbers) GetProduct(k int) int {
	toProduct := this.stream[len(this.stream)-k:]
	res := toProduct[0]
	for i := 1; i < len(toProduct); i++ {
		res = res * toProduct[i]
	}
	return res
}

func TestProductOfNumbers(t *testing.T) {
	var cases = []struct {
		actions  []string
		values   []int
		expected []int
	}{
		{
			actions:  []string{"ProductOfNumbers", "add", "add", "add", "add", "add", "getProduct", "getProduct", "getProduct", "add", "getProduct"},
			values:   []int{0, 3, 0, 2, 5, 4, 2, 3, 4, 8, 2},
			expected: []int{-1, 3, 0, 0, 0, 0, 20, 40, 0, 0, 32},
		},
	}
	for _, c := range cases {
		t.Run("", func(t *testing.T) {
			var obj ProductOfNumbers
			for i, action := range c.actions {
				switch action {
				case "ProductOfNumbers":
					obj = Constructor()
				case "add":
					obj.Add(c.values[i])
				case "getProduct":
					result := obj.GetProduct(c.values[i])
					require.Equal(t, c.expected[i], result)
				}
			}
		})
	}
}
