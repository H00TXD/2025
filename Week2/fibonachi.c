#include <stdio.h>


int baseCase(int x) {
    if (x < 0) {
        return -1;  // Handle invalid input (negative numbers)
    }
    if (x == 0) {
        return 0;  // Fibonacci base case: fib(0) = 0
    }
    if (x == 1) {
        return 1;  // Fibonacci base case: fib(1) = 1
    }
    return -1;  // Return an error for cases other than 0 and 1
}

int fib(int y) {
    if (y == 0 || y == 1) {  // Directly handle base cases
        return baseCase(y);
    } else {
        return (fib(y - 1) + fib(y - 2));  // Recursive case
    }
}
int main() {
    int num;

    printf("Enter  number: ");
      scanf("%d", &num);
    
    if(num <= 1 ){
        printf("the num you get is: %d ", baseCase(num));
        return 0;
    }else
    printf("the num you get is: %d ", fib(num));


    return 0;
}
