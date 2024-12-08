run: nasm
  go run main.go
  # FIXME compare reassembled, not deassembled
  for file in asm_decoded/*.asm; do echo Checking "$file"; diff -y --suppress-common-lines "$file" "${file/_decoded/}"; done

test:
  go test ./...
nasm:
  for file in asm/*.asm; do nasm "$file"; done
