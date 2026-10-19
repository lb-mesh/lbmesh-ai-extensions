package main

import (
	"context"
	"fmt"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type noArgs struct{}

type upArgs struct {
	Recipe string `json:"recipe" jsonschema:"name of the recipe to start"`
}

type downArgs struct {
	Recipe string `json:"recipe,omitempty" jsonschema:"name of the recipe to stop; omit to stop everything"`
}

// addTools registers the LBMesh tools. All of them are stubs for now.
func addTools(server *mcp.Server) {
	mcp.AddTool(server, &mcp.Tool{
		Name:        "lbmesh_status",
		Description: "View running services, scopes, and active port bindings.",
	}, notImplemented[noArgs]("lbmesh_status"))
	mcp.AddTool(server, &mcp.Tool{
		Name:        "lbmesh_up",
		Description: "Spin up any recipe (e.g., PostgreSQL, Redis, custom taps) with automatic port broker remapping.",
	}, notImplemented[upArgs]("lbmesh_up"))
	mcp.AddTool(server, &mcp.Tool{
		Name:        "lbmesh_down",
		Description: "Teardown active container services and prune bridges.",
	}, notImplemented[downArgs]("lbmesh_down"))
	mcp.AddTool(server, &mcp.Tool{
		Name:        "lbmesh_list_recipes",
		Description: "View available Cascading Recipe Registry templates.",
	}, notImplemented[noArgs]("lbmesh_list_recipes"))
}

// notImplemented returns a handler whose error the SDK reports as a tool
// result with isError set, not as a JSON-RPC protocol error.
func notImplemented[In any](name string) mcp.ToolHandlerFor[In, any] {
	return func(context.Context, *mcp.CallToolRequest, In) (*mcp.CallToolResult, any, error) {
		return nil, nil, fmt.Errorf("%s is not implemented yet", name)
	}
}
