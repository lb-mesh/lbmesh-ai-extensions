package main

import (
	"context"
	"encoding/json"
	"slices"
	"testing"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

var wantTools = map[string]string{
	"lbmesh_status":       "View running services, scopes, and active port bindings.",
	"lbmesh_up":           "Spin up any recipe (e.g., PostgreSQL, Redis, custom taps) with automatic port broker remapping.",
	"lbmesh_down":         "Teardown active container services and prune bridges.",
	"lbmesh_list_recipes": "View available Cascading Recipe Registry templates.",
}

// connect starts the server on an in-memory transport and returns a
// connected client session.
func connect(t *testing.T) *mcp.ClientSession {
	t.Helper()
	ctx := context.Background()
	serverTransport, clientTransport := mcp.NewInMemoryTransports()

	ss, err := newServer().Connect(ctx, serverTransport, nil)
	if err != nil {
		t.Fatalf("server connect: %v", err)
	}
	t.Cleanup(func() { ss.Close() })

	client := mcp.NewClient(&mcp.Implementation{Name: "test-client", Version: "0.0.0"}, nil)
	cs, err := client.Connect(ctx, clientTransport, nil)
	if err != nil {
		t.Fatalf("client connect: %v", err)
	}
	t.Cleanup(func() { cs.Close() })
	return cs
}

func TestInitialize(t *testing.T) {
	cs := connect(t)
	res := cs.InitializeResult()
	if res == nil || res.ServerInfo == nil {
		t.Fatal("no initialize result")
	}
	if res.ServerInfo.Name != "lbmesh" || res.ServerInfo.Version != "0.0.0" {
		t.Errorf("server info = %+v, want lbmesh 0.0.0", res.ServerInfo)
	}
	if res.Capabilities == nil || res.Capabilities.Tools == nil {
		t.Error("server does not advertise the tools capability")
	}
}

func TestListTools(t *testing.T) {
	cs := connect(t)
	res, err := cs.ListTools(context.Background(), nil)
	if err != nil {
		t.Fatalf("ListTools: %v", err)
	}
	if len(res.Tools) != len(wantTools) {
		t.Fatalf("got %d tools, want %d", len(res.Tools), len(wantTools))
	}
	for _, tool := range res.Tools {
		want, ok := wantTools[tool.Name]
		if !ok {
			t.Errorf("unexpected tool %q", tool.Name)
			continue
		}
		if tool.Description != want {
			t.Errorf("%s description = %q, want %q", tool.Name, tool.Description, want)
		}
	}

	required := map[string][]string{}
	for _, tool := range res.Tools {
		var schema struct {
			Type     string   `json:"type"`
			Required []string `json:"required"`
		}
		b, err := json.Marshal(tool.InputSchema)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(b, &schema); err != nil {
			t.Fatal(err)
		}
		if schema.Type != "object" {
			t.Errorf("%s input schema type = %q, want object", tool.Name, schema.Type)
		}
		required[tool.Name] = schema.Required
	}
	if !slices.Contains(required["lbmesh_up"], "recipe") {
		t.Errorf("lbmesh_up required = %v, want recipe", required["lbmesh_up"])
	}
	if slices.Contains(required["lbmesh_down"], "recipe") {
		t.Error("lbmesh_down marks recipe as required; it should be optional")
	}
}

func TestStubCalls(t *testing.T) {
	cs := connect(t)
	args := map[string]map[string]any{
		"lbmesh_status":       {},
		"lbmesh_up":           {"recipe": "postgres"},
		"lbmesh_down":         {},
		"lbmesh_list_recipes": {},
	}
	for name := range wantTools {
		t.Run(name, func(t *testing.T) {
			res, err := cs.CallTool(context.Background(), &mcp.CallToolParams{Name: name, Arguments: args[name]})
			if err != nil {
				t.Fatalf("CallTool returned protocol error: %v", err)
			}
			if !res.IsError {
				t.Error("isError = false, want true")
			}
			if len(res.Content) != 1 {
				t.Fatalf("got %d content items, want 1", len(res.Content))
			}
			text, ok := res.Content[0].(*mcp.TextContent)
			if !ok {
				t.Fatalf("content is %T, want *mcp.TextContent", res.Content[0])
			}
			if want := name + " is not implemented yet"; text.Text != want {
				t.Errorf("text = %q, want %q", text.Text, want)
			}
		})
	}
}

func TestUpRequiresRecipe(t *testing.T) {
	cs := connect(t)
	res, err := cs.CallTool(context.Background(), &mcp.CallToolParams{Name: "lbmesh_up", Arguments: map[string]any{}})
	if err == nil && !res.IsError {
		t.Fatal("lbmesh_up without recipe succeeded")
	}
	// The failure must come from validation, not from the stub handler.
	if err == nil {
		for _, c := range res.Content {
			if text, ok := c.(*mcp.TextContent); ok && text.Text == "lbmesh_up is not implemented yet" {
				t.Fatal("lbmesh_up without recipe reached the handler; want a validation failure")
			}
		}
	}
}
