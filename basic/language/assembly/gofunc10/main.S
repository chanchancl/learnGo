"".f1 STEXT nosplit size=20 args=0x10 locals=0x0
	0x0000 00000 (main.go:4)	TEXT	"".f1(SB), NOSPLIT|ABIInternal, $0-16
	0x0000 00000 (main.go:4)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0000 00000 (main.go:4)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0000 00000 (main.go:4)	MOVQ	$0, "".~r1+16(SP)
	0x0009 00009 (main.go:5)	MOVQ	"".a+8(SP), AX
	0x000e 00014 (main.go:5)	MOVQ	AX, "".~r1+16(SP)
	0x0013 00019 (main.go:5)	RET
	0x0000 48 c7 44 24 10 00 00 00 00 48 8b 44 24 08 48 89  H.D$.....H.D$.H.
	0x0010 44 24 10 c3                                      D$..
"".main STEXT size=79 args=0x0 locals=0x28
	0x0000 00000 (main.go:9)	TEXT	"".main(SB), ABIInternal, $40-0
	0x0000 00000 (main.go:9)	MOVQ	(TLS), CX
	0x0009 00009 (main.go:9)	CMPQ	SP, 16(CX)
	0x000d 00013 (main.go:9)	PCDATA	$0, $-2
	0x000d 00013 (main.go:9)	JLS	72
	0x000f 00015 (main.go:9)	PCDATA	$0, $-1
	0x000f 00015 (main.go:9)	SUBQ	$40, SP
	0x0013 00019 (main.go:9)	MOVQ	BP, 32(SP)
	0x0018 00024 (main.go:9)	LEAQ	32(SP), BP
	0x001d 00029 (main.go:9)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001d 00029 (main.go:9)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001d 00029 (main.go:10)	MOVL	$16, (SP)
	0x0024 00036 (main.go:10)	LEAQ	"".f1·f(SB), AX
	0x002b 00043 (main.go:10)	MOVQ	AX, 8(SP)
	0x0030 00048 (main.go:10)	MOVQ	$10, 16(SP)
	0x0039 00057 (main.go:10)	PCDATA	$1, $0
	0x0039 00057 (main.go:10)	CALL	runtime.newproc(SB)
	0x003e 00062 (main.go:11)	MOVQ	32(SP), BP
	0x0043 00067 (main.go:11)	ADDQ	$40, SP
	0x0047 00071 (main.go:11)	RET
	0x0048 00072 (main.go:11)	NOP
	0x0048 00072 (main.go:9)	PCDATA	$1, $-1
	0x0048 00072 (main.go:9)	PCDATA	$0, $-2
	0x0048 00072 (main.go:9)	CALL	runtime.morestack_noctxt(SB)
	0x004d 00077 (main.go:9)	PCDATA	$0, $-1
	0x004d 00077 (main.go:9)	JMP	0
	0x0000 64 48 8b 0c 25 00 00 00 00 48 3b 61 10 76 39 48  dH..%....H;a.v9H
	0x0010 83 ec 28 48 89 6c 24 20 48 8d 6c 24 20 c7 04 24  ..(H.l$ H.l$ ..$
	0x0020 10 00 00 00 48 8d 05 00 00 00 00 48 89 44 24 08  ....H......H.D$.
	0x0030 48 c7 44 24 10 0a 00 00 00 e8 00 00 00 00 48 8b  H.D$..........H.
	0x0040 6c 24 20 48 83 c4 28 c3 e8 00 00 00 00 eb b1     l$ H..(........
	rel 5+4 t=17 TLS+0
	rel 39+4 t=16 "".f1·f+0
	rel 58+4 t=8 runtime.newproc+0
	rel 73+4 t=8 runtime.morestack_noctxt+0
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
