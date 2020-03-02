Python
-
<br>2.7.10

[doc](https://www.python.org/)
[packages](https://pypi.org/)

Python - dynamic, ~~strongly typed~~, general-purpose, high-level,
object-oriented, multipurpose programming language.

Easier to ask for forgiveness than permission (EAFP) - `try` and `except` statements.
Look before you leap (LBYL) - tests for pre-conditions before making calls,
characterized by the presence of many `if` statements.

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

python3 # REPL
help("math")

PYTHONPATH # env var
# PATH=$PATH:~/ed/l/python/examples/pkg python3 pkg
````

````py
global myvar
nonlocal myvar # don't use var from enclosing scope

# partition
o, _, t = "v1:v2".partition(':')
print(o, t) # v1 v2

print('{a}->{b}'.format(a='foo', b='bar')) # foo->bar

# conditional statement
if cond:
  res = value1
else:
  res = value2
# conditional expression
res = value1 if cond else value2

while c != 0:
  print(c)
  c -= 1

for el in arr:
  print(el)

range(5, 10, 2) == [5, 7, 9]

t = (1,2,3)
print(t) # (1, 2, 3)
print(*t) # 1 2 3

def trace(f, *args, **kwargs):
  return f(*args, **kwargs)

yield val # works like in php

globals()
locals()
any([false, true]) # true
all([false, true]) # false
type(myvar) # type of var
isinstance(3, int)
myvar.mro() # info about class instance
hasattr(var, attr)
getattr(var, attr)
setattr(var, attr, val)
delattr(var, attr)
dir(var)
repr(var) # representation, for DBG
str(var)

if __name__ == '__main__'
__doc__
__path__
__closure__

__call__()
__str__(self)
__repr__(self)
__format__(self)
__next__()
__enter__() # enter context (with keyword)
__end__() # exit context (with keyword)

@staticmethod
@abstractmethod
@classmethod
@property # getter
@propertyName.setter # setter
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

````py
5/2  # 2.5
5//2 # 2
Decimal('0.8') - Decimal('0.7')

r'ok' # raw string

b'data' # bytes

md = {'foo': 1, "name": "Mark", "active": true} # dict - dictionary

t = ("f", 2) # tuple

s = {1, 2} # set

arr = ['a', 'b'] # list
len(arr)
del arr[1]
el in arr == true
````
#### Features

**comprehensions**.

**generators**.

**generator expression** like comprehensions but generator.

**module** - usually just a single file with code.

````py
import myModule             # module
from myModule import myFunc
from myModule import *
from .dir import *          # relative import
from . import common        # from common dir in same dir
from ..dir import *         # from parent dir
myModule.path               # gets path to module
````

**package** - module which contains other modules.

````py
# pkg
sys.path.append('pkg_dir')
import pkg
# package dir:
my_package
└── __init__.py # package init file

__main__.py # must be in executable dir
````

**exceptions**

````py
ValueError

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
````

**lambda**

````py
def first_name(name):
  return name.split()[0]

lambda name: name.split()[-1]
````
