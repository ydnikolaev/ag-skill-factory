.PHONY: install uninstall install-factory install-squads

# Paths
GLOBAL_SKILLS_DIR := $(HOME)/.gemini/antigravity/skills
FACTORY_SKILLS_DIR := $(shell pwd)/.agent/skills
SQUADS_DIR := $(shell pwd)/squads

# Install everything (factory + squads)
install: install-factory install-squads

# Install only the skill-creator (the factory itself)
install-factory:
	@echo "🏭 Installing Skill Factory (skill-creator)..."
	@mkdir -p $(GLOBAL_SKILLS_DIR)
	@for skill in $(FACTORY_SKILLS_DIR)/*; do \
		if [ -d "$$skill" ]; then \
			skill_name=$$(basename "$$skill"); \
			target_link="$(GLOBAL_SKILLS_DIR)/$$skill_name"; \
			echo "   🔗 Linking $$skill_name..."; \
			if [ -e "$$target_link" ]; then \
				rm -rf "$$target_link"; \
			fi; \
			ln -s "$$skill" "$$target_link"; \
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
				target_link="$(GLOBAL_SKILLS_DIR)/$$skill_name"; \
				echo "   🔗 Linking $$skill_name..."; \
				if [ -e "$$target_link" ]; then \
					rm -rf "$$target_link"; \
				fi; \
				ln -s "$$skill" "$$target_link"; \
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
			target_link="$(GLOBAL_SKILLS_DIR)/$$skill_name"; \
			if [ -L "$$target_link" ]; then \
				rm "$$target_link"; \
				echo "   Removed $$skill_name"; \
			fi \
		fi \
	done
	@if [ -d "$(SQUADS_DIR)" ]; then \
		for skill in $(SQUADS_DIR)/*; do \
			if [ -d "$$skill" ]; then \
				skill_name=$$(basename "$$skill"); \
				target_link="$(GLOBAL_SKILLS_DIR)/$$skill_name"; \
				if [ -L "$$target_link" ]; then \
					rm "$$target_link"; \
					echo "   Removed $$skill_name"; \
				fi \
			fi \
		done; \
	fi
	@echo "✅ Uninstall complete."
