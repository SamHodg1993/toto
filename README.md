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
- **Jira API token authentication** - Securely authenticate with Jira using API tokens
- Simple and intuitive command-line interface
- Build so you spend less time planning...

## LLM & AI Assistant Support

Toto includes comprehensive documentation specifically designed for Large Language Models and AI assistants like Claude Code, GitHub Copilot, and ChatGPT.

**Access the LLM documentation:**

```bash
# Display comprehensive usage guide in terminal
toto llm-help

# Save to a file for reference
toto llm-help > toto-guide.txt

# Or access the auto-generated file directly
cat ~/.config/toto/LLMs.txt
```

**What's included:**
- Complete command reference with examples
- Common workflows and usage patterns
- Flag documentation and combinations
- Jira integration workflows
- Tips for AI assistants working with toto
- Output format interpretation

**For AI assistants:** The LLM documentation provides everything needed to effectively use toto, including:
- How to pull Jira tickets and break them down with AI
- Bulk operations for managing multiple todos
- Directory-based project management patterns
- Best practices for todo refinement and cleanup

**File locations:**
- **In repository:** `/internal/embedded/LLMs.txt` (embedded in binary)
- **User system:** `~/.config/toto/LLMs.txt` (auto-updated on every run)

This makes toto particularly powerful when combined with AI assistants - they can pull Jira tickets, create todo lists, refine descriptions, and help manage your workflow with full context of how toto works.

## Installation

### For Users (Recommended)
```bash
# Install directly with Go (requires Go 1.21+)
# Works on Windows, Mac, and Linux
go install github.com/odgy8/toto@latest
```

### Ensuring `toto` is in your PATH

After installation, make sure `$GOPATH/bin` (typically `$HOME/go/bin`) is in your PATH.

**Check if it's already in your PATH:**
```bash
echo $PATH | grep go/bin
```

**If not, add it to your shell configuration:**

**Bash** (add to `~/.bashrc`):
```bash
export PATH="$HOME/go/bin:$PATH"
```

**Zsh** (add to `~/.zshrc`):
```bash
export PATH="$HOME/go/bin:$PATH"
```

**Fish** (add to `~/.config/fish/config.fish`):
```fish
set -gx PATH $HOME/go/bin $PATH
```

**Then reload your shell:**
```bash
source ~/.bashrc  # or source ~/.zshrc, or restart your terminal
```

The `toto` command will now be available globally.

### For Developers/Contributors

```bash
# Clone the repository
git clone https://github.com/odgy8/toto.git

# Navigate to the directory
cd toto

# Build locally
go build -o toto .
```

#### Windows
- Powershell
```powershell
# Run the install script
cd toto
./install.ps1
```

#### Linux & Mac
```bash
# Run the install script
cd toto
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

# Mark task as complete (single)
toto toggle-complete -i <task-id>
# Or the shorthand
toto comp -i <task-id>

# Mark multiple tasks as complete (bulk)
toto comp -I 1,2,3,4

# Clear the completed todos for the current project
toto remove-complete
# Or the shorthand
toto cls-comp

# Clean: clear screen, remove completed, and list remaining todos
toto clean
# With optional -r flag to reverse list order
toto clean -r

# Delete a task (single)
toto delete -i <task-id>
# Or the shorthand
toto del -i <task-id>

# Delete multiple tasks (bulk)
toto del -I 1,2,3,4

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
| `list`            | `ls`        | Show all tasks                                     | `-C`: clear terminal before render, `-r`: reverse list order, `-A`: show all projects |
| `list-long`       | `lsl`       | Show detailed task list                            | `-D`: get full date, `-C`: clear terminal before render, `-r`: reverse list order, `-A`: show all projects |
| -                 | `lsla`      | Show detailed task list for all projects          | `-D`: get full date, `-C`: clear terminal before render, `-r`: reverse list order |
| `edit`            | -           | Edit an existing task's title or description       | `-t`: text for title update, `-d`: description for update, `-i`: target todo id |
| `description`     | `desc`      | Get description for a single todo                  | `-i`: target todo id                                        |
| `toggle-complete` | `comp`      | Mark a task as complete                            | `-i`: single todo id, `-I`: comma-separated todo ids for bulk operations |
| `remove-complete` | `cls-comp`  | Remove all completed todos for the current project |                                                              |
| `clean`           | -           | Clear screen, remove completed todos, and show remaining | `-r`: reverse list order                                    |
| `delete`          | `del`       | Remove a task                                      | `-i`: single todo id, `-I`: comma-separated todo ids for bulk operations |
| `help`            | -           | Show help information                              |                                                              |
| `reset`           | -           | Reset the database to its initial state            |                                                              |
| `project-list`    | `proj-ls`   | Show all projects                                  | `-C`: clear terminal before render                          |
| `project-add`     | `proj-add`  | Add a new project                                  | `-t`: specify title, `-d`: specify description, `-f`: specify project filepath |
| `project-delete`  | `proj-del`  | Delete an existing project                         | `-i`: target project id                                     |
| -                 | `proj-edit` | Update a single project                            | `-t`: text for title update, `-f`: text for filepath update , `-i`: target project id, `-d`: text for description update |
| **Jira Integration** | -         | **Jira ticket management**                        | -                                                            |
| `jira-auth`       | -           | Authenticate with Jira using API token            | Prompts for Jira URL, email, and API token. Stores securely in OS keyring |
| `jira-pull`       | `jp`        | Pull a Jira ticket and create a linked todo       | `-i`: Jira ticket ID (e.g., PROJ-123)                       |
| `jira-pull-claude` | `jpc`      | Pull Jira ticket and break it into subtasks with AI | `-i`: Jira ticket ID (e.g., PROJ-123)                       |
| `jira-set-default-url` | -       | Update the global default Jira URL               | `-u`: Jira URL (e.g., mycompany.atlassian.net). Stored in keyring as fallback for all projects |
| `project-set-jira-url` | -       | Set project-specific Jira URL                    | `-p`: project ID, `-u`: Jira URL. Overrides default for specific project |
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

List tasks in reverse order (newest first):
```bash
toto ls -r
```

List detailed tasks with full dates:
```bash
toto lsl -D
```

Mark task as complete:
```bash
toto comp -i 1
```

Mark multiple tasks as complete (bulk):
```bash
toto comp -I 1,2,3,4
```

Clean up completed tasks and display remaining:
```bash
toto clean
```

Clean up and display remaining tasks in reverse order:
```bash
toto clean -r
```

**Jira Integration Examples:**

Authenticate with Jira:
```bash
toto jira-auth
```
You'll be prompted to enter your Jira URL, email, and API token (create one at https://id.atlassian.com/manage-profile/security/api-tokens). Your credentials are validated and securely stored in your OS keyring.

Pull a Jira ticket and create a todo:
```bash
toto jira-pull -i PROJ-123
```

Pull a Jira ticket and break it into subtasks with AI:
```bash
toto jira-pull-claude -i PROJ-123
```
This uses Claude AI to intelligently break down the ticket into actionable subtasks. If the description contains a bulleted list, it extracts each item. Otherwise, it analyzes the ticket and creates 3-8 subtasks.

Set a global default Jira URL (used as fallback for all projects):
```bash
toto jira-set-default-url -u mycompany.atlassian.net
```

Set a project-specific Jira URL (overrides the default):
```bash
toto project-set-jira-url -p 3 -u customjira.atlassian.net
```

## Roadmap

### Completed ✅
- ✅ **Jira integration** - OAuth authentication, ticket fetching, and AI-powered breakdown
- ✅ **LLM Integration (Phase 1)** - Claude API for Jira ticket breakdown into subtasks

### Planned Features
- **Terminal UI (TUI)** - Interactive browser for Jira tickets and todos (lazygit-style)
- **LLM Integration (Phase 2)** - Additional AI features:
  - Auto-generate titles from brief descriptions
  - Improve/rename existing todos
  - Expand descriptions with context
  - Suggest criticality levels and completion order
- **Jira Push/Sync** - Create Jira tickets from todos and sync status
- **Priority/Criticality System** - Add priority levels for todos
- **Monday.com integration** - Similar integration to Jira
- **github.com integration** - Similar integration to Jira

## Jira Integration Setup

Jira integration is fully functional with simple API token authentication and AI-powered ticket breakdown!

### Current Status
✅ **Completed:**
- API token authentication (no OAuth app creation needed!)
- Secure credential storage in OS keyring
- Direct Jira URL support (works with any Jira instance)
- **Project-specific Jira URLs** - Different projects can use different Jira instances
- Smart fallback system - Global default with per-project overrides
- Automatic URL protocol handling (accepts URLs with or without `https://`)
- Database schema for Jira tickets
- REST API client for fetching individual tickets
- ADF (Atlassian Document Format) description parsing (supports bullet lists, ordered lists, etc.)
- `jira-pull` command - Fetch ticket and create linked todo
- `jira-pull-claude` command - AI-powered ticket breakdown using Claude Sonnet
- `jira-set-default-url` command - Update global default Jira URL
- `project-set-jira-url` command - Set project-specific Jira URL

🚧 **Coming Soon:**
- `toto jira-push -i <todo-id>` - Create Jira ticket from todo
- `toto jira-sync` - Sync status between todos and Jira

### Setup

1. **Create a Jira API token:**
   - Go to: https://id.atlassian.com/manage-profile/security/api-tokens
   - Click "Create API token"
   - Give it a name (e.g., "Toto CLI")
   - Copy the token immediately (you won't see it again!)

2. **Authenticate with Jira:**
   ```bash
   toto jira-auth
   ```
   You'll be prompted to enter:
   - Your Jira URL (e.g., `https://mycompany.atlassian.net`)
   - Your email address
   - Your API token

   Credentials are validated with a test API call and securely stored in your OS keyring.

3. **Pull a Jira ticket:**
   ```bash
   toto jira-pull -i PROJ-123
   ```
   Fetches the ticket from Jira, saves it to the database, and creates a linked todo.

4. **Pull a Jira ticket with AI breakdown:**
   ```bash
   toto jira-pull-claude -i PROJ-123
   ```
   Uses Claude AI to intelligently break down the ticket into multiple subtasks.

   **Note:** Requires `CLAUDE_API_KEY` environment variable. Export it in your shell:
   ```bash
   export CLAUDE_API_KEY=your-claude-api-key
   ```

## Prerequisites

- Go 1.21 or higher
- Git (for development)
- SQLite3 (automatically handled by Go modules)
- **For Jira integration:** Jira account with API access
- **For AI features:** Anthropic API key (get one at https://console.anthropic.com/)

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## Author

Sam Hodgkinson- [odgy8](https://github.com/odgy8)

## Contributors 

- You could be the first!

## License

This project is licensed under the MIT License - see the LICENSE file for details.
