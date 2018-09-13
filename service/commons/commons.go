package commons

import (
	"strconv"
)

//String2int 函数
//提供string到int的转换功能，返回转换后的int.
func String2int(s string) (int, error) {
	return strconv.Atoi(s)
}
