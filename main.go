package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Brennon-Oliveira/flexar/compilation"
	"github.com/Brennon-Oliveira/flexar/program_data"
	"github.com/Brennon-Oliveira/flexar/utils"
)

func NavigateInProject(projectPath string) ([]string, error){
	files := []string{}

	// check if path is a flexar file (.fl)

	if filepath.Ext(projectPath) != ".fl" {
		// check if path is a directory
		if info, err := os.Stat(projectPath); err == nil && info.IsDir() {
			compilation.FindFiles(projectPath, &files)
		} else {
			return nil, fmt.Errorf("Invalid path")
		}
	} else {
		// check if file exists
		if _, err := os.Stat(projectPath); os.IsNotExist(err) {
			return nil, fmt.Errorf("File does not exist")
		}
		files = append(files, projectPath)
	}

	return files, nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("You must provide a path as argument")
		return
	}

	projectPath := os.Args[1]

	

	for _, file := range files {
		utils.SetCurrentFile(file)
		compilation.Compile(file)
	}

	fmt.Println("Compilation finished")
	fmt.Println("=====================================")

	namespaces := program_data.GetNamespaces()
	for _, namespace := range namespaces {
		fmt.Printf("Namespace %s: \n", namespace.Name)
		fmt.Println("Files:")
		for _, file := range namespace.Files {
			fmt.Printf("\t%s\n", file)
		}
		fmt.Println("Functions:")
		for _, function := range namespace.Functions {
			fmt.Printf("\t%s\n", function.Name)
			for _, param := range function.Params {
				fmt.Printf("\t\t%s %s\n", param.Name, param.Type)
			}
			for _, ret := range function.Returns {
				fmt.Printf("\t\t%s\n", ret)
			}
		}

		fmt.Println("Classes:")
		for _, class := range namespace.Classes {
			fmt.Printf("\t%s\n", class.Name)
			fmt.Println("\t\tAttributes:")
			for _, attr := range class.Attributes {
				fmt.Printf("\t\t\t%s\n", attr.GetString())
			}
			fmt.Println("\t\tMethods:")
			for _, method := range class.Methods {
				fmt.Printf("\t\t\t%s\n", method.Name)
				for _, param := range method.Params {
					fmt.Printf("\t\t\t\t%s %s\n", param.Name, param.Type)
				}
				for _, ret := range method.Returns {
					fmt.Printf("\t\t\t\t%s\n", ret)
				}
			}
		}
	}
}
