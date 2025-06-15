# Todo CLI in Go

A simple, fast, and extensible command-line Todo List application written in Go. This project helps you manage your daily tasks efficiently using a CSV file for persistent storage. Built with the power of Go!

## Features

- Add new tasks quickly from the command line
- List all tasks in a well-formatted table
- Delete tasks by description
- Delete all tasks at once using the `-a` flag
- Tasks are timestamped with their creation date
- Persistent storage to CSV file
- Uses file locking for safe concurrent access
- Modular code structure for easy extension

## Getting Started

### Prerequisites

- [Go](https://golang.org/dl/) 1.18 or later

### Installation

Clone the repository:

```sh
git clone https://github.com/yashpal2104/todo-list-in-go.git
cd todo-list-in-go
```

Build the CLI:

```sh
go build -o tasks
```

### Usage

Run the CLI application:

```sh
./tasks
```

#### Available Commands

| Command                | Description                          |
|------------------------|--------------------------------------|
| `add`                  | Add a new task                       |
| `list`                 | List all tasks                       |
| `delete`               | Delete a task by description         |
| `delete -a`            | Delete all tasks at once             |
| `help`                 | Show help for any command            |

Example usage:

- Add a new task:
  ```sh
  ./tasks add "Buy groceries"
  ```
- List all tasks:
  ```sh
  ./tasks list
  ```
- Delete a task:
  ```sh
  ./tasks delete "Buy groceries"
  ```
- Delete all tasks:
  ```sh
  ./tasks delete -a
  ```
  Output:
  ```
  All tasks deleted.

  There are currently no tasks. Please use the 'tasks add <description>' to add your tasks

  Thank you for using the Todo CLI application!
  ```

### How It Works

- Tasks are stored in `output.csv` with fields: ID, Description, CreatedAt.
- Each command is implemented as a Cobra subcommand (see `cmd/root.go`).
- Adding a task appends a new entry to the CSV file.
- Listing tasks reads and displays all tasks in a formatted table.
- Deleting tasks removes entries with matching descriptions from the CSV.
- Deleting all tasks with `delete -a` clears the tasks file.

### Project Structure

```
.
├── cmd/
│   ├── root.go         # CLI command definitions (Cobra)
│   ├── csv_record.go   # CSV read/write helpers
│   ├── file.go         # File utility functions (locking, safe IO)
│   └── writer.go       # CSV writer helpers
├── main.go             # Entry point, invokes CLI
├── go.mod, go.sum      # Go module files
├── output.csv          # Task storage (created on first run)
└── tasks               # (Built binary after compilation)
```

### Extending

- Add more commands by defining new Cobra subcommands in the `cmd/` directory.
- The `Item` struct in `cmd/root.go` can be extended to track more attributes (e.g., priority, due date).
- CSV helpers are modular for easy changes to storage format.

### Contributing

Contributions are welcome! Please open issues or submit pull requests for improvements or bug fixes.

### License

MIT License.

---

Thank you for using the Todo CLI application!  
Happy task management!
