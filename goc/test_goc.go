package goc

/*
#cgo CFLAGS: -I ./
#cgo LDFLAGS: -L ./ -ltest
#include <stdio.h>
#include <stdlib.h>
#include "test.h"
*/
import "C"
import "fmt"

func GoCallCTest() {
	fmt.Printf("go call c test count value=%d\n", C.count)
	C.test()
}
