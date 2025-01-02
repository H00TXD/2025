#include <iostream> // Include the input-output stream library

//math functions 


//add
float addFunc(float x, float y) {
    return x + y;
}

//sub
float subFunc(float x, float y) {
    return  x - y;
}

//divide
float divFunc(float x, float y) {
    return x / y;
}

//mul 
float mulFunc(float x, float y) {
    return x * y;
}

int main() {
    float x = 0;
    float y = 0;

    std::cout << "this should work!!!!" << std::endl; // Output message


    std::cout << "Enter the first number: ";
    std::cin >> x;

    std::cout << "Enter the second number: ";
    std::cin >> y;

      std::cout << addFunc(x,y) << std::endl; // Output message
        std::cout << subFunc(x,y) << std::endl; 
          std::cout << divFunc(x,y) << std::endl; 
            std::cout << mulFunc(x,y) << std::endl; 
    return 0; // Indicate successful program termination
}
