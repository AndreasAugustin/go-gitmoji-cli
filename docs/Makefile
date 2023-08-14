SHELL := /bin/bash
.DEFAULT_GOAL := help

###########################
# VARIABLES
###########################

TAPES := commit.tape config.tape help.tape list.tape version.tape

###########################
# MAPPINGS
###########################

###########################
# TARGETS
###########################

.PHONY: help
help:  ## help target to show available commands with information
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) |  awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'


.PHONY: vhs $(TAPES)
vhs: $(TAPES) ## create gif from tape

$(TAPES):
	vhs $@
