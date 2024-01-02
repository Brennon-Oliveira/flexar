package program_data

type Namespace struct {
	Name      string
	Files     []string
	Functions map[string]*Function
	Classes   map[string]*Class
}

func GetNamespace(name string) *Namespace {
	n, ok := namespaces[name]
	if !ok {
		return &Namespace{
			Name:      name,
			Files:     []string{},
			Functions: map[string]*Function{},
			Classes:   map[string]*Class{},
		}
	}

	return n
}

func (n *Namespace) AddFile(file string) {
	n.Files = append(n.Files, file)
	n.Functions = make(map[string]*Function)
	n.Classes = make(map[string]*Class)
}

func ExistsNamespace(name string) bool {
	_, ok := namespaces[name]
	return ok
}

func GetNamespaces() map[string]*Namespace {
	return namespaces
}
