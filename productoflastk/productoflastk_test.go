package productoflastk

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
