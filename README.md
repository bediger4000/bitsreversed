# Return a bits-reversed integer

Daily Coding Problem: Problem #655 [Easy]

This problem was asked by Facebook.

Given a 32-bit integer, return the number with its bits reversed.

For example, given the binary number 1111 0000 1111 0000 1111 0000 1111 0000,
return 0000 1111 0000 1111 0000 1111 0000 1111.

## Build and Run

```sh
$ go build .
$ ./bitsreversed 4042322160
0000000000000000000000000000000011110000111100001111000011110000
0000111100001111000011110000111100000000000000000000000000000000
```

It does 64-bit integers, while the problem statement seems to assume
32-bit integers.

The input is a base-10 representation of a number.
The outputs are base-2 representations.

## Analysis

My algorithm is fairly simple:

```go
    for i := 0; i < 64; i++ {
        reversed <<= 1
        bit := n & 0x1
        n >>= 1
        reversed |= bit
    }
```

1. Shift the current value of the reversed number 1 bit left.
2. Get the value of the right-most bit from the original (`n`)
3. Shift the original 1 bit right.
4. Bit-wise OR that bit value with the reversed number.
Because we just shifted the reversed number 1 bit left,
that OR causes the right-most bit to have the value of the bit
shifted right off the original.

This is a little fiddly to get correct.
You definitely have to shift-left the reversed value before
putting the right-most bit from the original.

The actual code causes the variable that held the original value
to end up as all-zero-bits.

There is a way to do this a byte at a time,
maybe with a lookup table to go from one of 256 byte values,
to the value that has all 8 bits of the byte reversed.
But that would take effort to prepare the lookup table.

As a programming problem, this emphasizes bit-wise operations.
There's only one loop, but there's left and right shifts,
along with bitwise AND and OR operations.
It's probably reasonable to rate this as "Easy",
and it's appropriate as an interview question for an
entry- or junior-level job.
As always, test cases would be a good place for a candidate to distiguish
themselves from others.
High-bit-set test cases, and low-bit set test cases are merited,
given some programming languages problems with interpretation of sign bits
on twos-complement number representations.
The candidate could also talk about 32- vs 64-bit variables,
and maybe even arbitrary bit-sized represenations.

If the interviewer is interested in verifying programming
language familiarity,
this might be a good easy question to ask.
Given the bitwise operation emphasis,
this might not be a good question for candidates for some jobs,
like front-end progammers, maybe.
