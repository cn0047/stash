from collections import namedtuple


Point = namedtuple('Point', ['x', 'y'])


def f1():
  p = Point(11, y=22)
  print('[f1]', p.x + p.y)
  print('[f1]', p[0] + p[1])
  x, y = p
  print('[f1]', x + y)


f1()
