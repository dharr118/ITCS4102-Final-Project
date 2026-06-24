import time

def is_prime(number):
    if number < 2:
        return False

    for divisor in range(2, int(number ** 0.5) + 1):
        if number % divisor == 0:
            return False

    return True


start_time = time.time()

primes = []
candidate = 2

while len(primes) < 10000:
    if is_prime(candidate):
        primes.append(candidate)
    candidate += 1

end_time = time.time()

print("Generated 10,000 primes")
print("Last prime:", primes[-1])
print("Elapsed time:", round(end_time - start_time, 4), "seconds")
