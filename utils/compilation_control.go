package utils

var currentFile string

func GetCurrentFile() string {
	return currentFile
}

func SetCurrentFile(file string) {
	currentFile = file
}
