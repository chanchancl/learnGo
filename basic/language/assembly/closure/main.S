"".withoutArg STEXT size=184 args=0x8 locals=0x28
	0x0000 00000 (main.go:5)	TEXT	"".withoutArg(SB), ABIInternal, $40-8
	0x0000 00000 (main.go:5)	MOVQ	(TLS), CX
	0x0009 00009 (main.go:5)	CMPQ	SP, 16(CX)
	0x000d 00013 (main.go:5)	PCDATA	$0, $-2
	0x000d 00013 (main.go:5)	JLS	174
	0x0013 00019 (main.go:5)	PCDATA	$0, $-1
	0x0013 00019 (main.go:5)	SUBQ	$40, SP
	0x0017 00023 (main.go:5)	MOVQ	BP, 32(SP)
	0x001c 00028 (main.go:5)	LEAQ	32(SP), BP
	0x0021 00033 (main.go:5)	PCDATA	$0, $-2
	0x0021 00033 (main.go:5)	PCDATA	$1, $-2
	0x0021 00033 (main.go:5)	FUNCDATA	$0, gclocals·272bd7e3f0f711ce8c35e9a618d8e620(SB)
	0x0021 00033 (main.go:5)	FUNCDATA	$1, gclocals·cc5cad6002ade825472fe3ad2912087d(SB)
	0x0021 00033 (main.go:5)	FUNCDATA	$2, gclocals·1e847dfe98917b7bcc7220b0d92d534f(SB)
	0x0021 00033 (main.go:5)	PCDATA	$0, $0
	0x0021 00033 (main.go:5)	PCDATA	$1, $0
	0x0021 00033 (main.go:5)	MOVQ	$0, "".~r0+48(SP)
	0x002a 00042 (main.go:6)	PCDATA	$0, $1
	0x002a 00042 (main.go:6)	LEAQ	type.int(SB), AX
	0x0031 00049 (main.go:6)	PCDATA	$0, $0
	0x0031 00049 (main.go:6)	MOVQ	AX, (SP)
	0x0035 00053 (main.go:6)	CALL	runtime.newobject(SB)
	0x003a 00058 (main.go:6)	PCDATA	$0, $1
	0x003a 00058 (main.go:6)	MOVQ	8(SP), AX
	0x003f 00063 (main.go:6)	PCDATA	$1, $1
	0x003f 00063 (main.go:6)	MOVQ	AX, "".&i+24(SP)
	0x0044 00068 (main.go:6)	PCDATA	$0, $0
	0x0044 00068 (main.go:6)	MOVQ	$0, (AX)
	0x004b 00075 (main.go:7)	PCDATA	$0, $1
	0x004b 00075 (main.go:7)	LEAQ	type.noalg.struct { F uintptr; "".i *int }(SB), AX
	0x0052 00082 (main.go:7)	PCDATA	$0, $0
	0x0052 00082 (main.go:7)	MOVQ	AX, (SP)
	0x0056 00086 (main.go:7)	CALL	runtime.newobject(SB)
	0x005b 00091 (main.go:7)	PCDATA	$0, $1
	0x005b 00091 (main.go:7)	MOVQ	8(SP), AX
	0x0060 00096 (main.go:7)	PCDATA	$1, $2
	0x0060 00096 (main.go:7)	MOVQ	AX, ""..autotmp_4+16(SP)
	0x0065 00101 (main.go:7)	LEAQ	"".withoutArg.func1(SB), CX
	0x006c 00108 (main.go:7)	PCDATA	$0, $0
	0x006c 00108 (main.go:7)	MOVQ	CX, (AX)
	0x006f 00111 (main.go:7)	PCDATA	$0, $1
	0x006f 00111 (main.go:7)	MOVQ	""..autotmp_4+16(SP), AX
	0x0074 00116 (main.go:7)	TESTB	AL, (AX)
	0x0076 00118 (main.go:7)	PCDATA	$0, $2
	0x0076 00118 (main.go:7)	PCDATA	$1, $3
	0x0076 00118 (main.go:7)	MOVQ	"".&i+24(SP), CX
	0x007b 00123 (main.go:7)	PCDATA	$0, $3
	0x007b 00123 (main.go:7)	LEAQ	8(AX), DI
	0x007f 00127 (main.go:7)	PCDATA	$0, $-2
	0x007f 00127 (main.go:7)	PCDATA	$1, $-2
	0x007f 00127 (main.go:7)	CMPL	runtime.writeBarrier(SB), $0
	0x0086 00134 (main.go:7)	JEQ	138
	0x0088 00136 (main.go:7)	JMP	164
	0x008a 00138 (main.go:7)	MOVQ	CX, 8(AX)
	0x008e 00142 (main.go:7)	JMP	144
	0x0090 00144 (main.go:7)	PCDATA	$0, $1
	0x0090 00144 (main.go:7)	PCDATA	$1, $0
	0x0090 00144 (main.go:7)	MOVQ	""..autotmp_4+16(SP), AX
	0x0095 00149 (main.go:7)	PCDATA	$0, $0
	0x0095 00149 (main.go:7)	PCDATA	$1, $4
	0x0095 00149 (main.go:7)	MOVQ	AX, "".~r0+48(SP)
	0x009a 00154 (main.go:7)	MOVQ	32(SP), BP
	0x009f 00159 (main.go:7)	ADDQ	$40, SP
	0x00a3 00163 (main.go:7)	RET
	0x00a4 00164 (main.go:7)	PCDATA	$0, $-2
	0x00a4 00164 (main.go:7)	PCDATA	$1, $-2
	0x00a4 00164 (main.go:7)	MOVQ	CX, AX
	0x00a7 00167 (main.go:7)	CALL	runtime.gcWriteBarrier(SB)
	0x00ac 00172 (main.go:7)	JMP	144
	0x00ae 00174 (main.go:7)	NOP
	0x00ae 00174 (main.go:5)	PCDATA	$1, $-1
	0x00ae 00174 (main.go:5)	PCDATA	$0, $-2
	0x00ae 00174 (main.go:5)	CALL	runtime.morestack_noctxt(SB)
	0x00b3 00179 (main.go:5)	PCDATA	$0, $-1
	0x00b3 00179 (main.go:5)	JMP	0
	0x0000 64 48 8b 0c 25 00 00 00 00 48 3b 61 10 0f 86 9b  dH..%....H;a....
	0x0010 00 00 00 48 83 ec 28 48 89 6c 24 20 48 8d 6c 24  ...H..(H.l$ H.l$
	0x0020 20 48 c7 44 24 30 00 00 00 00 48 8d 05 00 00 00   H.D$0....H.....
	0x0030 00 48 89 04 24 e8 00 00 00 00 48 8b 44 24 08 48  .H..$.....H.D$.H
	0x0040 89 44 24 18 48 c7 00 00 00 00 00 48 8d 05 00 00  .D$.H......H....
	0x0050 00 00 48 89 04 24 e8 00 00 00 00 48 8b 44 24 08  ..H..$.....H.D$.
	0x0060 48 89 44 24 10 48 8d 0d 00 00 00 00 48 89 08 48  H.D$.H......H..H
	0x0070 8b 44 24 10 84 00 48 8b 4c 24 18 48 8d 78 08 83  .D$...H.L$.H.x..
	0x0080 3d 00 00 00 00 00 74 02 eb 1a 48 89 48 08 eb 00  =.....t...H.H...
	0x0090 48 8b 44 24 10 48 89 44 24 30 48 8b 6c 24 20 48  H.D$.H.D$0H.l$ H
	0x00a0 83 c4 28 c3 48 89 c8 e8 00 00 00 00 eb e2 e8 00  ..(.H...........
	0x00b0 00 00 00 e9 48 ff ff ff                          ....H...
	rel 5+4 t=17 TLS+0
	rel 45+4 t=16 type.int+0
	rel 54+4 t=8 runtime.newobject+0
	rel 78+4 t=16 type.noalg.struct { F uintptr; "".i *int }+0
	rel 87+4 t=8 runtime.newobject+0
	rel 104+4 t=16 "".withoutArg.func1+0
	rel 129+4 t=16 runtime.writeBarrier+-1
	rel 168+4 t=8 runtime.gcWriteBarrier+0
	rel 175+4 t=8 runtime.morestack_noctxt+0
"".main STEXT size=433 args=0x0 locals=0xc0
	0x0000 00000 (main.go:13)	TEXT	"".main(SB), ABIInternal, $192-0
	0x0000 00000 (main.go:13)	MOVQ	(TLS), CX
	0x0009 00009 (main.go:13)	LEAQ	-64(SP), AX
	0x000e 00014 (main.go:13)	CMPQ	AX, 16(CX)
	0x0012 00018 (main.go:13)	PCDATA	$0, $-2
	0x0012 00018 (main.go:13)	JLS	423
	0x0018 00024 (main.go:13)	PCDATA	$0, $-1
	0x0018 00024 (main.go:13)	SUBQ	$192, SP
	0x001f 00031 (main.go:13)	MOVQ	BP, 184(SP)
	0x0027 00039 (main.go:13)	LEAQ	184(SP), BP
	0x002f 00047 (main.go:13)	PCDATA	$0, $-2
	0x002f 00047 (main.go:13)	PCDATA	$1, $-2
	0x002f 00047 (main.go:13)	FUNCDATA	$0, gclocals·7be4bbacbfdb05fb3044e36c22b41e8b(SB)
	0x002f 00047 (main.go:13)	FUNCDATA	$1, gclocals·9ee33c85bd820cba75d4d4aac78994e7(SB)
	0x002f 00047 (main.go:13)	FUNCDATA	$2, gclocals·74bf2e2f0b69fa7db29b8b87c766370b(SB)
	0x002f 00047 (main.go:13)	FUNCDATA	$3, "".main.stkobj(SB)
	0x002f 00047 (main.go:14)	PCDATA	$0, $0
	0x002f 00047 (main.go:14)	PCDATA	$1, $0
	0x002f 00047 (main.go:14)	CALL	"".withoutArg(SB)
	0x0034 00052 (main.go:14)	PCDATA	$0, $1
	0x0034 00052 (main.go:14)	MOVQ	(SP), DX
	0x0038 00056 (main.go:14)	PCDATA	$1, $1
	0x0038 00056 (main.go:14)	MOVQ	DX, "".f1+72(SP)
	0x003d 00061 (main.go:15)	MOVQ	(DX), AX
	0x0040 00064 (main.go:15)	PCDATA	$0, $0
	0x0040 00064 (main.go:15)	CALL	AX
	0x0042 00066 (main.go:15)	MOVQ	(SP), AX
	0x0046 00070 (main.go:15)	MOVQ	AX, "".a+64(SP)
	0x004b 00075 (main.go:16)	PCDATA	$0, $1
	0x004b 00075 (main.go:16)	MOVQ	"".f1+72(SP), DX
	0x0050 00080 (main.go:16)	MOVQ	(DX), AX
	0x0053 00083 (main.go:16)	PCDATA	$0, $0
	0x0053 00083 (main.go:16)	CALL	AX
	0x0055 00085 (main.go:16)	MOVQ	(SP), AX
	0x0059 00089 (main.go:16)	MOVQ	AX, "".b+56(SP)
	0x005e 00094 (main.go:17)	PCDATA	$0, $1
	0x005e 00094 (main.go:17)	PCDATA	$1, $0
	0x005e 00094 (main.go:17)	MOVQ	"".f1+72(SP), DX
	0x0063 00099 (main.go:17)	MOVQ	(DX), AX
	0x0066 00102 (main.go:17)	PCDATA	$0, $0
	0x0066 00102 (main.go:17)	CALL	AX
	0x0068 00104 (main.go:17)	MOVQ	(SP), AX
	0x006c 00108 (main.go:17)	MOVQ	AX, "".c+48(SP)
	0x0071 00113 (main.go:19)	MOVQ	"".a+64(SP), AX
	0x0076 00118 (main.go:19)	MOVQ	AX, (SP)
	0x007a 00122 (main.go:19)	CALL	runtime.convT64(SB)
	0x007f 00127 (main.go:19)	PCDATA	$0, $2
	0x007f 00127 (main.go:19)	MOVQ	8(SP), AX
	0x0084 00132 (main.go:19)	PCDATA	$0, $0
	0x0084 00132 (main.go:19)	PCDATA	$1, $2
	0x0084 00132 (main.go:19)	MOVQ	AX, ""..autotmp_5+104(SP)
	0x0089 00137 (main.go:19)	MOVQ	"".b+56(SP), AX
	0x008e 00142 (main.go:19)	MOVQ	AX, (SP)
	0x0092 00146 (main.go:19)	CALL	runtime.convT64(SB)
	0x0097 00151 (main.go:19)	PCDATA	$0, $2
	0x0097 00151 (main.go:19)	MOVQ	8(SP), AX
	0x009c 00156 (main.go:19)	PCDATA	$0, $0
	0x009c 00156 (main.go:19)	PCDATA	$1, $3
	0x009c 00156 (main.go:19)	MOVQ	AX, ""..autotmp_6+96(SP)
	0x00a1 00161 (main.go:19)	MOVQ	"".c+48(SP), AX
	0x00a6 00166 (main.go:19)	MOVQ	AX, (SP)
	0x00aa 00170 (main.go:19)	CALL	runtime.convT64(SB)
	0x00af 00175 (main.go:19)	PCDATA	$0, $2
	0x00af 00175 (main.go:19)	MOVQ	8(SP), AX
	0x00b4 00180 (main.go:19)	PCDATA	$0, $0
	0x00b4 00180 (main.go:19)	PCDATA	$1, $4
	0x00b4 00180 (main.go:19)	MOVQ	AX, ""..autotmp_7+88(SP)
	0x00b9 00185 (main.go:19)	PCDATA	$1, $5
	0x00b9 00185 (main.go:19)	XORPS	X0, X0
	0x00bc 00188 (main.go:19)	MOVUPS	X0, ""..autotmp_4+136(SP)
	0x00c4 00196 (main.go:19)	XORPS	X0, X0
	0x00c7 00199 (main.go:19)	MOVUPS	X0, ""..autotmp_4+152(SP)
	0x00cf 00207 (main.go:19)	XORPS	X0, X0
	0x00d2 00210 (main.go:19)	MOVUPS	X0, ""..autotmp_4+168(SP)
	0x00da 00218 (main.go:19)	PCDATA	$0, $2
	0x00da 00218 (main.go:19)	PCDATA	$1, $4
	0x00da 00218 (main.go:19)	LEAQ	""..autotmp_4+136(SP), AX
	0x00e2 00226 (main.go:19)	PCDATA	$1, $6
	0x00e2 00226 (main.go:19)	MOVQ	AX, ""..autotmp_9+80(SP)
	0x00e7 00231 (main.go:19)	TESTB	AL, (AX)
	0x00e9 00233 (main.go:19)	PCDATA	$0, $3
	0x00e9 00233 (main.go:19)	PCDATA	$1, $7
	0x00e9 00233 (main.go:19)	MOVQ	""..autotmp_5+104(SP), CX
	0x00ee 00238 (main.go:19)	PCDATA	$0, $4
	0x00ee 00238 (main.go:19)	LEAQ	type.int(SB), DX
	0x00f5 00245 (main.go:19)	PCDATA	$0, $3
	0x00f5 00245 (main.go:19)	MOVQ	DX, ""..autotmp_4+136(SP)
	0x00fd 00253 (main.go:19)	PCDATA	$0, $2
	0x00fd 00253 (main.go:19)	MOVQ	CX, ""..autotmp_4+144(SP)
	0x0105 00261 (main.go:19)	PCDATA	$0, $0
	0x0105 00261 (main.go:19)	TESTB	AL, (AX)
	0x0107 00263 (main.go:19)	PCDATA	$0, $2
	0x0107 00263 (main.go:19)	PCDATA	$1, $8
	0x0107 00263 (main.go:19)	MOVQ	""..autotmp_6+96(SP), AX
	0x010c 00268 (main.go:19)	PCDATA	$0, $3
	0x010c 00268 (main.go:19)	LEAQ	type.int(SB), CX
	0x0113 00275 (main.go:19)	PCDATA	$0, $2
	0x0113 00275 (main.go:19)	MOVQ	CX, ""..autotmp_4+152(SP)
	0x011b 00283 (main.go:19)	PCDATA	$0, $0
	0x011b 00283 (main.go:19)	MOVQ	AX, ""..autotmp_4+160(SP)
	0x0123 00291 (main.go:19)	PCDATA	$0, $2
	0x0123 00291 (main.go:19)	MOVQ	""..autotmp_9+80(SP), AX
	0x0128 00296 (main.go:19)	TESTB	AL, (AX)
	0x012a 00298 (main.go:19)	PCDATA	$0, $3
	0x012a 00298 (main.go:19)	PCDATA	$1, $9
	0x012a 00298 (main.go:19)	MOVQ	""..autotmp_7+88(SP), CX
	0x012f 00303 (main.go:19)	PCDATA	$0, $4
	0x012f 00303 (main.go:19)	LEAQ	type.int(SB), DX
	0x0136 00310 (main.go:19)	PCDATA	$0, $3
	0x0136 00310 (main.go:19)	MOVQ	DX, 32(AX)
	0x013a 00314 (main.go:19)	PCDATA	$0, $5
	0x013a 00314 (main.go:19)	LEAQ	40(AX), DI
	0x013e 00318 (main.go:19)	PCDATA	$0, $-2
	0x013e 00318 (main.go:19)	PCDATA	$1, $-2
	0x013e 00318 (main.go:19)	CMPL	runtime.writeBarrier(SB), $0
	0x0145 00325 (main.go:19)	JEQ	329
	0x0147 00327 (main.go:19)	JMP	413
	0x0149 00329 (main.go:19)	MOVQ	CX, 40(AX)
	0x014d 00333 (main.go:19)	JMP	335
	0x014f 00335 (main.go:19)	PCDATA	$0, $2
	0x014f 00335 (main.go:19)	PCDATA	$1, $0
	0x014f 00335 (main.go:19)	MOVQ	""..autotmp_9+80(SP), AX
	0x0154 00340 (main.go:19)	TESTB	AL, (AX)
	0x0156 00342 (main.go:19)	JMP	344
	0x0158 00344 (main.go:19)	MOVQ	AX, ""..autotmp_8+112(SP)
	0x015d 00349 (main.go:19)	MOVQ	$3, ""..autotmp_8+120(SP)
	0x0166 00358 (main.go:19)	MOVQ	$3, ""..autotmp_8+128(SP)
	0x0172 00370 (main.go:19)	PCDATA	$0, $0
	0x0172 00370 (main.go:19)	MOVQ	AX, (SP)
	0x0176 00374 (main.go:19)	MOVQ	$3, 8(SP)
	0x017f 00383 (main.go:19)	MOVQ	$3, 16(SP)
	0x0188 00392 (main.go:19)	CALL	fmt.Println(SB)
	0x018d 00397 (main.go:20)	MOVQ	184(SP), BP
	0x0195 00405 (main.go:20)	ADDQ	$192, SP
	0x019c 00412 (main.go:20)	RET
	0x019d 00413 (main.go:19)	PCDATA	$0, $-2
	0x019d 00413 (main.go:19)	PCDATA	$1, $-2
	0x019d 00413 (main.go:19)	MOVQ	CX, AX
	0x01a0 00416 (main.go:19)	CALL	runtime.gcWriteBarrier(SB)
	0x01a5 00421 (main.go:19)	JMP	335
	0x01a7 00423 (main.go:19)	NOP
	0x01a7 00423 (main.go:13)	PCDATA	$1, $-1
	0x01a7 00423 (main.go:13)	PCDATA	$0, $-2
	0x01a7 00423 (main.go:13)	CALL	runtime.morestack_noctxt(SB)
	0x01ac 00428 (main.go:13)	PCDATA	$0, $-1
	0x01ac 00428 (main.go:13)	JMP	0
	0x0000 64 48 8b 0c 25 00 00 00 00 48 8d 44 24 c0 48 3b  dH..%....H.D$.H;
	0x0010 41 10 0f 86 8f 01 00 00 48 81 ec c0 00 00 00 48  A.......H......H
	0x0020 89 ac 24 b8 00 00 00 48 8d ac 24 b8 00 00 00 e8  ..$....H..$.....
	0x0030 00 00 00 00 48 8b 14 24 48 89 54 24 48 48 8b 02  ....H..$H.T$HH..
	0x0040 ff d0 48 8b 04 24 48 89 44 24 40 48 8b 54 24 48  ..H..$H.D$@H.T$H
	0x0050 48 8b 02 ff d0 48 8b 04 24 48 89 44 24 38 48 8b  H....H..$H.D$8H.
	0x0060 54 24 48 48 8b 02 ff d0 48 8b 04 24 48 89 44 24  T$HH....H..$H.D$
	0x0070 30 48 8b 44 24 40 48 89 04 24 e8 00 00 00 00 48  0H.D$@H..$.....H
	0x0080 8b 44 24 08 48 89 44 24 68 48 8b 44 24 38 48 89  .D$.H.D$hH.D$8H.
	0x0090 04 24 e8 00 00 00 00 48 8b 44 24 08 48 89 44 24  .$.....H.D$.H.D$
	0x00a0 60 48 8b 44 24 30 48 89 04 24 e8 00 00 00 00 48  `H.D$0H..$.....H
	0x00b0 8b 44 24 08 48 89 44 24 58 0f 57 c0 0f 11 84 24  .D$.H.D$X.W....$
	0x00c0 88 00 00 00 0f 57 c0 0f 11 84 24 98 00 00 00 0f  .....W....$.....
	0x00d0 57 c0 0f 11 84 24 a8 00 00 00 48 8d 84 24 88 00  W....$....H..$..
	0x00e0 00 00 48 89 44 24 50 84 00 48 8b 4c 24 68 48 8d  ..H.D$P..H.L$hH.
	0x00f0 15 00 00 00 00 48 89 94 24 88 00 00 00 48 89 8c  .....H..$....H..
	0x0100 24 90 00 00 00 84 00 48 8b 44 24 60 48 8d 0d 00  $......H.D$`H...
	0x0110 00 00 00 48 89 8c 24 98 00 00 00 48 89 84 24 a0  ...H..$....H..$.
	0x0120 00 00 00 48 8b 44 24 50 84 00 48 8b 4c 24 58 48  ...H.D$P..H.L$XH
	0x0130 8d 15 00 00 00 00 48 89 50 20 48 8d 78 28 83 3d  ......H.P H.x(.=
	0x0140 00 00 00 00 00 74 02 eb 54 48 89 48 28 eb 00 48  .....t..TH.H(..H
	0x0150 8b 44 24 50 84 00 eb 00 48 89 44 24 70 48 c7 44  .D$P....H.D$pH.D
	0x0160 24 78 03 00 00 00 48 c7 84 24 80 00 00 00 03 00  $x....H..$......
	0x0170 00 00 48 89 04 24 48 c7 44 24 08 03 00 00 00 48  ..H..$H.D$.....H
	0x0180 c7 44 24 10 03 00 00 00 e8 00 00 00 00 48 8b ac  .D$..........H..
	0x0190 24 b8 00 00 00 48 81 c4 c0 00 00 00 c3 48 89 c8  $....H.......H..
	0x01a0 e8 00 00 00 00 eb a8 e8 00 00 00 00 e9 4f fe ff  .............O..
	0x01b0 ff                                               .
	rel 5+4 t=17 TLS+0
	rel 48+4 t=8 "".withoutArg+0
	rel 64+0 t=11 +0
	rel 83+0 t=11 +0
	rel 102+0 t=11 +0
	rel 123+4 t=8 runtime.convT64+0
	rel 147+4 t=8 runtime.convT64+0
	rel 171+4 t=8 runtime.convT64+0
	rel 241+4 t=16 type.int+0
	rel 271+4 t=16 type.int+0
	rel 306+4 t=16 type.int+0
	rel 320+4 t=16 runtime.writeBarrier+-1
	rel 393+4 t=8 fmt.Println+0
	rel 417+4 t=8 runtime.gcWriteBarrier+0
	rel 424+4 t=8 runtime.morestack_noctxt+0
type..eq.[3]interface {} STEXT dupok size=248 args=0x18 locals=0x60
	0x0000 00000 (<autogenerated>:1)	TEXT	type..eq.[3]interface {}(SB), DUPOK|ABIInternal, $96-24
	0x0000 00000 (<autogenerated>:1)	MOVQ	(TLS), CX
	0x0009 00009 (<autogenerated>:1)	CMPQ	SP, 16(CX)
	0x000d 00013 (<autogenerated>:1)	PCDATA	$0, $-2
	0x000d 00013 (<autogenerated>:1)	JLS	238
	0x0013 00019 (<autogenerated>:1)	PCDATA	$0, $-1
	0x0013 00019 (<autogenerated>:1)	SUBQ	$96, SP
	0x0017 00023 (<autogenerated>:1)	MOVQ	BP, 88(SP)
	0x001c 00028 (<autogenerated>:1)	LEAQ	88(SP), BP
	0x0021 00033 (<autogenerated>:1)	PCDATA	$0, $-2
	0x0021 00033 (<autogenerated>:1)	PCDATA	$1, $-2
	0x0021 00033 (<autogenerated>:1)	FUNCDATA	$0, gclocals·5e326a2f5498e528ce8fbe7bd86e7d21(SB)
	0x0021 00033 (<autogenerated>:1)	FUNCDATA	$1, gclocals·5363b3e3636eaaac400b9fb2466dc1f5(SB)
	0x0021 00033 (<autogenerated>:1)	FUNCDATA	$2, gclocals·6e8d7ea4abad763909b26991048ee1fe(SB)
	0x0021 00033 (<autogenerated>:1)	PCDATA	$0, $0
	0x0021 00033 (<autogenerated>:1)	PCDATA	$1, $0
	0x0021 00033 (<autogenerated>:1)	MOVB	$0, "".~r2+120(SP)
	0x0026 00038 (<autogenerated>:1)	MOVQ	$0, ""..autotmp_4+48(SP)
	0x002f 00047 (<autogenerated>:1)	MOVQ	$3, ""..autotmp_5+40(SP)
	0x0038 00056 (<autogenerated>:1)	JMP	58
	0x003a 00058 (<autogenerated>:1)	MOVQ	""..autotmp_5+40(SP), AX
	0x003f 00063 (<autogenerated>:1)	CMPQ	""..autotmp_4+48(SP), AX
	0x0044 00068 (<autogenerated>:1)	JLT	75
	0x0046 00070 (<autogenerated>:1)	JMP	223
	0x004b 00075 (<autogenerated>:1)	MOVQ	""..autotmp_4+48(SP), AX
	0x0050 00080 (<autogenerated>:1)	MOVQ	AX, "".i+32(SP)
	0x0055 00085 (<autogenerated>:1)	SHLQ	$4, AX
	0x0059 00089 (<autogenerated>:1)	PCDATA	$0, $1
	0x0059 00089 (<autogenerated>:1)	ADDQ	"".q+112(SP), AX
	0x005e 00094 (<autogenerated>:1)	PCDATA	$0, $2
	0x005e 00094 (<autogenerated>:1)	MOVQ	8(AX), CX
	0x0062 00098 (<autogenerated>:1)	PCDATA	$0, $3
	0x0062 00098 (<autogenerated>:1)	MOVQ	(AX), AX
	0x0065 00101 (<autogenerated>:1)	PCDATA	$1, $1
	0x0065 00101 (<autogenerated>:1)	MOVQ	AX, ""..autotmp_6+72(SP)
	0x006a 00106 (<autogenerated>:1)	PCDATA	$0, $0
	0x006a 00106 (<autogenerated>:1)	MOVQ	CX, ""..autotmp_6+80(SP)
	0x006f 00111 (<autogenerated>:1)	MOVQ	"".i+32(SP), AX
	0x0074 00116 (<autogenerated>:1)	SHLQ	$4, AX
	0x0078 00120 (<autogenerated>:1)	PCDATA	$0, $1
	0x0078 00120 (<autogenerated>:1)	ADDQ	"".p+104(SP), AX
	0x007d 00125 (<autogenerated>:1)	PCDATA	$0, $2
	0x007d 00125 (<autogenerated>:1)	MOVQ	8(AX), CX
	0x0081 00129 (<autogenerated>:1)	PCDATA	$0, $3
	0x0081 00129 (<autogenerated>:1)	MOVQ	(AX), AX
	0x0084 00132 (<autogenerated>:1)	MOVQ	AX, ""..autotmp_7+56(SP)
	0x0089 00137 (<autogenerated>:1)	MOVQ	CX, ""..autotmp_7+64(SP)
	0x008e 00142 (<autogenerated>:1)	CMPQ	""..autotmp_6+72(SP), AX
	0x0093 00147 (<autogenerated>:1)	JNE	151
	0x0095 00149 (<autogenerated>:1)	JMP	168
	0x0097 00151 (<autogenerated>:1)	PCDATA	$0, $-1
	0x0097 00151 (<autogenerated>:1)	PCDATA	$1, $-1
	0x0097 00151 (<autogenerated>:1)	JMP	153
	0x0099 00153 (<autogenerated>:1)	PCDATA	$0, $0
	0x0099 00153 (<autogenerated>:1)	PCDATA	$1, $2
	0x0099 00153 (<autogenerated>:1)	MOVB	$0, "".~r2+120(SP)
	0x009e 00158 (<autogenerated>:1)	MOVQ	88(SP), BP
	0x00a3 00163 (<autogenerated>:1)	ADDQ	$96, SP
	0x00a7 00167 (<autogenerated>:1)	RET
	0x00a8 00168 (<autogenerated>:1)	PCDATA	$0, $3
	0x00a8 00168 (<autogenerated>:1)	PCDATA	$1, $1
	0x00a8 00168 (<autogenerated>:1)	MOVQ	AX, (SP)
	0x00ac 00172 (<autogenerated>:1)	PCDATA	$0, $0
	0x00ac 00172 (<autogenerated>:1)	MOVQ	CX, 8(SP)
	0x00b1 00177 (<autogenerated>:1)	PCDATA	$0, $1
	0x00b1 00177 (<autogenerated>:1)	PCDATA	$1, $0
	0x00b1 00177 (<autogenerated>:1)	MOVQ	""..autotmp_6+80(SP), AX
	0x00b6 00182 (<autogenerated>:1)	PCDATA	$0, $0
	0x00b6 00182 (<autogenerated>:1)	MOVQ	AX, 16(SP)
	0x00bb 00187 (<autogenerated>:1)	CALL	runtime.efaceeq(SB)
	0x00c0 00192 (<autogenerated>:1)	CMPB	24(SP), $0
	0x00c5 00197 (<autogenerated>:1)	JNE	201
	0x00c7 00199 (<autogenerated>:1)	JMP	221
	0x00c9 00201 (<autogenerated>:1)	PCDATA	$0, $-1
	0x00c9 00201 (<autogenerated>:1)	PCDATA	$1, $-1
	0x00c9 00201 (<autogenerated>:1)	JMP	203
	0x00cb 00203 (<autogenerated>:1)	PCDATA	$0, $0
	0x00cb 00203 (<autogenerated>:1)	PCDATA	$1, $0
	0x00cb 00203 (<autogenerated>:1)	MOVQ	""..autotmp_4+48(SP), AX
	0x00d0 00208 (<autogenerated>:1)	INCQ	AX
	0x00d3 00211 (<autogenerated>:1)	MOVQ	AX, ""..autotmp_4+48(SP)
	0x00d8 00216 (<autogenerated>:1)	JMP	58
	0x00dd 00221 (<autogenerated>:1)	PCDATA	$0, $-1
	0x00dd 00221 (<autogenerated>:1)	PCDATA	$1, $-1
	0x00dd 00221 (<autogenerated>:1)	JMP	153
	0x00df 00223 (<autogenerated>:1)	PCDATA	$0, $0
	0x00df 00223 (<autogenerated>:1)	PCDATA	$1, $2
	0x00df 00223 (<autogenerated>:1)	MOVB	$1, "".~r2+120(SP)
	0x00e4 00228 (<autogenerated>:1)	MOVQ	88(SP), BP
	0x00e9 00233 (<autogenerated>:1)	ADDQ	$96, SP
	0x00ed 00237 (<autogenerated>:1)	RET
	0x00ee 00238 (<autogenerated>:1)	NOP
	0x00ee 00238 (<autogenerated>:1)	PCDATA	$1, $-1
	0x00ee 00238 (<autogenerated>:1)	PCDATA	$0, $-2
	0x00ee 00238 (<autogenerated>:1)	CALL	runtime.morestack_noctxt(SB)
	0x00f3 00243 (<autogenerated>:1)	PCDATA	$0, $-1
	0x00f3 00243 (<autogenerated>:1)	JMP	0
	0x0000 64 48 8b 0c 25 00 00 00 00 48 3b 61 10 0f 86 db  dH..%....H;a....
	0x0010 00 00 00 48 83 ec 60 48 89 6c 24 58 48 8d 6c 24  ...H..`H.l$XH.l$
	0x0020 58 c6 44 24 78 00 48 c7 44 24 30 00 00 00 00 48  X.D$x.H.D$0....H
	0x0030 c7 44 24 28 03 00 00 00 eb 00 48 8b 44 24 28 48  .D$(......H.D$(H
	0x0040 39 44 24 30 7c 05 e9 94 00 00 00 48 8b 44 24 30  9D$0|......H.D$0
	0x0050 48 89 44 24 20 48 c1 e0 04 48 03 44 24 70 48 8b  H.D$ H...H.D$pH.
	0x0060 48 08 48 8b 00 48 89 44 24 48 48 89 4c 24 50 48  H.H..H.D$HH.L$PH
	0x0070 8b 44 24 20 48 c1 e0 04 48 03 44 24 68 48 8b 48  .D$ H...H.D$hH.H
	0x0080 08 48 8b 00 48 89 44 24 38 48 89 4c 24 40 48 39  .H..H.D$8H.L$@H9
	0x0090 44 24 48 75 02 eb 11 eb 00 c6 44 24 78 00 48 8b  D$Hu......D$x.H.
	0x00a0 6c 24 58 48 83 c4 60 c3 48 89 04 24 48 89 4c 24  l$XH..`.H..$H.L$
	0x00b0 08 48 8b 44 24 50 48 89 44 24 10 e8 00 00 00 00  .H.D$PH.D$......
	0x00c0 80 7c 24 18 00 75 02 eb 14 eb 00 48 8b 44 24 30  .|$..u.....H.D$0
	0x00d0 48 ff c0 48 89 44 24 30 e9 5d ff ff ff eb ba c6  H..H.D$0.]......
	0x00e0 44 24 78 01 48 8b 6c 24 58 48 83 c4 60 c3 e8 00  D$x.H.l$XH..`...
	0x00f0 00 00 00 e9 08 ff ff ff                          ........
	rel 5+4 t=17 TLS+0
	rel 188+4 t=8 runtime.efaceeq+0
	rel 239+4 t=8 runtime.morestack_noctxt+0
"".withoutArg.func1 STEXT nosplit size=70 args=0x8 locals=0x10
	0x0000 00000 (main.go:7)	TEXT	"".withoutArg.func1(SB), NOSPLIT|NEEDCTXT|ABIInternal, $16-8
	0x0000 00000 (main.go:7)	SUBQ	$16, SP
	0x0004 00004 (main.go:7)	MOVQ	BP, 8(SP)
	0x0009 00009 (main.go:7)	LEAQ	8(SP), BP
	0x000e 00014 (main.go:7)	PCDATA	$0, $-2
	0x000e 00014 (main.go:7)	PCDATA	$1, $-2
	0x000e 00014 (main.go:7)	FUNCDATA	$0, gclocals·69c1753bd5f81501d95132d08af04464(SB)
	0x000e 00014 (main.go:7)	FUNCDATA	$1, gclocals·9fb7f0986f647f17cb53dda1484e0f7a(SB)
	0x000e 00014 (main.go:7)	FUNCDATA	$2, gclocals·ec2e8d6f928c0735ffaf01f3f360ba43(SB)
	0x000e 00014 (main.go:7)	PCDATA	$0, $2
	0x000e 00014 (main.go:7)	PCDATA	$1, $0
	0x000e 00014 (main.go:7)	MOVQ	8(DX), AX
	0x0012 00018 (main.go:7)	PCDATA	$0, $0
	0x0012 00018 (main.go:7)	PCDATA	$1, $1
	0x0012 00018 (main.go:7)	MOVQ	AX, "".&i(SP)
	0x0016 00022 (main.go:7)	MOVQ	$0, "".~r0+24(SP)
	0x001f 00031 (main.go:8)	PCDATA	$0, $2
	0x001f 00031 (main.go:8)	MOVQ	"".&i(SP), AX
	0x0023 00035 (main.go:8)	PCDATA	$0, $0
	0x0023 00035 (main.go:8)	MOVQ	(AX), AX
	0x0026 00038 (main.go:8)	PCDATA	$0, $3
	0x0026 00038 (main.go:8)	MOVQ	"".&i(SP), CX
	0x002a 00042 (main.go:8)	INCQ	AX
	0x002d 00045 (main.go:8)	PCDATA	$0, $0
	0x002d 00045 (main.go:8)	MOVQ	AX, (CX)
	0x0030 00048 (main.go:9)	PCDATA	$0, $2
	0x0030 00048 (main.go:9)	PCDATA	$1, $0
	0x0030 00048 (main.go:9)	MOVQ	"".&i(SP), AX
	0x0034 00052 (main.go:9)	PCDATA	$0, $0
	0x0034 00052 (main.go:9)	MOVQ	(AX), AX
	0x0037 00055 (main.go:9)	MOVQ	AX, "".~r0+24(SP)
	0x003c 00060 (main.go:9)	MOVQ	8(SP), BP
	0x0041 00065 (main.go:9)	ADDQ	$16, SP
	0x0045 00069 (main.go:9)	RET
	0x0000 48 83 ec 10 48 89 6c 24 08 48 8d 6c 24 08 48 8b  H...H.l$.H.l$.H.
	0x0010 42 08 48 89 04 24 48 c7 44 24 18 00 00 00 00 48  B.H..$H.D$.....H
	0x0020 8b 04 24 48 8b 00 48 8b 0c 24 48 ff c0 48 89 01  ..$H..H..$H..H..
	0x0030 48 8b 04 24 48 8b 00 48 89 44 24 18 48 8b 6c 24  H..$H..H.D$.H.l$
	0x0040 08 48 83 c4 10 c3                                .H....
go.cuinfo.packagename. SDWARFINFO dupok size=0
	0x0000 6d 61 69 6e                                      main
go.loc."".withoutArg SDWARFLOC size=0
go.info."".withoutArg SDWARFINFO size=64
	0x0000 03 22 22 2e 77 69 74 68 6f 75 74 41 72 67 00 00  ."".withoutArg..
	0x0010 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 01  ................
	0x0020 9c 00 00 00 00 01 0a 26 69 00 06 00 00 00 00 02  .......&i.......
	0x0030 91 68 0f 7e 72 30 00 01 05 00 00 00 00 01 9c 00  .h.~r0..........
	rel 0+0 t=24 type.*int+0
	rel 0+0 t=24 type.*struct { F uintptr; "".i *int }+0
	rel 0+0 t=24 type.func() int+0
	rel 15+8 t=1 "".withoutArg+0
	rel 23+8 t=1 "".withoutArg+184
	rel 33+4 t=30 gofile../home/zquxbow/go/learnGo/basic/language/assembly/closure/main.go+0
	rel 43+4 t=29 go.info.*int+0
	rel 57+4 t=29 go.info.func() int+0
go.range."".withoutArg SDWARFRANGE size=0
go.debuglines."".withoutArg SDWARFMISC size=27
	0x0000 04 02 13 0a cd f6 06 55 06 08 24 06 55 06 02 26  .......U..$.U..&
	0x0010 ff 06 41 06 08 17 04 01 03 7c 01                 ..A......|.
go.loc."".main SDWARFLOC size=0
go.info."".main SDWARFINFO size=82
	0x0000 03 22 22 2e 6d 61 69 6e 00 00 00 00 00 00 00 00  ."".main........
	0x0010 00 00 00 00 00 00 00 00 00 01 9c 00 00 00 00 01  ................
	0x0020 0a 63 00 11 00 00 00 00 03 91 e8 7e 0a 62 00 10  .c.........~.b..
	0x0030 00 00 00 00 03 91 f0 7e 0a 61 00 0f 00 00 00 00  .......~.a......
	0x0040 03 91 f8 7e 0a 66 31 00 0e 00 00 00 00 03 91 80  ...~.f1.........
	0x0050 7f 00                                            ..
	rel 0+0 t=24 type.*[3]interface {}+0
	rel 0+0 t=24 type.[3]interface {}+0
	rel 0+0 t=24 type.[]interface {}+0
	rel 0+0 t=24 type.func() int+0
	rel 0+0 t=24 type.int+0
	rel 0+0 t=24 type.unsafe.Pointer+0
	rel 9+8 t=1 "".main+0
	rel 17+8 t=1 "".main+433
	rel 27+4 t=30 gofile../home/zquxbow/go/learnGo/basic/language/assembly/closure/main.go+0
	rel 36+4 t=29 go.info.int+0
	rel 48+4 t=29 go.info.int+0
	rel 60+4 t=29 go.info.int+0
	rel 73+4 t=29 go.info.func() int+0
go.range."".main SDWARFRANGE size=0
go.debuglines."".main SDWARFMISC size=42
	0x0000 04 02 03 07 14 0a ff f6 06 41 06 6a 06 2d 06 7e  .........A.j.-.~
	0x0010 06 41 06 9c 06 41 06 9d 06 41 06 02 80 02 f6 06  .A...A...A......
	0x0020 ae 06 03 7e 6f 04 01 03 74 01                    ...~o...t.
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
go.loc.type..eq.[3]interface {} SDWARFLOC dupok size=0
go.info.type..eq.[3]interface {} SDWARFINFO dupok size=99
	0x0000 03 74 79 70 65 2e 2e 65 71 2e 5b 33 5d 69 6e 74  .type..eq.[3]int
	0x0010 65 72 66 61 63 65 20 7b 7d 00 00 00 00 00 00 00  erface {}.......
	0x0020 00 00 00 00 00 00 00 00 00 00 01 9c 00 00 00 00  ................
	0x0030 01 0a 69 00 01 00 00 00 00 03 91 b8 7f 0f 70 00  ..i...........p.
	0x0040 00 01 00 00 00 00 01 9c 0f 71 00 00 01 00 00 00  .........q......
	0x0050 00 02 91 08 0f 7e 72 32 00 01 01 00 00 00 00 02  .....~r2........
	0x0060 91 10 00                                         ...
	rel 0+0 t=24 type.*[3]interface {}+0
	rel 0+0 t=24 type.bool+0
	rel 0+0 t=24 type.int+0
	rel 0+0 t=24 type.interface {}+0
	rel 26+8 t=1 type..eq.[3]interface {}+0
	rel 34+8 t=1 type..eq.[3]interface {}+248
	rel 44+4 t=30 gofile..<autogenerated>+0
	rel 53+4 t=29 go.info.int+0
	rel 66+4 t=29 go.info.*[3]interface {}+0
	rel 77+4 t=29 go.info.*[3]interface {}+0
	rel 91+4 t=29 go.info.bool+0
go.range.type..eq.[3]interface {} SDWARFRANGE dupok size=0
go.debuglines.type..eq.[3]interface {} SDWARFMISC dupok size=30
	0x0000 04 01 0f 0a cd 06 cd 06 d7 06 41 06 02 42 ff 06  ..........A..B..
	0x0010 41 06 02 29 ff 06 41 06 73 04 01 03 00 01        A..)..A.s.....
type..eqfunc.[3]interface {} SRODATA dupok size=8
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 type..eq.[3]interface {}+0
type..namedata.*[3]interface {}- SRODATA dupok size=19
	0x0000 00 00 10 2a 5b 33 5d 69 6e 74 65 72 66 61 63 65  ...*[3]interface
	0x0010 20 7b 7d                                          {}
type.*[3]interface {} SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 b0 18 fe b9 08 08 08 36 00 00 00 00 00 00 00 00  .......6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64·f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[3]interface {}-+0
	rel 48+8 t=1 type.[3]interface {}+0
runtime.gcbits.2a SRODATA dupok size=1
	0x0000 2a                                               *
type.[3]interface {} SRODATA dupok size=72
	0x0000 30 00 00 00 00 00 00 00 30 00 00 00 00 00 00 00  0.......0.......
	0x0010 1d dd cf d9 02 08 08 11 00 00 00 00 00 00 00 00  ................
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 03 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 type..eqfunc.[3]interface {}+0
	rel 32+8 t=1 runtime.gcbits.2a+0
	rel 40+4 t=5 type..namedata.*[3]interface {}-+0
	rel 44+4 t=6 type.*[3]interface {}+0
	rel 48+8 t=1 type.interface {}+0
	rel 56+8 t=1 type.[]interface {}+0
go.loc."".withoutArg.func1 SDWARFLOC size=0
go.info."".withoutArg.func1 SDWARFINFO size=70
	0x0000 03 22 22 2e 77 69 74 68 6f 75 74 41 72 67 2e 66  ."".withoutArg.f
	0x0010 75 6e 63 31 00 00 00 00 00 00 00 00 00 00 00 00  unc1............
	0x0020 00 00 00 00 00 01 9c 00 00 00 00 01 0a 26 69 00  .............&i.
	0x0030 07 00 00 00 00 02 91 68 0f 7e 72 30 00 01 07 00  .......h.~r0....
	0x0040 00 00 00 01 9c 00                                ......
	rel 0+0 t=24 type.*int+0
	rel 0+0 t=24 type.int+0
	rel 21+8 t=1 "".withoutArg.func1+0
	rel 29+8 t=1 "".withoutArg.func1+70
	rel 39+4 t=30 gofile../home/zquxbow/go/learnGo/basic/language/assembly/closure/main.go+0
	rel 49+4 t=29 go.info.*int+0
	rel 63+4 t=29 go.info.int+0
go.range."".withoutArg.func1 SDWARFRANGE size=0
go.debuglines."".withoutArg.func1 SDWARFMISC size=22
	0x0000 04 02 0a 03 01 14 06 9b 06 ba 06 37 06 92 06 37  ...........7...7
	0x0010 04 01 03 78 06 01                                ...x..
""..inittask SNOPTRDATA size=32
	0x0000 00 00 00 00 00 00 00 00 01 00 00 00 00 00 00 00  ................
	0x0010 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	rel 24+8 t=1 fmt..inittask+0
type..namedata.*struct { F uintptr; i *int }- SRODATA dupok size=32
	0x0000 00 00 1d 2a 73 74 72 75 63 74 20 7b 20 46 20 75  ...*struct { F u
	0x0010 69 6e 74 70 74 72 3b 20 69 20 2a 69 6e 74 20 7d  intptr; i *int }
type..namedata..F- SRODATA dupok size=5
	0x0000 00 00 02 2e 46                                   ....F
type..namedata.i- SRODATA dupok size=4
	0x0000 00 00 01 69                                      ...i
type.noalg.struct { F uintptr; "".i *int } SRODATA dupok size=128
	0x0000 10 00 00 00 00 00 00 00 10 00 00 00 00 00 00 00  ................
	0x0010 b7 43 25 9a 02 08 08 19 00 00 00 00 00 00 00 00  .C%.............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 02 00 00 00 00 00 00 00 02 00 00 00 00 00 00 00  ................
	0x0050 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0060 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0070 00 00 00 00 00 00 00 00 10 00 00 00 00 00 00 00  ................
	rel 32+8 t=1 runtime.gcbits.02+0
	rel 40+4 t=5 type..namedata.*struct { F uintptr; i *int }-+0
	rel 44+4 t=6 type.*struct { F uintptr; "".i *int }+0
	rel 48+8 t=1 type..importpath."".+0
	rel 56+8 t=1 type.noalg.struct { F uintptr; "".i *int }+80
	rel 80+8 t=1 type..namedata..F-+0
	rel 88+8 t=1 type.uintptr+0
	rel 104+8 t=1 type..namedata.i-+0
	rel 112+8 t=1 type.*int+0
type.*struct { F uintptr; "".i *int } SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 23 f3 08 44 08 08 08 36 00 00 00 00 00 00 00 00  #..D...6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64·f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*struct { F uintptr; i *int }-+0
	rel 48+8 t=1 type.noalg.struct { F uintptr; "".i *int }+0
type..namedata.*func() int- SRODATA dupok size=14
	0x0000 00 00 0b 2a 66 75 6e 63 28 29 20 69 6e 74        ...*func() int
type.*func() int SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 51 b2 ed c6 08 08 08 36 00 00 00 00 00 00 00 00  Q......6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64·f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*func() int-+0
	rel 48+8 t=1 type.func() int+0
type.func() int SRODATA dupok size=64
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 e5 86 39 e0 02 08 08 33 00 00 00 00 00 00 00 00  ..9....3........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 01 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*func() int-+0
	rel 44+4 t=6 type.*func() int+0
	rel 56+8 t=1 type.int+0
type..importpath.fmt. SRODATA dupok size=6
	0x0000 00 00 03 66 6d 74                                ...fmt
gclocals·272bd7e3f0f711ce8c35e9a618d8e620 SRODATA dupok size=13
	0x0000 05 00 00 00 01 00 00 00 00 00 00 00 01           .............
gclocals·cc5cad6002ade825472fe3ad2912087d SRODATA dupok size=13
	0x0000 05 00 00 00 02 00 00 00 00 02 03 01 00           .............
gclocals·1e847dfe98917b7bcc7220b0d92d534f SRODATA dupok size=12
	0x0000 04 00 00 00 07 00 00 00 00 01 03 43              ...........C
gclocals·7be4bbacbfdb05fb3044e36c22b41e8b SRODATA dupok size=8
	0x0000 0a 00 00 00 00 00 00 00                          ........
gclocals·9ee33c85bd820cba75d4d4aac78994e7 SRODATA dupok size=28
	0x0000 0a 00 00 00 0e 00 00 00 00 00 01 00 10 00 18 00  ................
	0x0010 1c 00 1c 2a 1e 00 0e 00 06 00 02 00              ...*........
gclocals·74bf2e2f0b69fa7db29b8b87c766370b SRODATA dupok size=14
	0x0000 06 00 00 00 07 00 00 00 00 04 01 03 07 43        .............C
"".main.stkobj SRODATA size=24
	0x0000 01 00 00 00 00 00 00 00 d0 ff ff ff ff ff ff ff  ................
	0x0010 00 00 00 00 00 00 00 00                          ........
	rel 16+8 t=1 type.[3]interface {}+0
gclocals·5e326a2f5498e528ce8fbe7bd86e7d21 SRODATA dupok size=11
	0x0000 03 00 00 00 02 00 00 00 03 03 00                 ...........
gclocals·5363b3e3636eaaac400b9fb2466dc1f5 SRODATA dupok size=11
	0x0000 03 00 00 00 04 00 00 00 00 08 00                 ...........
gclocals·6e8d7ea4abad763909b26991048ee1fe SRODATA dupok size=12
	0x0000 04 00 00 00 02 00 00 00 00 01 03 02              ............
gclocals·69c1753bd5f81501d95132d08af04464 SRODATA dupok size=8
	0x0000 02 00 00 00 00 00 00 00                          ........
gclocals·9fb7f0986f647f17cb53dda1484e0f7a SRODATA dupok size=10
	0x0000 02 00 00 00 01 00 00 00 00 01                    ..........
gclocals·ec2e8d6f928c0735ffaf01f3f360ba43 SRODATA dupok size=12
	0x0000 04 00 00 00 03 00 00 00 00 04 01 02              ............
