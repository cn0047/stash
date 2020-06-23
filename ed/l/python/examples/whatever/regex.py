import re


def anum(s: str) -> bool:
  return re.search(r'^[\d\w]+$', s)


print(anum('1'))
print(anum('x'))
print(anum('x_'))
print(anum('x-'))
print(anum('='))
