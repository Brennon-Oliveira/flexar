package program_data

type FunctionParam struct {
	Name  string
	Type  string
	Order int
}

type Function struct {
	Name    string
	Returns []string
	Params  []FunctionParam
}

type Namespace struct {
	Name  string
	Files []string
}

var namespaces = make(map[string]Namespace)

func AddFileToNamespace(name string, file string) {
	namespace, ok := namespaces[name]
	if !ok {
		namespace = Namespace{Name: name, Files: []string{}}
	}
	namespace.Files = append(namespace.Files, file)
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
