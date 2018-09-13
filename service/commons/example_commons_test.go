package commons_test

import (
	"fmt"
)

func ExampleInt2string() {
	s := Int2string(6)
	fmt.Printf("return result=%s", s)
	// Output:
	// 6
}

func ExampleString2int() {
	i := String2int("6")
	fmt.Printf("return result=%d", i)
	// Output:
	// 6
}


func Example(){
	i := String2int("6")
	fmt.Printf("return result=%d", i)
	// Output:
	// 6
}
