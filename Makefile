.PHONY: install uninstall

# Paths
GLOBAL_SKILLS_DIR := $(HOME)/.gemini/antigravity/skills
LOCAL_SKILLS_DIR := $(shell pwd)/.agent/skills

install:
	@echo "🚀 Installing all skills from $(LOCAL_SKILLS_DIR) to $(GLOBAL_SKILLS_DIR)..."
	@mkdir -p $(GLOBAL_SKILLS_DIR)
	@for skill in $(LOCAL_SKILLS_DIR)/*; do \
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
	@echo "✅ Success! All skills are linked globally."

uninstall:
	@echo "🗑️  Uninstalling skills..."
	@for skill in $(LOCAL_SKILLS_DIR)/*; do \
		if [ -d "$$skill" ]; then \
			skill_name=$$(basename "$$skill"); \
			target_link="$(GLOBAL_SKILLS_DIR)/$$skill_name"; \
			if [ -L "$$target_link" ]; then \
				rm "$$target_link"; \
				echo "   Removed $$skill_name"; \
			fi \
		fi \
	done
