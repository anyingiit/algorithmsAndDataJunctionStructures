package Stack

// Stack 栈. 使用重命名的方式不清晰, 最终还是使用了组合的方式制作了Stack
type Stack struct {
	rawSlice []interface{}
}

func NewStack() *Stack {
	return &Stack{nil}
}

func (s *Stack) Push(e interface{}) {
	s.rawSlice = append(s.rawSlice, e)
}

func (s *Stack) Size() int {
	return len(s.rawSlice)
}

func (s *Stack) IsEmpty() bool {
	return s.Size() == 0
}

func (s *Stack) Pop() interface{} {
	last := s.rawSlice[s.Size()-1]
	s.rawSlice = s.rawSlice[:s.Size()-1]
	return last
}
