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

static constexpr std::uintmax_t SumPrimes(std::size_t until) {
    std::size_t count = 1;
    std::uintmax_t i = 3;
    std::uintmax_t sum = 2;

    while (count < until) {
        if (IsPrime(i)) {
            sum += i;
            count += 1;
        }
        i += 2;
    }
    return sum;
}

int main() {
    const auto start = std::chrono::system_clock::now();
    constexpr std::size_t Max = 10000;
    constexpr auto sum = SumPrimes(Max);

    std::ofstream file{};
    file.open("out.txt");
    file << sum << std::endl;
    file.close();

    const auto end = std::chrono::system_clock::now();
    std::cout << std::chrono::duration_cast<std::chrono::milliseconds>(end - start).count() << "ms" << std::endl;

    return 0;
}
