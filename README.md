# Code Geek

This repo is dedicated to store all of programing interview test from a very simple one to a very hardest one, algorithm implementation, mini console game like tetris or pacman maybe and all of geeky code. The aim of this repo is to help people to learn algorithm better and find solution that maybe useful in the future.

## Prerequisites

In order to run all example in this repository you need to have

- Linux

  - General
    - `make` command - GNU make utility to maintain groups of programs
  - C
    - C programmming language
    - `gcc` - _GNU Compiler Collections_ project C and C++ compiler
  - C++
    - C++ programming language with standard C++11 or above
    - `gcc` - _GNU Compiler Collections_ project C and C++ compiler
  - Go
    - [Go programming language](https://golang.org/doc/install)
    - `go` command
  - Java
    - [Java programming language](https://www.java.com/en/download/help/download_options.xml)
    - `javac` and `java` - Java project for compile and running

  - Python or Python3
    - Python Programming language
    - `python` or `python3` - Python command for executed python file.

- macOS

  - General
    - `brew` command - Brew or Homebrew calls itself The missing package manager for macOS and is an essential tool for any developer.
  - C/C++ Family Languages
    - `clang` - Xcode shipped with clang compiler for C, C++, Objective-C, and Swift for free.
    - Xcode required for most macOS development tools to run.
  - Go
    - `go` command
  - Python
    - `python3` command should be installed.
  - Java
    - JDK required for compile and running Java project

  macOS setup tutorial for development in various tools can be found in [macOS Setup - Sourabh Bajaj](http://sourabhbajaj.com/mac-setup/).

## Table of Contents

- C
  - Sort
    - [Sort in one loop](https://github.com/insomnius/programming-test-interview/blob/master/c/sort/sort-in-one-loop.c)
- C++
  - String
    - [FizzBuzz Program](/c++/string/FizzBuzz.cpp)
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
    - [Subsequence](https://github.com/insomnius/programming-test-interview/blob/master/go/string/subsequence.go)
    - [Longest Common Subsequence](https://github.com/insomnius/programming-test-interview/blob/master/go/string/longest_common_subsequence.go)
    - [Common Subsequence](https://github.com/insomnius/programming-test-interview/blob/master/go/string/common_subsequence.go)

- Java

  - Sort
    - [Bubble sort](https://github.com/insomnius/programming-test-interview/blob/master/java/sort/BubbleSort.java)
    - [Selection sort](https://github.com/insomnius/programming-test-interview/blob/master/java/sort/SelectionSort.java)
    - [Insertion sort](https://github.com/insomnius/programming-test-interview/blob/master/java/sort/InsertionSort.java)
  - Searching
    - [Sequential searching](https://github.com/insomnius/programming-test-interview/blob/master/java/searching/SequentialSearching.java)
    - [Binary searching](https://github.com/insomnius/programming-test-interview/blob/master/java/searching/BinarySearching.java)
  - Recursion
    - [Triangle recursion](https://github.com/insomnius/programming-test-interview/blob/master/java/recursion/TriangleRecursion.java)

- Python
  - dictionary
    - [Dynamically Importing And Convert into Dictionary]()
    - [Searching Dictionary Inside the List based on Keyword Dictionary]()
  - [Create Numbers Pattern using conditional statment and loop](/python/logic/numberPattern.py)

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

### C++

- String
  - use command `g++ -o c++/string/main c++/string/FizzBuzz.cpp` to generate compiled file and ther run it with `./c++/string/main` or just simply use `make cpp_fizzbuzz`

### Java

- use command `javac filename.java` to generate compiled file and then run it with `java filename.class`.

### Python

- use command `python filename.py arg arg ...` to execute the python file and run it with terminal or command prompt

## Contributing

Please read [CONTRIBUTING.md](https://github.com/insomnius/code-geek/blob/master/CONTRIBUTING.md) for details on our code of conduct, and the process for submitting pull requests to us. Check all of people that already contribute in this project in [here](https://github.com/insomnius/code-geek/blob/master/CONTRIBUTOR).

## Authors

- **Insomnius** - Initial work, main GO contributor.
- **Kzulfazriawan** - python contributor.

## License

This project is licensed under the MIT License - see the [LICENSE](https://github.com/insomnius/code-geek/blob/master/LICENSE) file for details.
