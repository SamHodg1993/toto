# Todo CLI App

A simple command-line interface (CLI) application for managing your daily tasks and todo items. This is currently and will always be (while under my control) a FOSS (Free and Open Source Software) application.

## Why?
- Why toto?
The reason is simple really, every developer at some point will write a todo app, probably more than one. For some reason, whenever I want to write todo, my fingers type toto. It's simple, it's (probably) unique, so I went with it. 
- Why another todo cli tool?
I just kinda wanted to. None of the other tools really fitted what I wanted it to be, plus, I wanted another Golang project.
- What makes toto special?
Right now, just that we base projects around the current working directory. Soon though, we will be adding Monday.com and Jira integrations!!!

## Features

- Add new tasks
- List all tasks
- List project tasks
- Mark tasks as complete
- Delete tasks
- Edit title and description information
- Reset the database
- Projects based on the current working directory
- Simple and intuitive command-line interface
- Build so you spend less time planning...

## Installation

### For Users (Recommended)
```bash
# Install directly with Go (requires Go 1.21+)
# Works on Windows, Mac, and Linux
go install github.com/samhodg1993/toto-todo-cli@latest
```

The `toto` command will be available in your terminal once `$GOPATH/bin` is in your PATH.

### For Developers/Contributors

```bash
# Clone the repository
git clone https://github.com/samhodg1993/toto-todo-cli.git

# Navigate to the directory
cd toto-todo-cli

# Build locally
go build -o toto .
```

#### Windows
- Powershell
```powershell
# Run the install script
cd toto-todo-cli
./install.ps1
```

#### Linux & Mac
```bash
# Run the install script
cd toto-todo-cli
./install.sh
```

## Uninstallation

### Windows
- Powershell
```powershell
# Run the install script 
./uninstall.ps1
```

### Linux
```bash
# Run the install script
./uninstall.sh
```

## Usage

When using these commands, by default, they will only get the todo's that exist for a project which has been linked to the current working directory unless you use either the lsla command or the list-long command with the -A flag!

```bash
# Add a new task
toto add -t <"Title of the todo" (required)> -d <"Description of the todo" (optional)> -c <"Created time" (optional)> -u <"updated time" (optional)>

# List all tasks
toto list
# OR the shorthand version 
toto ls

# List all tasks with more detail
toto list-long (with optional -A flag)
# OR the shorthand version 
toto lsl
# OR for all todos regardless of project 
toto lsla

# Edit a todo
toto edit -i <todo-id (required)> -t <"New title" (optional)> -d <"New description" (optional)> 

# Mark task as complete
toto toggle-complete <task-id>
# Or the shorthand
toto comp <task-id>

# Clear the completed todos for the current project
toto remove-complete
# Or the shorthand
toto cls-comp

# Delete a task
toto delete <task-id>
# Or the shorthand
toto del <task-id>

# Reset the database
# the confirm flag is optional, if it does not exist, you will be prompted to confirm the action
toto reset -confirm(optional) 

# List projects 
toto proj-ls 
# Or the shorthand 
toto prls 

# Add a new project 
toto project-add -t <title (required)> -d <description (optional)>
# Or the shorthand
toto proj-add -t <title (required)> -d <description (optional)>

# Show help
toto --help
```

## Commands and Flags

| Command           | Shorthand   | Description                                        | Flags                                                        |
| ----------------- | ----------- | -------------------------------------------------- | ------------------------------------------------------------ |
| `add`             | -           | Add a new task                                     | `-t`: specify title, `-d`: specify description, `-c`: specify created-at, `-u`: specify updated-at, `-p`: specify project-id |
| `list`            | `ls`        | Show all tasks                                     | `-C`: clear terminal before render                          |
| `list-long`       | `lsl`       | Show detailed task list                            | `-D`: get full date, `-C`: clear terminal before render      |
| -                 | `lsla`      | Show detailed task list for all projects          | `-D`: get full date, `-C`: clear terminal before render      |
| `edit`            | -           | Edit an existing task's title or description       | `-t`: text for title update, `-d`: description for update, `-i`: target todo id |
| `description`     | `desc`      | Get description for a single todo                  | `-i`: target todo id                                        |
| `toggle-complete` | `comp`      | Mark a task as complete                            |                                                              |
| `remove-complete` | `cls-comp`  | Remove all completed todos for the current project |                                                              |
| `clean`           | -           | Clear screen, remove completed todos, and show remaining | -                                                            |
| `delete`          | `del`       | Remove a task                                      |                                                              |
| `help`            | -           | Show help information                              |                                                              |
| `reset`           | -           | Reset the database to its initial state            |                                                              |
| `project-list`    | `proj-ls`   | Show all projects                                  | `-C`: clear terminal before render                          |
| `project-add`     | `proj-add`  | Add a new project                                  | `-t`: specify title, `-d`: specify description, `-f`: specify project filepath |
| `project-delete`  | `proj-del`  | Delete an existing project                         | `-i`: target project id                                     |
| -                 | `proj-edit` | Update a single project                            | `-t`: text for title update, `-f`: text for filepath update , `-i`: target project id, `-d`: text for description update |
| `completion`      | -           | Generate autocompletion script for specified shell | Run `toto completion --help` for shell options               |

## Examples

Add a new task with a description:
```bash
toto add -t "Implement login feature" -d "Add user authentication to the API"
```

List tasks in current project:
```bash
toto ls
```

Mark task as complete:
```bash
toto comp 1
```

## Roadmap 
- Complete projects/workspaces
- Add monday.com integration 
- Add Jira integration

## Prerequisites

- Go 1.21 or higher
- Git (for development)
- SQLite3 (automatically handled by Go modules)

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## Author

Sam Hodgkinson- [samhodg1993](https://github.com/samhodg1993)

## Contributors 

- You could be the first!

## License

This project is licensed under the MIT License - see the LICENSE file for details.
