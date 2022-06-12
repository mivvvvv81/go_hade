package cobra

import (
	"project/framework"
)

func (c *Command) SetContainer(container framework.Container) {
	c.container = container
}
func (c *Command) GetContainer() framework.Container {
	return c.container
}
