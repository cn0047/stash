# @todo: fix issue here ⚠️

import asyncio
from concurrent.futures import ThreadPoolExecutor


def cr1():
    print('done')


async def f1():
    loop = asyncio.get_event_loop()
    executor = ThreadPoolExecutor(max_workers=1)
    await loop.run_in_executor(executor, cr1)


if __name__ == '__main__':
    asyncio.run(f1())
