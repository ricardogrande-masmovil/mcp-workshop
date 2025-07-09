# Workshop Navigation Tool

This repository includes a workshop navigation tool that helps you move through different chapters of the MCP workshop.

## Installation

You can install the workshop tool globally using Go:

```bash
go install github.com/ricardogrande-masmovil/mcp-workshop/cmd/workshop@latest
```

Or build it locally:

```bash
go build -o workshop ./cmd/workshop
```

## Usage

Once installed, you can use the `workshop` command from anywhere in the repository:

```bash
# Move to the next chapter
workshop next

# Show current status and progress
workshop status

# Jump to a specific chapter
workshop jump 002-mcp-weather-definition

# Reset to main branch (start over)
workshop reset

# Show help
workshop help
```

## Available Commands

- `workshop next` - Move to the next chapter in the workshop sequence
- `workshop status` - Show current branch and progress through the workshop
- `workshop jump <chapter-name>` - Jump directly to a specific chapter
- `workshop reset` - Reset to main branch and start over from the beginning
- `workshop help` - Display help information

## Workshop Chapters

1. `main`
2. `001-initial-setup`
3. `002-mcp-weather-definition`
4. `002-mcp-weather-complete`
5. `003-mcp-zoo`
6. `003-mcp-zoo-complete`

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
