package slice

import (
	"fmt"
	"testing"
)

func TestRemoveElem(t *testing.T) {
	slices := []interface{}{1, 2, 3, 4, 5}
	res := RemoveElement(slices, 2)
	fmt.Println(res)
}
