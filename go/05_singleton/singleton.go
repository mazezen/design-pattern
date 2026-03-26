package singleton

import "sync"

// Singleton 单利模式接口 导出的
// 通过返回接口类型，可以隐藏具体实现类型，避免外部依赖包私有结构体
type Singleton interface {
	foo()
}

// singleton 单利模式类 包私有的
type singleton struct {}

func (s singleton) foo() {}

var (
	instance *singleton
	once sync.Once
)

// GetInstance 获取单利模式对象
func GetInstance() Singleton {
	once.Do(func ()  {
		instance = &singleton{}
	})

	return instance
}