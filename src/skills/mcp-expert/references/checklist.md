# MCP Server Quality Checklist

## 🚨 Document Persistence (MANDATORY)

> [!CAUTION]
> **BEFORE handing off to another skill, you MUST:**

- [ ] **Final document exists in `docs/`** at the path defined in Artifact Ownership
- [ ] **ARTIFACT_REGISTRY.md updated** with status ✅ Done and Last Updated date
- [ ] **Artifact synced** — if you used an artifact, copy final content to `docs/`

**Why?** Artifacts don't persist between sessions. Without `docs/` file, the next skill cannot continue.

Use this checklist when building or reviewing an MCP server.

## 1. Structure & Config
- [ ] **Absolute paths**: Does `mcp_config.json` use absolute paths?
- [ ] **Environment variables**: Are secrets in env vars, not hardcoded?
- [ ] **Logging to stderr**: Is all logging going to stderr (not stdout)?

## 2. Tools
- [ ] **Naming**: Are tool names `snake_case`? (e.g., `get_user`, `create_file`)
- [ ] **Descriptions**: Does each tool have a clear description for LLM selection?
- [ ] **Input validation**: Are inputs validated before processing?
- [ ] **Error handling**: Do errors return informative messages (not panics)?
- [ ] **Typed I/O**: Are input/output structs using `jsonschema` tags?

## 3. Resources
- [ ] **URI scheme**: Are resource URIs using clear patterns? (e.g., `config://app`)
- [ ] **MIME types**: Are MIME types specified for resources?
- [ ] **Read-only**: Are resources truly read-only (no side effects)?

## 4. Transport
- [ ] **Stdio default**: Is stdio transport used for IDE integration?
- [ ] **Graceful shutdown**: Does the server handle context cancellation?

## 5. Go-Specific (Official SDK)
- [ ] **Import path**: Using `github.com/modelcontextprotocol/go-sdk/mcp`?
- [ ] **Handler signature**: Tool handlers return `(*CallToolResult, Output, error)`?
- [ ] **Struct tags**: Input structs use `json` and `jsonschema` tags?

## 6. Testing
- [ ] **Standalone run**: Can the server run with `go run server.go`?
- [ ] **MCP Inspector**: Tested with `npx @anthropic/mcp-inspector`?
- [ ] **Integration test**: Tested with actual Antigravity/IDE connection?

## 7. Security
- [ ] **No secrets in logs**: Are API keys/tokens excluded from stderr?
- [ ] **Input sanitization**: Are file paths and user inputs validated?
- [ ] **Minimal permissions**: Does the server request only needed capabilities?
