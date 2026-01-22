.PHONY: install uninstall install-factory install-squads generate-team validate validate-all

# Paths
GLOBAL_SKILLS_DIR := $(HOME)/.gemini/antigravity/global_skills
FACTORY_SKILLS_DIR := $(shell pwd)/.agent/skills
SQUADS_DIR := $(shell pwd)/squads
VALIDATOR := $(FACTORY_SKILLS_DIR)/skill-creator/scripts/validate_skill.py

# Install everything (factory + squads)
install: generate-team validate-all install-factory install-squads

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
	@echo "✅ Uninstall complete."
