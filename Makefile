project = cleancode
.Phony : install
install:
	@go build -o ${project} cmd/main.go