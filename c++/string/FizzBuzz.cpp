// 
// Copyright (c) 2019 Yunindyo Prabowo
// This code is licensed under MIT license (see LICENSE.txt for details)
// 


/*
 *   TASK :
 * -----------------------------------------------------------------------------
 *    Write a program that prints the numbers from 1 to 100.
 *    But for multiples of three print “Fizz” instead of the number,
 *    and for the multiples of five print “Buzz”. 
 *    For numbers which are multiples of both three and five print “FizzBuzz"
 * -----------------------------------------------------------------------------
 */

#include <iostream>
#include <algorithm>
#include <vector>
#include <numeric>

void logicTest(int i){
        if (i%3==0)
                if(i%5==0)
                        std::cout << " FizzBuzz";
                else
                        std::cout << " Fizz";
        else if (i%5==0)
                std::cout << " Buzz";
        else
                std::cout << " " <<i;
}

int main(void){
        // initializing vector array, sizes 100
        std::vector<int> myvec(100);
        
        // generate array 1 to 100 using iota
        // function template from numeric header
        std::iota(myvec.begin(), myvec.end(), 1);
        
        std::cout<<"Test Logic"<<std::endl;

        // Applies function fn to each of the elements in the range [first,last).
        // in this case. logicTest Function to choose conditions that are in accordance
        // with task
        std::for_each(myvec.begin(),myvec.end(),logicTest);
        std::cout<<"\n";
        return 0;
}            