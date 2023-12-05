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

type Namespace struct {
	Name      string
	Files     []string
	Functions map[string]Function
}

var namespaces = make(map[string]Namespace)

func AddFileToNamespace(name string, file string) {
	namespace, ok := namespaces[name]
	if !ok {
		namespace = Namespace{Name: name, Files: []string{}}
	}
	namespace.Files = append(namespace.Files, file)
	namespace.Functions = make(map[string]Function)
	namespaces[name] = namespace
}

func GetNamespace(name string) Namespace {
	return namespaces[name]
}

func ExistsNamespace(name string) bool {
	_, ok := namespaces[name]
	return ok
}

func GetNamespaces() map[string]Namespace {
	return namespaces
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
