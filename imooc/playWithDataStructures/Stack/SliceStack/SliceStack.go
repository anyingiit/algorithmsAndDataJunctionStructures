package SliceStack

// SliceStack 栈. 使用重命名的方式不清晰, 最终还是使用了组合的方式制作了Stack
type SliceStack struct {
	rawSlice []interface{}
}

func NewSliceStack() *SliceStack {
	return &SliceStack{nil}
}

func (s *SliceStack) Enqueue(e interface{}) {
	s.rawSlice = append(s.rawSlice, e)
}

func (s *SliceStack) Size() int {
	return len(s.rawSlice)
}

func (s *SliceStack) IsEmpty() bool {
	return s.Size() == 0
}

func (s *SliceStack) Dequeue() interface{} {
	last := s.rawSlice[s.Size()-1]
	s.rawSlice = s.rawSlice[:s.Size()-1]
	return last
}
