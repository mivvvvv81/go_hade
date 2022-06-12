package demo

import (
	"fmt"
	"project/framework"
)

type DemoServiceProvider struct {
}

func (sp *DemoServiceProvider) Register(container framework.Container) framework.NewInstance {
	return NewDemoService
}

func (sp *DemoServiceProvider) Boot(container framework.Container) error {
	fmt.Println("demo service boot")
	return nil
}

func (sp *DemoServiceProvider) IsDefer() bool {
	return true
}

func (sp *DemoServiceProvider) Params(container framework.Container) []interface{} {
	return []interface{}{container}
}

func (sp *DemoServiceProvider) Name() string {
	return Key
}
