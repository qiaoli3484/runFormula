package runformula

import (
	"math"
	"strconv"
	"strings"
)

//计算公式
//ceil 小数位数,-1不限
func Run(script string, ceil int) string {

	s1 := filter(script)
	tan := strings.Index(s1, "TAN(")
	for tan > 0 {
		pos := strings.Index(s1[:tan], "TAN(")
		if pos < 0 {
			break
		}
		num := pos - tan + 1
		s2 := mid(s1, tan, num)

		s3 := math.Tan(tofloat(compute(mid(s2, 4, num-4), ceil)) * math.Pi / 180)

		s1 = strings.Replace(s1, s2, strconv.Itoa(int(s3)), -1) //精度需要更改

		tan = strings.Index(s1, "TAN(")
	}
	if strings.Contains(s1, ";") || strings.Contains(s1, "IF(") {
		arr := strings.Split(s1, ";")
		for i := range arr {
			if res, ok := parseIF(arr[i], ceil); ok {
				return res
			}
		}
	} else {
		if res, ok := parseIF(s1, ceil); ok {
			return res
		}
	}
	return "0"
}

//过滤特殊符号
func filter(script string) string {

	s1 := strings.ToUpper(script)

	s1 = strings.TrimSpace(s1)
	s1 = strings.ReplaceAll(s1, "（", "(")

	s1 = strings.ReplaceAll(s1, "（", "(")
	s1 = strings.ReplaceAll(s1, "）", ")")
	s1 = strings.ReplaceAll(s1, "；", ";")
	s1 = strings.ReplaceAll(s1, ";;", ";")

	s1 = strings.ReplaceAll(s1, "判断", "IF")
	s1 = strings.ReplaceAll(s1, "!", "=") //替代等号
	return s1
}

func parseIF(str string, ceil int) (string, bool) {
	if str[:3] == "IF(" {
		pos := strings.Index(str, ")")
		s1 := mid(str, 3, pos-3)
		s2 := mid(str, pos+1, len(str)-pos-1)
		if strings.Contains(s1, "或") {
			arr := strings.Split(s1, "或")
			if aabb(arr[0], ceil) || aabb(arr[1], ceil) {
				return compute(s2, ceil), true
			}
		} else if strings.Contains(s1, "且") {
			arr := strings.Split(s1, "且")
			if aabb(arr[0], ceil) && aabb(arr[1], ceil) {
				return compute(s2, ceil), true
			}
		} else if aabb(s1, ceil) {
			return compute(s2, ceil), true
		}
		return "0", false
	}
	return compute(str, ceil), true
}

//计算结果
//ceil 位数
func compute(script string, ceil int) string {
	ars := Stack{}
	var aa, rr string
	_, arr := suffixFormula(script, -1)
	//fmt.Println(arr)
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
			rr = floatto(tofloat(ars.Top())+tofloat(aa), ceil)

			ars.Pop()
			ars.Push(rr)
		}
		if arr[i] == "-" {
			aa, _ = ars.Pop()
			rr = floatto(tofloat(ars.Top())-tofloat(aa), ceil)

			ars.Pop()
			ars.Push(rr)
		}
		if arr[i] == "*" {
			aa, _ = ars.Pop()
			rr = floatto(tofloat(ars.Top())*tofloat(aa), ceil)

			ars.Pop()
			ars.Push(rr)
		}
		if arr[i] == "/" {
			aa, _ = ars.Pop()
			rr = floatto(tofloat(ars.Top())/tofloat(aa), ceil)

			ars.Pop()
			ars.Push(rr)
		}
	}
	return rr
}

//转后缀公式
func suffixFormula(script string, pos int) (int, []string) {
	n := len(script)
	var cc, tt string
	ars := Stack{pos: -1, str: [20]string{}}
	arr := make([]string, 0)
	for {

		pos++
		if pos >= n {
			break
		}
		aa := mid(script, pos, 1)
		if []byte(aa)[0] >= 48 || []byte(aa)[0] == 46 { // 区分小数点
			cc += aa
			continue
		}
		//fmt.Println(ars, "后缀", arr, cc)
		if cc != "" {
			arr = append(arr, cc)
			cc = ""
		}

		if aa == "(" {
			var ab []string
			pos, ab = suffixFormula(script, pos)
			arr = append(arr, ab...)
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
			return pos, arr
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
			ars.Push(aa)
			continue
		}
		if level == 2 { //' 栈顶大
			ars.Push(aa)
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
	return pos, arr
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
func aabb(str string, ceil int) bool {
	var arr []string
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
	arr[0] = compute(arr[0], ceil)
	arr[1] = compute(arr[1], ceil)

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
