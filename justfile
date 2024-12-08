run: nasm
  go run main.go
  for listing in 37 38 39; do \
    echo Checking re-assembled listing "$listing"; \
    nasm asm_decoded/$listing.asm; \
    diff -q asm_decoded/$listing asm/$listing || \
    diff -y --suppress-common-lines "asm_decoded/$listing.asm" "asm/$listing.asm"; \
  done

test:
  go test ./...
nasm:
  for file in asm/*.asm; do nasm "$file"; done
