.PHONY: help run-nats run-sender run-receiver run-lint run-tests

###############################################################################
#
#	Utils
#
###############################################################################

# COLORS
GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
WHITE  := $(shell tput -Txterm setaf 7)
RESET  := $(shell tput -Txterm sgr0)


TARGET_MAX_CHAR_NUM=25
help:
	@echo ''
	@echo 'Usage:'
	@echo '  ${YELLOW}make${RESET} ${GREEN}<target>${RESET}'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
		helpMessage = match(lastLine, /^## (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 3, RLENGTH); \
			printf "  ${YELLOW}%-$(TARGET_MAX_CHAR_NUM)s${RESET} ${GREEN}%s${RESET}\n", helpCommand, helpMessage; \
		} \
	} \

###############################################################################
#
#	Commands
#
###############################################################################

# Run a NATS server (with Docker).
run-nats:
	docker run -p 4222:4222 -ti nats:latest

# Run a cloudevents sender with NATS.
run-sender:
	go run ./cmd/sender/main.go

# Run a cloudevents receiver with NATS.
run-receiver:
	go run ./cmd/receiver/main.go

# Run the linter (static code analysis with `golangci-lint`).
run-lint:
	$(warning This command requires `golangci-lint`)
	golangci-lint run

## Run all tests.
run-tests: run-unit-tests

## Run all unit tests.
run-unit-tests:
	$(warning Not implemented yet)