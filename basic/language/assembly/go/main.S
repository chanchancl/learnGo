"".f1 STEXT nosplit size=49 args=0x20 locals=0x0
	0x0000 00000 (main.go:4)	TEXT	"".f1(SB), NOSPLIT|ABIInternal, $0-32
	0x0000 00000 (main.go:4)	PCDATA	$0, $-2
	0x0000 00000 (main.go:4)	PCDATA	$1, $-2
	0x0000 00000 (main.go:4)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0000 00000 (main.go:4)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0000 00000 (main.go:4)	FUNCDATA	$2, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0000 00000 (main.go:4)	PCDATA	$0, $0
	0x0000 00000 (main.go:4)	PCDATA	$1, $0
	0x0000 00000 (main.go:4)	MOVQ	$0, "".~r2+24(SP)
	0x0009 00009 (main.go:4)	MOVQ	$0, "".~r3+32(SP)
	0x0012 00018 (main.go:5)	MOVQ	"".a+8(SP), AX
	0x0017 00023 (main.go:5)	ADDQ	"".b+16(SP), AX
	0x001c 00028 (main.go:5)	MOVQ	AX, "".~r2+24(SP)
	0x0021 00033 (main.go:5)	MOVQ	"".a+8(SP), AX
	0x0026 00038 (main.go:5)	SUBQ	"".b+16(SP), AX
	0x002b 00043 (main.go:5)	MOVQ	AX, "".~r3+32(SP)
	0x0030 00048 (main.go:5)	RET
	0x0000 48 c7 44 24 18 00 00 00 00 48 c7 44 24 20 00 00  H.D$.....H.D$ ..
	0x0010 00 00 48 8b 44 24 08 48 03 44 24 10 48 89 44 24  ..H.D$.H.D$.H.D$
	0x0020 18 48 8b 44 24 08 48 2b 44 24 10 48 89 44 24 20  .H.D$.H+D$.H.D$ 
	0x0030 c3                                               .
"".main STEXT size=88 args=0x0 locals=0x38
	0x0000 00000 (main.go:9)	TEXT	"".main(SB), ABIInternal, $56-0
	0x0000 00000 (main.go:9)	MOVQ	(TLS), CX
	0x0009 00009 (main.go:9)	CMPQ	SP, 16(CX)
	0x000d 00013 (main.go:9)	PCDATA	$0, $-2
	0x000d 00013 (main.go:9)	JLS	81
	0x000f 00015 (main.go:9)	PCDATA	$0, $-1
	0x000f 00015 (main.go:9)	SUBQ	$56, SP
	0x0013 00019 (main.go:9)	MOVQ	BP, 48(SP)
	0x0018 00024 (main.go:9)	LEAQ	48(SP), BP
	0x001d 00029 (main.go:9)	PCDATA	$0, $-2
	0x001d 00029 (main.go:9)	PCDATA	$1, $-2
	0x001d 00029 (main.go:9)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001d 00029 (main.go:9)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001d 00029 (main.go:9)	FUNCDATA	$2, gclocals·9fb7f0986f647f17cb53dda1484e0f7a(SB)
	0x001d 00029 (main.go:10)	PCDATA	$0, $0
	0x001d 00029 (main.go:10)	PCDATA	$1, $0
	0x001d 00029 (main.go:10)	MOVL	$32, (SP)
	0x0024 00036 (main.go:10)	PCDATA	$0, $1
	0x0024 00036 (main.go:10)	LEAQ	"".f1·f(SB), AX
	0x002b 00043 (main.go:10)	PCDATA	$0, $0
	0x002b 00043 (main.go:10)	MOVQ	AX, 8(SP)
	0x0030 00048 (main.go:10)	MOVQ	$10, 16(SP)
	0x0039 00057 (main.go:10)	MOVQ	$11, 24(SP)
	0x0042 00066 (main.go:10)	CALL	runtime.newproc(SB)
	0x0047 00071 (main.go:11)	MOVQ	48(SP), BP
	0x004c 00076 (main.go:11)	ADDQ	$56, SP
	0x0050 00080 (main.go:11)	RET
	0x0051 00081 (main.go:11)	NOP
	0x0051 00081 (main.go:9)	PCDATA	$1, $-1
	0x0051 00081 (main.go:9)	PCDATA	$0, $-2
	0x0051 00081 (main.go:9)	CALL	runtime.morestack_noctxt(SB)
	0x0056 00086 (main.go:9)	PCDATA	$0, $-1
	0x0056 00086 (main.go:9)	JMP	0
	0x0000 64 48 8b 0c 25 00 00 00 00 48 3b 61 10 76 42 48  dH..%....H;a.vBH
	0x0010 83 ec 38 48 89 6c 24 30 48 8d 6c 24 30 c7 04 24  ..8H.l$0H.l$0..$
	0x0020 20 00 00 00 48 8d 05 00 00 00 00 48 89 44 24 08   ...H......H.D$.
	0x0030 48 c7 44 24 10 0a 00 00 00 48 c7 44 24 18 0b 00  H.D$.....H.D$...
	0x0040 00 00 e8 00 00 00 00 48 8b 6c 24 30 48 83 c4 38  .......H.l$0H..8
	0x0050 c3 e8 00 00 00 00 eb a8                          ........
	rel 5+4 t=17 TLS+0
	rel 39+4 t=16 "".f1·f+0
	rel 67+4 t=8 runtime.newproc+0
	rel 82+4 t=8 runtime.morestack_noctxt+0
go.cuinfo.packagename. SDWARFINFO dupok size=0
	0x0000 6d 61 69 6e                                      main
go.loc."".f1 SDWARFLOC size=0
go.info."".f1 SDWARFINFO size=82
	0x0000 03 22 22 2e 66 31 00 00 00 00 00 00 00 00 00 00  ."".f1..........
	0x0010 00 00 00 00 00 00 00 01 9c 00 00 00 00 01 0f 61  ...............a
	0x0020 00 00 04 00 00 00 00 01 9c 0f 62 00 00 04 00 00  ..........b.....
	0x0030 00 00 02 91 08 0f 7e 72 32 00 01 04 00 00 00 00  ......~r2.......
	0x0040 02 91 10 0f 7e 72 33 00 01 04 00 00 00 00 02 91  ....~r3.........
	0x0050 18 00                                            ..
	rel 0+0 t=24 type.int+0
	rel 7+8 t=1 "".f1+0
	rel 15+8 t=1 "".f1+49
	rel 25+4 t=30 gofile../home/zquxbow/go/learnGo/basic/language/assembly/go/main.go+0
	rel 35+4 t=29 go.info.int+0
	rel 46+4 t=29 go.info.int+0
	rel 60+4 t=29 go.info.int+0
	rel 74+4 t=29 go.info.int+0
go.range."".f1 SDWARFRANGE size=0
go.debuglines."".f1 SDWARFMISC size=15
	0x0000 04 02 12 06 69 06 6a 06 41 04 01 03 7c 06 01     ....i.j.A...|..
go.loc."".main SDWARFLOC size=0
go.info."".main SDWARFINFO size=33
	0x0000 03 22 22 2e 6d 61 69 6e 00 00 00 00 00 00 00 00  ."".main........
	0x0010 00 00 00 00 00 00 00 00 00 01 9c 00 00 00 00 01  ................
	0x0020 00                                               .
	rel 9+8 t=1 "".main+0
	rel 17+8 t=1 "".main+88
	rel 27+4 t=30 gofile../home/zquxbow/go/learnGo/basic/language/assembly/go/main.go+0
go.range."".main SDWARFRANGE size=0
go.debuglines."".main SDWARFMISC size=19
	0x0000 04 02 03 03 14 0a a5 9c 06 55 06 08 7e 71 04 01  .........U..~q..
	0x0010 03 78 01                                         .x.
""..inittask SNOPTRDATA size=24
	0x0000 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0010 00 00 00 00 00 00 00 00                          ........
"".f1·f SRODATA dupok size=8
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 "".f1+0
gclocals·33cdeccccebe80329f1fdbee7f5874cb SRODATA dupok size=8
	0x0000 01 00 00 00 00 00 00 00                          ........
gclocals·9fb7f0986f647f17cb53dda1484e0f7a SRODATA dupok size=10
	0x0000 02 00 00 00 01 00 00 00 00 01                    ..........
