from functools import partial


def mul_1(x, y):
  return x * y


def mul_2(x, y, z):
  return x * y *z


f1 = partial(mul_1, 2)
print(f1(3)) # 6

f2 = partial(mul_2, 2, 10)
print(f2(3)) # 60
