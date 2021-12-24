package runFormula

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
	return s.pos == -1

}
