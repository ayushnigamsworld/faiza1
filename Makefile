
.PHONY: run
run: main
	./$<

main: *.go go.mod
	go build -o $@ .
	chmod +x $@

.PHONY: test
test:
	go test ./...

.PHONY: all
all: main
