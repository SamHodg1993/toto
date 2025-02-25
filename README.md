# Todo CLI App

A simple command-line interface (CLI) application for managing your daily tasks and todo items. 

## Why?
- Why toto?
The reason is simple really, every developer at some point will write a todo app, probably more than one. For some reason, whenever I want to write todo, my fingers type toto. It's simple, it's (probably) unique, so I went with it. 
- Why another todo cli tool?
I just kinda wanted to. None of the other tools really fitted what I wanted it to be, plus, I wanted another Golang project.

## Features

- Add new tasks
- List all tasks
- Mark tasks as complete
- Delete tasks
- Edit title and description information
- Reset the database
- Simple and intuitive command-line interface

## Installation

```bash
# Clone the repository
git clone https://github.com/samhodg1993/toto-todo-cli.git

# Navigate to the directory
cd toto-todo-cli
```

### Windows
- Powershell
```powershell
# Run the install script 
./install.ps1
```
- Git Bash 
```bash
# Run the install script
./install.sh
```

### Linux
```bash
# Run the install script
./install.sh
```

## Usage

```bash
# Add a new task
toto add -t <"Title of the todo" (required)> -d <"Description of the todo" (optional)> -c <"Created time" (optional)> -u <"updated time" (optional)>

# List all tasks
toto list
# OR the shorthand version 
toto ls

# List all tasks with more detail
toto list-long
# OR the shorthand version 
toto lsla

# Edit a todo
toto edit -i <todo-id (required)> -t <"New title" (optional)> -d <"New description" (optional)> 

# Mark task as complete
toto toggle-complete <task-id>
# Or the shorthand
toto comp <task-id>

# Delete a task
toto delete <task-id>
# Or the shorthand
toto del <task-id>

# Reset the database
# the confirm flag is optional, if it does not exist, you will be prompted to confirm the action
toto reset -confirm(optional) 

# Show help
toto --help
```

## Commands

| Command | Shorthand | Description |
|---------|-----------|-------------|
| `add`   | -         | Add a new task |
| `list`  | `ls`      | Show all tasks |
| `list-long` | `lsla` | Show detailed task list |
| `edit`  | -        | Edit an existing task's title or description |
| `toggle-complete` | `comp` | Mark a task as complete |
| `delete` | `del`    | Remove a task |
| `help`   | -        | Show help information |
| `reset`  | -        | Reset the database to its initial state |

## Roadmap 
- Introduce projects/workspaces
- Add monday.com integration 
- Add Jira integration

## Prerequisites

- Go 1.x or higher
- Git

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## Author

Sam Hodgkinson- [samhodg1993](https://github.com/samhodg1993)

## Acknowledgments

- Golang Cobra cli tool [cobra](https://github.com/spf13/cobra)
