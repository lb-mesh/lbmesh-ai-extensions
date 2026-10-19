# LBMesh AI Extensions

AI-assistant integrations that let Cursor, Claude Code, GitHub Copilot, and Model Context Protocol (MCP) clients drive LBMesh local container orchestration.

## Integrations

| Integration | Directory | Status | Guide |
| :--- | :--- | :--- | :--- |
| MCP Server | [`mcp-server/`](mcp-server/) | Planned | [mcp-server/README.md](mcp-server/README.md) |
| Cursor | [`cursor/`](cursor/) | Rules available | [cursor/README.md](cursor/README.md) |
| GitHub Copilot | [`copilot/`](copilot/) | Planned | [copilot/README.md](copilot/README.md) |
| Claude Code | [`claude-code/`](claude-code/) | Instructions available | [claude-code/system-instructions.md](claude-code/system-instructions.md) |

## Repository layout

```
lbmesh-ai-extensions/
├── LICENSE
├── README.md
├── mcp-server/
│   └── README.md
├── cursor/
│   ├── .cursorrules
│   └── README.md
├── copilot/
│   ├── manifest.json
│   └── README.md
└── claude-code/
    └── system-instructions.md
```

## License

Licensed under the [Apache License 2.0](LICENSE).
