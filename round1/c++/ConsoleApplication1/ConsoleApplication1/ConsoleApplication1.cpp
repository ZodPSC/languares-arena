#include <iostream>
#include <chrono>

int main()
{
    const int countNumbers = 100000000;
    auto start = std::chrono::system_clock::now();
    int *nums = new int [countNumbers];
    for (int i = 0; i < countNumbers; i++) {
        nums[i] = rand() % 10;
    }
    auto end = std::chrono::system_clock::now();

    auto diff = std::chrono::duration_cast<std::chrono::milliseconds>(end - start).count();

    delete[] nums;
    std::cout << "diff = " << diff<<"\n"; 
    //1791 ms = 1 791 000 000 nanosec
    //1831 ms = 1 831 000 000 nanosec
    std::cout << "end";

}
