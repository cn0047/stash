import asyncio


async def worker(name, queue):
  print('start worker: ', name)
  while True:
    await asyncio.sleep(1)
    s = await queue.get()
    print(s, end='', flush=True)
    queue.task_done()


async def cr1():
  print('[f1]')
  queue = asyncio.Queue()

  queue.put_nowait('a')
  queue.put_nowait('b')
  queue.put_nowait('c')

  task1 = asyncio.create_task(worker('w1', queue))
  task2 = asyncio.create_task(worker('w2', queue))

  await asyncio.sleep(7)

  queue.put_nowait('d')
  queue.put_nowait('e')
  queue.put_nowait('f')

  print('\njoin')
  await queue.join() # Block until all items in the queue have been received and processed.

  task1.cancel()
  task2.cancel()


def f1():
  asyncio.run(cr1())


if __name__ == '__main__':
  f1()
