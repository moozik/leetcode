package p5186

import (
	"fmt"
	"testing"
)

func TestConstructor(t *testing.T) {
	ret := Constructor([]int{1, 1, 1, 2, 2})

	fmt.Printf("count:%+v\n", ret.Query(0, 1, 2))
}
