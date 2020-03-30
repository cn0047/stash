from typing import Tuple

def t1():
  return 'foo1', 'bar1'

def t2():
  return ('foo2', 'bar2')

def t3() -> Tuple[str, str]:
  return 'foo3', 'bar3'

f, b = t1()
print(f'f = {f}, b = {b}')
(f, b) = t2()
print(f'f = {f}, b = {b}')
f, b = t3()
print(f'f = {f}, b = {b}')
