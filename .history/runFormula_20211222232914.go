package runFormula

import (
	"strconv"
	"strings"
)

func Run(script string) {

	s1 := filter(script)

	// TAN计算
	/*
		tan ＝ 寻找文本 (参公式, “TAN(”, , 假)
		.判断循环首 (tan ＞ 0)
			pos ＝ 寻找文本 (参公式, “)”, tan, 假)
			.如果真 (pos ＜ 0)
				跳出循环 ()
			.如果真结束
			len ＝ pos － tan ＋ 1
			str ＝ 取文本中间 (参公式, tan, len)
			str_ ＝ 到文本 (求正切 (到数值 (计算公式 (取文本中间 (str, 5, len － 5))) × #pi ÷ 180))
			参公式 ＝ 子文本替换 (参公式, str, str_, , , 真)
			tan ＝ 寻找文本 (参公式, “TAN(”, , 假)
		.判断循环尾 ()


		.如果 (寻找文本 (参公式, “;”, , 假) ＞ 0 或 取文本左边 (参公式, 3) ＝ “IF(”)
			公式组 ＝ 分割文本 (参公式, “;”, )
			.计次循环首 (取数组成员数 (公式组), n)
				.如果真 (if解析 (公式组 [n], rest))
					返回 (rest)
				.如果真结束

			.计次循环尾 ()

		.否则

			.如果真 (if解析 (参公式, rest))
				返回 (rest)
			.如果真结束

		.如果结束
		返回 (“0”)
	*/

}

//过滤特殊符号
func filter(script string) string {
	s1 := strings.ToUpper(script)
	s1 = strings.ReplaceAll(s1, "（", "(")

	s1 = strings.ReplaceAll(s1, "（", "(")
	s1 = strings.ReplaceAll(s1, "）", ")")
	s1 = strings.ReplaceAll(s1, "；", ";")
	s1 = strings.ReplaceAll(s1, ";;", ";")

	/*
		参文本 ＝ 到大写 (参文本)
		参文本 ＝ 子文本替换 (参文本, 到文本 ({ 10 }),"", , , 真)
		参文本 ＝ 子文本替换 (参文本, 到文本 ({ 13 }),"", , , 真)
		参文本 ＝ 子文本替换 (参文本,"（","(", , , 真)
		参文本 ＝ 子文本替换 (参文本,"）",")", , , 真)
		参文本 ＝ 子文本替换 (参文本,"；",";", , , 真)
		参文本 ＝ 子文本替换 (参文本,";;",";", , , 真)
		参文本 ＝ 子文本替换 (参文本,"判断","IF", , , 真)
		参文本 ＝ 子文本替换 (参文本,"!","=", , , 真)  ' 替代等号
	*/
	return s1
}

func parseIF(str string) (string, bool) {
	if str[:3] == "IF(" {
		pos := strings.Index(str, ")")
		s1 := mid(str, 4, pos-4)
		s2 := mid(str, pos+1, len(str)-pos)
		if strings.Contains(s1, "或") {
			arr := strings.Split(s1, "或")
			if aabb(arr[0]) || aabb(arr[1]) {
				return compute(s2), true
			}
		} else if strings.Contains(s1, "且") {
			arr := strings.Split(s1, "且")
			if aabb(arr[0]) && aabb(arr[1]) {
				return compute(s2), true
			}
		} else if aabb(s1) {
			return compute(s2), true
		}
		return "0", false
	}
	return compute(str), true
	/*
		.如果真 (取文本左边 (参文本, 3) ＝ “IF(”)
			位置a ＝ 寻找文本 (参文本, “)”, 3, 假)
			局文本 ＝ 取文本中间 (参文本, 4, 位置a － 4)
			局结果 ＝ 取文本中间 (参文本, 位置a ＋ 1, 取文本长度 (参文本) － 位置a)

			.判断开始 (寻找文本 (局文本, “或”, , 假) ＞ 0)
				_文组 ＝ 分割文本 (局文本, “或”, )
				.如果真 (判断大小 (_文组 [1]) 或 判断大小 (_文组 [2]))
					res ＝ 计算公式 (局结果)
					返回 (真)
				.如果真结束

			.判断 (寻找文本 (局文本, “且”, , 假) ＞ 0)
				_文组 ＝ 分割文本 (局文本, “且”, )
				.如果真 (判断大小 (_文组 [1]) 且 判断大小 (_文组 [2]))
					res ＝ 计算公式 (局结果)
					返回 (真)
				.如果真结束

			.默认

				.如果真 (判断大小 (局文本))
					res ＝ 计算公式 (局结果)
					返回 (真)
				.如果真结束

			.判断结束
			返回 (假)
		.如果真结束
		res ＝ 计算公式 (参文本)
		返回 (真)
	*/
}

//计算结果
func compute(script string) string {
	arr := make([]string, 0)
	ars := Stack{}
	var aa string
	SuffixFormula(script, 0, arr)
	if len(arr) == 1 { //只有一个数不用计算
		return arr[0]
	}

	if len(arr)%2 == 0 { //能被2整除的 多一个符号
		ars.Push("0")
	}

	for i := range arr {

		if ([]byte(arr[i]))[0] >= 48 {
			ars.Push(arr[i])
			continue
		}
		if arr[i] == "+" {
			aa, _ = ars.Pop()
			rr := ars.Top() + aa

			ars.Pop()
			ars.Push(rr)
		}
		if arr[i] == "-" {
			aa, _ = ars.Pop()
			rr := ars.Top() + aa

			ars.Pop()
			ars.Push(rr)
		}
		if arr[i] == "*" {
			aa, _ = ars.Pop()
			rr := ars.Top() + aa

			ars.Pop()
			ars.Push(rr)
		}
		if arr[i] == "/" {
			aa, _ = ars.Pop()
			rr := ars.Top() + aa

			ars.Pop()
			ars.Push(rr)
		}
	}
	return ars.Top()
}

//转后缀公式
func SuffixFormula(script string, pos int, arr []string) int {
	n := len(script)
	var cc, tt string
	ars := Stack{pos: -1, str: [20]string{}}

	for pos < n {
		pos++
		aa := mid(script, pos, 1)

		if []byte(aa)[0] > 48 || []byte(aa)[0] == 46 { // 区分小数点
			cc += aa
			continue
		}
		if cc != "" {
			arr = append(arr, cc)
			cc = ""
		}

		if aa == "(" {
			pos = SuffixFormula(script, pos, arr)
			continue
		}

		if aa == ")" {
			for {
				var ok bool
				if tt, ok = ars.Pop(); !ok {
					break
				}
				arr = append(arr, tt)
				if tt == "" {
					break
				}
			}
			return pos
		}

		if ars.Empty() {
			ars.Push(aa)
			continue
		}
		level := priority(aa, ars.Top())
		if level == 0 { //' 同级别 先栈低?
			arr = append(arr, ars.Top())
			ars.Pop()
			ars.Push(aa)
			continue
		}
		if level == 1 { //' 栈低大
			arr = append(arr, ars.Top())
			ars.Pop()
			for { //' 遇到高级 全部出完
				var ok bool
				if tt, ok = ars.Pop(); !ok {
					break
				}
				arr = append(arr, tt)
				if tt == "" {
					break
				}
			}
			arr = append(arr, aa)
			continue
		}
		if level == 2 { //' 栈顶大
			arr = append(arr, aa)
		}
	}
	if cc != "" {
		arr = append(arr, cc)
	}

	for {
		var ok bool
		if tt, ok = ars.Pop(); !ok {
			break
		}
		arr = append(arr, tt)
		if tt == "" {
			break
		}
	}

	return n
}

//取文本中间
func mid(str string, pos, num int) string {
	return str[pos : pos+num]
}

//符号优先级
func priority(a, b string) int {

	if (a == "+" && b == "-") || (a == "-" && b == "+") {
		return 0
	}
	if (a == "/" && b == "*") || (a == "*" && b == "/") {
		return 0
	}

	if (a == "+" && b == "*") || (a == "+" && b == "/") {
		return 1
	}

	if (a == "-" && b == "*") || (a == "-" && b == "/") {
		return 1
	}

	if (a == "*" && b == "+") || (a == "*" && b == "-") {
		return 2
	}

	if (a == "/" && b == "+") || (a == "/" && b == "-") {
		return 2
	}
	return 0
}

//判断大小
func aabb(str string) bool {
	arr := make([]string, 0)
	sep := "0"
	if strings.Contains(str, ">=") {
		arr = strings.Split(str, ">=")
		sep = ">="
	} else if strings.Contains(str, "<=") {
		arr = strings.Split(str, "<=")
		sep = "<="
	} else if strings.Contains(str, ">") {
		arr = strings.Split(str, ">")
		sep = ">"
	} else if strings.Contains(str, "<") {
		arr = strings.Split(str, "<")
		sep = "<"
	} else if strings.Contains(str, "==") {
		arr = strings.Split(str, "==")
		sep = "=="
	} else {
		sep = "0"
		return false
	}
	arr[0] = compute(arr[0])
	arr[1] = compute(arr[1])

	if sep == "<" {
		return tofloat(arr[0]) < tofloat(arr[1])
	} else if sep == ">" {
		return tofloat(arr[0]) > tofloat(arr[1])
	} else if sep == ">=" {
		return tofloat(arr[0]) >= tofloat(arr[1])
	} else if sep == "<=" {
		return tofloat(arr[0]) <= tofloat(arr[1])
	} else if sep == "==" {
		return tofloat(arr[0]) == tofloat(arr[1])
	}
	return false
}

func tofloat(val string) float64 {
	n, _ := strconv.ParseFloat(val, 32)
	return n
}

type Stack struct {
	pos int
	str [20]string
}

func newStack() *Stack {
	return &Stack{pos: -1, str: [20]string{}}
}

//入栈
func (s *Stack) Push(val string) {
	s.pos++
	if s.pos > 19 {
		return
	}
	s.str[s.pos] = val
}

//出栈
func (s *Stack) Pop() (string, bool) {
	if s.pos == -1 {
		return "", false
	}
	defer func() { s.pos-- }()
	return s.str[s.pos], true
}

func (s *Stack) Clear() {
	s.pos = 0
}

func (s *Stack) Top() string {
	if s.pos > 19 || s.pos <= -1 {
		return ""
	}
	return s.str[s.pos]
}

func (s *Stack) Empty() bool {
	if s.pos == -1 {
		return true
	}
	return false
}
