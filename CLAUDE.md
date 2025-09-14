# Claude Code Knowledge Base - Toto CLI Project

## Project Overview
Toto is a command-line todo application written in Go that manages tasks based on the current working directory. It uses SQLite for storage and Cobra for CLI framework.

## Key Architecture
- **Language**: Go 1.23.4
- **CLI Framework**: Cobra (`github.com/spf13/cobra`)
- **Database**: SQLite3 (`github.com/mattn/go-sqlite3`) with automatic timestamp triggers
- **UI**: Table output using `github.com/olekukonko/tablewriter`
- **Colors**: `github.com/fatih/color`

## Project Structure
```
/home/sam/coding/toto/
├── main.go                     # Entry point, calls cmd.Execute()
├── cmd/
│   ├── root.go                 # Root command setup, database init
│   ├── reset.go                # Database reset command
│   ├── todo/
│   │   ├── add.go              # Add new todos
│   │   ├── delete.go           # Delete todos
│   │   ├── list.go             # List commands (ls, list, lsl, etc.)
│   │   ├── toggleComplete.go   # Mark todos complete/incomplete
│   │   ├── removeComplete.go   # Remove completed todos
│   │   ├── update.go           # Edit todo title/description
│   │   ├── description.go      # Get todo description
│   │   └── todo.go             # Todo service setup
│   └── projects/
│       ├── projectAdd.go       # Add projects
│       ├── projectDelete.go    # Delete projects
│       ├── projectList.go      # List projects
│       ├── projectUpdate.go    # Edit projects
│       └── projects.go         # Project service setup
├── internal/
│   ├── db/db.go               # Database initialization with auto-timestamp triggers
│   ├── service/               # Business logic
│   ├── models/                # Data structures (projects.go, todo.go)
│   └── utilities/general.go   # Helper functions (ClearScreen)
└── test-directory/            # Test directory (in .gitignore)
```

## Command Structure & Issues

### Working Commands
- `add` - Add new todo with -t (title) and -d (description)
- `ls`/`list` - Basic todo list for current project
- `lsl`/`list-long` - Detailed todo list with dates, descriptions
- `lsla` - All todos regardless of project
- `comp`/`toggle-complete` - Mark todos complete/incomplete
- `edit` - Update todo with -i (id), -t (title), -d (description)
- `del`/`delete` - Delete specific todo
- `cls-comp`/`remove-complete` - Remove completed todos for project
- `description`/`desc` - Get description for single todo with -i (id)
- `proj-add`/`project-add` - Add new project
- `proj-ls`/`project-list` - List all projects
- `proj-edit` - Edit project details
- `reset` - Reset database

### Known Issues (as of analysis)
1. **Project creation prompt bug**: When user selects option 2 to create new project, system loops without actually creating the project (internal/service/todo.go:42-44)

## Key Files for Bug Fixes

### Flag Issues Location
- **File**: `/home/sam/coding/toto/cmd/todo/list.go`
- **Lines 392-395**: Flag registration only for long commands
- **Missing**: Flag registration for `LsCmd` and `GetCmd`

### Flag Variables
```go
var (
    fullDate  bool = false  // -D flag for full timestamps
    allTodos  bool = false  // -A flag for all projects
    clearTerm bool = false  // -C flag for clearing terminal
)
```

## Testing Workflow

### Setup Test Environment
```bash
# Create test directory (already in .gitignore)
mkdir test-directory && cd test-directory

# Build project
cd /home/sam/coding/toto && go build -o toto .

# Add test project
/home/sam/coding/toto/toto proj-add -t "Test Project" -d "Testing"

# Add test todos
/home/sam/coding/toto/toto add -t "Test Todo" -d "Description"
```

### Testing Flag Issues
```bash
# These work (should clear screen):
/home/sam/coding/toto/toto lsl -C
/home/sam/coding/toto/toto list-long -C

# These fail:
/home/sam/coding/toto/toto ls -C    # Error: unknown shorthand flag: 'C'
/home/sam/coding/toto/toto list -C  # Error: unknown shorthand flag: 'C'
```

## Adding Todos to Project List

### Standard Workflow for Claude Code
When finding bugs or implementing features, add todos to the main project:

```bash
cd /home/sam/coding/toto
/home/sam/coding/toto/toto add -t "Bug/Feature Title" -d "Detailed description with file locations and specifics"
```

### Current Project Context
- **Main Project ID**: 3 (Toto project)
- **Main Project Path**: `/home/sam/coding/toto`
- **Database Location**: `todo.db` (in .gitignore)

## Common Tasks for Claude Code

### 1. Bug Investigation
- Build project: `go build -o toto .`
- Test commands in test-directory
- Check help: `/home/sam/coding/toto/toto [command] --help`
- Add findings to project todos

### 2. Flag Issues
- Check `cmd/todo/list.go` for flag registration
- Look for `init()` function at bottom of file
- Ensure flags added to correct commands

### 3. Documentation Updates
- Main README: `/home/sam/coding/toto/README.md`
- Commands table starts around line 124
- Check for typos in root command description

### 4. Command Aliases
- Defined in `cmd/root.go` lines 42-64
- Short and long versions should match README

## Build & Test Commands
```bash
# Build
go build -o toto .

# Test basic functionality
./toto --help
./toto add -t "Test" -d "Testing"
./toto ls

# Test flags (in test directory)
cd test-directory
/home/sam/coding/toto/toto ls -C    # Should fail currently
/home/sam/coding/toto/toto lsl -C   # Should work (clears screen)
```

## IMPORTANT: Claude Code Role & Boundaries

**DO NOT WRITE CODE OR MAKE CODE CHANGES**
- Your role is to ANALYZE, INVESTIGATE, and SUGGEST only
- Do NOT edit any .go files or make code modifications
- Do NOT implement fixes directly
- Do NOT run git commands (commit, push, etc.)
- Focus on understanding problems and providing detailed guidance

**Your Tasks:**
- Investigate bugs and issues by reading code and testing
- Add detailed todos to the project list with specific guidance
- Provide suggestions on how to fix issues (file locations, code patterns, etc.)
- Help manage the todo list and project organization
- Analyze code structure and identify root causes
- Keep this file up to date, command, bugs etc...
- Keep the README file up to date, commands, instructions etc...
- Any testing of command should be done in the ./test-directory/

## Quick Start for Claude Code Sessions

1. **Understand the issue**: Read current todos with `./toto ls`
2. **Test the problem**: Use test-directory for safe testing
3. **Locate the code**: Use file structure above to find relevant files
4. **Add detailed todos**: Add todos to main project with specific implementation guidance
5. **Provide analysis**: Give detailed suggestions on how to fix, but don't fix directly

## Tips
- The project uses directory-based project management
- Always test from test-directory to avoid polluting main project
- Flag variables are shared across commands (potential for bugs)
- ClearScreen utility is in `/home/sam/coding/toto/internal/utilities/general.go`
- Database connection is passed down from root to subcommands
