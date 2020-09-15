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

## Analysis
