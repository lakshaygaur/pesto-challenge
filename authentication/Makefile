
# update GOARCH and GOOS accordingly
compile:
	echo "Compiling for every OS and Platform"
	GOOS=freebsd GOARCH=386 go build -o bin/main-freebsd-386 main.go
	GOOS=linux GOARCH=386 go build -o bin/main-linux-386 main.go
	GOOS=windows GOARCH=386 go build -o bin/main-windows-386 main.go

build:
	echo "Building bin"
	go build -o build/

clean:
	rm -rf build

all: build

go-run:
	AUTH_CONFIG_YAML=/home/lakshay/go-playground/pesto-challenge/authentication/config \
	go run cmd/main.go