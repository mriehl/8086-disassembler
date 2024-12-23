run: nasm
  go run main.go
  for listing in 37 38 39 40; do \
    echo Checking re-assembled listing "$listing"; \
    nasm asm_decoded/$listing.asm; \
    xxd asm/$listing > asm/$listing.hex; \
    xxd asm_decoded/$listing > asm_decoded/$listing.hex; \
    diff -q asm_decoded/$listing asm/$listing || \
    diff -y "asm_decoded/$listing.asm" "asm/$listing.asm" || \
    diff -y "asm_decoded/$listing.hex" "asm/$listing.hex"; \
  done

test:
  go test ./...
nasm:
  for file in asm/*.asm; do nasm "$file"; done
