all: clean build

build:
	go build gotftpd.go

test:
	go test -cover -race -v

win64: clean
	env GOOS=windows GOARCH=amd64 go build -a -o gotftpd.exe gotftpd.go

clean:
	go clean -i

