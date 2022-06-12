package demo


// Key 服务的key
const Key = "hade:demo"

// Service 服务的接口
type Service interface {
	GetFoo()Foo
}

// Foo 服务接口定义的一个数据结构
type Foo struct {
	Name string
}