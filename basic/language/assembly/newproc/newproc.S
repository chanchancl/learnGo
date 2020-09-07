  
  // dump command
  // go tool objdump -s ^runtime.newproc$ binary
  // cause -s using regex, so add a $ at the end of newproc
  // binary can be any program build by go
  
  // remove hex code
  // (^.*go:\d{4})(\s*\w*\s*\w*\s*)(.*)
  // $1 $3

  proc.go:3523 SUBQ $0x40, SP
  // Stack size 0x40
  // 7 locals and arguments, 1 return address
  //
  // (0x38)SP, BP, save BP, and treate it as new BP
  proc.go:3523 MOVQ BP, 0x38(SP)
  proc.go:3523 LEAQ 0x38(SP), BP
  proc.go:3524 MOVQ FS:0xfffffff8, AX
  // MOVQ FS:0xfffffff8 -> AX
  proc.go:3527 XORPS X0, X0
  proc.go:3527 MOVUPS X0, 0x8(SP)
  proc.go:3527 MOVUPS X0, 0x18(SP)
  proc.go:3527 MOVUPS X0, 0x28(SP)
  // clear X0 
  // and then clear 0x8(SP) - 0x28(SP), size 0x20, 32 bytes
  // mov address of func1 to CX
  //
  // Current Stack
  //-(0x08)SP will be return address
  // (0x00)SP Raw 0x08(SP), newproc.func1
  // (0x08)SP newproc.func1 , who will run newproc1
  // (0x10)SP Raw 0x50(SP)
  // (0x18)SP Raw 0x58(SP)
  // (0x20)SP Raw 0x48(SP) MOVL?  CX
  // (0x28)SP AX
  // (0x30)SP Raw 0x40(SP)
  // (0x38)SP 

  // Previous Stack (From gofunc/main.S)
  // (0x40)SP              -> AX
  // (0x48)SP 
  // (0x50)SP 
  // (0x58)SP 
  proc.go:3527 LEAQ runtime.newproc.func1(SB), CX
  proc.go:3527 MOVQ CX, 0x8(SP)
  proc.go:3527 LEAQ 0x50(SP), CX
  proc.go:3527 MOVQ CX, 0x10(SP)
  proc.go:3527 LEAQ 0x58(SP), CX
  proc.go:3527 MOVQ CX, 0x18(SP)
  proc.go:3527 MOVL 0x48(SP), CX
  proc.go:3527 MOVL CX, 0x20(SP)
  proc.go:3527 MOVQ AX, 0x28(SP)
  proc.go:3527 MOVQ 0x40(SP), AX
  proc.go:3527 MOVQ AX, 0x30(SP)
  proc.go:3527 LEAQ 0x8(SP), AX
  proc.go:3527 MOVQ AX, 0(SP)
  proc.go:3527 CALL runtime.systemstack(SB)
  proc.go:3537 MOVQ 0x38(SP), BP
  proc.go:3537 ADDQ $0x40, SP
  proc.go:3537 RET


  // with hex code
  proc.go:3523          0x437ba0                4883ec40                SUBQ $0x40, SP
  proc.go:3523          0x437ba4                48896c2438              MOVQ BP, 0x38(SP)
  proc.go:3523          0x437ba9                488d6c2438              LEAQ 0x38(SP), BP
  proc.go:3524          0x437bae                64488b0425f8ffffff      MOVQ FS:0xfffffff8, AX
  proc.go:3527          0x437bb7                0f57c0                  XORPS X0, X0
  proc.go:3527          0x437bba                0f11442408              MOVUPS X0, 0x8(SP)
  proc.go:3527          0x437bbf                0f11442418              MOVUPS X0, 0x18(SP)
  proc.go:3527          0x437bc4                0f11442428              MOVUPS X0, 0x28(SP)
  proc.go:3527          0x437bc9                488d0dd0ec0100          LEAQ runtime.newproc.func1(SB), CX
  proc.go:3527          0x437bd0                48894c2408              MOVQ CX, 0x8(SP)
  proc.go:3527          0x437bd5                488d4c2450              LEAQ 0x50(SP), CX
  proc.go:3527          0x437bda                48894c2410              MOVQ CX, 0x10(SP)
  proc.go:3527          0x437bdf                488d4c2458              LEAQ 0x58(SP), CX
  proc.go:3527          0x437be4                48894c2418              MOVQ CX, 0x18(SP)
  proc.go:3527          0x437be9                8b4c2448                MOVL 0x48(SP), CX
  proc.go:3527          0x437bed                894c2420                MOVL CX, 0x20(SP)
  proc.go:3527          0x437bf1                4889442428              MOVQ AX, 0x28(SP)
  proc.go:3527          0x437bf6                488b442440              MOVQ 0x40(SP), AX
  proc.go:3527          0x437bfb                4889442430              MOVQ AX, 0x30(SP)
  proc.go:3527          0x437c00                488d442408              LEAQ 0x8(SP), AX
  proc.go:3527          0x437c05                48890424                MOVQ AX, 0(SP)
  proc.go:3527          0x437c09                e8120c0200              CALL runtime.systemstack(SB)
  proc.go:3537          0x437c0e                488b6c2438              MOVQ 0x38(SP), BP
  proc.go:3537          0x437c13                4883c440                ADDQ $0x40, SP
  proc.go:3537          0x437c17                c3                      RET