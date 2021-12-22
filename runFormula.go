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
		参文本 ＝ 子文本替换 (参文本, 到文本 ({ 10 }), “”, , , 真)
		参文本 ＝ 子文本替换 (参文本, 到文本 ({ 13 }), “”, , , 真)
		参文本 ＝ 子文本替换 (参文本, “（”, “(”, , , 真)
		参文本 ＝ 子文本替换 (参文本, “）”, “)”, , , 真)
		参文本 ＝ 子文本替换 (参文本, “；”, “;”, , , 真)
		参文本 ＝ 子文本替换 (参文本, “;;”, “;”, , , 真)
		参文本 ＝ 子文本替换 (参文本, “判断”, “IF”, , , 真)
		参文本 ＝ 子文本替换 (参文本, “!”, “=”, , , 真)  ' 替代等号
	*/
	return s1
}

func parse() {

}

//计算结果
func compute(script string) {

}

//转后缀公式
func SuffixFormula(script string, pos int) {
	n := len(script)
	for pos < n {
		pos++

	}

}
