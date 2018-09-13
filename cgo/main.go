package cgo

import "C"
import "fmt"

//export go_print
func go_print(value string) {
	fmt.Println("go print " + value)
}

//export hello
func hello() string {
	return "hello"
}
func main() { //main函数是必须的 有main函数才能让cgo编译器去把包编译成C的库
}
