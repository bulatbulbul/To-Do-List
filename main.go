package main

import (
	"bufio"
	"fmt"
	"github.com/google/uuid"
	"os"
	"strconv"
	"strings"
)

func GetInputString() string {
	scanner := bufio.NewScanner(os.Stdin)
	var input string
	for {
		if scanner.Scan() {
			input = scanner.Text()
			if len(strings.TrimSpace(input)) > 0 {
				break
			} else {
				fmt.Println("Ввод не может быть пустым. Повторите еще раз.")
			}
		} else {
			fmt.Println("Ошибка при вводе. Повторите еще раз.")
		}
	}
	return input
}

func GetInputInt() int {
	scanner := bufio.NewScanner(os.Stdin)
	var input int
	for {
		if scanner.Scan() {
			userInput := scanner.Text()
			number, err := strconv.Atoi(userInput)
			if err == nil {
				input = number
				break
			} else {
				fmt.Println("Введите корректное число. Повторите еще раз.")
			}
		} else {
			fmt.Println("Ошибка при вводе. Повторите еще раз.")
		}
	}
	return input
}

func menu() {
	fmt.Println("Меню:")
	fmt.Println("1. Добавить задачу")
	fmt.Println("2. Удалить задачу")
	fmt.Println("3. Отметить выполнение задачи")
	fmt.Println("4. Редактировать задачу")
	fmt.Println("5. Показать все задачи")
	fmt.Println("6. Загрузить данные")
	fmt.Println("7. Сохранить данные")
	fmt.Println("8. Выход")
}

type Task struct {
	ID   string
	Name string
	//Description string
	Status bool
	//Priority    int
	//Deadline    time.Time
	//CreatedAt time.Time
}

var taskList []Task

func getNextID() string {
	return uuid.New().String()
}

func AddTask() {
	fmt.Println("Добавление задачи")
	fmt.Println("Введеите название задачи")
	name := GetInputString()
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
	fmt.Println("Введите ID задачи")
	id := GetInputString()
	indexToRemove := -1
	for i, task := range taskList {
		if task.ID == id {
			indexToRemove = i
			break
		}
	}
	if indexToRemove == -1 {
		fmt.Printf("Задача с ID: %s не существует\n", id)
	} else {
		taskList = append(taskList[:indexToRemove], taskList[indexToRemove+1:]...)
		fmt.Printf("Задача с ID: %s удалена\n", id)
	}
}

func MaskTask() {
	fmt.Println("Отмечаем выполненную задачу")
	fmt.Println("Введите ID задачи")
	id := GetInputString()
	indexToEdit := -1
	for i, task := range taskList {
		if task.ID == id {
			indexToEdit = i
			break
		}
	}
	if indexToEdit == -1 {
		fmt.Printf("Задача с ID: %s не существует\n", id)
	} else {
		taskList[indexToEdit].Status = true
		fmt.Printf("Задача с ID: %s выполнена\n", id)
	}
}

func EditTask() {
	fmt.Println("Редактируем задачу")
	fmt.Println("Введите ID задачи")
	id := GetInputString()
	indexToEdit := -1
	for i, task := range taskList {
		if task.ID == id {
			indexToEdit = i
			break
		}
	}

	if indexToEdit == -1 {
		fmt.Printf("Задача с ID: %s не существует\n", id)
	} else {
		fmt.Println("Выберите что хотите редактировать:")
		fmt.Println("1. Name")
		fmt.Println("2. Status")
		choiceEdits := GetInputInt()
		switch choiceEdits {
		case 1:
			fmt.Println("Редактируем имя")
			newName := GetInputString()
			taskList[indexToEdit].Name = newName
			fmt.Printf("Задача с ID: %s редактирована\n", id)
		case 2:
			fmt.Println("Редактируем status")
			fmt.Println("1. Выполнена")
			fmt.Println("2. Не выполнена")
			statusChoice := GetInputInt()
			if statusChoice == 1 {
				taskList[indexToEdit].Status = true
			} else {
				taskList[indexToEdit].Status = false
			}
			fmt.Printf("Задача с ID: %s редактирована\n", id)
		default:
			fmt.Println("Неверный выбор. Попробуй еще раз")
		}
	}
}

func ShowTasks() {
	fmt.Println("Задачи:")
	for _, task := range taskList {
		fmt.Printf("ID: %s\n", task.ID)
		fmt.Printf("Name: %s\n", task.Name)
		fmt.Printf("Status: %t\n", task.Status)
	}
}

func SaveTasks(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Ошибка создания файла:", err)
		return
	}
	var closeErr error
	defer func() {
		closeErr = file.Close()
		if closeErr != nil {
			fmt.Println("Ошибка при закрытии файла:", err)
		}
	}()
	writer := bufio.NewWriter(file)
	for _, task := range taskList {
		line := task.ID + "," + task.Name + "," + strconv.FormatBool(task.Status) + "\n"
		_, err = writer.WriteString(line)
		if err != nil {
			fmt.Println("Ошибка записи в файла: ", err)
			return
		}
	}

	err = writer.Flush()
	if err != nil {
		fmt.Println("Ошибка сохранения в файл: ", err)
		return
	}

	fmt.Println("Данные успешно сохранены")
}

func LoadTasks(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Ошибка открытия файла:", err)
		return
	}
	var closeErr error
	defer func() {
		closeErr = file.Close()
		if closeErr != nil {
			fmt.Println("Ошибка при закрытии файла:", err)
		}
	}()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		pars := strings.Split(line, ",")
		if len(pars) != 3 {
			continue
		}
		id := pars[0]
		name := pars[1]
		status, _ := strconv.ParseBool(pars[2])
		taskList = append(taskList,
			Task{
				ID:     id,
				Name:   name,
				Status: status,
			})
	}
	if err = scanner.Err(); err != nil {
		fmt.Println("Ошибка при чтении файла:", err)
	}
	fmt.Println("Данные успешно заргужены")
}

func main() {
	for {
		menu()
		fmt.Print("Введите номер выбора: ")
		choice := GetInputInt()
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
			LoadTasks("tasks.txt")
		case 7:
			SaveTasks("tasks.txt")
		case 8:
			fmt.Println("Выход из программы")
			return
		default:
			fmt.Println("Неверный выбор. Попробуй еще раз")
		}
	}
}
