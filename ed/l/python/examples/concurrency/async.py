import asyncio

async def f() -> None:
  print('hop')
  await asyncio.sleep(1)
  print('hey')


l = asyncio.get_event_loop()
l.run_until_complete(f())
l.close()
