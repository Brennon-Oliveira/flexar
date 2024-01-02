package program_data

type FunctionParam struct {
	Name     string
	Type     string
	Order    int
	Optional bool
}

type Function struct {
	Name     string
	Returns  []string
	Params   []*FunctionParam
	Exported bool
}

func (n *Namespace) AddFunction(function *Function) *Function {
	n.Functions[function.Name] = function
	return function
}

func (n *Namespace) GetFunction(name string) *Function {
	return n.Functions[name]
}

func (n *Namespace) ExistsFunction(name string) bool {
	_, ok := n.Functions[name]
	return ok
}
