package main

import (
	"fmt"
	"github.com/Brennon-Oliveira/flexar/compilation"
	"github.com/Brennon-Oliveira/flexar/program_data"
	"github.com/Brennon-Oliveira/flexar/utils"
	"os"
	"path/filepath"
)

func main() {
	// receive file path as argument

	if len(os.Args) < 2 {
		fmt.Println("You must provide a path as argument")
		return
	}

	projectPath := os.Args[1]

	files := []string{}

	// check if path is a flexar file (.fl)

	if filepath.Ext(projectPath) != ".fl" {
		// check if path is a directory
		if info, err := os.Stat(projectPath); err == nil && info.IsDir() {
			compilation.FindFiles(projectPath, &files)
		} else {
			fmt.Println("Invalid path")
			return
		}
	} else {
		// check if file exists
		if _, err := os.Stat(projectPath); os.IsNotExist(err) {
			fmt.Println("File does not exist")
			return
		}
		files = append(files, projectPath)
	}

	for _, file := range files {
		utils.SetCurrentFile(file)
		compilation.Compile(file)
	}

	fmt.Println("Compilation finished")
	fmt.Println("=====================================")

	namespaces := program_data.GetNamespaces()
	for _, namespace := range namespaces {
		fmt.Printf("Namespace %s: \n", namespace.Name)
		for _, file := range namespace.Files {
			fmt.Printf("\t%s\n", file)
		}
	}
}