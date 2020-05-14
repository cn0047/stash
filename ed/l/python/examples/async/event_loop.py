import asyncio


async def cr1():
  print('cr1')


async def printStr(s: str):
  while True:
    await asyncio.sleep(1)
    print(s, end='', flush=True)


def f1():
  print('[f1]')
  loop = asyncio.get_event_loop()
  try:
      loop.run_until_complete(cr1())
  finally:
      loop.close()


def f2():
  print('[f2]')
  loop = asyncio.new_event_loop()
  asyncio.set_event_loop(loop)
  try:
      loop.run_until_complete(cr1())
  finally:
      loop.close()


def f3():
  print('[f3]')
  loop = asyncio.get_event_loop()
  try:
    asyncio.ensure_future(printStr('.'))
    asyncio.ensure_future(printStr('*'))
    loop.run_forever()
  except KeyboardInterrupt:
    pass
  finally:
    loop.close()


if __name__ == '__main__':
  # f1()
  # f2()
  f3()
