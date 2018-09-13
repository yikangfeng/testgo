package goc

/*
#include <stdio.h>
#include <stdlib.h>
void c_print(char *str) {
 printf("%s\n", str);
}
int c_return() {
  return 6;
}
*/
import "C" //import “C” 必须单起一行，并且紧跟在注释行之后
import "unsafe"
import "fmt"

func GoCallC() {
	s := "Hello Cgo"
	cs := C.CString(s)               //字符串映射
	C.c_print(cs)                    //调用C函数
	defer C.free(unsafe.Pointer(cs)) //释放内存
	fmt.Printf("c return value=%d\n", C.c_return())
}
