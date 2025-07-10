# Workshop Navigation Tool

This repository includes a workshop navigation tool that helps you move through different chapters of the MCP workshop.

## Installation

You can install the workshop tool globally using Go:

```bash
go install github.com/ricardogrande-masmovil/mcp-workshop/cmd/workshop@latest
```

### Optional: Install as "ws" alias

If you prefer a shorter command name, you can install it as "ws":

```bash
# Install with custom binary name
go build -o ws ./cmd/workshop
# Move to your PATH (example for macOS/Linux)
sudo mv ws /usr/local/bin/
```

Or build it locally:

```bash
go build -o workshop ./cmd/workshop
```

## Usage

Once installed, you can use the `workshop` command (or `ws` if you installed the alias) from anywhere in the repository:

```bash
# Move to the next chapter
workshop next
# or with the ws alias:
ws next

# Show current status and progress
workshop status
# or: ws status

# Jump to a specific chapter
workshop jump mcp-weather-definition
# or: ws jump mcp-weather-definition

# Reset to main branch (start over)
workshop reset
# or: ws reset

# Show help
workshop help
# or: ws help
```

## Available Commands

The following commands are available (replace `workshop` with `ws` if you installed the alias):

- `workshop next` - Move to the next chapter in the workshop sequence
- `workshop status` - Show current branch and progress through the workshop
- `workshop jump <chapter-name>` - Jump directly to a specific chapter
- `workshop reset` - Reset to main branch and start over from the beginning
- `workshop help` - Display help information

## Workshop Chapters

1. `main`
2. `initial-setup`
3. `mcp-weather-definition`
4. `mcp-weather-complete`
5. `mcp-zoo`
6. `mcp-zoo-complete`

## Important Notes

- The tool will discard all uncommitted changes when switching chapters
- Make sure you're in a Git repository when using the tool
- The tool creates new branches if they don't exist when switching chapters

## Migration from Shell Script

If you were previously using `./workshop.sh`, you can now use the Go binary instead:

- `./workshop.sh next` → `workshop next`
- `./workshop.sh status` → `workshop status`
- `./workshop.sh jump <chapter>` → `workshop jump <chapter>`
- `./workshop.sh reset` → `workshop reset`
- `./workshop.sh help` → `workshop help`
