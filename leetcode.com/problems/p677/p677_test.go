package p677

import (
	"fmt"
	"testing"
)

func Test_testFunc(t *testing.T) {
	mapSum := Constructor()
	mapSum.Insert("apple", 3)
	fmt.Print(mapSum.Sum("ap"))

	mapSum.Insert("app", 2)
	mapSum.Insert("apple", 2)
	fmt.Print(mapSum.Sum("ap"))

	fmt.Printf("%+v\n", mapSum)
}
