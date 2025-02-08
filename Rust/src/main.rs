use std::fs::File;
use std::io::prelude::*;

fn is_prime(n: u64) -> bool {
    if n < 2 {
        return false;
    }

    let mut i = 2;
    loop {
        if n == i {
            return true;
        }

        if n % i == 0 {
            return false;
        }
        i += 1;
    }
}

fn sum_primes(until: u64) -> u64 {
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
    let primes: u64 = sum_primes(MAX);

    let mut file = File::create("out.txt")?;
    file.write_all(primes.to_string().as_bytes())?;

    let elapsed = now.elapsed();
    println!("{}ms", elapsed.as_millis());
    Ok(())
}
