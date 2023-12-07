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

// Classes

type Method struct {
	Name      string
	Returns   []string
	Params    []FunctionParam
	modifiers []string
}

type Attribute struct {
	Name      string
	Type      string
	modifiers []string
}

type Class struct {
	Name       string
	Methods    map[string]Function
	Attributes map[string]Attribute
}

type Namespace struct {
	Name      string
	Files     []string
	Functions map[string]Function
	Classes   map[string]Class
}

var namespaces = make(map[string]Namespace)

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
