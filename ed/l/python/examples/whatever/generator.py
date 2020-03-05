def gen123():
    yield 1
    yield 2
    yield 3


g = gen123()
print(next(g))
print(next(g))
print(next(g))
