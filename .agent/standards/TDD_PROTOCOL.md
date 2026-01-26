# Antigravity TDD Protocol

> [!IMPORTANT]
> **This is a Hard Stop Protocol.**
> No implementation code is allowed without a failing test.
> No bug fix is allowed without a failing reproduction case.

## 1. The Golden Rule: Red-Green-Refactor

Agents must strictly follow the TDD loop. Breaking this loop is a violation of the Antigravity engineering standard.

### Phase 1: RED (The Hard Blocker)
- **Goal**: Demonstrate the missing feature or the bug.
- **Action**: Write a test case that fails.
- **Verification**: Run the test. **It MUST fail.**
- **Constraint**: You are NOT ALLOWED to write implementation code yet.

### Phase 2: GREEN (The Minimal Pass)
- **Goal**: Make the test pass.
- **Action**: Write the *minimum* amount of implementation code to satisfy the test.
- **Verification**: Run the test. **It MUST pass.**
- **Constraint**: Do not optimize yet. Do not over-engineer.

### Phase 3: REFACTOR (Cleanup)
- **Goal**: Improve code quality without changing behavior.
- **Action**: Clean up credentials, variable names, duplication.
- **Verification**: Run the test again. **It MUST still pass.**

## 2. Debugging Protocol: "Reproduction First"

When debugging an issue, agents are forbidden from "guessing" the fix.

1.  **Stop**: Do not touch the implementation code.
2.  **Reproduction**: Create a standalone test case that reproduces the bug.
3.  **Confirm**: Run the test. It should fail (Red).
4.  **Fix**: Apply the fix.
5.  **Verify**: Run the test. It should pass (Green).

## 3. Planning Protocol

Planning skills (`product-analyst`, `bmad-architect`, `tech-spec-writer`) must include a **Verification Strategy** in their outputs.

- **Product Analyst**: Must define "Acceptance Criteria" (What does 'Done' look like?).
- **Architect**: Must define "Test Boundaries" (What to mock? What to integration test?).
- **Tech Spec Writer**: Must define "Test Cases" (What to test? How to test?).

## 4. Interaction with User

If an agent is asked to "just fix it quickly" without tests:
- ** Polite Refusal**: "I must follow the TDD Protocol to ensure stability. I will write a reproduction test first."
- **Exception**: Only if the user explicitly overrides with "SKIP TDD" (but the agent should warn about risks).
