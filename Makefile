build:
	@go build

run: build
	@go run main.go run

clean_default_theme:
	@rm -fr ${XDG_CONFIG_HOME}/kibe/theme.yaml