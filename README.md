# Lox language
## TODOs
  - (Partially complete) unreachable code. if a "return" has been found in a local scope and we encounter other code directly following then we error
  - values never read/used

# RISCV-Meta-Assembler
RISC-V meta assembler that adds quality of life features to assembly

# Memory map (pg 40 of TheReader)
- sp = 0xbfffffff
- 0->128K is reserved
- code starts at 0x00010001
- static data follows
- then Dynamic data
- and finally the Stack which grows downward.

# Code example
```

import {
    "stdio" as std
}

const {
    GPIO_BASE = 0x10012000,  // GPIO base
    GPIO_RED  = 0x400000
}

// Program code section.
// align code to 2^2 bytes = word alignment
// declares "main" symbol as global
// positions code at address 0x00010001
[alignTo word, global, at 0x00010001]
code main {
    use std

    addi sp,sp,-16  // allocate stack frame
    sw   ra,12(sp)    // save return address

    // Setup function parameters    
    lui  a0,%hi(hello)  // compute address of hello upper 20
    addi a0,a0,%lo(hello) // lower 12
    lui  a1,%hi(world)
    addi a1,a1,%lo(world)

    call printf     // call std's printf
    
    lw   ra,12(sp) // restore return address
    addi sp,sp,16  // deallocate stack frame
    
    li   a0,0 // load return value 0
    ret
}

[readOnly, alignTo bytes<4>, at 0x10000000]
data {
    string hello "Hello, %s\n", // null terminated
    string world "world",
    char ['a', 'b', 'c'],
    byte [0x20, 0x0f, 0x40, 0x0a]
    half [0xff01,0xab02]
    word [0xffff001],
}

[readWrite, alignTo bytes(4)]
data {
    global int<4> count,    // 4 byte integer = word
}
```

# Links
https://github.com/gonzispina/golox
