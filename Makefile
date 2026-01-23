.PHONY: install uninstall install-factory install-squads generate-team validate validate-all build-skills install-skills test lint

# Paths
GLOBAL_SKILLS_DIR := $(HOME)/.gemini/antigravity/global_skills
FACTORY_SKILLS_DIR := $(shell pwd)/.agent/skills
SQUADS_DIR := $(shell pwd)/squads
VALIDATOR := $(FACTORY_SKILLS_DIR)/skill-creator/scripts/validate_skill.py
BIN_DIR := $(shell pwd)/bin
SKILLS_BIN := $(BIN_DIR)/skills
SYMLINK_PATH := /usr/local/bin/skills
SHELL_RC := $(HOME)/.zshrc

# Install everything (factory + squads + skills CLI)
install: generate-team validate-all install-factory install-squads build-skills install-skills install-completions
	@echo ""
	@echo "🎉 Installation complete!"
	@echo "   Run 'skills --help' from anywhere to get started."

# Validate a single skill: make validate SKILL=mcp-expert
validate:
	@if [ -z "$(SKILL)" ]; then \
		echo "Usage: make validate SKILL=<skill-name>"; \
		exit 1; \
	fi
	@python3 $(VALIDATOR) $(SQUADS_DIR)/$(SKILL)

# Validate all squads
validate-all:
	@echo "🔍 Validating all skills..."
	@failed=0; \
	for skill in $(SQUADS_DIR)/*/; do \
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

# Generate TEAM.md from squads/ directory
generate-team:
	@echo "📝 Generating squads/TEAM.md..."
	@echo "# Squad Team" > $(SQUADS_DIR)/TEAM.md
	@echo "" >> $(SQUADS_DIR)/TEAM.md
	@echo "Auto-generated from squads/ directory. Do not edit manually." >> $(SQUADS_DIR)/TEAM.md
	@echo "" >> $(SQUADS_DIR)/TEAM.md
	@echo "## Team Roster" >> $(SQUADS_DIR)/TEAM.md
	@echo "" >> $(SQUADS_DIR)/TEAM.md
	@echo "| Skill | Description |" >> $(SQUADS_DIR)/TEAM.md
	@echo "|-------|-------------|" >> $(SQUADS_DIR)/TEAM.md
	@for skill in $(SQUADS_DIR)/*/; do \
		if [ -f "$$skill/SKILL.md" ]; then \
			skill_name=$$(basename "$$skill"); \
			desc=$$(grep -m1 "^description:" "$$skill/SKILL.md" | sed 's/description: *//' | cut -c1-60); \
			echo "| \`$$skill_name\` | $$desc |" >> $(SQUADS_DIR)/TEAM.md; \
		fi \
	done
	@echo "" >> $(SQUADS_DIR)/TEAM.md
	@echo "## Usage" >> $(SQUADS_DIR)/TEAM.md
	@echo "" >> $(SQUADS_DIR)/TEAM.md
	@echo "Reference skills in Team Collaboration sections:" >> $(SQUADS_DIR)/TEAM.md
	@echo "" >> $(SQUADS_DIR)/TEAM.md
	@echo "\`\`\`markdown" >> $(SQUADS_DIR)/TEAM.md
	@echo "## Team Collaboration" >> $(SQUADS_DIR)/TEAM.md
	@echo "- **Role**: \`@skill-name\` (Description)" >> $(SQUADS_DIR)/TEAM.md
	@echo "\`\`\`" >> $(SQUADS_DIR)/TEAM.md
	@echo "✅ Generated TEAM.md with $$(ls -d $(SQUADS_DIR)/*/ 2>/dev/null | wc -l | tr -d ' ') skills."

# Install only the skill-creator (the factory itself)
install-factory:
	@echo "🏭 Installing Skill Factory (skill-creator)..."
	@mkdir -p $(GLOBAL_SKILLS_DIR)
	@for skill in $(FACTORY_SKILLS_DIR)/*; do \
		if [ -d "$$skill" ]; then \
			skill_name=$$(basename "$$skill"); \
			target_dir="$(GLOBAL_SKILLS_DIR)/$$skill_name"; \
			echo "   📦 Copying $$skill_name..."; \
			if [ -e "$$target_dir" ]; then \
				rm -rf "$$target_dir"; \
			fi; \
			cp -r "$$skill" "$$target_dir"; \
		fi \
	done
	@echo "✅ Skill Factory installed."

# Install the generated squad skills
install-squads:
	@echo "👥 Installing Squad skills from $(SQUADS_DIR)..."
	@if [ -d "$(SQUADS_DIR)" ]; then \
		for skill in $(SQUADS_DIR)/*; do \
			if [ -d "$$skill" ]; then \
				skill_name=$$(basename "$$skill"); \
				target_dir="$(GLOBAL_SKILLS_DIR)/$$skill_name"; \
				echo "   📦 Copying $$skill_name..."; \
				if [ -e "$$target_dir" ]; then \
					rm -rf "$$target_dir"; \
				fi; \
				cp -r "$$skill" "$$target_dir"; \
			fi \
		done; \
		echo "✅ Squad skills installed."; \
	else \
		echo "⚠️  No squads/ folder found. Skipping."; \
	fi
	@echo "📚 Installing Standards..."
	@if [ -d "$(SQUADS_DIR)/_standards" ]; then \
		target_dir="$(GLOBAL_SKILLS_DIR)/_standards"; \
		if [ -e "$$target_dir" ]; then \
			rm -rf "$$target_dir"; \
		fi; \
		cp -r "$(SQUADS_DIR)/_standards" "$$target_dir"; \
		echo "✅ Standards installed to $$target_dir"; \
	fi

uninstall:
	@echo "🗑️  Uninstalling all skills..."
	@for skill in $(FACTORY_SKILLS_DIR)/*; do \
		if [ -d "$$skill" ]; then \
			skill_name=$$(basename "$$skill"); \
			target_dir="$(GLOBAL_SKILLS_DIR)/$$skill_name"; \
			if [ -e "$$target_dir" ]; then \
				rm -rf "$$target_dir"; \
				echo "   Removed $$skill_name"; \
			fi \
		fi \
	done
	@if [ -d "$(SQUADS_DIR)" ]; then \
		for skill in $(SQUADS_DIR)/*; do \
			if [ -d "$$skill" ]; then \
				skill_name=$$(basename "$$skill"); \
				target_dir="$(GLOBAL_SKILLS_DIR)/$$skill_name"; \
				if [ -e "$$target_dir" ]; then \
					rm -rf "$$target_dir"; \
					echo "   Removed $$skill_name"; \
				fi \
			fi \
		done; \
	fi
	@# Remove symlink
	@if [ -L "$(SYMLINK_PATH)" ]; then \
		rm -f "$(SYMLINK_PATH)" 2>/dev/null || sudo rm -f "$(SYMLINK_PATH)"; \
		echo "   Removed skills symlink"; \
	fi
	@echo "✅ Uninstall complete."

# Build skills CLI binary to bin/
build-skills:
	@echo "🔨 Building skills CLI..."
	@mkdir -p $(BIN_DIR)
	@go build -o $(SKILLS_BIN) .
	@echo "✅ Built $(SKILLS_BIN)"

# Install skills CLI symlink to /usr/local/bin
install-skills:
	@echo "📦 Installing skills CLI symlink..."
	@if [ -f "$(SKILLS_BIN)" ]; then \
		if [ -L "$(SYMLINK_PATH)" ] || [ -e "$(SYMLINK_PATH)" ]; then \
			rm -f "$(SYMLINK_PATH)" 2>/dev/null || sudo rm -f "$(SYMLINK_PATH)"; \
		fi; \
		ln -s "$(SKILLS_BIN)" "$(SYMLINK_PATH)" 2>/dev/null || sudo ln -s "$(SKILLS_BIN)" "$(SYMLINK_PATH)"; \
		echo "✅ Symlinked $(SKILLS_BIN) -> $(SYMLINK_PATH)"; \
	else \
		echo "⚠️  skills binary not found. Run 'make build-skills' first."; \
	fi

# Generate and install shell completions
install-completions:
	@echo "🔧 Installing shell completions..."
	@mkdir -p $(BIN_DIR)
	@$(SKILLS_BIN) completion zsh > $(BIN_DIR)/skills.zsh 2>/dev/null || true
	@$(SKILLS_BIN) completion bash > $(BIN_DIR)/skills.bash 2>/dev/null || true
	@# Add to zshrc if not already there
	@if [ -f "$(SHELL_RC)" ]; then \
		if ! grep -q "skills.zsh" "$(SHELL_RC)"; then \
			echo "" >> "$(SHELL_RC)"; \
			echo "# ag-skills CLI completion" >> "$(SHELL_RC)"; \
			echo "source $(BIN_DIR)/skills.zsh 2>/dev/null || true" >> "$(SHELL_RC)"; \
			echo "✅ Added completion to $(SHELL_RC)"; \
		else \
			echo "✅ Completion already in $(SHELL_RC)"; \
		fi \
	fi

# Run tests
test:
	@echo "🧪 Running tests..."
	@go test ./internal/... ./cmd/... -v
	@echo "✅ All tests passed"

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
	@rm -f skills
	@echo "✅ Clean complete"
