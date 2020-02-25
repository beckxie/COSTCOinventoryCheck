.PHONY: all
all: COSTIC-linux-amd64 COSTIC-windows-amd64

.PHONY: COSTIC-linux-amd64
COSTIC-linux-amd64:
        @echo "Building linux binary..."
        @CGO_ENABLED=0 \
        GOOS=linux \
        GOARCH=amd64 \
        go build -mod=vendor -o bin/COSTIC_linux_amd64

.PHONY: COSTIC-windows-amd64
COSTIC-windows-amd64:
        @echo "Building windows binary..."
        @CGO_ENABLED=0 \
        GOOS=windows \
        GOARCH=amd64 \
        go build -mod=vendor -o bin/COSTIC_windows_amd64
