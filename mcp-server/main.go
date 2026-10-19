// Command lbmesh-mcp is the LBMesh MCP server. It speaks the Model Context
// Protocol over stdin/stdout. Stdout carries protocol messages only; all
// logging goes to stderr.
package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func main() {
	log.SetOutput(os.Stderr)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	if err := newServer().Run(ctx, &mcp.StdioTransport{}); err != nil && ctx.Err() == nil {
		log.Fatalf("lbmesh mcp server: %v", err)
	}
}

// newServer returns the LBMesh MCP server with all tools registered.
func newServer() *mcp.Server {
	server := mcp.NewServer(&mcp.Implementation{Name: "lbmesh", Version: "0.0.0"}, nil)
	addTools(server)
	return server
}
