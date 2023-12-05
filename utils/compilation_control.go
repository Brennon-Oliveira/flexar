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
