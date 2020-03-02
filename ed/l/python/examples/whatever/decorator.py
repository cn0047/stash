def dec(f):
  def wrap(*args, **kwargs):
    r = f(*args, **kwargs)
    print('[dec]', r)
    return r

  return wrap


class Dec2:


  def __call__(self, f):
    def w(*a, **kwa):
      r = f(*a, **kwa)
      print('[Dec2]', r)
      return r
    return w


d2 = Dec2()

@dec
@dec
@d2
def f(s):
  print('handle func')
  return 'foo' + s


print(f('bar'))
