Python
-
<br>2.7.10

[doc](https://www.python.org/)
[packages](https://pypi.org/)

Python - dynamic, ~~strongly typed~~, general-purpose, high-level,
object-oriented, multipurpose programming language.

````sh
python -V
python3 -V

pip -V
pip list
pip show requests
pip search requests
pip install requests
pip install -r requirements.txt

pip install requests virtualenv
````

````sh
python3 # REPL
help("math")
````

#### Data types

* int
* float (64-bit)
* bool
* none

* str
* bytes
* list
* dict
* tuple
* range
* set

Everything is an object.

**str** strings immutable.

**tuple** immutable sequence of arbitrary objects.

**set** unordered collection on unique elements.

#### Features

**comprehensions**.

**generators**.

**generator expression** like comprehensions but generator.

````py
5/2  # 2.5
5//2 # 2

r'ok' # raw string

b'data' # bytes

d = {'foo': 1} # dict - dictionary
md = {"name": "Mark", "active": true}

t = ("f", 2) # tuple

s = {1, 2} # set

arr = ['a', 'b'] # list
len(arr)
del arr[1]
el in arr == true
````

````py
global myvar
nonlocal myvar

any([false, true]) # true
all([false, true]) # false

# partition
o, _, t = "v1:v2".partition(':')
print(o, t) # v1 v2

while c != 0:
  print(c)
  c -= 1

for el in arr:
  print(el)

range(5, 10, 2) == [5, 7, 9]

try:
  f()
except (MyError, NotMyError) as err:
  print(err)
  raise # raise error again
except KeyError as err:
  print(err)
except KeyError2 as err:
  pass
except KeyError3 as err:
  raise
except KeyError4 as err:
  raise IndexError()
finally:
  # ...

yield val # works like in php

import myModule
from myModule import myFunc
from myModule import *
myModule.path # gets path to module

type(myvar) # type of var
isinstance(3, int)
myvar.mro() # info about class instance
hasattr(var, attr)
getattr(var, attr)
setattr(var, attr, val)
delattr(var, attr)
dir(var)

globals()
locals()

if __name__ == '__main__'

__call__()
__str__()
__next__()
__enter__() # enter context (with keyword)
__end__() # exit context (with keyword)

@staticmethod
@abstractmethod
@property # getter
@p.setter # setter
````

````
# package:
my_package
└── __init__.py # package init file
````
