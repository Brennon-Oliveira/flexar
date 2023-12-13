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
	Params   []FunctionParam
	Exported bool
}

func AddFunctionToNamespace(namespace string, function Function) {
	namespaces[namespace].Functions[function.Name] = function
}

func GetFunction(namespace string, name string) Function {
	return namespaces[namespace].Functions[name]
}

func ExistsFunction(namespace string, name string) bool {
	_, ok := namespaces[namespace].Functions[name]
	return ok
}

func GetFunctions(namespace string) map[string]Function {
	return namespaces[namespace].Functions
}
