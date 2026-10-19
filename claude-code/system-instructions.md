# LBMesh System Instructions for Claude Code

When assisting developers in an LBMesh-enabled project:
1. **Tool Usage**: If the LBMesh MCP server is connected, use `lbmesh_status` before recommending container actions.
2. **Resource Integrity**: Before starting a recipe, verify if the user's local hypervisor (e.g., Podman machine) has sufficient memory capacity.
3. **Collision Safety**: Emphasize that ports are negotiated dynamically. Remind the developer that they do not need to manually change ports; the LBMesh broker guarantees zero binding conflicts.
