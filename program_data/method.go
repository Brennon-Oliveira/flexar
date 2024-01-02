package program_data

type Method struct {
	Name      string
	Returns   []string
	Params    []*FunctionParam
	modifiers []string
}

func (c *Class) AddMethod(method *Method) *Method {
	c.Methods[method.Name] = method
	return method
}

func (c *Class) GetMethod(name string) *Method {
	return c.Methods[name]
}

func (c *Class) ExistsMethod(name string) bool {
	_, ok := c.Methods[name]
	return ok
}
