package ds

import "fmt"

func ExampleTup_D() {
	t := Tup[string, int]{"str", 5}
	a, b := t.D()
	fmt.Println(a, b)
	// Output: str 5
}
