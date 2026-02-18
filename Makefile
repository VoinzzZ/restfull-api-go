.PHONY: run build start clean
# Jalankan server (development)
run:
	go run ./cmd/api/main.go
# Build binary ke folder bin/
build:
	go build -o bin/api ./cmd/api/main.go
# Jalankan binary hasil build
start:
	./bin/api
# Hapus binary
clean:
	rm -rf bin/