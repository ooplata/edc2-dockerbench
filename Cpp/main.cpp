#include <chrono>
#include <cstdint>
#include <fstream>
#include <iostream>

static constexpr bool IsPrime(std::uintmax_t n) {
    if (n <= 1) {
        return false;
    }

    if (n == 2 || n == 3) {
        return true;
    }

    if (n % 2 == 0 || n % 3 == 0) {
        return false;
    }

    if ((n - 1) % 6 != 0 && (n + 1) % 6 != 0) {
        return false;
    }

    for (std::uintmax_t i = 5; i * i <= n; i += 6) {
        if (n % i == 0 || n % (i + 2) == 0) {
            return false;
        }
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
    constexpr std::uintmax_t sum = SumPrimes(Max);

    std::ofstream file{};
    file.open("out.txt");
    file << sum << std::endl;
    file.close();

    const auto end = std::chrono::system_clock::now();
    std::cout << std::chrono::duration_cast<std::chrono::milliseconds>(end - start).count() << "ms" << std::endl;

    return 0;
}
