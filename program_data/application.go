package program_data

type Application struct {
	Namespaces map[string]*Namespace
	Errors     []error
}

func NewApplication() *Application {
	return &Application{
		Namespaces: make(map[string]*Namespace),
		Errors:     make([]error, 0),
	}
}

func (a *Application) GetNamespace(name string) *Namespace {
	if ns, ok := a.Namespaces[name]; ok {
		return ns
	}

	ns := &Namespace{
		Name:      name,
		Files:     []string{},
		Functions: map[string]*Function{},
		Classes:   map[string]*Class{},
	}

	a.Namespaces[name] = ns

	return ns
}

func (a *Application) ExistsNamespace(name string) bool {
	_, ok := a.Namespaces[name]
	return ok
}

func (a *Application) GetNamespaces() map[string]*Namespace {
	return a.Namespaces
}
