def dec(f):
  def wrap(*args, **kwargs):
    r = f(*args, **kwargs)
    print('[dec]')
    return r

  return wrap

@dec
@dec
def f(s):
  return 'foo' + s

print(f('bar'))
