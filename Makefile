init:
	go mod init ktool

build:
	go build -o ktool main.go

run:
	go run main.go

clean:
	rm go.mod
	rm go.sum
	rm ktool
