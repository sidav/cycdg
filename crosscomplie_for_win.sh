GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc go build -ldflags="-s -w" -o cycdg$1$(date '+%d-%m-%Y_%H').exe *.go
