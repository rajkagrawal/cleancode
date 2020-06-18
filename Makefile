project = cleancode
.Phony : install
install:
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${project} cmd/main.go

#lint-prepare:
#	@echo "Installing golangci-lint"
#	curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh| sh -s latest
#
#lint:
#	./bin/golangci-lint run ./...
