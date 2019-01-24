c_sort:
	@gcc -o c/sort/main c/sort/main.c
	@./c/sort/main

cpp_fizzbuzz:
	@g++ -o c++/string/main c++/string/FizzBuzz.cpp
	@./c++/string/main

go_sort:
	@go run go/sort/*

go_recursion:
	@go run go/recursion/*

go_manipulation:
	@go run go/manipulation/*

go_string:
	@go run go/string/*