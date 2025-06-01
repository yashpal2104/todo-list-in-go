package main

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/yashpal2104/todo-list-in-go/cmd"
)

func main() {

	cmd.Execute() // Execute the root command from the cmd package

	fmt.Println("Thank you for using the Todo CLI application!")
	// os.Exit(0) // Exit the application gracefully
}

func tabWriterExample() {
	fmt.Println("Todo list in GO")
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "ID\tTask\tStatus\tPriority\tDue Date\t")
	fmt.Fprintln(w, "1\tLearn Go\tIn Progress\tHigh\t2023-10-01\t")
	fmt.Fprintln(w, "2\tBuild CLI App\tNot Started\tMedium\t2023-10-15\t")
	fmt.Fprintln(w, "3\tWrite Tests\tNot Started\tLow\t2023-10-20\t")
	w.Flush()

	fmt.Println("Use the 'todo' command to manage your tasks.")
	fmt.Println("Available commands:")
	fmt.Println("  - add: Add a new task")
	fmt.Println("  - list: List all tasks")
	fmt.Println("  - update: Update an existing task")
	fmt.Println("  - delete: Delete a task")
	fmt.Println("  - filter: Filter tasks by status or priority")
	fmt.Println("  - help: Show this help message")
	fmt.Println("  - exit: Exit the application")
	fmt.Println("To get started, use the 'add' command to create your first task.")
	fmt.Println("You can also use the 'list' command to view all tasks.")
	fmt.Println("For more information on a specific command, use 'todo <command> --help'.")
	fmt.Println("Happy task management!")
	fmt.Println("This is a simple CLI application to manage your tasks.")
	fmt.Println("You can add, list, update, and delete tasks.")
	fmt.Println("Use the 'filter' command to filter tasks by status or priority.")
	fmt.Println("For more information, visit the project's GitHub repository.")
	fmt.Println("You can also use the 'help' command to see a list of available commands.")
	fmt.Println("To exit the application, use the 'exit' command or press Ctrl+C.")
	fmt.Println("Thank you for using the Todo CLI application!")
	fmt.Println("This application is built with Go and uses Cobra for command-line interface.")
}
