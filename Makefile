friends-management:
	echo "friends-management"

build:
	go build -o bin/main main.go

run:
	go run main.go

#If you wanted to cross-compile your application to run on every OS \
and every architecture available but didn’t want to manually set the GOOS and GOARCH variables for every command.
compile:
	echo "Compiling for every OS and Platform"
	GOOS=freebsd GOARCH=386 go build -o bin/main-freebsd-386 main.go
	GOOS=linux GOARCH=386 go build -o bin/main-linux-386 main.go
	GOOS=windows GOARCH=386 go build -o bin/main-windows-386 main.go

all: friends-management build run