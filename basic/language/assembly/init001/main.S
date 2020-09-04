"".f1 STEXT nosplit size=49 args=0x20 locals=0x0
	0x0000 00000 (main.go:3)	TEXT	"".f1(SB), NOSPLIT|ABIInternal, $0-32
	0x0000 00000 (main.go:3)	PCDATA	$0, $-2
	0x0000 00000 (main.go:3)	PCDATA	$1, $-2
	0x0000 00000 (main.go:3)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0000 00000 (main.go:3)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0000 00000 (main.go:3)	FUNCDATA	$2, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0000 00000 (main.go:3)	PCDATA	$0, $0
	0x0000 00000 (main.go:3)	PCDATA	$1, $0
	// (gdb) x/g 0xc00003e750
	// 0xc00003e750:   0x00000000004583f3  //  0(SP) return address, and also 0(SP)
	// 0xc00003e758:   0x0000000000000001  //  8(SP) argument 1
	// 0xc00003e760:   0x0000000000000002  // 16(SP) argument 2
	// 0xc00003e768:   0x0000000000461ce0  // 24(SP) return value 1
	// 0xc00003e770:   0x000000c000094058  // 32(SP) return value 2   (stack size = (1 return address + 4 arguments) * 8 = 40 Byte
	// above are allocated for call f1
	// 0xc00003e778:   0x000000c00003e7d0  // 40(SP) stack of main
	// 0xc00003e780:   0x000000000042c00a
	0x0000 00000 (main.go:3)	MOVQ	$0, "".~r2+24(SP)
	0x0009 00009 (main.go:3)	MOVQ	$0, "".~r3+32(SP)
	0x0012 00018 (main.go:4)	MOVQ	"".a+8(SP), AX
	0x0017 00023 (main.go:4)	ADDQ	"".b+16(SP), AX
	0x001c 00028 (main.go:4)	MOVQ	AX, "".~r2+24(SP)
	0x0021 00033 (main.go:4)	MOVQ	"".a+8(SP), AX
	0x0026 00038 (main.go:4)	SUBQ	"".b+16(SP), AX
	0x002b 00043 (main.go:4)	MOVQ	AX, "".~r3+32(SP)
	0x0030 00048 (main.go:4)	RET
	0x0000 48 c7 44 24 18 00 00 00 00 48 c7 44 24 20 00 00  H.D$.....H.D$ ..
	0x0010 00 00 48 8b 44 24 08 48 03 44 24 10 48 89 44 24  ..H.D$.H.D$.H.D$
	0x0020 18 48 8b 44 24 08 48 2b 44 24 10 48 89 44 24 20  .H.D$.H+D$.H.D$ 
	0x0030 c3                                               .
"".main STEXT size=90 args=0x0 locals=0x28
	0x0000 00000 (main.go:8)	TEXT	"".main(SB), ABIInternal, $40-0
	0x0000 00000 (main.go:8)	MOVQ	(TLS), CX
	0x0009 00009 (main.go:8)	CMPQ	SP, 16(CX)
	0x000d 00013 (main.go:8)	PCDATA	$0, $-2
	0x000d 00013 (main.go:8)	JLS	83
	0x000f 00015 (main.go:8)	PCDATA	$0, $-1
	0x000f 00015 (main.go:8)	SUBQ	$40, SP
	// Alloc 40 Byte Size
	// 
	// SP + 32,  BP
	//
	// SP +  8,  $2
	// SP +  0,  $1
	0x0013 00019 (main.go:8)	MOVQ	BP, 32(SP)
	0x0018 00024 (main.go:8)	LEAQ	32(SP), BP
	0x001d 00029 (main.go:8)	PCDATA	$0, $-2
	0x001d 00029 (main.go:8)	PCDATA	$1, $-2
	0x001d 00029 (main.go:8)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001d 00029 (main.go:8)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001d 00029 (main.go:8)	FUNCDATA	$2, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001d 00029 (main.go:9)	PCDATA	$0, $0
	0x001d 00029 (main.go:9)	PCDATA	$1, $0
	0x001d 00029 (main.go:9)	MOVQ	$1, (SP)
	0x0025 00037 (main.go:9)	MOVQ	$2, 8(SP)
	// Before call, SP
	// sp             0xc00003e758     0xc00003e758
	0x002e 00046 (main.go:9)	CALL	"".f1(SB)
	// In f1, SP
	// sp             0xc00003e750     0xc00003e750
	
	// 0xc00003e750 is added by CALL instruction

	// (gdb) x/64xu 0xc00003e750
	// 0xc00003e750:   243     131     69      0       0       0       0       0   // in call f1
	// 0xc00003e758:   1       0       0       0       0       0       0       0   // before call
	// 0xc00003e760:   2       0       0       0       0       0       0       0
	// 0xc00003e768:   224     28      70      0       0       0       0       0
	// 0xc00003e770:   88      64      9       0       192     0       0       0   // MOVQ BP 32(SP), BP Value

	// (gdb) x/g 0xc00003e750
	// 0xc00003e750:   0x00000000004583f3  // return address, and also 0(SP)
	// 0xc00003e758:   0x0000000000000001  // argument 1
	// 0xc00003e760:   0x0000000000000002  // argument 2
	// 0xc00003e768:   0x0000000000461ce0
	// 0xc00003e770:   0x000000c000094058
	// 0xc00003e778:   0x000000c00003e7d0
	// 0xc00003e780:   0x000000000042c00a
	//
	0x0033 00051 (main.go:11)	MOVQ	$2, (SP)
	0x003b 00059 (main.go:11)	MOVQ	$3, 8(SP)
	0x0044 00068 (main.go:11)	CALL	"".f1(SB)
	0x0049 00073 (main.go:12)	MOVQ	32(SP), BP
	0x004e 00078 (main.go:12)	ADDQ	$40, SP
	0x0052 00082 (main.go:12)	RET
	0x0053 00083 (main.go:12)	NOP
	0x0053 00083 (main.go:8)	PCDATA	$1, $-1
	0x0053 00083 (main.go:8)	PCDATA	$0, $-2
	0x0053 00083 (main.go:8)	CALL	runtime.morestack_noctxt(SB)
	0x0058 00088 (main.go:8)	PCDATA	$0, $-1
	0x0058 00088 (main.go:8)	JMP	0
	0x0000 64 48 8b 0c 25 00 00 00 00 48 3b 61 10 76 44 48  dH..%....H;a.vDH
	0x0010 83 ec 28 48 89 6c 24 20 48 8d 6c 24 20 48 c7 04  ..(H.l$ H.l$ H..
	0x0020 24 01 00 00 00 48 c7 44 24 08 02 00 00 00 e8 00  $....H.D$.......
	0x0030 00 00 00 48 c7 04 24 02 00 00 00 48 c7 44 24 08  ...H..$....H.D$.
	0x0040 03 00 00 00 e8 00 00 00 00 48 8b 6c 24 20 48 83  .........H.l$ H.
	0x0050 c4 28 c3 e8 00 00 00 00 eb a6                    .(........
	rel 5+4 t=17 TLS+0
	rel 47+4 t=8 "".f1+0
	rel 69+4 t=8 "".f1+0
	rel 84+4 t=8 runtime.morestack_noctxt+0
go.cuinfo.packagename. SDWARFINFO dupok size=0
	0x0000 6d 61 69 6e                                      main
go.loc."".f1 SDWARFLOC size=0
go.info."".f1 SDWARFINFO size=82
	0x0000 03 22 22 2e 66 31 00 00 00 00 00 00 00 00 00 00  ."".f1..........
	0x0010 00 00 00 00 00 00 00 01 9c 00 00 00 00 01 0f 61  ...............a
	0x0020 00 00 03 00 00 00 00 01 9c 0f 62 00 00 03 00 00  ..........b.....
	0x0030 00 00 02 91 08 0f 7e 72 32 00 01 03 00 00 00 00  ......~r2.......
	0x0040 02 91 10 0f 7e 72 33 00 01 03 00 00 00 00 02 91  ....~r3.........
	0x0050 18 00                                            ..
	rel 0+0 t=24 type.int+0
	rel 7+8 t=1 "".f1+0
	rel 15+8 t=1 "".f1+49
	rel 25+4 t=30 gofile../home/zquxbow/go/learnGo/basic/language/assembly/init001/main.go+0
	rel 35+4 t=29 go.info.int+0
	rel 46+4 t=29 go.info.int+0
	rel 60+4 t=29 go.info.int+0
	rel 74+4 t=29 go.info.int+0
go.range."".f1 SDWARFRANGE size=0
go.debuglines."".f1 SDWARFMISC size=15
	0x0000 04 02 11 06 69 06 6a 06 41 04 01 03 7d 06 01     ....i.j.A...}..
go.loc."".main SDWARFLOC size=0
go.info."".main SDWARFINFO size=33
	0x0000 03 22 22 2e 6d 61 69 6e 00 00 00 00 00 00 00 00  ."".main........
	0x0010 00 00 00 00 00 00 00 00 00 01 9c 00 00 00 00 01  ................
	0x0020 00                                               .
	rel 9+8 t=1 "".main+0
	rel 17+8 t=1 "".main+90
	rel 27+4 t=30 gofile../home/zquxbow/go/learnGo/basic/language/assembly/init001/main.go+0
go.range."".main SDWARFRANGE size=0
go.debuglines."".main SDWARFMISC size=23
	0x0000 04 02 03 02 14 0a a5 9c 06 5f 06 9d 06 5f 06 69  ........._..._.i
	0x0010 42 6f 04 01 03 79 01                             Bo...y.
""..inittask SNOPTRDATA size=24
	0x0000 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0010 00 00 00 00 00 00 00 00                          ........
gclocals·33cdeccccebe80329f1fdbee7f5874cb SRODATA dupok size=8
	0x0000 01 00 00 00 00 00 00 00                          ........
