import asyncio


async def f1():
    s = asyncio.Semaphore(value=2)
    await s.acquire()
    print(1)
    await s.acquire()
    print(2)
    await s.acquire()
    print(3)


asyncio.run(f1())
