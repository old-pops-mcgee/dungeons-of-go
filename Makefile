setup-desktop:
	rm -f go.mod && cp desktopgo.mod go.mod
	go mod tidy

setup-web:
	rm -f go.mod && cp webgo.mod go.mod
	go mod tidy

build:
	make setup-desktop
	GOOS=linux GOARCH=amd64 go build

build-web:
	make setup-web
	GOOS=js GOARCH=wasm go build -tags web -o ./Raylib-Go-Wasm/index/main.wasm .