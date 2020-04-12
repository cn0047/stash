g = 500


def x():
    return 1, 2


def f1(i: int, s: str='none') -> str:
    return "[f1]: {0} {1}".format(i, s)


def f0(*args):
    print(args)


def f(n=0, **args):
    print("[f]: g+n=", g + n)
    print("[f] args:", args)


def f3(src, ksize, sigmaX, dst=None, sigmaY=None, borderType=None):
    print("[f3] {}, {}, sigmaX = {}, sigmaY = {}".format(src, ksize, sigmaX, sigmaY))

f2 = lambda x: x * 2

print(x())
f0(1)
print(f1(17))
f0.myVal = 100
print(f0.myVal)
f(n=204, type="test")
n = input("Set n value:")
print(f2(n))
f3('z.png', '1920x980', 20, sigmaY=40)

"""
(1, 2)
(1,)
[f1]: 17 none
100
[f]: g+n= 704
[f] args: {'type': 'test'}
Set n value:2
22
"""
