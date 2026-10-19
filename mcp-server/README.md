# LBMesh MCP Server

Expose LBMesh container orchestration tools directly to your local AI clients (Cursor, Claude Desktop, Claude Code).

> **Status:** The server skeleton exists, but all tools are stubs that return "not implemented yet". This README documents the intended interface.

## Installation

> **Note:** Every snippet below launches the server with `lbmesh mcp start`, which ships in the main `lbmesh` CLI from release v1.0.0. Run `lbmesh mcp --help` to check that your installed `lbmesh` has it. If it does not, upgrade `lbmesh` or use the standalone binary described in [Building the standalone server](#building-the-standalone-server).

### Claude Desktop

Add the following configuration block to your Claude Desktop configuration file:
*   **macOS**: `~/Library/Application Support/Claude/claude_desktop_config.json`
*   **Windows**: `%APPDATA%\Claude\claude_desktop_config.json`

```json
{
  "mcpServers": {
    "lbmesh": {
      "command": "lbmesh",
      "args": ["mcp", "start"]
    }
  }
}
```

### Claude Code

Register the server with the `claude` CLI:

```bash
claude mcp add lbmesh -- lbmesh mcp start
```

Use `--scope` (`-s`) to choose where the registration is stored: `local` (the default, private to you in the current project), `project` (shared with the team through a `.mcp.json` file in the project root), or `user` (available to you in every project).

Confirm the registration:

```bash
claude mcp list
```

### Cursor

Add the same `mcpServers` block to a Cursor `mcp.json` file:
*   **Project**: `.cursor/mcp.json` in your project root
*   **Global**: `~/.cursor/mcp.json`

```json
{
  "mcpServers": {
    "lbmesh": {
      "command": "lbmesh",
      "args": ["mcp", "start"]
    }
  }
}
```

## Building the standalone server

The Go server in this directory builds as a standalone binary. It speaks MCP over stdio and takes no arguments. Build it to a path of your choice:

```bash
go build -o /usr/local/bin/lbmesh-mcp .
```

Then point your client's `command` at the binary and omit `args`, for example in Claude Desktop or Cursor:

```json
{
  "mcpServers": {
    "lbmesh": {
      "command": "/usr/local/bin/lbmesh-mcp"
    }
  }
}
```

## Available AI Tools

- `lbmesh_status`: View running services, scopes, and active port bindings.
- `lbmesh_up`: Spin up any recipe (e.g., PostgreSQL, Redis, custom taps) with automatic port broker remapping.
- `lbmesh_down`: Teardown active container services and prune bridges.
- `lbmesh_list_recipes`: View available Cascading Recipe Registry templates.

## Development

The server speaks MCP over stdio (JSON-RPC on stdin/stdout). From this directory:

```bash
go build ./...
go test ./...
```
