# Feature Specification Template

Use this template when writing technical specifications.

## Feature: [Feature Name]

### Overview
Brief description of what this feature does.

### Behavior
Natural language description of what happens:
> "When user does X, system does Y, then returns Z."

### Edge Cases

| Input | Expected Output | Notes |
|-------|-----------------|-------|
| Empty input | 400 Bad Request | Validation error |
| Invalid format | 400 Bad Request | Include format hint |
| Not found | 404 Not Found | Log for monitoring |
| Server error | 500 Internal Error | Retry logic applies |

### Pseudocode

```text
FUNCTION feature_name(input):
  IF input is empty:
    RETURN error 400
  
  result = process(input)
  
  IF result not found:
    RETURN error 404
  
  RETURN result
```

### API Examples

**Request:**
```json
POST /api/v1/feature
{
  "field": "value"
}
```

**Response 200:**
```json
{
  "id": "abc123",
  "status": "success"
}
```

**Response 400:**
```json
{
  "error": "validation_failed",
  "details": "field is required"
}
```

### TDD Hints

- **Mock**: External service calls, database
- **Test first**: Happy path with valid input
- **Assert**: Response status, body structure, side effects

### Dependencies

- **Requires**: Authentication service, Database connection
- **Blocks**: Feature Y depends on this
