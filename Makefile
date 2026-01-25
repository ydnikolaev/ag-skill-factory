.PHONY: install uninstall build-factory install-factory install-completions generate-team validate validate-all validate-blueprint test lint clean check-loc changelog

# Paths
BLUEPRINT_DIR := $(shell pwd)/blueprint
FACTORY_SKILLS_DIR := $(shell pwd)/.agent/skills
VALIDATOR := $(FACTORY_SKILLS_DIR)/skill-creator/scripts/validate_skill.py
BLUEPRINT_VALIDATOR := $(FACTORY_SKILLS_DIR)/skill-creator/scripts/validate_blueprint.py
BIN_DIR := $(shell pwd)/bin
FACTORY_BIN := $(BIN_DIR)/factory
SYMLINK_PATH := /usr/local/bin/factory
SHELL_RC := $(HOME)/.zshrc

# Install the factory CLI (no global_skills anymore)
install: validate-all build-factory install-factory install-completions
	@echo ""
	@echo "🎉 Installation complete!"
	@echo "   Run 'factory install' in any project to get started."

# Validate a single skill: make validate SKILL=mcp-expert
validate:
	@if [ -z "$(SKILL)" ]; then \
		echo "Usage: make validate SKILL=<skill-name>"; \
		exit 1; \
	fi
	@python3 $(VALIDATOR) $(BLUEPRINT_DIR)/skills/$(SKILL)

# Validate all blueprint skills
validate-all:
	@echo "🔍 Validating all skills..."
	@failed=0; \
	for skill in $(BLUEPRINT_DIR)/skills/*/; do \
		if [ -f "$$skill/SKILL.md" ]; then \
			skill_name=$$(basename "$$skill"); \
			if ! python3 $(VALIDATOR) "$$skill" > /dev/null 2>&1; then \
				echo "❌ $$skill_name failed validation"; \
				failed=1; \
			else \
				echo "✅ $$skill_name"; \
			fi \
		fi \
	done; \
	if [ $$failed -eq 1 ]; then \
		echo ""; \
		echo "⚠️  Some skills failed validation. Run 'make validate SKILL=<name>' for details."; \
	else \
		echo "✅ All skills validated!"; \
	fi

# Validate blueprint consistency (presets, TEAM.md sync)
validate-blueprint:
	@python3 $(BLUEPRINT_VALIDATOR) $(BLUEPRINT_DIR)

# Auto-bump skill versions based on staged changes
bump-versions:
	@python3 $(FACTORY_SKILLS_DIR)/skill-creator/scripts/bump_versions.py

# Generate TEAM.md from blueprint/skills/ directory
generate-team:
	@echo "📝 Generating blueprint/rules/TEAM.md..."
	@echo "# Team Roster" > $(BLUEPRINT_DIR)/rules/TEAM.md
	@echo "" >> $(BLUEPRINT_DIR)/rules/TEAM.md
	@echo "Auto-generated from blueprint/skills/. Do not edit manually." >> $(BLUEPRINT_DIR)/rules/TEAM.md
	@echo "" >> $(BLUEPRINT_DIR)/rules/TEAM.md
	@echo "| Skill | Description |" >> $(BLUEPRINT_DIR)/rules/TEAM.md
	@echo "|-------|-------------|" >> $(BLUEPRINT_DIR)/rules/TEAM.md
	@for skill in $(BLUEPRINT_DIR)/skills/*/; do \
		if [ -f "$$skill/SKILL.md" ]; then \
			skill_name=$$(basename "$$skill"); \
			desc=$$(grep -m1 "^description:" "$$skill/SKILL.md" | sed 's/description: *//' | cut -c1-60); \
			echo "| \`$$skill_name\` | $$desc |" >> $(BLUEPRINT_DIR)/rules/TEAM.md; \
		fi \
	done
	@echo "" >> $(BLUEPRINT_DIR)/rules/TEAM.md
	@echo "## Usage" >> $(BLUEPRINT_DIR)/rules/TEAM.md
	@echo "" >> $(BLUEPRINT_DIR)/rules/TEAM.md
	@echo "Reference skills with \`@skill-name\` in skill collaboration sections." >> $(BLUEPRINT_DIR)/rules/TEAM.md
	@echo "✅ Generated TEAM.md with $$(ls -d $(BLUEPRINT_DIR)/skills/*/ 2>/dev/null | wc -l | tr -d ' ') skills."

# Build factory CLI binary
build-factory:
	@echo "🔨 Building factory CLI..."
	@mkdir -p $(BIN_DIR)
	@go build -o $(FACTORY_BIN) .
	@echo "✅ Built $(FACTORY_BIN)"

# Install factory CLI symlink to /usr/local/bin
install-factory:
	@echo "📦 Installing factory CLI symlink..."
	@if [ -f "$(FACTORY_BIN)" ]; then \
		if [ -L "$(SYMLINK_PATH)" ] || [ -e "$(SYMLINK_PATH)" ]; then \
			rm -f "$(SYMLINK_PATH)" 2>/dev/null || sudo rm -f "$(SYMLINK_PATH)"; \
		fi; \
		ln -s "$(FACTORY_BIN)" "$(SYMLINK_PATH)" 2>/dev/null || sudo ln -s "$(FACTORY_BIN)" "$(SYMLINK_PATH)"; \
		echo "✅ Symlinked $(FACTORY_BIN) -> $(SYMLINK_PATH)"; \
	else \
		echo "⚠️  factory binary not found. Run 'make build-factory' first."; \
	fi

# Generate and install shell completions
install-completions:
	@echo "🔧 Installing shell completions..."
	@mkdir -p $(BIN_DIR)
	@$(FACTORY_BIN) completion zsh > $(BIN_DIR)/factory.zsh 2>/dev/null || true
	@$(FACTORY_BIN) completion bash > $(BIN_DIR)/factory.bash 2>/dev/null || true
	@if [ -f "$(SHELL_RC)" ]; then \
		if ! grep -q "factory.zsh" "$(SHELL_RC)"; then \
			echo "" >> "$(SHELL_RC)"; \
			echo "# Antigravity Factory CLI completion" >> "$(SHELL_RC)"; \
			echo "source $(BIN_DIR)/factory.zsh 2>/dev/null || true" >> "$(SHELL_RC)"; \
			echo "✅ Added completion to $(SHELL_RC)"; \
		else \
			echo "✅ Completion already in $(SHELL_RC)"; \
		fi \
	fi

# Uninstall factory CLI
uninstall:
	@echo "🗑️  Uninstalling factory CLI..."
	@if [ -L "$(SYMLINK_PATH)" ]; then \
		rm -f "$(SYMLINK_PATH)" 2>/dev/null || sudo rm -f "$(SYMLINK_PATH)"; \
		echo "   Removed factory symlink"; \
	fi
	@echo "✅ Uninstall complete."

# Run tests
test:
	@echo "🧪 Running tests..."
	@go test ./internal/... ./cmd/... -v
	@echo "✅ All tests passed"

# Check file sizes (max 300 LOC per file)
check-loc:
	@echo "📏 Checking file sizes..."
	@failed=0; \
	for f in $$(find . -name "*.go" ! -name "*_test.go" ! -path "./vendor/*" ! -path "./blueprint/*"); do \
		lines=$$(wc -l < "$$f"); \
		if [ $$lines -gt 300 ]; then \
			echo "⚠️  $$f: $$lines lines (max 300)"; \
			failed=1; \
		fi \
	done; \
	if [ $$failed -eq 1 ]; then \
		echo "❌ Some files exceed 300 LOC limit"; \
		exit 1; \
	else \
		echo "✅ All files under 300 LOC"; \
	fi

# Run linter (MAXIMUM STRICTNESS)
lint:
	@echo "🔍 Running linter (FASCIST MODE)..."
	@if command -v golangci-lint >/dev/null 2>&1; then \
		golangci-lint run ./cmd/... ./internal/... --config .golangci.yml; \
	else \
		echo "⚠️  golangci-lint not installed. Install with:"; \
		echo "   brew install golangci-lint"; \
		echo "   or: go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest"; \
	fi

# Clean build artifacts
clean:
	@echo "🧹 Cleaning..."
	@rm -rf $(BIN_DIR)
	@rm -f factory
	@echo "✅ Clean complete"

# Generate CHANGELOG.md from git history using git-cliff
changelog:
	@echo "📝 Generating CHANGELOG.md..."
	@if command -v git-cliff >/dev/null 2>&1; then \
		git-cliff -o CHANGELOG.md; \
		echo "✅ CHANGELOG.md updated"; \
	else \
		echo "⚠️  git-cliff not installed. Install with:"; \
		echo "   brew install git-cliff"; \
	fi

