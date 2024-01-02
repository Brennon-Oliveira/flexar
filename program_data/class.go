package program_data

type Class struct {
	Name       string
	Methods    map[string]*Method
	Attributes map[string]*Attribute
}

func (n *Namespace) AddClass(name string) *Class {
	class := Class{Name: name, Methods: make(map[string]*Method), Attributes: make(map[string]*Attribute)}
	n.Classes[name] = &class
	return &class
}

func (n *Namespace) GetClass(name string) *Class {
	return n.Classes[name]
}

func (n *Namespace) ExistsClass(name string) bool {
	_, ok := n.Classes[name]
	return ok
}
