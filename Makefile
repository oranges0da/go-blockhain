new:
	go run main.go

fresh:
	rm -rf tmp
dev:
	rm -rf tmp && go run main.go