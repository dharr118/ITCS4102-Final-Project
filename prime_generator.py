"""Prime number generator used for programming language comparison."""

import time


def is_prime(number: int) -> bool:
    if number < 2:
        return False
    for divisor in range(2, int(number ** 0.5) + 1):
        if number % divisor == 0:
            return False
    return True


def generate_primes(limit: int) -> list[int]:
    primes = []
    candidate = 2
    while len(primes) < limit:
        if is_prime(candidate):
            primes.append(candidate)
        candidate += 1
    return primes


if __name__ == "__main__":
    start = time.perf_counter()
    result = generate_primes(10000)
    elapsed = time.perf_counter() - start
    print(f"Generated {len(result)} primes")
    print(f"Last prime: {result[-1]}")
    print(f"Elapsed seconds: {elapsed:.6f}")
