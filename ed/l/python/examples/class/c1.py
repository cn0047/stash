class Student0:
    pass


class Student1:
    @staticmethod
    def f3():
        print("static")


class Student:
    classProperty = "foo"

    def __init__(self):  # initializer not constructor
        print("initialize")

    def f(self, v):
        print("Got:", v)
        Student1.f3()


class Student2(Student, Student0):  # Student2 extends Student & Student0
    # only Student initializer will be called.

    def f2(self, v):
        super().f(v)

    def __call__(self, arg):
        print('__call__ got arg = ', arg)


student = Student2()
student.f2(100)
student('foo')
