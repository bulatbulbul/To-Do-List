package main

import (
	"bufio"
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
	"strconv"
	"strings"
)

var db *sql.DB

func handleError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %v", msg, err)
	}
}

func initDB() {
	var err error
	connStr := "user=todo_user password=Nc0zt_oEa4cBLuT dbname=todo_db sslmode=disable"
	db, err = sql.Open("postgres", connStr)
	handleError(err, "Ошибка с базой данных:")
	err = db.Ping()
	handleError(err, "Не удалось подключиться к базе данных:")
	fmt.Println("Успешное подключение к базе данных!")
}

func createTable() {
	query := `
	CREATE TABLE IF NOT EXISTS tasks (
	    id SERIAL PRIMARY KEY,
	    name TEXT NOT NULL,
	    status BOOLEAN NOT NULL DEFAULT FALSE
	);`
	_, err := db.Exec(query)
	handleError(err, "Не удалось создать таблицу:")
	fmt.Println("Таблица tasks создана или уже существует.")

}

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
	fmt.Println("6. Выход")
}

func AddTask() {
	fmt.Println("Добавление задачи")
	fmt.Println("Введите название задачи")
	name := GetInputString()
	query := `
	INSERT INTO tasks(name, status)
	VALUES ($1, $2)`
	_, err := db.Exec(query, name, false)
	handleError(err, "Ошибка при добавлении задачи:")
	fmt.Println("Задача успешно добавлена")
}

func DeleteTask() {
	fmt.Println("Удаление задачи")
	fmt.Println("Введите ID задачи")
	id := GetInputInt()
	query := `
	DELETE FROM tasks 
	WHERE id = $1`
	_, err := db.Exec(query, id)
	handleError(err, "Ошибка при удалении задачи:")
	fmt.Printf("Задача с ID: %d удалена\n", id)
}

func MaskTask() {
	fmt.Println("Отмечаем выполненную задачу")
	fmt.Println("Введите ID задачи")
	id := GetInputInt()
	query := `
	UPDATE tasks 
	SET status = TRUE
	WHERE id = $1`
	_, err := db.Exec(query, id)
	handleError(err, "Ошибка при обновлении задачи:")
	fmt.Printf("Задача с ID: %d отмечена как выполнена\n", id)
}

func EditTask() {
	fmt.Println("Редактируем задачу")
	fmt.Println("Введите ID задачи")
	id := GetInputInt()
	var name string
	var status bool
	query := `
	SELECT name, status FROM tasks 
	WHERE id = $1
	`
	row := db.QueryRow(query, id)
	err := row.Scan(&name, &status)
	if errors.Is(err, sql.ErrNoRows) {
		fmt.Printf("Задача с ID: %d не существует\n", id)
		return
	} else if err != nil {
		log.Fatal("Ошибка при получении запроса:", err)
	}

	fmt.Println("Выберите что хотите редактировать:")
	fmt.Println("1. Name")
	fmt.Println("2. Status")
	choiceEdits := GetInputInt()
	switch choiceEdits {
	case 1:
		fmt.Println("Редактируем имя")
		newName := GetInputString()
		query = `
		UPDATE tasks SET name = $1
		WHERE id = $2
		`
		_, err = db.Exec(query, newName, id)
		handleError(err, "Ошибка при обновлении имени:")
		fmt.Printf("Задача с ID: %d редактирована\n", id)
	case 2:
		fmt.Println("Редактируем status")
		fmt.Println("1. Выполнена")
		fmt.Println("2. Не выполнена")
		statusChoice := GetInputInt()
		newStatus := statusChoice == 1
		query = `
		UPDATE tasks SET status = $1
		WHERE id = $2
		`
		_, err = db.Exec(query, newStatus, id)
		handleError(err, "Ошибка при обновлении статуса:")
		fmt.Printf("Задача с ID: %d редактирована\n", id)
	default:
		fmt.Println("Неверный выбор. Попробуй еще раз")
	}
}

func ShowTasks() {
	query := `
	SELECT id, name, status FROM tasks
	`
	rows, err := db.Query(query)
	handleError(err, "Ошибка при запросе задач:")
	defer func() {
		if err = rows.Close(); err != nil {
			log.Printf("Ошибка при закрытии rows: %v", err)
		}
	}()

	fmt.Println("Задачи: ")
	for rows.Next() {
		var id int
		var name string
		var status bool
		err = rows.Scan(&id, &name, &status)
		handleError(err, "Ошибка с базой данных:")
		fmt.Printf("ID: %d | Name: %s | Status: %t\n", id, name, status)
	}
}

func main() {
	initDB()
	defer func() {
		if err := db.Close(); err != nil {
			log.Printf("Ошибка при закрытии rows: %v", err)
		}
	}()
	createTable()
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
			fmt.Println("Выход из программы")
			return
		default:
			fmt.Println("Неверный выбор. Попробуй еще раз")
		}
	}
}
