package program_data

type Namespace struct {
	Name      string
	Files     []string
	Functions map[string]*Function
	Classes   map[string]*Class
}

func (n *Namespace) AddFile(file string) {
	n.Files = append(n.Files, file)
	// TODO: CRITICAL BUG. This wipes out existing functions and classes when adding a new file to an existing namespace.
	// Initialize these only if they are nil.
	n.Functions = make(map[string]*Function)
	n.Classes = make(map[string]*Class)
}
