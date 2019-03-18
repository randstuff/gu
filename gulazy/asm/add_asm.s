// add_asm.s
TEXT add(SB),$0
        MOVL x+0(FP), BX
        MOVL y+4(FP), BP
        ADDL BP, BX
        MOVL BX, ret+8(FP)
        RET
