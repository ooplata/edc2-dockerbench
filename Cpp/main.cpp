#include <array>
#include <chrono>
#include <cstdint>
#include <fstream>
#include <iostream>

static constexpr bool IsPrime(std::uintmax_t n) {
    for (std::uintmax_t i = 2; i * i <= n; i++) {
        if (n % i == 0)
            return false;
    }
    return true;
}

template<std::size_t N>
static constexpr std::uintmax_t SumPrimes() {
    std::size_t count = 0;
    std::uintmax_t i = 2;
    std::uintmax_t sum = 0;

    while (count < N) {
        if (IsPrime(i)) {
            sum += i;
            count += 1;
        }
        i++;
    }
    return sum;
}

int main() {
    const auto start = std::chrono::system_clock::now();
    constexpr std::size_t Max = 10000;
    constexpr auto sum = SumPrimes<Max>();

    std::ofstream file{};
    file.open("result.txt");
    file << sum;
    file.close();

    const auto end = std::chrono::system_clock::now();
    std::cout << std::chrono::duration_cast<std::chrono::milliseconds>(end - start).count() << "ms";

    return 0;
}
