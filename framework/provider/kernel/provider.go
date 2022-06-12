package kernel

import (
	"project/framework"
	"project/framework/contract"
	"project/framework/gin"
)

// HadeKernelProvider 提供web引擎
type HadeKernelProvider struct {
	HttpEngine *gin.Engine
}

// Register 注册服务提供者
func (provider *HadeKernelProvider) Register(container framework.Container) framework.NewInstance {
	return NewHadeKernelService
}


func (provider *HadeKernelProvider) Boot(container framework.Container) error {
	if provider.HttpEngine == nil {
		provider.HttpEngine = gin.Default()
	}
	provider.HttpEngine.SetContainer(container)
	return nil
}

func (provider *HadeKernelProvider) IsDefer() bool {
	return false
}

func (provider *HadeKernelProvider) Params(container framework.Container) []interface{} {
	return []interface{}{provider.HttpEngine}
}

func (provider *HadeKernelProvider) Name() string {
	return contract.KernelKey
}
