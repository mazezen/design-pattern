package iterator

// NumberIterator 数字迭代器
type NumberIterator struct {
	numbers []int
	index int
}

// NewNumberIterator 数字迭代器
// 遍历数字集合
func NewNumberIterator(nums []int) *NumberIterator {
	return &NumberIterator{numbers: nums}
}

func (nt *NumberIterator) HasNext() bool {
	return nt.index < len(nt.numbers)
}

func (nt *NumberIterator) Next() interface{} {
	if !nt.HasNext() {
		return 0;
	}

	n := nt.numbers[nt.index]
	nt.index++
	return n
}