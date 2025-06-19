# Todo-CLI

A powerful and interactive command-line application for managing your tasks with a clean interface and user authentication.

![Todo-CLI Demo](https://example.com/todo-cli-demo.gif) <!-- Consider adding a demo GIF here -->

## Features

- **User Authentication**: Secure login and signup functionality
- **Session Management**: Remembers logged-in users between sessions
- **Task Management**:
  - Add tasks with title, description, and priority levels
  - List tasks with a beautiful interactive interface
  - Complete/uncomplete tasks
  - Update task details
  - Delete tasks
- **Interactive UI**: Built with [Bubble Tea](https://github.com/charmbracelet/bubbletea) for a smooth terminal UI experience
- **Data Persistence**: All data is stored locally in JSON files

## Installation

### Using Go Install

If you have Go installed (version 1.18+), you can install Todo-CLI directly:

```bash
go install github.com/Blaze5333/todo-cli@latest
```

### From Source

```bash
# Clone the repository
git clone https://github.com/Blaze5333/todo-cli.git

# Navigate to the project directory
cd todo-cli

# Build the application
go build -o todo-cli

# Move the binary to your path (optional)
sudo mv todo-cli /usr/local/bin/
```

## Usage

### User Management

#### Sign Up
```bash
todo-cli signup
```
Follow the interactive prompts to create a new user account.

#### Login
```bash
todo-cli login
```
Enter your username and password when prompted.

#### Logout
```bash
todo-cli logout
```

### Task Management

#### Add a New Task
```bash
todo-cli add
```
Follow the prompts to enter task title, description, and priority level.

#### List All Tasks
```bash
todo-cli list
```
This will open an interactive interface where you can view and manage your tasks.

#### Complete a Task
In the interactive list view, select a task and press the appropriate key to mark it as complete.

#### Update a Task
In the interactive list view, select a task and press the appropriate key to update its details.

#### Delete a Task
In the interactive list view, select a task and press the appropriate key to delete it.

### Interactive UI Commands

When in the interactive task list view:
- **Arrow keys**: Navigate through tasks
- **Enter**: Select a task
- **a**: Add a new task
- **d**: Delete the selected task
- **c**: Mark the selected task as complete/incomplete
- **u**: Update the selected task
- **q**: Quit the interactive view

## Project Structure

```
todo-cli/
├── cmd/                    # Command definitions using Cobra
│   ├── add.go              # Add task command
│   ├── complete.go         # Complete task command
│   ├── delete.go           # Delete task command
│   ├── list.go             # List tasks command
│   ├── login.go            # Login command
│   ├── root.go             # Root command
│   ├── signup.go           # Signup command
│   └── update.go           # Update task command
├── internal/               # Internal packages
│   ├── bubbletea/          # Interactive UI components
│   │   └── initialization.go # BubbleTea initialization and UI logic
│   ├── storage/            # Data persistence layer
│   │   └── storage.go      # Generic JSON storage implementation
│   ├── todo/               # Todo business logic
│   │   └── todo.go         # Todo operations
│   └── user/               # User management
│       ├── session.go      # User session management
│       └── user.go         # User authentication
├── utils/                  # Utility functions
│   └── showmessage.go      # Message display utilities
├── go.mod                  # Go module definition
├── go.sum                  # Go module checksums
├── main.go                 # Application entry point
├── LICENSE                 # License information
├── todos.json              # Todo data storage
└── users.json              # User data storage
```

## Data Storage

Todo-CLI uses local JSON files for data storage:

- `users.json`: Stores user credentials
- `todos.json`: Stores todo items for all users
- `session.json`: Stores the current user session

## Dependencies

Todo-CLI relies on the following key libraries:

- [Cobra](https://github.com/spf13/cobra): Command-line interface framework
- [Bubble Tea](https://github.com/charmbracelet/bubbletea): Terminal UI framework
- [Prompt UI](https://github.com/manifoldco/promptui): Interactive prompts
- [Table](https://github.com/aquasecurity/table): Table rendering in terminal

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the terms of the license included in the repository.

## Acknowledgments

- [Cobra](https://github.com/spf13/cobra) for the CLI framework
- [Bubble Tea](https://github.com/charmbracelet/bubbletea) for the terminal UI
- All contributors and users of this project
