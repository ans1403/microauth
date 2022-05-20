package shell

import (
	"fmt"
	"golang-devcontainer/src/config"
	"golang-devcontainer/src/service"
)

type CalculateShell interface {
	Execute()
}

func NewCalculateShell() CalculateShell {
	return &CalculateShellImpl{
		service.NewCalculateService(),
		config.NewConfig(),
	}
}

type CalculateShellImpl struct {
	CalculateService service.CalculateService
	Cfg              config.Config
}

func (c *CalculateShellImpl) Execute() {
	maxPrintCount := c.Cfg.GetMaxPrintCount()
	result := c.CalculateService.Service(maxPrintCount)
	fmt.Println(result)
}
