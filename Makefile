NO_COLOR=\033[0m
OK_COLOR=\033[32;01m
ERROR_COLOR=\033[31;01m
WARN_COLOR=\033[33;01m
DEPS = $(go list -f '{{range .TestImports}}{{.}} {{end}}' ./...)

deps:
	@echo "$(OK_COLOR)==> Installing dependencies$(NO_COLOR)"
	@go get -d -v ./...
	@echo $(DEPS) | xargs -n1 go get -d

updatedeps:
	@echo "$(OK_COLOR)==> Updating all dependencies$(NO_COLOR)"
	@go get -d -v -u ./...
	@echo $(DEPS) | xargs -n1 go get -d -u

format:
	go fmt ./...

test: deps
	@echo "$(OK_COLOR)==> Testing s3lite...$(NO_COLOR)"
	go test ./...
