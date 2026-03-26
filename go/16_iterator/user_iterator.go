package iterator

// User 用户实体
type User struct {
	Name string
	Age int
}

// UserIterator 用户迭代器
// 遍历用户结构体
type UserIterator struct {
	users []User
	index int
}

// NewUserIterator 创建用户迭代器
func NewUserIterator(users []User) *UserIterator {
	return &UserIterator{users: users}
}

func (ut *UserIterator) HasNext() bool {
	return ut.index < len(ut.users)
}

func (ut *UserIterator) Next() interface{} {
	if !ut.HasNext() {
		return nil
	}
	user := ut.users[ut.index]
	ut.index++
	return user
}