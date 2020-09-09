"".f1 STEXT size=71 args=0x0 locals=0x18
	0x0000 00000 (main.go:4)	TEXT	"".f1(SB), ABIInternal, $24-0
	0x0000 00000 (main.go:4)	MOVQ	(TLS), CX
	0x0009 00009 (main.go:4)	CMPQ	SP, 16(CX)
	0x000d 00013 (main.go:4)	PCDATA	$0, $-2
	0x000d 00013 (main.go:4)	JLS	64
	0x000f 00015 (main.go:4)	PCDATA	$0, $-1
	0x000f 00015 (main.go:4)	SUBQ	$24, SP
	0x0013 00019 (main.go:4)	MOVQ	BP, 16(SP)
	0x0018 00024 (main.go:4)	LEAQ	16(SP), BP
	// This shows that function don't allocate space for return address
	// 00(SP) 0, arguents size
	// 08(SP) address of f1
	// 16(SP) Raw BP value
	0x001d 00029 (main.go:4)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001d 00029 (main.go:4)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001d 00029 (main.go:5)	MOVL	$0, (SP)
	0x0024 00036 (main.go:5)	LEAQ	"".f1·f(SB), AX
	0x002b 00043 (main.go:5)	MOVQ	AX, 8(SP)
	0x0030 00048 (main.go:5)	PCDATA	$1, $0
	0x0030 00048 (main.go:5)	CALL	runtime.newproc(SB)
	0x0035 00053 (main.go:6)	MOVQ	16(SP), BP
	0x003a 00058 (main.go:6)	ADDQ	$24, SP
	0x003e 00062 (main.go:6)	RET
	0x003f 00063 (main.go:6)	NOP
	0x003f 00063 (main.go:4)	PCDATA	$1, $-1
	0x003f 00063 (main.go:4)	PCDATA	$0, $-2
	0x003f 00063 (main.go:4)	NOP
	0x0040 00064 (main.go:4)	CALL	runtime.morestack_noctxt(SB)
	0x0045 00069 (main.go:4)	PCDATA	$0, $-1
	0x0045 00069 (main.go:4)	JMP	0
	0x0000 64 48 8b 0c 25 00 00 00 00 48 3b 61 10 76 31 48  dH..%....H;a.v1H
	0x0010 83 ec 18 48 89 6c 24 10 48 8d 6c 24 10 c7 04 24  ...H.l$.H.l$...$
	0x0020 00 00 00 00 48 8d 05 00 00 00 00 48 89 44 24 08  ....H......H.D$.
	0x0030 e8 00 00 00 00 48 8b 6c 24 10 48 83 c4 18 c3 90  .....H.l$.H.....
	0x0040 e8 00 00 00 00 eb b9                             .......
	rel 5+4 t=17 TLS+0
	rel 39+4 t=16 "".f1·f+0
	rel 49+4 t=8 runtime.newproc+0
	rel 65+4 t=8 runtime.morestack_noctxt+0
"".main STEXT size=71 args=0x0 locals=0x18
	0x0000 00000 (main.go:9)	TEXT	"".main(SB), ABIInternal, $24-0
	0x0000 00000 (main.go:9)	MOVQ	(TLS), CX
	0x0009 00009 (main.go:9)	CMPQ	SP, 16(CX)
	0x000d 00013 (main.go:9)	PCDATA	$0, $-2
	0x000d 00013 (main.go:9)	JLS	64
	0x000f 00015 (main.go:9)	PCDATA	$0, $-1
	0x000f 00015 (main.go:9)	SUBQ	$24, SP
	0x0013 00019 (main.go:9)	MOVQ	BP, 16(SP)
	0x0018 00024 (main.go:9)	LEAQ	16(SP), BP
	0x001d 00029 (main.go:9)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001d 00029 (main.go:9)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001d 00029 (main.go:10)	MOVL	$0, (SP)
	0x0024 00036 (main.go:10)	LEAQ	"".f1·f(SB), AX
	0x002b 00043 (main.go:10)	MOVQ	AX, 8(SP)
	0x0030 00048 (main.go:10)	PCDATA	$1, $0
	0x0030 00048 (main.go:10)	CALL	runtime.newproc(SB)
	0x0035 00053 (main.go:11)	MOVQ	16(SP), BP
	0x003a 00058 (main.go:11)	ADDQ	$24, SP
	0x003e 00062 (main.go:11)	RET
	0x003f 00063 (main.go:11)	NOP
	0x003f 00063 (main.go:9)	PCDATA	$1, $-1
	0x003f 00063 (main.go:9)	PCDATA	$0, $-2
	0x003f 00063 (main.go:9)	NOP
	0x0040 00064 (main.go:9)	CALL	runtime.morestack_noctxt(SB)
	0x0045 00069 (main.go:9)	PCDATA	$0, $-1
	0x0045 00069 (main.go:9)	JMP	0
	0x0000 64 48 8b 0c 25 00 00 00 00 48 3b 61 10 76 31 48  dH..%....H;a.v1H
	0x0010 83 ec 18 48 89 6c 24 10 48 8d 6c 24 10 c7 04 24  ...H.l$.H.l$...$
	0x0020 00 00 00 00 48 8d 05 00 00 00 00 48 89 44 24 08  ....H......H.D$.
	0x0030 e8 00 00 00 00 48 8b 6c 24 10 48 83 c4 18 c3 90  .....H.l$.H.....
	0x0040 e8 00 00 00 00 eb b9                             .......
	rel 5+4 t=17 TLS+0
	rel 39+4 t=16 "".f1·f+0
	rel 49+4 t=8 runtime.newproc+0
	rel 65+4 t=8 runtime.morestack_noctxt+0
go.cuinfo.packagename. SDWARFINFO dupok size=0
	0x0000 6d 61 69 6e                                      main
""..inittask SNOPTRDATA size=24
	0x0000 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0010 00 00 00 00 00 00 00 00                          ........
"".f1·f SRODATA dupok size=8
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 "".f1+0
gclocals·33cdeccccebe80329f1fdbee7f5874cb SRODATA dupok size=8
	0x0000 01 00 00 00 00 00 00 00                          ........
