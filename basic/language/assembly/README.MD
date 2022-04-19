
# As a conclusion (Without optimize)

1. function will allocate the space for locals and function return value
2. instruction call will push return address (push (1)ipc )
3. instruction retq will jmp to ((0)SP) and pop it
4. Base on 2 and 3, a function should allocate and free stack before return


## 

```
// (gdb) x/g 0xc00003e750
// 0xc00003e750:   0x00000000004583f3  //  0(SP) return address, and also 0(SP)
// 0xc00003e758:   0x0000000000000001  //  8(SP) argument 1
// 0xc00003e760:   0x0000000000000002  // 16(SP) argument 2
// 0xc00003e768:   0x0000000000461ce0  // 24(SP) return value 1
// 0xc00003e770:   0x000000c000094058  // 32(SP) return value 2   (stack size = (1 return address + 4 arguments) * 8 = 40 Byte
// above are allocated for call f1
// 0xc00003e778:   0x000000c00003e7d0  // 40(SP) stack of main
// 0xc00003e780:   0x000000000042c00a
```

```
// After enter a new function, SP refers the return address
//
// Alloc 40 Byte Size
// 5 elements
//
// Before
// 0x00(SP) $1
// 0x08(SP) $2
// 0x10(SP) return value1
// 0x18(SP) return value2
// 0x20(SP) Raw BP value
```

```
// Before call, SP
// rsp             0xc00003e758     0xc00003e758
// space 0xc00003e758 - 0xc00003e778 are stack for main
// space to save return addres don't own any function (at least do not need to allocate and free)
// (gdb) x/g $rsp
// 00(SP) 0xc00003e758:   0x0000000000000001  // argument 1
// 08(SP) 0xc00003e760:   0x0000000000000002  // argument 2
// 16(SP) 0xc00003e768:   0x0000000000000202  // space for return value 1
// 24(SP) 0xc00003e770:   0x000000c00003e7d0  // space for return value 2
// 32(SP) 0xc00003e778:   0x000000c00003e7d0	// Raw BP value

// 0xc00003e778:   0x000000c00003e7d0  // this maybe previous return address?
// 0xc00003e780:   0x000000000042fb29
0x002e 00046 (main.go:9)	CALL	"".f1(SB)
// In function f11
// 1.value of SP sub by 8 (push return address)
// rsp             0xc00003e750     0xc00003e750
// 2.stack info
// (gdb) x/g $rsp
// 00(SP) 0xc00003e750:   0x000000000045dab3  // return address
// 08(SP) 0xc00003e758:   0x0000000000000001  // argument 1
// 16(SP) 0xc00003e760:   0x0000000000000002  // argument 2
// 24(SP) 0xc00003e768:   0x0000000000000202  // space for return value 1
// 32(SP) 0xc00003e770:   0x000000c00003e7d0  // space for return value 2
// 40(SP) 0xc00003e778:   0x000000c00003e7d0	// Raw BP value
//

// from rip we know 0x45dab3 is return address
// rip            0x45dab3 0x45dab3 <main.main+51>

// After function f11 and return to main
// 1. value of SP
// i r rsp
// rsp            0xc00003e758     0xc00003e758
// 2. stack info
// (gdb) x/g $rsp
// 00(SP) 0xc00003e758:   0x0000000000000001  // argument 1
// 08(SP) 0xc00003e760:   0x0000000000000002  // argument 2
// 16(SP) 0xc00003e768:   0x0000000000000003  // return value 1 (1 + 2 ==  3)
// 24(SP) 0xc00003e770:   0xffffffffffffffff  // return value 2 (1 - 2 == -1)
// 32(SP) 0xc00003e778:   0x000000c00003e7d0	// Raw BP value
```