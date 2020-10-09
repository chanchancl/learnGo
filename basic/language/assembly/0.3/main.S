"".main STEXT size=517 args=0x0 locals=0xd0
	0x0000 00000 (main.go:5)	TEXT	"".main(SB), ABIInternal, $208-0
	0x0000 00000 (main.go:5)	MOVQ	(TLS), CX
	0x0009 00009 (main.go:5)	LEAQ	-80(SP), AX
	0x000e 00014 (main.go:5)	CMPQ	AX, 16(CX)
	0x0012 00018 (main.go:5)	PCDATA	$0, $-2
	0x0012 00018 (main.go:5)	JLS	503
	0x0018 00024 (main.go:5)	PCDATA	$0, $-1
	0x0018 00024 (main.go:5)	SUBQ	$208, SP
	0x001f 00031 (main.go:5)	MOVQ	BP, 200(SP)
	0x0027 00039 (main.go:5)	LEAQ	200(SP), BP
	0x002f 00047 (main.go:5)	FUNCDATA	$0, gclocals·69c1753bd5f81501d95132d08af04464(SB)
	0x002f 00047 (main.go:5)	FUNCDATA	$1, gclocals·47999a83046ca050150d29c4e3145d66(SB)
	0x002f 00047 (main.go:5)	FUNCDATA	$3, "".main.stkobj(SB)
	0x002f 00047 (main.go:6)	XORPS	X0, X0
	0x0032 00050 (main.go:6)	MOVUPS	X0, ""..autotmp_2+112(SP)
	0x0037 00055 (main.go:6)	LEAQ	""..autotmp_2+112(SP), AX
	0x003c 00060 (main.go:6)	MOVQ	AX, ""..autotmp_4+104(SP)
	0x0041 00065 (main.go:6)	TESTB	AL, (AX)
	0x0043 00067 (main.go:6)	LEAQ	type.float64(SB), CX
	0x004a 00074 (main.go:6)	MOVQ	CX, ""..autotmp_2+112(SP)
	0x004f 00079 (main.go:6)	LEAQ	""..stmp_0(SB), CX
	0x0056 00086 (main.go:6)	MOVQ	CX, ""..autotmp_2+120(SP)
	0x005b 00091 (main.go:6)	TESTB	AL, (AX)
	0x005d 00093 (main.go:6)	JMP	95
	0x005f 00095 (main.go:6)	MOVQ	AX, ""..autotmp_3+176(SP)
	0x0067 00103 (main.go:6)	MOVQ	$1, ""..autotmp_3+184(SP)
	0x0073 00115 (main.go:6)	MOVQ	$1, ""..autotmp_3+192(SP)
	0x007f 00127 (main.go:6)	MOVQ	AX, (SP)
	0x0083 00131 (main.go:6)	MOVQ	$1, 8(SP)
	0x008c 00140 (main.go:6)	MOVQ	$1, 16(SP)
	0x0095 00149 (main.go:6)	PCDATA	$1, $0
	0x0095 00149 (main.go:6)	CALL	fmt.Println(SB)
	0x009a 00154 (main.go:7)	MOVSD	$f64.3fb999999999999a(SB), X0
	0x00a2 00162 (main.go:7)	MOVSD	X0, "".a+72(SP)
	0x00a8 00168 (main.go:8)	MOVSD	$f64.3fc999999999999a(SB), X0
	0x00b0 00176 (main.go:8)	MOVSD	X0, "".b+64(SP)
	0x00b6 00182 (main.go:9)	XORPS	X0, X0
	0x00b9 00185 (main.go:9)	MOVUPS	X0, ""..autotmp_2+112(SP)
	0x00be 00190 (main.go:9)	LEAQ	""..autotmp_2+112(SP), AX
	0x00c3 00195 (main.go:9)	MOVQ	AX, ""..autotmp_6+96(SP)
	0x00c8 00200 (main.go:9)	MOVSD	"".a+72(SP), X0
	0x00ce 00206 (main.go:9)	ADDSD	"".b+64(SP), X0
	0x00d4 00212 (main.go:9)	MOVSD	X0, (SP)
	0x00d9 00217 (main.go:9)	PCDATA	$1, $1
	0x00d9 00217 (main.go:9)	CALL	runtime.convT64(SB)
	0x00de 00222 (main.go:9)	MOVQ	8(SP), AX
	0x00e3 00227 (main.go:9)	MOVQ	AX, ""..autotmp_7+88(SP)
	0x00e8 00232 (main.go:9)	MOVQ	""..autotmp_6+96(SP), CX
	0x00ed 00237 (main.go:9)	TESTB	AL, (CX)
	0x00ef 00239 (main.go:9)	LEAQ	type.float64(SB), DX
	0x00f6 00246 (main.go:9)	MOVQ	DX, (CX)
	0x00f9 00249 (main.go:9)	LEAQ	8(CX), DI
	0x00fd 00253 (main.go:9)	PCDATA	$0, $-2
	0x00fd 00253 (main.go:9)	CMPL	runtime.writeBarrier(SB), $0
	0x0104 00260 (main.go:9)	JEQ	267
	0x0106 00262 (main.go:9)	JMP	493
	0x010b 00267 (main.go:9)	MOVQ	AX, 8(CX)
	0x010f 00271 (main.go:9)	JMP	273
	0x0111 00273 (main.go:9)	PCDATA	$0, $-1
	0x0111 00273 (main.go:9)	MOVQ	""..autotmp_6+96(SP), AX
	0x0116 00278 (main.go:9)	TESTB	AL, (AX)
	0x0118 00280 (main.go:9)	JMP	282
	0x011a 00282 (main.go:9)	MOVQ	AX, ""..autotmp_5+152(SP)
	0x0122 00290 (main.go:9)	MOVQ	$1, ""..autotmp_5+160(SP)
	0x012e 00302 (main.go:9)	MOVQ	$1, ""..autotmp_5+168(SP)
	0x013a 00314 (main.go:9)	MOVQ	AX, (SP)
	0x013e 00318 (main.go:9)	MOVQ	$1, 8(SP)
	0x0147 00327 (main.go:9)	MOVQ	$1, 16(SP)
	0x0150 00336 (main.go:9)	PCDATA	$1, $0
	0x0150 00336 (main.go:9)	CALL	fmt.Println(SB)
	0x0155 00341 (main.go:10)	XORPS	X0, X0
	0x0158 00344 (main.go:10)	MOVUPS	X0, ""..autotmp_2+112(SP)
	0x015d 00349 (main.go:10)	LEAQ	""..autotmp_2+112(SP), AX
	0x0162 00354 (main.go:10)	MOVQ	AX, ""..autotmp_9+80(SP)
	0x0167 00359 (main.go:10)	TESTB	AL, (AX)
	0x0169 00361 (main.go:10)	LEAQ	type.float64(SB), CX
	0x0170 00368 (main.go:10)	MOVQ	CX, ""..autotmp_2+112(SP)
	0x0175 00373 (main.go:10)	LEAQ	""..stmp_1(SB), CX
	0x017c 00380 (main.go:10)	MOVQ	CX, ""..autotmp_2+120(SP)
	0x0181 00385 (main.go:10)	TESTB	AL, (AX)
	0x0183 00387 (main.go:10)	JMP	389
	0x0185 00389 (main.go:10)	MOVQ	AX, ""..autotmp_8+128(SP)
	0x018d 00397 (main.go:10)	MOVQ	$1, ""..autotmp_8+136(SP)
	0x0199 00409 (main.go:10)	MOVQ	$1, ""..autotmp_8+144(SP)
	0x01a5 00421 (main.go:10)	LEAQ	go.string."%.54f\n"(SB), AX
	0x01ac 00428 (main.go:10)	MOVQ	AX, (SP)
	0x01b0 00432 (main.go:10)	MOVQ	$6, 8(SP)
	0x01b9 00441 (main.go:10)	MOVQ	""..autotmp_8+128(SP), AX
	0x01c1 00449 (main.go:10)	MOVQ	AX, 16(SP)
	0x01c6 00454 (main.go:10)	MOVQ	$1, 24(SP)
	0x01cf 00463 (main.go:10)	MOVQ	$1, 32(SP)
	0x01d8 00472 (main.go:10)	CALL	fmt.Printf(SB)
	0x01dd 00477 (main.go:11)	MOVQ	200(SP), BP
	0x01e5 00485 (main.go:11)	ADDQ	$208, SP
	0x01ec 00492 (main.go:11)	RET
	0x01ed 00493 (main.go:9)	PCDATA	$0, $-2
	0x01ed 00493 (main.go:9)	CALL	runtime.gcWriteBarrier(SB)
	0x01f2 00498 (main.go:9)	JMP	273
	0x01f7 00503 (main.go:9)	NOP
	0x01f7 00503 (main.go:5)	PCDATA	$1, $-1
	0x01f7 00503 (main.go:5)	PCDATA	$0, $-2
	0x01f7 00503 (main.go:5)	CALL	runtime.morestack_noctxt(SB)
	0x01fc 00508 (main.go:5)	PCDATA	$0, $-1
	0x01fc 00508 (main.go:5)	NOP
	0x0200 00512 (main.go:5)	JMP	0
	0x0000 64 48 8b 0c 25 00 00 00 00 48 8d 44 24 b0 48 3b  dH..%....H.D$.H;
	0x0010 41 10 0f 86 df 01 00 00 48 81 ec d0 00 00 00 48  A.......H......H
	0x0020 89 ac 24 c8 00 00 00 48 8d ac 24 c8 00 00 00 0f  ..$....H..$.....
	0x0030 57 c0 0f 11 44 24 70 48 8d 44 24 70 48 89 44 24  W...D$pH.D$pH.D$
	0x0040 68 84 00 48 8d 0d 00 00 00 00 48 89 4c 24 70 48  h..H......H.L$pH
	0x0050 8d 0d 00 00 00 00 48 89 4c 24 78 84 00 eb 00 48  ......H.L$x....H
	0x0060 89 84 24 b0 00 00 00 48 c7 84 24 b8 00 00 00 01  ..$....H..$.....
	0x0070 00 00 00 48 c7 84 24 c0 00 00 00 01 00 00 00 48  ...H..$........H
	0x0080 89 04 24 48 c7 44 24 08 01 00 00 00 48 c7 44 24  ..$H.D$.....H.D$
	0x0090 10 01 00 00 00 e8 00 00 00 00 f2 0f 10 05 00 00  ................
	0x00a0 00 00 f2 0f 11 44 24 48 f2 0f 10 05 00 00 00 00  .....D$H........
	0x00b0 f2 0f 11 44 24 40 0f 57 c0 0f 11 44 24 70 48 8d  ...D$@.W...D$pH.
	0x00c0 44 24 70 48 89 44 24 60 f2 0f 10 44 24 48 f2 0f  D$pH.D$`...D$H..
	0x00d0 58 44 24 40 f2 0f 11 04 24 e8 00 00 00 00 48 8b  XD$@....$.....H.
	0x00e0 44 24 08 48 89 44 24 58 48 8b 4c 24 60 84 01 48  D$.H.D$XH.L$`..H
	0x00f0 8d 15 00 00 00 00 48 89 11 48 8d 79 08 83 3d 00  ......H..H.y..=.
	0x0100 00 00 00 00 74 05 e9 e2 00 00 00 48 89 41 08 eb  ....t......H.A..
	0x0110 00 48 8b 44 24 60 84 00 eb 00 48 89 84 24 98 00  .H.D$`....H..$..
	0x0120 00 00 48 c7 84 24 a0 00 00 00 01 00 00 00 48 c7  ..H..$........H.
	0x0130 84 24 a8 00 00 00 01 00 00 00 48 89 04 24 48 c7  .$........H..$H.
	0x0140 44 24 08 01 00 00 00 48 c7 44 24 10 01 00 00 00  D$.....H.D$.....
	0x0150 e8 00 00 00 00 0f 57 c0 0f 11 44 24 70 48 8d 44  ......W...D$pH.D
	0x0160 24 70 48 89 44 24 50 84 00 48 8d 0d 00 00 00 00  $pH.D$P..H......
	0x0170 48 89 4c 24 70 48 8d 0d 00 00 00 00 48 89 4c 24  H.L$pH......H.L$
	0x0180 78 84 00 eb 00 48 89 84 24 80 00 00 00 48 c7 84  x....H..$....H..
	0x0190 24 88 00 00 00 01 00 00 00 48 c7 84 24 90 00 00  $........H..$...
	0x01a0 00 01 00 00 00 48 8d 05 00 00 00 00 48 89 04 24  .....H......H..$
	0x01b0 48 c7 44 24 08 06 00 00 00 48 8b 84 24 80 00 00  H.D$.....H..$...
	0x01c0 00 48 89 44 24 10 48 c7 44 24 18 01 00 00 00 48  .H.D$.H.D$.....H
	0x01d0 c7 44 24 20 01 00 00 00 e8 00 00 00 00 48 8b ac  .D$ .........H..
	0x01e0 24 c8 00 00 00 48 81 c4 d0 00 00 00 c3 e8 00 00  $....H..........
	0x01f0 00 00 e9 1a ff ff ff e8 00 00 00 00 0f 1f 40 00  ..............@.
	0x0200 e9 fb fd ff ff                                   .....
	rel 5+4 t=17 TLS+0
	rel 70+4 t=16 type.float64+0
	rel 82+4 t=16 ""..stmp_0+0
	rel 150+4 t=8 fmt.Println+0
	rel 158+4 t=16 $f64.3fb999999999999a+0
	rel 172+4 t=16 $f64.3fc999999999999a+0
	rel 218+4 t=8 runtime.convT64+0
	rel 242+4 t=16 type.float64+0
	rel 255+4 t=16 runtime.writeBarrier+-1
	rel 337+4 t=8 fmt.Println+0
	rel 364+4 t=16 type.float64+0
	rel 376+4 t=16 ""..stmp_1+0
	rel 424+4 t=16 go.string."%.54f\n"+0
	rel 473+4 t=8 fmt.Printf+0
	rel 494+4 t=8 runtime.gcWriteBarrier+0
	rel 504+4 t=8 runtime.morestack_noctxt+0
go.cuinfo.packagename. SDWARFINFO dupok size=0
	0x0000 6d 61 69 6e                                      main
go.string."%.54f\n" SRODATA dupok size=6
	0x0000 25 2e 35 34 66 0a                                %.54f.
runtime.nilinterequal·f SRODATA dupok size=8
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 runtime.nilinterequal+0
runtime.memequal64·f SRODATA dupok size=8
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 runtime.memequal64+0
runtime.gcbits.01 SRODATA dupok size=1
	0x0000 01                                               .
type..namedata.*interface {}- SRODATA dupok size=16
	0x0000 00 00 0d 2a 69 6e 74 65 72 66 61 63 65 20 7b 7d  ...*interface {}
type.*interface {} SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 4f 0f 96 9d 08 08 08 36 00 00 00 00 00 00 00 00  O......6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64·f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*interface {}-+0
	rel 48+8 t=1 type.interface {}+0
runtime.gcbits.02 SRODATA dupok size=1
	0x0000 02                                               .
type.interface {} SRODATA dupok size=80
	0x0000 10 00 00 00 00 00 00 00 10 00 00 00 00 00 00 00  ................
	0x0010 e7 57 a0 18 02 08 08 14 00 00 00 00 00 00 00 00  .W..............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	rel 24+8 t=1 runtime.nilinterequal·f+0
	rel 32+8 t=1 runtime.gcbits.02+0
	rel 40+4 t=5 type..namedata.*interface {}-+0
	rel 44+4 t=6 type.*interface {}+0
	rel 56+8 t=1 type.interface {}+80
type..namedata.*[]interface {}- SRODATA dupok size=18
	0x0000 00 00 0f 2a 5b 5d 69 6e 74 65 72 66 61 63 65 20  ...*[]interface 
	0x0010 7b 7d                                            {}
type.*[]interface {} SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 f3 04 9a e7 08 08 08 36 00 00 00 00 00 00 00 00  .......6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64·f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[]interface {}-+0
	rel 48+8 t=1 type.[]interface {}+0
type.[]interface {} SRODATA dupok size=56
	0x0000 18 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 70 93 ea 2f 02 08 08 17 00 00 00 00 00 00 00 00  p../............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[]interface {}-+0
	rel 44+4 t=6 type.*[]interface {}+0
	rel 48+8 t=1 type.interface {}+0
type..namedata.*[1]interface {}- SRODATA dupok size=19
	0x0000 00 00 10 2a 5b 31 5d 69 6e 74 65 72 66 61 63 65  ...*[1]interface
	0x0010 20 7b 7d                                          {}
type.*[1]interface {} SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 bf 03 a8 35 08 08 08 36 00 00 00 00 00 00 00 00  ...5...6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64·f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[1]interface {}-+0
	rel 48+8 t=1 type.[1]interface {}+0
type.[1]interface {} SRODATA dupok size=72
	0x0000 10 00 00 00 00 00 00 00 10 00 00 00 00 00 00 00  ................
	0x0010 50 91 5b fa 02 08 08 11 00 00 00 00 00 00 00 00  P.[.............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 01 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.nilinterequal·f+0
	rel 32+8 t=1 runtime.gcbits.02+0
	rel 40+4 t=5 type..namedata.*[1]interface {}-+0
	rel 44+4 t=6 type.*[1]interface {}+0
	rel 48+8 t=1 type.interface {}+0
	rel 56+8 t=1 type.[]interface {}+0
""..inittask SNOPTRDATA size=32
	0x0000 00 00 00 00 00 00 00 00 01 00 00 00 00 00 00 00  ................
	0x0010 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	rel 24+8 t=1 fmt..inittask+0
""..stmp_0 SRODATA size=8
	0x0000 33 33 33 33 33 33 d3 3f                          333333.?
""..stmp_1 SRODATA size=8
	0x0000 33 33 33 33 33 33 d3 3f                          333333.?
type..importpath.fmt. SRODATA dupok size=6
	0x0000 00 00 03 66 6d 74                                ...fmt
gclocals·69c1753bd5f81501d95132d08af04464 SRODATA dupok size=8
	0x0000 02 00 00 00 00 00 00 00                          ........
gclocals·47999a83046ca050150d29c4e3145d66 SRODATA dupok size=12
	0x0000 02 00 00 00 0f 00 00 00 00 00 04 00              ............
"".main.stkobj SRODATA size=24
	0x0000 01 00 00 00 00 00 00 00 a8 ff ff ff ff ff ff ff  ................
	0x0010 00 00 00 00 00 00 00 00                          ........
	rel 16+8 t=1 type.[1]interface {}+0
