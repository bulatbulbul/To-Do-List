package main

import (
	"fmt"
)

func menu() {
	fmt.Println("Меню:")
	fmt.Println("1. Добавить задачу")
	fmt.Println("2. Удалить задачу")
	fmt.Println("3. Отметить выполнение задачи")
	fmt.Println("4. Редактировать задачу")
	fmt.Println("5. Показать все задачи")
	fmt.Println("6. Выход")
}

type Task struct {
	ID   int
	Name string
	//Description string
	Status bool
	//Priority    int
	//Deadline    time.Time
	//CreatedAt time.Time
}

var currentID int
var taskList []Task

func getNextID() int {
	currentID++
	return currentID
}
func AddTask() {
	fmt.Println("Добавление задачи")
	var name string
	fmt.Println("Введеите название задачи")
	fmt.Scan(&name)
	fmt.Println(name)
	newTask := Task{
		ID:     getNextID(),
		Name:   name,
		Status: false,
	}
	taskList = append(taskList, newTask)
	fmt.Println("Задача успешно добавлена")
}

func DeleteTask() {
	fmt.Println("Удаление задачи")
	var id int
	fmt.Println("Введите ID задачи")
	fmt.Scan(&id)
	indexToRemove := -1
	for i, task := range taskList {
		if task.ID == id {
			indexToRemove = i
			break
		}
	}
	if indexToRemove == -1 {
		fmt.Printf("Задача с ID: %d не существует", id)
		fmt.Print("\n")
	} else {
		taskList = append(taskList[:indexToRemove], taskList[indexToRemove+1:]...)
		fmt.Printf("Задача с ID: %d удалена", id)
		fmt.Print("\n")
	}
}

func MaskTask() {
	fmt.Println("Отмечаем выполненную задачу")
	var id int
	fmt.Println("Введите ID задачи")
	fmt.Scan(&id)
	indexToEdit := -1
	for i, task := range taskList {
		if task.ID == id {
			indexToEdit = i
			break
		}
	}
	if indexToEdit == -1 {
		fmt.Printf("Задача с ID: %d не существует", id)
		fmt.Print("\n")
	} else {
		taskList[indexToEdit].Status = true
		fmt.Printf("Задача с ID: %d выполнена", id)
		fmt.Print("\n")
	}
}

func EditTask() {
	fmt.Println("Редактируем задачу")
	var id int
	fmt.Println("Введите ID задачи")
	fmt.Scan(&id)
	indexToEdit := -1
	for i, task := range taskList {
		if task.ID == id {
			indexToEdit = i
			break
		}
	}

	if indexToEdit == -1 {
		fmt.Printf("Задача с ID: %d не существует", id)
		fmt.Print("\n")
	} else {
		fmt.Println("Выберите что хотите редактировать:")
		fmt.Println("1. Name")
		fmt.Println("2. Status")
		var choiceEdits int
		fmt.Scan(&choiceEdits)
		switch choiceEdits {
		case 1:
			fmt.Println("Редактируем имя")
			var newName string
			fmt.Scan(&newName)
			taskList[indexToEdit].Name = newName
		case 2:
			fmt.Println("Редактируем status")
			fmt.Println("1. Выполнена")
			fmt.Println("2. Не выполнена")
			var statusChoice int
			fmt.Scan(&statusChoice)
			if statusChoice == 1 {
				taskList[indexToEdit].Status = true
			} else {
				taskList[indexToEdit].Status = false
			}
		}
		fmt.Printf("Задача с ID: %d редактирована", id)
		fmt.Print("\n")
	}
}

func ShowTasks() {
	fmt.Println("Задачи:")
	for _, task := range taskList {
		fmt.Printf("ID: %d\n", task.ID)
		fmt.Printf("Name: %s\n", task.Name)
		fmt.Printf("Status: %t\n", task.Status)
		fmt.Print("\n")
	}
}

func main() {
	var choice int
	flag := true
	for flag == true {
		menu()
		fmt.Scan(&choice)
		switch choice {
		case 1:
			AddTask()
		case 2:
			DeleteTask()
		case 3:
			MaskTask()
		case 4:
			EditTask()
		case 5:
			ShowTasks()
		case 6:
			flag = false
		}
	}

}
