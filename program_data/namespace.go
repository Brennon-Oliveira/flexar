package program_data

type Namespace struct {
	Name      string
	Files     []string
	Functions map[string]Function
	Classes   map[string]Class
}

func AddFileToNamespace(name string, file string) {
	namespace, ok := namespaces[name]
	if !ok {
		namespace = Namespace{Name: name, Files: []string{}}
	}
	namespace.Files = append(namespace.Files, file)
	namespace.Functions = make(map[string]Function)
	namespace.Classes = make(map[string]Class)
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
