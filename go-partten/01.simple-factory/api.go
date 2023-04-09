package simple_factory

type API interface {
	Say(name string) string
}

func NewAPI(str string) API {
	if str == "en" {
		return &English{}
	} else if str == "zh" {
		return &Chinese{}
	}
	return nil
}

type Chinese struct {
}

func (c *Chinese) Say(name string) string {
	return "你好" + name
}

type English struct {
}

func (e *English) Say(name string) string {
	return "hello " + name
}
