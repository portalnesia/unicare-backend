APP=unicare

build:
	CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -ldflags "-s -w -X main.build=production" -o build/${APP} main.go
	upx --brute build/${APP}