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
2. Go to Settings > Tool Providers
3. Click "Add Tool Provider"
4. Configure the tool provider:
   - Name: `FlyMCP`
   - Command: `/full/path/to/your/flymcp` (e.g., `/Users/username/code/flymcp/flymcp`)
   - Working Directory: (optional) The directory containing your `flymcp` binary

### Authentication

Before using the tools, make sure you're authenticated with Fly.io:

```bash
flyctl auth login
```

## Available Tools

### fly-logs
Returns logs for a Fly.io app or specific machine.

Parameters:
- `app` (string, required): The name of the app
- `machine` (string, optional): The specific machine ID

Example:
```json
{
  "name": "fly-logs",
  "arguments": {
    "app": "my-app",
    "machine": "123456"
  }
}
```

### fly-status
Returns the status of a Fly.io app in JSON format.

Parameters:
- `app` (string, required): The name of the app

Example:
```json
{
  "name": "fly-status",
  "arguments": {
    "app": "my-app"
  }
}
```

## Development

The server is built using the [MCP Go library](https://github.com/mark3labs/mcp-go) and communicates with Claude Desktop through standard input/output streams.

## License

MIT 