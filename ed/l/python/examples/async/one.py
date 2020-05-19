import asyncio


async def cr1():
  print('hello')
  await asyncio.sleep(1)
  print('world')


def f1():
  asyncio.run(cr1())


async def cr2() -> None:
  print('hop')
  await asyncio.sleep(1)
  print('hey')


def f2():
  l = asyncio.get_event_loop()
  l.run_until_complete(cr2())
  l.close()


if __name__ == '__main__':
  # f1()
  f2()
