//go:build ignore

// MCP server example using the official modelcontextprotocol/go-sdk
// This is the official SDK maintained by Anthropic.
//
// Install: go get github.com/modelcontextprotocol/go-sdk
// Run: go run server.go

package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// Input/Output structs for tools (uses jsonschema tags for descriptions)
type ReadFileInput struct {
	Path string `json:"path" jsonschema:"required,description=Absolute path to the file"`
}

type ReadFileOutput struct {
	Content string `json:"content" jsonschema:"description=File contents"`
	Size    int    `json:"size" jsonschema:"description=File size in bytes"`
}

type ListDirInput struct {
	Path string `json:"path" jsonschema:"required,description=Absolute path to the directory"`
}

type ListDirOutput struct {
	Entries []string `json:"entries" jsonschema:"description=Directory entries"`
}

// Tool handlers
func ReadFile(ctx context.Context, req *mcp.CallToolRequest, input ReadFileInput) (
	*mcp.CallToolResult,
	ReadFileOutput,
	error,
) {
	content, err := os.ReadFile(input.Path)
	if err != nil {
		return mcp.NewError(fmt.Sprintf("failed to read file: %v", err)), ReadFileOutput{}, nil
	}

	return nil, ReadFileOutput{
		Content: string(content),
		Size:    len(content),
	}, nil
}

func ListDir(ctx context.Context, req *mcp.CallToolRequest, input ListDirInput) (
	*mcp.CallToolResult,
	ListDirOutput,
	error,
) {
	entries, err := os.ReadDir(input.Path)
	if err != nil {
		return mcp.NewError(fmt.Sprintf("failed to list directory: %v", err)), ListDirOutput{}, nil
	}

	var names []string
	for _, e := range entries {
		name := e.Name()
		if e.IsDir() {
			name += "/"
		}
		names = append(names, name)
	}

	return nil, ListDirOutput{Entries: names}, nil
}

func main() {
	// Create a server
	server := mcp.NewServer(&mcp.Implementation{
		Name:    "my-mcp-server",
		Version: "v1.0.0",
	}, nil)

	// Add tools
	mcp.AddTool(server, &mcp.Tool{
		Name:        "read_file",
		Description: "Read contents of a file",
	}, ReadFile)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "list_dir",
		Description: "List contents of a directory",
	}, ListDir)

	// Run the server over stdio
	if err := server.Run(context.Background(), &mcp.StdioTransport{}); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
