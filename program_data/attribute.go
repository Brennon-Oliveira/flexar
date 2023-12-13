package program_data

type Attribute struct {
	Name      string
	Type      string
	modifiers []string
}

func AddAttributeToClass(namespace string, class string, attribute Attribute) {
	namespaces[namespace].Classes[class].Attributes[attribute.Name] = attribute
}

func GetAttribute(namespace string, class string, name string) Attribute {
	return namespaces[namespace].Classes[class].Attributes[name]
}

func ExistsAttribute(namespace string, class string, name string) bool {
	_, ok := namespaces[namespace].Classes[class].Attributes[name]
	return ok
}

func GetAttributes(namespace string, class string) map[string]Attribute {
	return namespaces[namespace].Classes[class].Attributes
}
