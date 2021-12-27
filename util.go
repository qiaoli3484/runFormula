package runformula

import (
	"strconv"
	"unsafe"
)

func tofloat(val string) float64 {
	n, _ := strconv.ParseFloat(val, 32)
	return n
}

//float32到文本
//ceil 位数
func floatto(val float64, ceil int) string {
	return strconv.FormatFloat(val, 'f', ceil, 32)
}

//取文本中间
func mid(str string, pos, num int) string {
	s1 := make([]byte, num)
	//fmt.Println(pos, pos+num)
	copy(s1, str[pos:pos+num])

	return Byte2Str(s1)
}

func Byte2Str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
