package simplefactory

import "fmt"

// API接口
type API interface {
	Say(name string) string
}

// NewAPI 根据类型创建API实例
func NewAPI(t int) API {
	if t == 1 {
		return &hiAPI{}
	} else if t == 2 {
		return &helloAPI{}
	} else {
		return nil
	}
}

// hiAPI 实现API接口
type hiAPI struct{}

func (*hiAPI) Say(name string) string {
	return fmt.Sprintf("Hi, %s", name)
}

// helloAPI 实现API接口
type helloAPI struct{}

func (*helloAPI) Say(name string) string {
	return fmt.Sprintf("Hello, %s", name)
}
