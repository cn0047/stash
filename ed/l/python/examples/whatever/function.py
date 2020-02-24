g = 500


def f0(*args):
  print(args)


def f(n = 0, **args):
  print("g+n=", g+n)
  print("args:", args)


f2 = lambda x: x * 2

f0(1)
f0.myVal = 100
print(f0.myVal)
f(n = 204, type = "test")
n = input("Set n value:")
print(f2(n))
"""
(1,)
100
g+n= 704
args: {'type': 'test'}
Set n value:2
22
"""
