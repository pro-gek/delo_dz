package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("-----------------------------")
	fmt.Println("| Консольный менеджер задач |")
	fmt.Println("-----------------------------")

	for {
		repeat, err := loop()
		if err != nil {
			fmt.Println("Error:", err)
			break
		}
		if repeat == false {
			break
		}
	}

	fmt.Println("Завершение прорграммы")
}

func loop() (bool, error) {
	folders, err := getSliceFolders()
	if err != nil {
		return false, errors.New("ошибка получения папок")
	}

	chooseFolder, err := chooseFolder(folders)
	if err != nil {
		return false, errors.New("ошибка выбора папки")
	}
	if chooseFolder == "Выход" {
		return false, nil
	}

	tasks, err := openFolder(chooseFolder)
	if err != nil {
		return false, errors.New("ошибка открытия папки")
	}

	printTasks(tasks)

	command, err := chooseCommand()
	if err != nil {
		return false, errors.New("ошибка выбора команды")
	}
	switch command {
	case "Выход":
		return false, nil
	case "Назад":
		return true, nil
	default:
		return false, errors.New("эта функция в разработке")
	}

}
