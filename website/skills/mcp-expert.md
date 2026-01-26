# mcp-expert

> Expert on Model Context Protocol (MCP) servers. Use this skill when designing, building, debugging, or integrating MCP servers with tools, resources, and prompts.

**Version:** 1.2.0

---


# MCP Expert

Expert-level guidance for building MCP servers. **Primary language: Go** (official SDK). Also covers Python and TypeScript.

## When to Use This Skill

- **Trigger**: Create an MCP server or add tools/resources
- **Trigger**: Integrate a service via MCP (databases, APIs)
- **Trigger**: Debug MCP connection or transport issues
- **Anti-pattern**: Do NOT use for general API development without MCP context

## Decision Tree

1.  **IF** creating a new MCP server:
    - **Use Go** (primary) — see `examples/go-server-official.go`
    - Library: `modelcontextprotocol/go-sdk` (official, recommended)
    - Transport: `stdio` (IDE) or HTTP (web services)

2.  **IF** adding to existing Go server:
    - Use `s.AddTool()` with handler function
    - Use `s.AddResource()` for read-only data

3.  **IF** debugging:
    - Check `mcp_config.json` paths (must be absolute)
    - Use Inspector: `npx @anthropic/mcp-inspector`
    - Logs go to stderr (stdout reserved for JSON-RPC)

## MCP Primitives

| Primitive | Purpose | Example |
|-----------|---------|---------|
| **Tool** | Execute actions with side effects | Run builds, create issues |
| **Resource** | Expose read-only data | DB schemas, config files |
| **Prompt** | Reusable prompt templates | Code review patterns |

## Go Quick Reference (Official SDK)

```go
server := mcp.NewServer(&mcp.Implementation{
    Name: "my-server", Version: "v1.0.0",
}, nil)

mcp.AddTool(server, &mcp.Tool{
    Name: "my_tool",
    Description: "Brief description",
}, HandleMyTool)

server.Run(ctx, &mcp.StdioTransport{})
```
**Full example**: See `examples/go-server-official.go`

## Go Libraries

| Library | Install | Notes |
|---------|---------|-------|
| **Official SDK** | `go get github.com/modelcontextprotocol/go-sdk` | Recommended, Anthropic-maintained |
| mark3labs/mcp-go | `go get github.com/mark3labs/mcp-go` | Alternative, good DX |

## Antigravity Config (`mcp_config.json`)

```json
{
  "mcpServers": {
    "my-server": {
      "command": "/absolute/path/to/server",
      "args": [],
      "env": { "API_KEY": "your-key" }
    }
  }
}
```

**Key rules:**
1. **Always use absolute paths**
2. **Logs to stderr** — stdout is for JSON-RPC only
3. **Descriptions matter** — LLM uses them for tool selection

## Debugging Workflow

1. Check `mcp_config.json` paths
2. Test standalone: `go run server.go`
3. Use Inspector: `npx @anthropic/mcp-inspector`
4. Check stderr in terminal

## Best Practices

### Tool Design
- Use `snake_case` names: `get_user`, `create_issue`
- Write rich descriptions
- Validate inputs, return informative errors

### Security
- Never log secrets to stdout/stderr
- Use environment variables for API keys

## TDD Protocol (Hard Stop)

> [!CAUTION]
> **NO CODE WITHOUT FAILING TEST.**
> - **Tools**: Write a test that calls the tool with mock input -> Assert output.
> - **Resources**: Write a test that reads the resource URI -> Assert content.
>
> **Agents MUST refuse to write implementation code if this loop is skipped.**

## Tech Debt Protocol (Hard Stop)

> [!CAUTION]
> **Follow `../standards/TECH_DEBT_PROTOCOL.md`.**
> When creating workarounds:
> 1. Add `// TODO(TD-XXX): description` in code
> 2. Register in `project/docs/TECH_DEBT.md`
>
> **Forbidden:** Untracked TODOs, undocumented hardcoded values.

## Git Protocol (Hard Stop)

> [!CAUTION]
> **Follow `../standards/GIT_PROTOCOL.md`.**
> 1. **Branch**: Work in `feat/<name>` or `fix/<name>`.
> 2. **Commit**: Use Conventional Commits (`feat:`, `fix:`).
> 3. **Atomic**: One commit = One logical change.
>
> **Reject**: "wip", "update", "fix" as commit messages.

## Language Requirements

> All skill files must be in English. See [LANGUAGE.md](file://blueprint/rules/LANGUAGE.md).

## Team Collaboration

- **Backend**: `@backend-go-expert` (Integrates MCP into Go services)
- **DevOps**: `@devops-sre` (MCP server deployment, systemd units)
- **CLI**: `@cli-architect` (MCP tools for CLI applications)
- **TMA**: `@tma-expert` (MCP integration with Telegram Mini Apps)

## When to Delegate

- ✅ **Delegate to `@backend-go-expert`** when: MCP server needs integration into larger Go service
- ✅ **Delegate to `@devops-sre`** when: MCP server ready for deployment

- ⬅️ **Return to `@systems-analyst`** if: Requirements unclear for tool design


## Iteration Protocol (Ephemeral → Persistent)

> [!IMPORTANT]
> **Phase 1: Draft in Brain** — Create MCP Server Config as artifact. Iterate via `notify_user`.
> **Phase 2: Persist on Approval** — ONLY after "Looks good" → write to `project/docs/active/mcp/`

## Document Lifecycle

> **Protocol**: [`DOCUMENT_STRUCTURE_PROTOCOL.md`](../standards/DOCUMENT_STRUCTURE_PROTOCOL.md)

| Operation | Document | Location | Trigger |
|-----------|----------|----------|---------|
| 🔵 Creates | server-config.md | `active/mcp/` | MCP server design complete |
| 📖 Reads | api-contracts.yaml | `active/architecture/` | On activation |
| 📝 Updates | ARTIFACT_REGISTRY.md | `project/docs/` | On create, on complete |
| 🟡 To Review | server-config.md | `review/mcp/` | Ready for implementation |
| ✅ Archive | — | `closed/<work-unit>/` | @doc-janitor on final approval |

## Pre-Handoff Validation (Hard Stop)

> [!CAUTION]
> **MANDATORY self-check before `notify_user` or delegation.**

| # | Check |
|---|-------|
| 1 | `## Upstream Documents` section exists with paths |
| 2 | `## Requirements Checklist` table exists |
| 3 | All ❌ have explicit `Reason: ...` |
| 4 | Document in `review/` folder |
| 5 | `ARTIFACT_REGISTRY.md` updated |

**If ANY unchecked → DO NOT PROCEED.**

## Handoff Protocol


> [!CAUTION]
> **BEFORE handoff:**
> 1. Save final document to `project/docs/` path
> 2. Change file status from `Draft` to `Approved` in header/frontmatter
> 3. Update `project/docs/ARTIFACT_REGISTRY.md` status to ✅ Done
> 4. Use `notify_user` for final approval
> 5. THEN delegate to next skill

## Examples

| File | Description |
|------|-------------|
| `examples/go-server-official.go` | Go server (official SDK) — recommended |
| `examples/go-server-mcp-go.go` | Go server (mark3labs/mcp-go) — alternative |
| `examples/python-server.py` | Python server (FastMCP) |
| `examples/typescript-server.ts` | TypeScript server |
| `examples/mcp_config.json` | Antigravity config |

## Resources

- **Docs**: [modelcontextprotocol.io](https://modelcontextprotocol.io)
- **Go (mcp-go)**: [github.com/mark3labs/mcp-go](https://github.com/mark3labs/mcp-go)
- **Go (official)**: [github.com/modelcontextprotocol/go-sdk](https://github.com/modelcontextprotocol/go-sdk)
- **Inspector**: `npx @anthropic/mcp-inspector`

> [!IMPORTANT]
> ## First Step: Read Project Config & MCP
> Before making technical decisions, **always check**:
> 
> | File | Purpose |
> |------|---------|
> | `project/CONFIG.yaml` | Stack versions, modules, architecture |
> | `mcp.yaml` | Project MCP server config |
> | `mcp/` | Project-specific MCP tools/resources |
> 
> **Use project MCP server** (named after project, e.g. `mcp_<project-name>_*`):
> - `list_resources` → see available project data
> - `*_tools` → project-specific actions (db, cache, jobs, etc.)
> 
> **Use `mcp_context7`** for library docs:
> - Check `mcp.yaml → context7.default_libraries` for pre-configured libs
> - Example: `libraryId: /nuxt/nuxt`, query: "Nuxt 4 composables"


