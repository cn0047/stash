def gen123():
    yield 1
    yield 2
    yield 3

def f1():
  g = gen123()
  f1b(g)

def f1b(g):
  print(next(g))
  print(next(g))
  print(next(g))

def f2():
  g = gen123()
  for k, v in enumerate(g):
    print(v)

def f2b():
  g = gen123()
  for v in g:
    print(v)

def f4():
  g = gen123()
  print('=====')
  for k, v in enumerate(g):
    print(v)
  print('=====')
  for v in g:
    print(v)


f4()
