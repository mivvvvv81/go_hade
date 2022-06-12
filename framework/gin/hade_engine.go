package gin

import "project/framework"

func (engine *Engine) SetContainer(container framework.Container) {
	engine.container = container
}

func (engine *Engine) Bind(provider framework.ServiceProvider) error{
	return engine.container.Bind(provider)
}
