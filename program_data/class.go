package program_data

type Class struct {
	Name       string
	Methods    map[string]Function
	Attributes map[string]Attribute
}

func AddClass(name string, namespace string, file string) {
	class := Class{Name: name, Methods: make(map[string]Function), Attributes: make(map[string]Attribute)}
	namespaces[namespace].Classes[name] = class
}

func GetClass(namespace string, name string) Class {
	return namespaces[namespace].Classes[name]
}

func ExistsClass(namespace string, name string) bool {
	_, ok := namespaces[namespace].Classes[name]
	return ok
}

func GetClasses(namespace string) map[string]Class {
	return namespaces[namespace].Classes
}
