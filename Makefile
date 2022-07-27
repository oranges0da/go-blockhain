new:
	go run main.go
clean:
	rm -rf tmp
dev:
	rm -rf tmp && go run main.go
build:
	go build main.go