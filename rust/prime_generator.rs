use std::time::Instant;

fn is_prime(number: u64) -> bool {
    if number < 2 {
        return false;
    }

    let limit = (number as f64).sqrt() as u64;

    for divisor in 2..=limit {
        if number % divisor == 0 {
            return false;
        }
    }

    true
}

fn main() {
    let start_time = Instant::now();

    let mut primes: Vec<u64> = Vec::new();
    let mut candidate: u64 = 2;

    while primes.len() < 10000 {
        if is_prime(candidate) {
            primes.push(candidate);
        }
        candidate += 1;
    }

    let elapsed = start_time.elapsed();

    println!("Generated 10,000 primes");
    println!("Last prime: {}", primes[primes.len() - 1]);
    println!("Elapsed time: {:.4} seconds", elapsed.as_secs_f64());
}
