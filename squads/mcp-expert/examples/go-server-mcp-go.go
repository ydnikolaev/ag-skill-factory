// Full-featured MCP server example using mark3labs/mcp-go
// This is the most popular Go MCP library with excellent DX.
//
// Install: go get github.com/mark3labs/mcp-go
// Run: go run server.go

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func main() {
	// Create server with capabilities
	s := server.NewMCPServer(
		"My MCP Server",
		"1.0.0",
		server.WithToolCapabilities(true),
		server.WithResourceCapabilities(false, true),
	)

	// ============================================================
	// TOOLS - Execute actions with side effects
	// ============================================================

	// Tool: Read file
	s.AddTool(
		mcp.NewTool("read_file",
			mcp.WithDescription("Read contents of a file"),
			mcp.WithString("path",
				mcp.Required(),
				mcp.Description("Absolute path to the file"),
			),
		),
		handleReadFile,
	)

	// Tool: List directory
	s.AddTool(
		mcp.NewTool("list_dir",
			mcp.WithDescription("List contents of a directory"),
			mcp.WithString("path",
				mcp.Required(),
				mcp.Description("Absolute path to the directory"),
			),
		),
		handleListDir,
	)

	// Tool: Get current time
	s.AddTool(
		mcp.NewTool("get_time",
			mcp.WithDescription("Get the current time"),
			mcp.WithString("format",
				mcp.Description("Time format (RFC3339, Unix)"),
				mcp.DefaultString("RFC3339"),
			),
		),
		handleGetTime,
	)

	// ============================================================
	// RESOURCES - Expose read-only data
	// ============================================================

	s.AddResource(
		mcp.NewResource(
			"config://app",
			"Application Configuration",
			mcp.WithResourceDescription("Current application configuration"),
			mcp.WithMIMEType("application/json"),
		),
		handleConfig,
	)

	// Start the server (stdio transport)
	if err := server.ServeStdio(s); err != nil {
		fmt.Fprintf(os.Stderr, "Server error: %v\n", err)
		os.Exit(1)
	}
}

// ============================================================
// HANDLERS
// ============================================================

func handleReadFile(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	path, err := req.RequireString("path")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	if !filepath.IsAbs(path) {
		return mcp.NewToolResultError("path must be absolute"), nil
	}

	content, err := os.ReadFile(path)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("failed to read file: %v", err)), nil
	}

	return mcp.NewToolResultText(string(content)), nil
}

func handleListDir(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	path, err := req.RequireString("path")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	entries, err := os.ReadDir(path)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("failed to list directory: %v", err)), nil
	}

	var names []string
	for _, e := range entries {
		name := e.Name()
		if e.IsDir() {
			name += "/"
		}
		names = append(names, name)
	}

	result, _ := json.MarshalIndent(names, "", "  ")
	return mcp.NewToolResultText(string(result)), nil
}

func handleGetTime(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	format := req.GetString("format", "RFC3339")

	var result string
	switch format {
	case "Unix":
		result = fmt.Sprintf("%d", time.Now().Unix())
	default:
		result = time.Now().Format(time.RFC3339)
	}

	return mcp.NewToolResultText(result), nil
}

func handleConfig(ctx context.Context, req mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
	config := map[string]any{
		"version":     "1.0.0",
		"environment": os.Getenv("ENV"),
		"debug":       os.Getenv("DEBUG") == "true",
	}

	data, _ := json.MarshalIndent(config, "", "  ")
	return []mcp.ResourceContents{
		mcp.TextResourceContents{
			URI:      "config://app",
			MIMEType: "application/json",
			Text:     string(data),
		},
	}, nil
}
