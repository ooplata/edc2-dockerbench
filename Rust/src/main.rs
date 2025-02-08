use std::fs::File;
use std::io::prelude::*;

const fn is_prime(n: u64) -> bool {
    if n <= 1 {
        return false;
    }

    if n == 2 || n == 3 {
        return true;
    }

    if n % 2 == 0 || n % 3 == 0 {
        return false;
    }

    if (n - 1) % 6 != 0 && (n + 1) % 6 != 0 {
        return false;
    }

    let mut i: u64 = 5;
    loop {
        if i * i > n {
            return true;
        }

        if n % i == 0 || n % (i + 2) == 0 {
            return false;
        }
        i += 6;
    }
}

const fn sum_primes(until: u64) -> u64 {
    let mut count: u64 = 1;
    let mut i: u64 = 3;
    let mut sum = 2;

    while count < until {
        if is_prime(i) {
            sum += i;
            count += 1;
        }
        i += 2;
    }
    return sum;
}

fn main() -> std::io::Result<()> {
    use std::time::Instant;
    let now = Instant::now();

    const MAX: u64 = 10000;
    const PRIMES: u64 = sum_primes(MAX);

    let mut file = File::create("out.txt")?;
    file.write_all(PRIMES.to_string().as_bytes())?;

    let elapsed = now.elapsed();
    println!("{}ms", elapsed.as_millis());
    Ok(())
}
