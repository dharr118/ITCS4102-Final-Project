"""Concurrent task example using Python asyncio."""

import asyncio
import time


async def simulated_task(task_id: int) -> str:
    await asyncio.sleep(1)
    return f"Task {task_id} finished"


async def main() -> None:
    start = time.perf_counter()
    results = await asyncio.gather(*(simulated_task(i) for i in range(1, 6)))
    elapsed = time.perf_counter() - start
    for result in results:
        print(result)
    print(f"Elapsed seconds: {elapsed:.3f}")


if __name__ == "__main__":
    asyncio.run(main())
