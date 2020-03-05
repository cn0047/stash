from functools import  singledispatch

class Foo:
  def __init__(self, x: int):
    self.x = x

@singledispatch
def printX(inst):
  raise TypeError("Not implemented for {0}".format(inst))

@printX.register(Foo)
def _(inst):
  print(inst.x)

f = Foo(7)
printX(f)
