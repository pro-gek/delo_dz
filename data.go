package main

import "fmt"

func getSliceFolders() ([]string, error) {
	result := []string{"Входящие", "Работа", "Личное"}
	return result, nil
}

func openFolder(folder string) ([]string, error) {
	fmt.Printf("[%s]\n", folder)
	result := []string{}
	return result, nil
}
