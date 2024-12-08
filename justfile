run: nasm
  go run main.go
  for file in asm_decoded/*.asm; do echo Checking "$file"; diff -y --suppress-common-lines "$file" "${file/_decoded/}"; done

test:
  go test ./...
nasm:
  nasm asm/37.asm
  nasm asm/38.asm
