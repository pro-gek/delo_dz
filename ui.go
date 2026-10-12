package main

import (
	"errors"
	"fmt"
)

func printFolders(folders []string) error {
	fmt.Println("\n=== папки ===")
	for i, str := range folders {
		fmt.Printf("%d. %s\n", i+1, str)
	}
	fmt.Println("0. Выход")
	return nil
}


func printTasks(tasks []string) error{
	if len(tasks) == 0 {
		fmt.Println("задач пока нет")
	} else {
		for _, str := range tasks {
			fmt.Printf("[ ] %s\n", str)
		}
	}
	return nil
}

func printCommands(commands []string) error {
	for i, str := range commands {
		fmt.Printf("%d. %s\n", i+1, str)
	}
	fmt.Println("0. Выход")
	return nil
}



func chooseFolder(folders []string) (string, error) {
	printFolders(folders)

	input, err := chooseInt(0, len(folders))
	if err != nil {
		return "", errors.New("ошибка выбора папки")
	}

	if input == 0{
		return "Выход",nil
	}

	return folders[input-1],nil 
}

func chooseCommand() (string, error) {
	commands := []string{"Создать папку", "Удалить папку", "Переименовать папку", "Назад"}

	printCommands(commands)

	input, err := chooseInt(0, 4)
	if err != nil {
		return "", errors.New("ошибка выбора папки")
	}

	if input == 0{
		return "Выход",nil
	}

	return commands[input-1], nil 
}

