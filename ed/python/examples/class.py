class Student0:
  pass

class Student:
  classProperty = "foo"
  def __init__(self):
    print("construct")
  def f(self, v):
    print("Got:",v)

class Student2(Student): # Student2 extends Student
  def f2(self, v):
    super().f(v)

student = Student2()
student.f2(100)
