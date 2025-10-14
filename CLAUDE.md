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
│   ├── todo/
│   │   ├── add.go              # Add new todos
│   │   ├── delete.go           # Delete todos
│   │   ├── list.go             # List commands (ls, list, lsl, etc.)
│   │   ├── toggleComplete.go   # Mark todos complete/incomplete
│   │   ├── removeComplete.go   # Remove completed todos
│   │   ├── update.go           # Edit todo title/description
│   │   ├── description.go      # Get todo description
│   │   └── todo.go             # Todo service setup
│   ├── projects/
│   │   ├── projectAdd.go       # Add projects
│   │   ├── projectDelete.go    # Delete projects
│   │   ├── projectList.go      # List projects
│   │   ├── projectUpdate.go    # Edit projects
│   │   └── projects.go         # Project service setup
│   └── utilityCommands/
│       ├── clean.go            # Clean command (clear, remove completed, list)
│       ├── reset.go            # Database reset command
│       └── utilityCommands.go  # Utility service setup
├── internal/
│   ├── db/db.go               # Database initialization with auto-timestamp triggers
│   ├── service/               # Business logic
│   │   ├── todo.go            # Todo operations with input sanitization
│   │   ├── projects.go        # Project operations with input sanitization
│   │   ├── db.go              # Database utilities (reset)
│   │   └── utilityCommandsService.go # Utility command operations
│   ├── models/                # Data structures (projects.go, todo.go)
│   └── utilities/             # Helper functions
│       ├── general.go         # ClearScreen functionality
│       └── sanitisation.go   # Input sanitization (ANSI escape prevention)
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
- `clean` - Clear screen, remove completed todos, and display remaining todos
- `proj-add`/`project-add` - Add new project
- `proj-ls`/`project-list` - List all projects
- `proj-edit` - Edit project details
- `reset` - Reset database

### Recent Fixes
1. ~~**Project creation prompt bug**: Fixed - now properly calls AddNewProjectWithPrompt() when user selects option 2~~
2. ~~**Project delete commands**: Updated to use -i flag consistently instead of positional arguments~~
3. ~~**Completed_at timestamp**: Fixed ToggleComplete to properly set/unset completed_at timestamps~~
4. ~~**Input sanitization**: Added ANSI escape sequence prevention in titles and descriptions~~
5. ~~**Flag consistency**: All commands confirmed to use -i flag correctly, no positional arguments remain~~

### Known Issues (as of analysis)
- All major known issues have been resolved

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

## Database Schema

### Current Tables
**todos table:**
- id INTEGER PRIMARY KEY
- title VARCHAR(255) NOT NULL
- description TEXT
- project_id INTEGER NOT NULL
- jira_ticket_id INTEGER (nullable - links to jira_tickets.id)
- completed BOOLEAN NOT NULL DEFAULT FALSE
- completed_at DATETIME (nullable)
- created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
- updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP

**projects table:**
- id INTEGER PRIMARY KEY
- title VARCHAR(255) NOT NULL
- description TEXT
- archived BOOLEAN NOT NULL DEFAULT FALSE
- filepath VARCHAR(255) NOT NULL
- created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
- updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP

**jira_tickets table:** (NEW - for Jira integration)
- id INTEGER PRIMARY KEY
- jira_key VARCHAR(50) NOT NULL UNIQUE (e.g., "PROJ-123")
- title VARCHAR(500) NOT NULL
- status VARCHAR(50) NOT NULL (e.g., "To Do", "In Progress", "Done")
- project_key VARCHAR(50) (e.g., "PROJ", "DEV")
- issue_type VARCHAR(50) NOT NULL (e.g., "Story", "Bug", "Task")
- url VARCHAR(500) NOT NULL (full Jira URL)
- last_synced_at DATETIME (nullable - tracks last sync)
- created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
- updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP

### Relationships
- **One-to-Many**: jira_tickets → todos (one Jira ticket can have multiple todos)
- **Many-to-One**: todos → projects (multiple todos belong to one project)
- **Foreign Keys**: todos.jira_ticket_id → jira_tickets.id, todos.project_id → projects.id

### Auto-Update Triggers
- All tables have UPDATE triggers that automatically set updated_at to CURRENT_TIMESTAMP
- Allows tracking when records were last modified

## Jira Integration Progress

### Completed ✅
1. **OAuth 2.0 Authentication** (`cmd/jira/jiraAuth.go`)
   - Browser-based OAuth flow with Atlassian
   - Callback server implementation with state validation (port 8989)
   - Token exchange (authorization code → access token + refresh token)
   - Automatic token refresh with expiry tracking
   - Secure token storage using OS keyring (`github.com/zalando/go-keyring`)
   - Environment variable support via `.env` file (`github.com/joho/godotenv`)

2. **Cloud ID Management** (`internal/utilities/jira_util.go`, `cmd/jira/jiraAuth.go`)
   - `GetUsersJiraCloudId()` - Automatic cloud ID fetch and storage
   - `jira-set-cloud-id` command - Manual cloud ID override
   - Automatic fallback if cloud ID is missing

3. **Jira REST API Client** (`internal/service/jira.go`)
   - `GetSingleJiraTicket()` - Fetch individual tickets by issue key
   - Automatic token refresh integration
   - Automatic cloud ID fallback
   - Proper error handling with user-friendly messages

4. **Data Models** (`internal/models/jira.go`)
   - `JiraTicket` model with validation methods (database storage)
   - `JiraBasedTicket` model for API responses with ADF support
   - `GetDescriptionText()` - Extracts plain text from Atlassian Document Format
   - `JiraConfig` model for configuration management
   - `TokenResponse` model for OAuth responses

5. **Commands** (`cmd/jira/`)
   - `jira-auth` - OAuth authentication
   - `jira-set-cloud-id` - Manual cloud ID configuration
   - `jira-pull -i <issue-key>` - Pull Jira ticket, save to database, create linked todo

6. **Infrastructure**
   - Database schema for `jira_tickets` table
   - Callback server service (`internal/service/jira.go`)
   - Browser opening utility (`internal/utilities/general.go`)
   - Token session management (`HandleJiraSessionBeforeCall()`)

### Environment Setup
Required `.env` file in project root:
```
JIRA_CLIENT_ID=your-client-id
JIRA_CLIENT_SECRET=your-client-secret
```

### OAuth Scopes
- `read:jira-work` - Read Jira data
- `write:jira-work` - Create/update Jira tickets
- `offline_access` - Get refresh token for long-term access

### Keyring Storage Keys
- Service: `toto-cli`
- Keys: `jira-access-token`, `jira-refresh-token`, `access-token-expiry`, `jira-cloud-id`

### Pending Implementation
- ~~Token refresh functionality (when access token expires)~~ ✅ Complete
- ~~Jira REST API client (fetch tickets)~~ ✅ Complete
- ~~`jira-pull` command~~ ✅ Complete - Fetches Jira ticket, saves to database, creates linked todo
- `jira-push` command - Push local todo to Jira as new ticket
- `jira-sync` command - Sync status between todos and Jira
- Display Jira keys in list commands
- Tests for Jira functionality

### Implementation Strategy
**Phase 1 (Current):** Simple CLI commands with ticket ID flags
- `jira-pull -i TICKET-123` - Pull specific ticket by ID
- `jira-push -i <todo-id>` - Push todo to Jira
- `jira-sync` - Sync all linked tickets

**Phase 2 (Future):** Interactive TUI for browsing and managing
- `jira-pull` (no ID) - Opens interactive browser
- `jira-browse` - Dedicated ticket browser with filtering
- Built with [Bubbletea](https://github.com/charmbracelet/bubbletea) for cross-platform support

## Future Enhancements

### TUI (Terminal User Interface)
Plan to add an interactive terminal UI (similar to lazygit) for:
- Browsing and selecting Jira tickets
- Exploring and managing todos with rich context
- Interactive filtering and search
- Multi-select operations

**Tech Stack:** [Bubbletea](https://github.com/charmbracelet/bubbletea) - Cross-platform TUI framework

### LLM Integration (Claude API)
Plan to integrate Claude API for AI-assisted todo management:
- **Title generation** - Generate descriptive titles from brief inputs
- **Renaming** - Improve existing todo titles
- **Description editing** - Expand or refine todo descriptions
- **Criticality assignment** - Suggest priority levels (requires database schema update)
- **Completion order** - Suggest optimal order of completion (future feature)

**Implementation Notes:**
- Initial integration with Claude API via Anthropic's official SDK
- Criticality and order features depend on database schema additions
- Will require API key storage (similar to Jira token storage in keyring)

## Tips
- The project uses directory-based project management
- Always test from test-directory to avoid polluting main project
- Flag variables are shared across commands (potential for bugs)
- ClearScreen utility is in `/home/sam/coding/toto/internal/utilities/general.go`
- Database connection is passed down from root to subcommands
