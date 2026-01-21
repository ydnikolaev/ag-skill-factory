.PHONY: install uninstall

# Paths
SKILLS_DIR := $(HOME)/.gemini/antigravity/skills
SKILL_NAME := skill-creator
SOURCE_DIR := $(shell pwd)/.agent/skills/$(SKILL_NAME)
TARGET_LINK := $(SKILLS_DIR)/$(SKILL_NAME)

install:
	@echo "🚀 Installing $(SKILL_NAME) to $(SKILLS_DIR)..."
	@mkdir -p $(SKILLS_DIR)
	@# Remove existing directory or symlink to ensure clean state
	@if [ -e "$(TARGET_LINK)" ]; then \
		echo "   Removing existing installation..."; \
		rm -rf "$(TARGET_LINK)"; \
	fi
	@# Create symbolic link for SSOT
	@ln -s "$(SOURCE_DIR)" "$(TARGET_LINK)"
	@echo "✅ Success! $(SKILL_NAME) is now linked globally."
	@echo "   Source: $(SOURCE_DIR)"
	@echo "   Link:   $(TARGET_LINK)"

uninstall:
	@rm -rf "$(TARGET_LINK)"
	@echo "🗑️  Uninstalled $(SKILL_NAME)"
