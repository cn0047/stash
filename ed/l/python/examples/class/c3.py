class MyClass:

  b = "on class" # class attribute

  def __init__(self):
    self.a = "on instance" # instance attribute
    print("[init]", MyClass.b)

  @property
  def val2(self):
    return self.v2

  @val2.setter
  def val2(self, v):
    self.v2 = v


def f1():
  o = MyClass()
  o.val2 = "foo"
  print("[f1]", o.val2)


f1()
