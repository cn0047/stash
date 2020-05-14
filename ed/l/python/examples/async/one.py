import asyncio


async def cr1():
  print('hello')
  await asyncio.sleep(1)
  print('world')


def f1():
  asyncio.run(cr1())


if __name__ == '__main__':
  f1()
