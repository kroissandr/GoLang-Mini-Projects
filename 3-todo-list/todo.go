package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var tasks []string

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("ToDo-List options:")
		fmt.Println("1. View tasks")
		fmt.Println("2. Add a task")
		fmt.Println("3. Delete a task")
		fmt.Println("4. Exit")
		fmt.Println("Please choose an option:")

		scanner.Scan()
		option := scanner.Text()

		fmt.Println("")

		switch option {
		case "1":
			viewTasks()
		case "2":
			addTask(scanner)
		case "3":
			deleteTask(scanner)
		case "4":
			fmt.Println("Exiting from the program...")
			return
		default:
			fmt.Println("Invalid option. Please write 1, 2, 3 or 4.")
		}

		fmt.Println("")
	}
}

func viewTasks() {
	if len(tasks) == 0 {
		fmt.Println("No tasks found")
	} else {
		for i, task := range tasks {
			fmt.Printf("%d: %s\n", i+1, task)
		}
	}
}

func addTask(scanner *bufio.Scanner) {
	fmt.Println("Please add your task:")
	scanner.Scan()
	task := scanner.Text()
	tasks = append(tasks, task)
	fmt.Println("\nYour task has been successfully added to the list")
}

func deleteTask(scanner *bufio.Scanner) {
	fmt.Println("Please write a number of the task that you want to delete:")
	scanner.Scan()
	taskNum := scanner.Text()
	taskNumInt, err := strconv.Atoi(taskNum)
	if err != nil {
		fmt.Println("\nSome error occured")
	}

	tasks = append(tasks[:taskNumInt-1], tasks[taskNumInt:]...)
	fmt.Println("\nYour task has been successfully deleted")
}
