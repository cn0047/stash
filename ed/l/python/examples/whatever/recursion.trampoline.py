import types


def tramp(gen, *args, **kwargs):
  g = gen(*args, **kwargs)
  while isinstance(g, types.GeneratorType):
    g = next(g)
  return g


def f(n, curr=0, next=1):
  if n == 0:
    yield curr
  else:
    yield f(n - 1, next, curr + next)


print(tramp(f, 30))
