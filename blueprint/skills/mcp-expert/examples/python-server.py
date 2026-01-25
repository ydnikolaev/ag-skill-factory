#!/usr/bin/env python3
"""
Full-featured MCP server example with tools, resources, and prompts.
Run with: python server.py
"""
from mcp.server.fastmcp import FastMCP
import os
import json

mcp = FastMCP("MyServer", json_response=True)


# ============================================================
# TOOLS - Execute actions with side effects
# ============================================================

@mcp.tool()
def read_file(path: str) -> str:
    """Read contents of a file at the given path.
    
    Args:
        path: Absolute path to the file to read
    """
    if not os.path.isabs(path):
        return f"Error: Path must be absolute. Got: {path}"
    try:
        with open(path, 'r') as f:
            return f.read()
    except FileNotFoundError:
        return f"Error: File not found: {path}"
    except Exception as e:
        return f"Error reading file: {e}"


@mcp.tool()
def list_directory(path: str) -> str:
    """List contents of a directory.
    
    Args:
        path: Absolute path to the directory
    """
    try:
        entries = os.listdir(path)
        return json.dumps(entries, indent=2)
    except Exception as e:
        return f"Error: {e}"


# ============================================================
# RESOURCES - Expose read-only data via URI templates
# ============================================================

@mcp.resource("env://{name}")
def get_env_var(name: str) -> str:
    """Get value of an environment variable."""
    return os.environ.get(name, f"Environment variable {name} not set")


@mcp.resource("config://app")
def get_app_config() -> str:
    """Get application configuration as JSON."""
    return json.dumps({
        "version": "1.0.0",
        "environment": os.environ.get("ENV", "development"),
        "debug": os.environ.get("DEBUG", "false").lower() == "true"
    }, indent=2)


# ============================================================
# PROMPTS - Reusable prompt templates
# ============================================================

@mcp.prompt()
def code_review(code: str, language: str = "python") -> str:
    """Generate a code review prompt for the given code."""
    return f"""Please review the following {language} code:

```{language}
{code}
```

Consider:
1. Code quality and readability
2. Potential bugs or edge cases
3. Performance implications
4. Security concerns
5. Suggestions for improvement"""


@mcp.prompt()
def explain_error(error: str, context: str = "") -> str:
    """Generate a prompt to explain an error message."""
    prompt = f"Please explain this error and how to fix it:\n\n```\n{error}\n```"
    if context:
        prompt += f"\n\nContext:\n{context}"
    return prompt


# ============================================================
# MAIN
# ============================================================

if __name__ == "__main__":
    import sys
    print("Starting MCP server...", file=sys.stderr)
    mcp.run()  # Uses stdio transport by default
