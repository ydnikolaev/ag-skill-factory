# Project-Bro: Decision Flow

How project-bro processes user questions:

```mermaid
graph TD
    A[User asks about project] --> B{What type of question?}
    B -->|Status| C[Read ARTIFACT_REGISTRY.md]
    B -->|Progress| D[Compare roadmap vs artifacts]
    B -->|Technical| E[Analyze code structure]
    B -->|Next steps| F[Check pipeline phase]
    C --> G[Summarize state]
    D --> G
    E --> G
    F --> G
```

## Question Types

| Question Pattern | Action |
|-----------------|--------|
| "where are we?" / "status" | Read `docs/ARTIFACT_REGISTRY.md` |
| "what's left?" / "progress" | Compare roadmap vs artifacts |
| "show code" / "what's in backend?" | Analyze codebase structure |
| "what's next?" / "next step" | Check pipeline phase, recommend skill |
