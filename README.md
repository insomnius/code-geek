# Code Geek

This repo is dedicated to store all of programing interview test from a very simple one to a very hardest one, algorithm implementation, mini console game like tetris or pacman maybe and all of geeky code. The aim of this repo is to help people to learn algorithm better and find solution that maybe useful in the future.

## Prerequisites

In order to run all example in this repository you need to have

- Linux
  - General
    - `make` command - GNU make utility to maintain groups of programs
  - C
    - C programmming language
    - `gcc` - GNU project C and C++ compiler
  - Go
    - [Go programming language](https://golang.org/doc/install)
    - `go` command
    
- macOS
  - General
    - `brew` command - Brew or Homebrew calls itself The missing package manager for macOS and is an essential tool for any developer.
  - C Family Languages
    - `clang` - Xcode shipped with clang compiler for C, C++, Objective-C, and Swift for free.
    - Xcode required for most macOS development tools to run.
  - Go
    - `go` command
  
  macOS setup tutorial for development in various tools can be found in [macOS Setup - Sourabh Bajaj](http://sourabhbajaj.com/mac-setup/).

## Table of Contents

- C
  - Sort
    - [Sort in one loop](https://github.com/insomnius/programming-test-interview/blob/master/c/sort/sort-in-one-loop.c)

- Go
  - Sort
    - [Sort in one loop](https://github.com/insomnius/programming-test-interview/blob/master/go/sort/sort-in-one-loop.go)
    - [Bubble sort](https://github.com/insomnius/programming-test-interview/blob/master/go/sort/bubble-sort.go)
    - [Selection sort](https://github.com/insomnius/programming-test-interview/blob/master/go/sort/selection-sort.go)
    - [Insertion sort](https://github.com/insomnius/programming-test-interview/blob/master/go/sort/insertion-sort.go)
  - Manipulation
    - [Remove duplicate from sorted array](https://github.com/insomnius/programming-test-interview/blob/master/go/manipulation/remove_duplicate_from_sorted_array.go)
    - [Remove duplicate from unsorted array](https://github.com/insomnius/programming-test-interview/blob/master/go/manipulation/remove_duplicate_from_unsorted_array.go)
    - [Reverse random array](https://github.com/insomnius/programming-test-interview/blob/master/go/manipulation/reverse_random_array.go)
    - [Pig Latin](https://github.com/insomnius/programming-test-interview/blob/master/go/manipulation/pig_latin.go.go)
  - Recursion
    - [Block and Iron](https://github.com/insomnius/programming-test-interview/blob/master/go/recursion/block_and_iron.go)
  - String
    - [Reverse String](https://github.com/insomnius/programming-test-interview/blob/master/go/string/reverse_string.go)

## How to run

### Go

- Sort
  - To run all of go sort example please run `go run go/sort/*` and you will get all of the output sample or just simply use `make go_sort`.
- Manipulation
  - To run all of go sort example please run `go run go/manipulation/*` and you will get all of the output sample or just simply use `make go_manipulation`.
- Recursion
  - To run all of go sort example please run `go run go/recursion/*` and you will get all of the output sample or just simply use `make go_recursion`.
- String
  - To run all of go sort example please run `go run go/string/*` and you will get all of the output sample or just simply use `make go_string`.

### C

- Sort
  - use command `gcc -o c/sort/main c/sort/main.c` to generate compiled file and then run it with `./c/sort/main` or just simply use `make c_sort`

## Contributing

Please read [CONTRIBUTING.md](https://github.com/insomnius/code-geek/blob/master/CONTRIBUTING.md) for details on our code of conduct, and the process for submitting pull requests to us. Check all of people that already contribute in this project in [here](https://github.com/insomnius/code-geek/blob/master/CONTRIBUTOR).

## Authors

- **Insomnius** - Initial work, main GO contributor.

## License

This project is licensed under the MIT License - see the [LICENSE](https://github.com/insomnius/code-geek/blob/master/LICENSE) file for details.
