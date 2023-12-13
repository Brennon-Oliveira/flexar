package program_data

type Method struct {
	Name      string
	Returns   []string
	Params    []FunctionParam
	modifiers []string
}

func AddMethodToClass(namespace string, class string, method Function) {
	namespaces[namespace].Classes[class].Methods[method.Name] = method
}

func GetMethod(namespace string, class string, name string) Function {
	return namespaces[namespace].Classes[class].Methods[name]
}

func ExistsMethod(namespace string, class string, name string) bool {
	_, ok := namespaces[namespace].Classes[class].Methods[name]
	return ok
}

func GetMethods(namespace string, class string) map[string]Function {
	return namespaces[namespace].Classes[class].Methods
}
