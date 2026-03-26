package iterator

type FilterFunc func(interface{}) bool

type FilterIterator struct {
	iterator Iterator
	filter FilterFunc
	nextItem interface{}
	hasNext bool
}

// NewFilterIterator 创建过滤迭代器
func NewFilterIterator(it Iterator, filter FilterFunc) *FilterIterator {
	f := &FilterIterator{
		iterator: it,
		filter: filter,
	}
	f.advance()
	return f
}

func (f *FilterIterator) advance() {
	f.hasNext = false
	for f.iterator.HasNext() {
		item := f.iterator.Next()
		if f.filter(item) {
			f.nextItem = item
			f.hasNext = true
			return
		}
	}
}

func (f *FilterIterator) HasNext() bool {
	return f.hasNext
}

func (f *FilterIterator) Next() interface{} {
	if !f.hasNext {
		return nil
	}

	result := f.nextItem
	f.advance()
	return result
}