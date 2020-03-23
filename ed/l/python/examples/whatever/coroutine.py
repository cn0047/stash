def coroutine(func):
  def start(*a, **kwa):
    cr = func(*a, **kwa)
    next(cr)
    return cr
  return start


@coroutine
def my_coroutine(a):
  print(f'--> Started with {a}')
  b = yield
  print(f'--> But continues with {b}')


m = my_coroutine(2)
m.send(5)
