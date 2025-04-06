package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func main() {
	// Create MCP server
	srv := server.NewMCPServer(
		"FlyMCP 🚀",
		"1.0.0",
	)

	// Register logs capability
	srv.AddTool(mcp.NewTool("fly-logs",
		mcp.WithDescription("Get logs for a Fly.io app or specific machine"),
		mcp.WithString("app",
			mcp.Required(),
			mcp.Description("Name of the app"),
		),
		mcp.WithString("machine",
			mcp.Description("Specific machine ID"),
		),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		app, ok := request.Params.Arguments["app"].(string)
		if !ok {
			return nil, fmt.Errorf("app parameter is required and must be a string")
		}

		var machine string
		if machineVal, ok := request.Params.Arguments["machine"]; ok && machineVal != nil {
			if machine, ok = machineVal.(string); !ok {
				return nil, fmt.Errorf("machine parameter must be a string")
			}
		}

		args := []string{"logs", "--no-tail"}
		if app != "" {
			args = append(args, "-a", app)
		}
		if machine != "" {
			args = append(args, "--machine", machine)
		}

		fmt.Fprintf(os.Stderr, "Executing flyctl command: %v\n", args)
		cmd := exec.Command("flyctl", args...)
		output, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error executing flyctl: %v\nOutput: %s\n", err, string(output))
			return nil, fmt.Errorf("failed to get logs: %v\nOutput: %s", err, string(output))
		}

		return mcp.NewToolResultText(string(output)), nil
	})

	// Register status capability
	srv.AddTool(mcp.NewTool("fly-status",
		mcp.WithDescription("Get status of a Fly.io app"),
		mcp.WithString("app",
			mcp.Required(),
			mcp.Description("Name of the app"),
		),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		app, ok := request.Params.Arguments["app"].(string)
		if !ok {
			return nil, fmt.Errorf("app parameter is required and must be a string")
		}

		args := []string{"status", "--json"}
		if app != "" {
			args = append(args, "-a", app)
		}

		fmt.Fprintf(os.Stderr, "Executing flyctl command: %v\n", args)
		cmd := exec.Command("flyctl", args...)
		output, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error executing flyctl: %v\nOutput: %s\n", err, string(output))
			return nil, fmt.Errorf("failed to get status: %v\nOutput: %s", err, string(output))
		}

		return mcp.NewToolResultText(string(output)), nil
	})

	fmt.Fprintf(os.Stderr, "Starting MCP server...\n")
	if err := server.ServeStdio(srv); err != nil {
		fmt.Fprintf(os.Stderr, "Server error: %v\n", err)
	}
}
