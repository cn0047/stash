class Base1:

  def __init__(self):
    print('Base1.init')


class Base2:

  def __init__(self):
    print('Base2.init')


class Sub(Base1, Base2):
  pass


def f1():
  s = Sub()
  print('[f1]', Sub.__bases__)


f1()
