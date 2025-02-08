import timeit


def is_prime(n: int) -> bool:
    if n <= 1:
        return False

    if n == 2 or n == 3:
        return True

    if n % 2 == 0 or n % 3 == 0:
        return False

    if (n - 1) % 6 != 0 and (n + 1) % 6 != 0:
        return False

    i: int = 5
    while True:
        if i * i > n:
            return True

        if n % i == 0 or n % (i + 2) == 0:
            return False

        i += 6


def sum_primes(until: int) -> int:
    count: int = 1
    i: int = 3
    sum: int = 2

    while count < until:
        if is_prime(i):
            sum += i
            count += 1

        i += 2

    return sum


def to_bench():
    MAX: int = 10000
    PRIMES: int = sum_primes(MAX)

    with open("out.txt", "w") as file:
        file.write(str(PRIMES))


elapsed = timeit.timeit("to_bench()", globals=locals(), number=1)
print(f"{elapsed * 1000}ms")
