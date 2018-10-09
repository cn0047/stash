g = 500

def f0(*args):
  print(args)

def f(n = 0, **args):
  print("g+n=", g+n)
  print("args:", args)

f2 = lambda x: x * 2

f(n = 204)
n = input("Set n value:")
print(f2(n))
