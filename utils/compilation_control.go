package utils

var currentFile string

func GetCurrentFile() string {
	return currentFile
}

func SetCurrentFile(file string) {
	currentFile = file
}

var currentNamespace string

func GetCurrentNamespace() string {
	return currentNamespace
}

func SetCurrentNamespace(namespace string) {
	currentNamespace = namespace
}

var currentClass string

func GetCurrentClass() string {
	return currentClass
}

func SetCurrentClass(class string) {
	currentClass = class
}
