package app

import (
	"errors"
	"flag"
	"path/filepath"
	"project/framework"
	"project/framework/util"
)

type HadeApp struct {
	container  framework.Container
	baseFolder string
}

func (h HadeApp) Version() string {
	return "0.0.3"
}

func (h HadeApp) BaseFolder() string {
	if h.baseFolder != "" {
		return h.baseFolder
	}
	var baseFolder string
	flag.StringVar(&baseFolder, "base_folder", "", "base_fold参数，默认是当前路径")
	flag.Parse()
	if baseFolder != "" {
		return baseFolder
	}
	return util.GetExecDirectory()
}

func (h HadeApp) ConfigFolder() string {
	return filepath.Join(h.BaseFolder(), "config")
}

func (h HadeApp) LogFolder() string {
	return filepath.Join(h.BaseFolder(), "log")
}

func (h HadeApp) ProviderFolder() string {
	return filepath.Join(h.BaseFolder(), "provider")
}

func (h HadeApp) MiddlewareFolder() string {
	return filepath.Join(h.BaseFolder(), "middleware")

}

func (h HadeApp) CommandFolder() string {
	return filepath.Join(h.BaseFolder(), "command")
}

func (h HadeApp) RuntimeFolder() string {
	return filepath.Join(h.BaseFolder(), "runtime")
}

func (h HadeApp) TestFolder() string {
	return filepath.Join(h.BaseFolder(), "test")
}

func NewHadeApp(params ...interface{}) (interface{}, error) {
	if len(params) != 2 {
		return nil, errors.New("params error")
	}
	// 两个参数，一个是容器，一个是baseFolder
	container := params[0].(framework.Container)
	baseFolder := params[1].(string)
	return &HadeApp{container: container, baseFolder: baseFolder}, nil
}
