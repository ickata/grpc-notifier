# gRPC Notifier

This is a demo desktop service for scheduling system notifications based on Go, gRPC and React.

## Setup

1. [Install Go](https://golang.org/dl/)
2. [Install Docker](https://docs.docker.com/get-docker/)
3. Clone this repo in your `$GOPATH/src` (unix) or `%GOPATH%/src` (windows)
4. Generate gRPC stubs `cd api && ./build.sh` (unix) or `cd api && build.bat` (windows)
5. Install UI dependencies `cd ui && yarn` (yarn) or `cd ui && npm i` (npm)
6. Build UI `cd ui && yarn build` (yarn) or `cd ui && npm run build` (npm)
7. Install [statik CLI](https://github.com/rakyll/statik) `go get github.com/rakyll/statik`
8. Build static UI `cd service && statik -src=../ui/build/`
9. Build service `go build -o notifier ./service`

## Run

Run `./notifier` (unix) or `notifier.exe` (windows)
