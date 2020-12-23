from typing import List, Set, Dict, Tuple, Optional


def t1():
  return 'foo1', 'bar1'

def t2():
  return ('foo2', 'bar2')

def t3() -> Tuple[str, str]:
  return 'foo3', 'bar3'

def t():
  f, b = t1()
  print(f'f = {f}, b = {b}')
  (f, b) = t2()
  print(f'f = {f}, b = {b}')
  f, b = t3()
  print(f'f = {f}, b = {b}')


x: int = 1
x: float = 1.0
x: bool = True
x: str = "test"
x: bytes = b"test"

x: List[int] = [1]
x: Set[int] = {6, 7}
x: Dict[str, float] = {'field': 2.0}
x: Tuple[int, ...] = (1, 2, 3)
x: Tuple[int, str, float] = (3, "yes", 7.5)


print(x)
