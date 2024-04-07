build:
	@go build

run: build
	@go run main.go run

test-cmd: build
	@go run main.go test

reset-theme:
	@rm -fr ${XDG_CONFIG_HOME}/kibe/theme.yaml

debug-logs:
	@jq debug.log

reset-debug-logs:
	@echo "" > debug.log

stream-debug-logs: reset-debug-logs
	@tail -f debug.log | jq .