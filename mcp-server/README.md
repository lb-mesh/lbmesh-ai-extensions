# LBMesh MCP Server

Expose LBMesh container orchestration tools directly to your local AI clients (Cursor, Claude Desktop, Claude Code).

> **Status:** The server implementation is not yet released. This README documents the intended interface.

## Installation

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

## Available AI Tools

- `lbmesh_status`: View running services, scopes, and active port bindings.
- `lbmesh_up`: Spin up any recipe (e.g., PostgreSQL, Redis, custom taps) with automatic port broker remapping.
- `lbmesh_down`: Teardown active container services and prune bridges.
- `lbmesh_list_recipes`: View available Cascading Recipe Registry templates.
