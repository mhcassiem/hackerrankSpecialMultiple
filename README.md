# hackerrankSpecialMultiple
Using TDD and Golang I'll attempt to solve this problem: https://www.hackerrank.com/challenges/special-multiple/problem

## My solution:

main.go -> runs GetMultiple

ConvertToNineBin -> helper function to convert decimal number to binary and then multiply the result by 9, this is much quicker than iterating through every possible LCM that only contains 9's and 0's.

GetMultiple -> function that returns the LCM of only 9's and 0's. Increments a variable by 1 per cycle, converts to binary and then multiplies by 9. Then checks if result is an LCM of the input variable.