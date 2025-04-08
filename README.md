# FlyMCP

A simple MCP server that wraps the `flyctl` CLI to provide Fly.io capabilities to Claude Desktop.

## Prerequisites

- Go 1.21 or later
- `flyctl` CLI installed and available in PATH
- Claude Desktop installed

## Installation

```bash
# Clone the repository
git clone https://github.com/superfly/flymcp.git
cd flymcp

# Install dependencies
go mod download

# Build the binary
go build -o flymcp
```

## Configuration

### Claude Desktop Setup

1. Open Claude Desktop
2. Go to the Claude menu and select "Settings..."
3. Click on "Developer" in the left-hand bar
4. Click on "Edit Config"

This will create or open a configuration file at:
- macOS: `~/Library/Application Support/Claude/claude_desktop_config.json`
- Windows: `%APPDATA%\Claude\claude_desktop_config.json`

5. Add the following configuration to the file:

```json
{
  "mcpServers": {
    "flymcp": {
      "command": "/full/path/to/your/flymcp",
      "args": []
    }
  }
}
```

Make sure to replace `/full/path/to/your/flymcp` with the absolute path to your flymcp binary.

6. Save the file and restart Claude Desktop

### Authentication

Before using the tools, make sure you're authenticated with Fly.io:

```bash
flyctl auth login
```

## Troubleshooting

If you encounter issues with the MCP server:

1. Check the logs in:
   - macOS: `~/Library/Logs/Claude/mcp-server-flymcp.log`
   - Windows: `%APPDATA%\Claude\logs\mcp-server-flymcp.log`

2. Make sure the path to the flymcp binary is correct and absolute
3. Ensure you have the necessary permissions to execute the binary
4. Verify that flyctl is installed and accessible in your PATH

## Development

The server is built using the [MCP Go library](https://github.com/mark3labs/mcp-go) and communicates with Claude Desktop through standard input/output streams.

## License

MIT 