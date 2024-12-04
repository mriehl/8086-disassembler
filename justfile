run: nasm
  go run main.go
test:
  go test ./...
nasm:
  nasm asm/37.asm
  nasm asm/38.asm
