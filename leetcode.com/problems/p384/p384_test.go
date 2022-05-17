package p384

import (
	"fmt"
	"testing"
)

func TestConstructor(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6}
	obj := Constructor(nums)
	param_1 := obj.Reset()
	fmt.Printf("param_1:%+v\n", param_1)
	param_2 := obj.Shuffle()
	fmt.Printf("param_2:%+v\n", param_2)
	param_2 = obj.Shuffle()
	fmt.Printf("param_2:%+v\n", param_2)
	param_2 = obj.Shuffle()
	fmt.Printf("param_2:%+v\n", param_2)
}
