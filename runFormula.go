package runFormula

import "strings"

func Run(script string) {

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

func parse() {

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

		if ([]byte()(arr[i]))[0] >=48{
			ars.Push(arr[i])
			continue
		}
        if arr[i]=="+"{
			aa,_=ars.Pop()
			rr:=ars.Top()+aa

			ars.Pop()
			ars.Push(rr)
		}
 /*
    .如果真 (参公式 [n] ＝ “+”)
        符号_栈.pop (aa)
        rr ＝ 到数值 (符号_栈.top ()) ＋ 到数值 (aa)
        符号_栈.pop ()
        符号_栈.push (到文本 (rr))
    .如果真结束
    .如果真 (参公式 [n] ＝ “-”)
        符号_栈.pop (aa)
        rr ＝ 到数值 (符号_栈.top ()) － 到数值 (aa)
        符号_栈.pop ()
        符号_栈.push (到文本 (rr))
    .如果真结束
    .如果真 (参公式 [n] ＝ “*”)
        符号_栈.pop (aa)
        rr ＝ 到数值 (符号_栈.top ()) × 到数值 (aa)
        符号_栈.pop ()
        符号_栈.push (到文本 (rr))
    .如果真结束
    .如果真 (参公式 [n] ＝ “/”)
        符号_栈.pop (aa)
        rr ＝ 到数值 (符号_栈.top ()) ÷ 到数值 (aa)
        符号_栈.pop ()
        符号_栈.push (到文本 (rr))

	}
*/
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
			pos = SuffixFormula(script, pos)
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

type Stack struct {
	pos int
	str [20]string
}


func newStack() *Stack{
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
