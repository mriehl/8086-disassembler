; ========================================================================
; LISTING 40
; ========================================================================

bits 16

; Signed displacements
mov ax, [bx + di - 37]
mov [si - 300], cx
mov dx, [bx - 32]

; Explicit sizes
mov [bp + di], byte 7
mov [di + 901], word 347

; Direct address
mov bp, [5]
mov bx, [3458]

; Memory-to-accumulator test
mov ax, [2555]
mov ax, [16]

; Accumulator-to-memory test
mov [2554], ax
mov [15], ax


; for completeness
; rm to sr
;mov [42], cs
;mov cs, [999]
;mov cx, ds
;mov [di+915], ss
; sr to rm
;mov cs, [42]
;mov cs, [999]
;mov ds, cx
;mov [di+915], ss
;mov ss, [di+915]
