from abc import ABCMeta, abstractmethod

# python3
class AValuable(metaclass=ABCMeta):
# python2
# class AValuable(object):
#   __metaclass__ = ABCMeta


  @abstractmethod
  def get_value(self):
    """ Gets value. """


class Foo(AValuable):


  def get_value(self):
    return "foo"


def f1():
  f = Foo()
  print(f.get_value())


f1()
