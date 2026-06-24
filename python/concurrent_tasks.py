import asyncio
import time


async def task(number):
    print(f"Task {number} started")
    await asyncio.sleep(1)
    print(f"Task {number} finished")


async def main():
    start_time = time.time()

    await asyncio.gather(
        task(1),
        task(2),
        task(3),
        task(4),
        task(5)
    )

    elapsed = time.time() - start_time
    print("Elapsed time:", round(elapsed, 2), "seconds")


asyncio.run(main())
