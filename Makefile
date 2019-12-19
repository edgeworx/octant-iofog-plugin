PLUGIN_NAME=iofog
BUILD_DIR ?= bin
PACKAGE_DIR = cmd

ifdef XDG_CONFIG_HOME
	OCTANT_PLUGINSTUB_DIR ?= ${XDG_CONFIG_HOME}/octant/plugins
# Determine in on windows
else ifeq ($(OS),Windows_NT) 
	OCTANT_PLUGINSTUB_DIR ?= ${LOCALAPPDATA}/octant/plugins
else
	OCTANT_PLUGINSTUB_DIR ?= ${HOME}/.config/octant/plugins
endif

build:
	@go build -o $(BUILD_DIR)/$(PLUGIN_NAME) -v ./$(PACKAGE_DIR)

install: build
	@echo Installing to $(OCTANT_PLUGINSTUB_DIR)/$(PLUGIN_NAME)
	@mkdir -p $(OCTANT_PLUGINSTUB_DIR)
	@cp $(BUILD_DIR)/$(PLUGIN_NAME) $(OCTANT_PLUGINSTUB_DIR)/$(PLUGIN_NAME)
	chmod a+x $(OCTANT_PLUGINSTUB_DIR)/$(PLUGIN_NAME)

dep:
	@dep ensure -v -vendor-only

