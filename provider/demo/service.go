package demo

import (
	"fmt"
	"project/framework"
)

type DemoService struct {
	Service
	c framework.Container
}

func (d DemoService) GetFoo() Foo {
	return Foo{
		Name: "i am foo",
	}
}

func NewDemoService(params ...interface{}) (interface{}, error) {
	c := params[0].(framework.Container)
	fmt.Println("new demo service")
	return &DemoService{c: c}, nil
}
