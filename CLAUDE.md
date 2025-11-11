# Claude Code Knowledge Base - Toto CLI Project

## Project Overview
Toto is a command-line todo application written in Go that manages tasks based on the current working directory. It uses SQLite for storage and Cobra for CLI framework.

## Key Architecture
- **Language**: Go 1.23.4
- **CLI Framework**: Cobra (`github.com/spf13/cobra`)
- **Database**: SQLite3 (`github.com/mattn/go-sqlite3`) with automatic timestamp triggers
- **UI**: Table output using `github.com/olekukonko/tablewriter`
- **Colors**: `github.com/fatih/color`
- **Service Architecture**: Sub-package by domain pattern for scalability and maintainability

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
│       ├── llmHelp.go          # LLM documentation command
│       ├── defaultJiraUrl.go   # Set default Jira URL
│       └── utilityCommands.go  # Utility service setup
├── internal/
│   ├── db/db.go               # Database initialization with auto-timestamp triggers
│   ├── embedded/              # Embedded files (go:embed)
│   │   ├── llms.go            # Embeds LLMs.txt
│   │   └── LLMs.txt           # Comprehensive usage guide for LLMs
│   ├── service/               # Business logic (refactored to sub-packages)
│   │   ├── todo/              # Todo service package
│   │   │   ├── service.go     # Service struct + interfaces
│   │   │   ├── queries.go     # GetTodos, list operations
│   │   │   ├── crud.go        # Add, Update, Delete
│   │   │   ├── complete.go    # Toggle/Remove complete
│   │   │   └── generateList.go # Helper functions for table formatting
│   │   ├── project/           # Project service package
│   │   │   ├── service.go     # Service struct
│   │   │   ├── queries.go     # List, GetById operations
│   │   │   ├── crud.go        # Add, Update, Delete
│   │   │   └── prompts.go     # User interaction logic
│   │   ├── jira/              # Jira service package
│   │   │   ├── service.go     # Service struct
│   │   │   ├── client.go      # API calls
│   │   │   ├── storage.go     # Database operations
│   │   │   ├── auth.go        # OAuth/callback
│   │   │   └── pull.go        # Business logic handlers
│   │   ├── claude/            # Claude AI service package
│   │   │   └── service.go     # AI ticket breakdown
│   │   ├── utility/           # Utility commands service
│   │   │   └── service.go     # Clean command logic
│   │   └── database/          # Database utilities service
│   │       └── service.go     # Reset operations
│   ├── models/                # Data structures (projects.go, todo.go, jira.go)
│   └── utilities/             # Helper functions
│       ├── general.go         # ClearScreen, browser opening
│       ├── sanitisation.go    # Input sanitization (ANSI escape prevention)
│       └── jira_util.go       # Jira session management
└── test-directory/            # Test directory (in .gitignore)
```

## Command Structure & Issues

### Working Commands
- `add` - Add new todo with -t (title) and -d (description)
- `ls`/`list` - Basic todo list for current project
- `lsl`/`list-long` - Detailed todo list with dates, descriptions
- `lsla` - All todos regardless of project
- `comp`/`toggle-complete` - Mark todos complete/incomplete with -i (single id) or -I (bulk comma-separated ids)
- `edit` - Update todo with -i (id), -t (title), -d (description)
- `del`/`delete` - Delete specific todo with -i (single id) or -I (bulk comma-separated ids)
- `cls-comp`/`remove-complete` - Remove completed todos for project
- `description`/`desc` - Get description for single todo with -i (id)
- `clean` - Clear screen, remove completed todos, and display remaining todos
- `proj-add`/`project-add` - Add new project
- `proj-ls`/`project-list` - List all projects
- `proj-edit` - Edit project details
- `reset` - Reset database
- `llm-help`/`llm` - Display comprehensive LLM-focused usage documentation

### Recent Fixes
1. ~~**Project creation prompt bug**: Fixed - now properly calls AddNewProjectWithPrompt() when user selects option 2~~
2. ~~**Project delete commands**: Updated to use -i flag consistently instead of positional arguments~~
3. ~~**Completed_at timestamp**: Fixed ToggleComplete to properly set/unset completed_at timestamps~~
4. ~~**Input sanitization**: Added ANSI escape sequence prevention in titles and descriptions~~
5. ~~**Flag consistency**: All commands confirmed to use -i flag correctly, no positional arguments remain~~
6. ~~**List command refactoring**: All list commands now use helper functions (FormatTodoTableRow, FormatTodoTableRowLong) to reduce code duplication~~
7. ~~**Reverse list support**: All list commands now support -r flag using slices.Reverse() for cleaner code~~
8. ~~**Flag registration**: All list commands (ls, list, lsl, list-long, lsla) now have proper flag registration for -C, -r, -A, -D flags~~
9. ~~**Clean command refactoring**: Updated clean command to use FormatTodoTableRow helper and support -r flag for reverse order~~
10. ~~**Bulk operations**: Added -I flag for bulk complete and delete operations with comma-separated IDs~~
11. ~~**Jira command aliases**: Added `jp` (jira-pull) and `jpc` (jira-pull-claude) shorthand commands~~

### Known Issues (as of analysis)
- All major known issues have been resolved

## Key Files for Bug Fixes

### List Commands Architecture
- **File**: `/home/sam/coding/toto/cmd/todo/list.go`
- **Helper Functions**: `/home/sam/coding/toto/internal/service/todo/generateList.go`
  - `FormatTodoTableRow()` - Formats simple 3-column rows (ID, Todo, Status)
  - `FormatTodoTableRowLong()` - Formats detailed 8-column rows with dates and descriptions
- **All commands now support**: `-r` (reverse), `-C` (clear terminal), `-A` (all todos), `-D` (full dates for long commands)

### Flag Variables
```go
var (
    fullDate    bool = false  // -D flag for full timestamps
    allTodos    bool = false  // -A flag for all projects
    clearTerm   bool = false  // -C flag for clearing terminal
    reverseList bool = false  // -r flag for reversing list order
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

### Testing List Commands
```bash
# All list commands now support flags:
/home/sam/coding/toto/toto ls -C        # Clear screen before listing
/home/sam/coding/toto/toto list -r      # Reverse order
/home/sam/coding/toto/toto lsl -C -r    # Clear screen and reverse
/home/sam/coding/toto/toto list-long -A # Show all todos regardless of project
/home/sam/coding/toto/toto lsl -D       # Show full date timestamps
/home/sam/coding/toto/toto lsla -C -r -D # All flags combined

# Clean command also supports -r flag:
/home/sam/coding/toto/toto clean        # Clear, remove completed, list remaining
/home/sam/coding/toto/toto clean -r     # Same but in reverse order

# Bulk operations:
/home/sam/coding/toto/toto comp -i 1    # Toggle single todo
/home/sam/coding/toto/toto comp -I 1,2,3,4 # Toggle multiple todos
/home/sam/coding/toto/toto del -i 1     # Delete single todo
/home/sam/coding/toto/toto del -I 1,2,3,4  # Delete multiple todos

# Jira shortcuts:
/home/sam/coding/toto/toto jp -i PROJ-123    # Same as jira-pull
/home/sam/coding/toto/toto jpc -i PROJ-123   # Same as jira-pull-claude

# Jira URL management:
/home/sam/coding/toto/toto jira-set-default-url -u sta2020.atlassian.net  # Set global default
/home/sam/coding/toto/toto project-set-jira-url -p 3 -u custom.atlassian.net  # Set project-specific URL
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
- Keep this file (CLAUDE.md) up to date, commands, bugs etc...
- Keep the README.md file up to date, commands, instructions etc...
- Keep the LLMs.txt file (internal/embedded/LLMs.txt) up to date with new commands, workflows, and usage patterns
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
- criticality INTEGER NOT NULL DEFAULT 0 (reserved for future TUI/AI features - not yet exposed in CLI)

**projects table:**
- id INTEGER PRIMARY KEY
- title VARCHAR(255) NOT NULL
- description TEXT
- archived BOOLEAN NOT NULL DEFAULT FALSE
- filepath VARCHAR(255) NOT NULL
- jira_url VARCHAR(500) (nullable - project-specific Jira URL)
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
1. **API Token Authentication** (`cmd/jira/jiraAuth.go`)
   - Simple prompt-based authentication flow
   - User provides: Jira URL, email, and API token
   - Credentials validated with test API call before storage
   - Secure storage using OS keyring (`github.com/zalando/go-keyring`)
   - No OAuth app creation required - users create tokens at https://id.atlassian.com/manage-profile/security/api-tokens

2. **Jira REST API Client** (`internal/service/jira/client.go`)
   - `GetSingleJiraTicket()` - Fetch individual tickets by issue key
   - Uses Basic Auth with email:api_token
   - Direct URL construction (no cloud ID lookup needed)
   - Proper error handling with user-friendly messages
   - Duplicate ticket handling - UPDATE existing records on re-pull

3. **Data Models** (`internal/models/jira.go`)
   - `JiraTicket` model with validation methods (database storage)
   - `JiraBasedTicket` model for API responses with ADF support
   - `ADFNode` - Recursive structure for Atlassian Document Format parsing
   - `extractTextFromADF()` - Recursive ADF parser supporting bulletList, orderedList, listItem, paragraph, text, hardBreak
   - `GetDescriptionText()` - Extracts plain text from complex ADF structures including bullet points

4. **Commands** (`cmd/jira/`)
   - `jira-auth` - API token authentication
   - `jira-pull` / `jp -i <issue-key>` - Pull Jira ticket, save to database, create linked todo
   - `jira-pull-claude` / `jpc -i <issue-key>` - Pull Jira ticket and use Claude AI to break it into subtasks

5. **Claude AI Integration** (`internal/service/claude/service.go`)
   - `BreakdownJiraTicketWithClaude()` - Uses Claude Sonnet to intelligently break down Jira tickets
   - Smart prompt that detects explicit task lists vs high-level descriptions
   - Extracts bullet-pointed lists directly or breaks down complex tasks into 3-8 subtasks
   - API key from `CLAUDE_API_KEY` environment variable
   - Uses `anthropic-sdk-go` with model alias `claude-sonnet-4-5` (automatically uses latest version)

6. **Infrastructure**
   - Database schema for `jira_tickets` table
   - Session management (`HandleJiraSessionBeforeCall()` in `internal/utilities/jira_util.go`)
   - Credential validation during authentication

### Environment Setup
Optional `.env` file in project root:
```
S_DEV_MODE=true

# Optional: Claude API key for AI-powered ticket breakdown
# Can also be set via: export CLAUDE_API_KEY=your-key
# CLAUDE_API_KEY=your-claude-api-key
```

### Authentication Flow
1. User runs `toto jira-auth`
2. Prompted for:
   - Jira URL (e.g., https://mycompany.atlassian.net)
   - Email address
   - API token (created at https://id.atlassian.com/manage-profile/security/api-tokens)
3. Credentials validated with test API call
4. Stored securely in OS keyring for future use

### Keyring Storage Keys
- Service: `toto-cli`
- Keys: `jiraURL`, `jiraEmail`, `jiraApiKey`

### Project-Specific Jira URL Architecture ✅

**Overview:**
The Jira URL system uses a smart two-tier fallback strategy: project-specific URLs stored in the database override a global default stored in the keyring.

**Storage Strategy:**
1. **Database (projects.jira_url)** - Project-specific Jira URL (overrides default)
2. **Keyring (jiraURL)** - Global default Jira URL (fallback for all projects)

**Fallback Logic Flow:**
```
User runs jira command (jp, jpc, etc.)
  ↓
Check project's jira_url field in database
  ↓ (if NULL/empty)
Check keyring for default jiraURL
  ↓ (if found)
Prompt: "Use default [URL] for this project? (y/n)"
  - Yes → Save to project's jira_url, use it
  - No → Prompt for new URL, save to project's jira_url, use it
  ↓ (if no default)
Prompt for URL, save as both project-specific AND default (for first-time setup)
```

**Key Functions:**
- `GetProjectJiraURL()` (`internal/service/project/queries.go:49-76`) - Implements fallback logic with user prompts
- `HandleSetProjectJiraURL()` (`internal/service/project/crud.go:21-68`) - Handles user interaction for choosing/entering Jira URL
- `EnsureHTTPS()` (`internal/utilities/jira_util.go`) - Automatically adds `https://` protocol to URLs without it
- `UpdateProjectsJiraUrl()` (`internal/service/project/crud.go:264-281`) - Updates project's jira_url field

**New Commands:**
- `jira-set-default-url -u <url>` (`cmd/utilityCommands/defaultJiraUrl.go`) - Update global default Jira URL in keyring
- `project-set-jira-url -p <id> -u <url>` (`cmd/projects/setProjectJiraUrl.go`) - Set project-specific Jira URL manually

**Benefits:**
- Multi-tenant support (different projects can use different Jira instances)
- User-friendly prompts on first use
- Flexible: works for users with one or multiple Jira instances
- Automatic protocol handling (accepts URLs with or without `https://`)

### Pending Implementation
- ~~Token refresh functionality (when access token expires)~~ ✅ Complete
- ~~Jira REST API client (fetch tickets)~~ ✅ Complete
- ~~`jira-pull` command~~ ✅ Complete - Fetches Jira ticket, saves to database, creates linked todo
- ~~`jira-pull-claude` command~~ ✅ Complete - Uses Claude AI to break down tickets into subtasks
- ~~ADF (Atlassian Document Format) parser~~ ✅ Complete - Supports bullet lists, ordered lists, paragraphs, etc.
- `jira-push` command - Push local todo to Jira as new ticket
- `jira-sync` command - Sync status between todos and Jira
- Display Jira keys in list commands
- Tests for Jira functionality

### Implementation Strategy
**Phase 1 (Completed):** Simple CLI commands with ticket ID flags
- ✅ `jira-pull -i TICKET-123` - Pull specific ticket by ID
- ✅ `jira-pull-claude -i TICKET-123` - Pull ticket and break down with Claude AI
- 🚧 `jira-push -i <todo-id>` - Push todo to Jira (pending)
- 🚧 `jira-sync` - Sync all linked tickets (pending)

**Phase 2 (Future):** Interactive TUI for browsing and managing
- `jira-pull` (no ID) - Opens interactive browser
- `jira-browse` - Dedicated ticket browser with filtering
- Built with [Bubbletea](https://github.com/charmbracelet/bubbletea) for cross-platform support

## LLM Documentation & AI Assistant Support

### Overview
Toto includes comprehensive documentation specifically designed for Large Language Models and AI assistants. This makes it easy for AI tools like Claude Code, GitHub Copilot, and ChatGPT to effectively use toto commands.

### Implementation (Completed ✅)

**Hybrid Approach:**
1. **Embedded in binary** - LLMs.txt file is embedded using `go:embed` directive
2. **Command access** - `toto llm-help` (alias: `toto llm`) outputs the full documentation
3. **Auto-updated file** - Every time toto runs, it updates `~/.config/toto/LLMs.txt` with the latest version

**File Locations:**
- **Source:** `/internal/embedded/LLMs.txt` (embedded via `/internal/embedded/llms.go`)
- **User system:** `~/.config/toto/LLMs.txt` (auto-updated on every run in `cmd/root.go`)
- **Command:** `cmd/utilityCommands/llmHelp.go`

**Documentation Contents:**
- Complete command reference with examples
- Common workflows and usage patterns
- All flags and their combinations
- Jira integration workflows (including `jpc` AI breakdown)
- Tips for AI assistants working with toto
- Output format interpretation guide
- Database schema and relationships

**Architecture:**
```
internal/embedded/
├── llms.go           # Embeds LLMs.txt using go:embed
└── LLMs.txt          # Comprehensive usage guide

cmd/root.go           # ensureLLMUsageFile() - auto-generates ~/.config/toto/LLMs.txt
cmd/utilityCommands/
└── llmHelp.go        # Command implementation (uses embedded.LLMUsageDoc)
```

**Usage:**
```bash
# Display in terminal
toto llm-help

# Save to file
toto llm-help > toto-guide.txt

# Access auto-generated copy
cat ~/.config/toto/LLMs.txt
```

**Benefits:**
- AI assistants can quickly understand all toto capabilities
- Works immediately after `go install` (embedded in binary)
- Persistent reference available at `~/.config/toto/LLMs.txt`
- Reduces need for AI to search documentation or guess command syntax
- Particularly powerful with `jpc` command for Jira ticket breakdown + refinement workflows

## Future Enhancements

### TUI (Terminal User Interface)
Plan to add an interactive terminal UI (similar to lazygit) for:
- Browsing and selecting Jira tickets
- Exploring and managing todos with rich context
- Interactive filtering and search
- Multi-select operations
- **Priority/Criticality management** - Visual sliders and drag-to-reorder by priority

**Tech Stack:** [Bubbletea](https://github.com/charmbracelet/bubbletea) - Cross-platform TUI framework

**Note:** The `criticality` database field has been added (INTEGER NOT NULL DEFAULT 0) but CLI commands are intentionally postponed until TUI is implemented. Managing priorities via CLI flags is too tedious - TUI will make this feature delightful and actually useful.

### LLM Integration (Claude API)
✅ **Completed:** Jira ticket breakdown using Claude Sonnet
- `jira-pull-claude` command breaks down Jira tickets into subtasks
- Uses `anthropic-sdk-go` with model alias `claude-sonnet-4-5`
- Smart detection of explicit task lists vs high-level descriptions

🚧 **Future AI Features:**
- **Title generation** - Generate descriptive titles from brief inputs
- **Renaming** - Improve existing todo titles
- **Description editing** - Expand or refine todo descriptions
- **Criticality assignment** - Suggest priority levels (database field exists, awaiting TUI/AI implementation)
- **Completion order** - Suggest optimal order of completion

## Service Architecture (Refactored October 2024)

The service layer has been refactored from monolithic files to **sub-packages by domain**:

### Benefits
- **Modular Organization**: Each domain has its own package
- **Logical File Separation**: No more 400+ line service files
- **Easy Navigation**: Clear file names indicate purpose
- **Scalability**: Easy to add new features within each domain
- **Testability**: Easier to unit test smaller, focused files
- **Consistent Pattern**: All services follow the same structure

### Service Package Structure
Each service package typically contains:
- `service.go` - Service struct, constructor (`New()`), and interfaces
- `queries.go` - Read/List operations
- `crud.go` - Create, Update, Delete operations
- Domain-specific files (e.g., `complete.go` for todo completion, `prompts.go` for user interaction)

### Dependency Injection
Services that depend on other services use **interface-based dependency injection**:
- Interfaces defined in the consuming service
- Dependencies injected via `SetDependencies()` or `SetProjectService()` methods
- Wired together in `cmd/root.go` after all services are initialized

Example from `root.go`:
```go
// Jira service needs todo and project services
jira.JiraService.SetDependencies(todo.TodoService, projects.ProjectService)
// Todo service needs project service
todo.TodoService.SetProjectService(projects.ProjectService)
// Utility service needs todo and project services
utilityCommands.UtilityService.SetDependencies(todo.TodoService, projects.ProjectService)
```

## Tips
- The project uses directory-based project management
- Always test from test-directory to avoid polluting main project
- Flag variables are shared across commands (potential for bugs)
- ClearScreen utility is in `/home/sam/coding/toto/internal/utilities/general.go`
- Database connection is passed down from root to subcommands
- Services are organized by domain in sub-packages under `internal/service/`
