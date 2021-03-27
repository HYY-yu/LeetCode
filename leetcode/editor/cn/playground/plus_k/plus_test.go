package plus_k

import (
	"fmt"
	"testing"
)

func TestInt2IntArr(t *testing.T) {
	a := []int{1, 1, 3, 4, 9}
	r := IntArr2Int(a)
	fmt.Println(r)
	x := Int2IntArr(r)
	fmt.Println(x)

	p := plusK(a, 9)
	fmt.Println(p)
}
