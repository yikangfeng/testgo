/*
  commons 包提供项目中通用函数方法
*/
package commons

import (
	"strconv"
)

 // Int2string 函数
 // 提供int到string的转换功能，返回转换后的string.
func Int2string(i int) string {
	return strconv.Itoa(i)
}
